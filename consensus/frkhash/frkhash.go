// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package frkhash implements the frkhash proof-of-work consensus engine.
package frkhash

import (
	"errors"
	"math/big"
	"math/rand"
	"sync"
	"time"
	"unsafe"

	"github.com/expanse-org/go-expanse/consensus"
	"github.com/expanse-org/go-expanse/consensus/ethash"
	"github.com/expanse-org/go-expanse/log"
	"github.com/expanse-org/go-expanse/metrics"
	"github.com/expanse-org/go-expanse/rpc"
)

var ErrInvalidDumpMagic = errors.New("invalid dump magic")

var (
	// two256 is a big integer representing 2^256
	two256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

	// sharedFrkhash is a full instance that can be shared between multiple users.
	sharedFrkhash *Frkhash
)

func init() {
	sharedConfig := Config{
		PowMode: ModeNormal,
	}
	sharedFrkhash = New(sharedConfig, nil, false)
}

// isLittleEndian returns whether the local system is running in little or big
// endian byte order.
func isLittleEndian() bool {
	n := uint32(0x01020304)
	return *(*byte)(unsafe.Pointer(&n)) == 0x04
}

const (
	ModeNormal ethash.Mode = iota
	ModeShared
	ModeTest
	ModeFake
	ModeFullFake
)

// Config are the configuration parameters of the frkhash.
type Config struct {
	PowMode ethash.Mode

	// When set, notifications sent by the remote sealer will
	// be block header JSON objects instead of work package arrays.
	NotifyFull bool

	Log log.Logger `toml:"-"`
}

// Frkhash is a consensus engine based on proof-of-work implementing the frkhash
// algorithm.
type Frkhash struct {
	config Config

	// Mining related fields
	rand     *rand.Rand    // Properly seeded random source for nonces
	threads  int           // Number of threads to mine on if mining
	update   chan struct{} // Notification channel to update mining parameters
	hashrate metrics.Meter // Meter tracking the average hashrate
	remote   *remoteSealer

	// The fields below are hooks for testing
	shared    *Frkhash      // Shared PoW verifier to avoid cache regeneration
	fakeFail  uint64        // Block number which fails PoW check even in fake mode
	fakeDelay time.Duration // Time delay to sleep for before returning from verify

	lock      sync.Mutex // Ensures thread safety for the in-memory caches and mining fields
	closeOnce sync.Once  // Ensures exit channel will not be closed twice.
}

// New creates a full sized frkhash PoW scheme and starts a background thread for
// remote mining, also optionally notifying a batch of remote services of new work
// packages.
func New(config Config, notify []string, noverify bool) *Frkhash {
	if config.Log == nil {
		config.Log = log.Root()
	}

	frkhash := &Frkhash{
		config:   config,
		update:   make(chan struct{}),
		hashrate: metrics.NewMeterForced(),
	}
	if config.PowMode == ModeShared {
		frkhash.shared = sharedFrkhash
	}
	frkhash.remote = startRemoteSealer(frkhash, notify, noverify)
	return frkhash
}

// NewTester creates a small sized frkhash PoW scheme useful only for testing
// purposes.
func NewTester(notify []string, noverify bool) *Frkhash {
	return New(Config{PowMode: ModeTest}, notify, noverify)
}

// NewFaker creates a frkhash consensus engine with a fake PoW scheme that accepts
// all blocks' seal as valid, though they still have to conform to the Ethereum
// consensus rules.
func NewFaker() *Frkhash {
	return &Frkhash{
		config: Config{
			PowMode: ModeFake,
			Log:     log.Root(),
		},
	}
}

// NewFakeFailer creates a frkhash consensus engine with a fake PoW scheme that
// accepts all blocks as valid apart from the single one specified, though they
// still have to conform to the Ethereum consensus rules.
func NewFakeFailer(fail uint64) *Frkhash {
	return &Frkhash{
		config: Config{
			PowMode: ModeFake,
			Log:     log.Root(),
		},
		fakeFail: fail,
	}
}

// NewFakeDelayer creates a frkhash consensus engine with a fake PoW scheme that
// accepts all blocks as valid, but delays verifications by some time, though
// they still have to conform to the Ethereum consensus rules.
func NewFakeDelayer(delay time.Duration) *Frkhash {
	return &Frkhash{
		config: Config{
			PowMode: ModeFake,
			Log:     log.Root(),
		},
		fakeDelay: delay,
	}
}

// NewFullFaker creates an frkhash consensus engine with a full fake scheme that
// accepts all blocks as valid, without checking any consensus rules whatsoever.
func NewFullFaker() *Frkhash {
	return &Frkhash{
		config: Config{
			PowMode: ModeFullFake,
			Log:     log.Root(),
		},
	}
}

// NewShared creates a full sized frkhash PoW shared between all requesters running
// in the same process.
func NewShared() *Frkhash {
	return &Frkhash{shared: sharedFrkhash}
}

// Close closes the exit channel to notify all backend threads exiting.
func (frkhash *Frkhash) Close() error {
	frkhash.closeOnce.Do(func() {
		// Short circuit if the exit channel is not allocated.
		if frkhash.remote == nil {
			return
		}
		close(frkhash.remote.requestExit)
		<-frkhash.remote.exitCh
	})
	return nil
}

// Threads returns the number of mining threads currently enabled. This doesn't
// necessarily mean that mining is running!
func (frkhash *Frkhash) Threads() int {
	frkhash.lock.Lock()
	defer frkhash.lock.Unlock()

	return frkhash.threads
}

// SetThreads updates the number of mining threads currently enabled. Calling
// this method does not start mining, only sets the thread count. If zero is
// specified, the miner will use all cores of the machine. Setting a thread
// count below zero is allowed and will cause the miner to idle, without any
// work being done.
func (frkhash *Frkhash) SetThreads(threads int) {
	frkhash.lock.Lock()
	defer frkhash.lock.Unlock()

	// If we're running a shared PoW, set the thread count on that instead
	if frkhash.shared != nil {
		frkhash.shared.SetThreads(threads)
		return
	}
	// Update the threads and ping any running seal to pull in any changes
	frkhash.threads = threads
	select {
	case frkhash.update <- struct{}{}:
	default:
	}
}

// Hashrate implements PoW, returning the measured rate of the search invocations
// per second over the last minute.
// Note the returned hashrate includes local hashrate, but also includes the total
// hashrate of all remote miner.
func (frkhash *Frkhash) Hashrate() float64 {
	// Short circuit if we are run the frkhash in normal/test mode.
	if frkhash.config.PowMode != ModeNormal && frkhash.config.PowMode != ModeTest {
		return frkhash.hashrate.Rate1()
	}
	var res = make(chan uint64, 1)

	select {
	case frkhash.remote.fetchRateCh <- res:
	case <-frkhash.remote.exitCh:
		// Return local hashrate only if frkhash is stopped.
		return frkhash.hashrate.Rate1()
	}

	// Gather total submitted hash rate of remote sealers.
	return frkhash.hashrate.Rate1() + float64(<-res)
}

// APIs implements consensus.Engine, returning the user facing RPC APIs.
func (frkhash *Frkhash) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	// In order to ensure backward compatibility, we exposes frkhash RPC APIs
	// to both eth and frkhash namespaces.
	return []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   &API{frkhash},
			Public:    true,
		},
		{
			Namespace: "frkhash",
			Version:   "1.0",
			Service:   &API{frkhash},
			Public:    true,
		},
	}
}
