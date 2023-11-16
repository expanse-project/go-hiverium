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

package frkhash

import (
	"math/big"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/expanse-org/go-expanse/common"
	"github.com/expanse-org/go-expanse/common/hexutil"
	"github.com/expanse-org/go-expanse/core/types"
)

// Tests that frkhash works correctly in test mode.
func TestFrkhashTestMode(t *testing.T) {
	header := &types.Header{Number: big.NewInt(1), Difficulty: big.NewInt(100)}

	frkhash := NewTester(nil, false)
	defer frkhash.Close()

	results := make(chan *types.Block)
	err := frkhash.Seal(nil, types.NewBlockWithHeader(header), results, nil)
	if err != nil {
		t.Fatalf("failed to seal block: %v", err)
	}
	select {
	case block := <-results:
		header.Nonce = types.EncodeNonce(block.Nonce())
		header.MixDigest = block.MixDigest()
		if err := frkhash.verifySeal(nil, header, false); err != nil {
			t.Fatalf("unexpected verification error: %v", err)
		}
	case <-time.NewTimer(4 * time.Second).C:
		t.Error("sealing result timeout")
	}
}

func verifyFrkhashTest(wg *sync.WaitGroup, e *Frkhash, workerIndex, epochs int) {
	defer wg.Done()

	const wiggle = 4 * epochLength
	r := rand.New(rand.NewSource(int64(workerIndex)))
	for epoch := 0; epoch < epochs; epoch++ {
		block := int64(epoch)*epochLength - wiggle/2 + r.Int63n(wiggle)
		if block < 0 {
			block = 0
		}
		header := &types.Header{Number: big.NewInt(block), Difficulty: big.NewInt(100)}
		e.verifySeal(nil, header, false)
	}
}

func TestFrkhashRemoteSealer(t *testing.T) {
	frkhash := NewTester(nil, false)
	defer frkhash.Close()

	api := &API{frkhash}
	if _, err := api.GetWork(); err != errNoMiningWork {
		t.Error("expect to return an error indicate there is no mining work")
	}
	header := &types.Header{Number: big.NewInt(1), Difficulty: big.NewInt(100)}
	block := types.NewBlockWithHeader(header)
	sealhash := frkhash.SealHash(header)

	// Push new work.
	results := make(chan *types.Block)
	frkhash.Seal(nil, block, results, nil)

	var (
		work [4]string
		err  error
	)
	if work, err = api.GetWork(); err != nil || work[0] != sealhash.Hex() {
		t.Error("expect to return a mining work has same hash")
	}

	if res := api.SubmitWork(types.BlockNonce{}, sealhash, common.Hash{}); res {
		t.Error("expect to return false when submit a fake solution")
	}
	// Push new block with same block number to replace the original one.
	header = &types.Header{Number: big.NewInt(1), Difficulty: big.NewInt(1000)}
	block = types.NewBlockWithHeader(header)
	sealhash = frkhash.SealHash(header)
	frkhash.Seal(nil, block, results, nil)

	if work, err = api.GetWork(); err != nil || work[0] != sealhash.Hex() {
		t.Error("expect to return the latest pushed work")
	}
}

func TestFrkhashHashrate(t *testing.T) {
	var (
		hashrate = []hexutil.Uint64{100, 200, 300}
		expect   uint64
		ids      = []common.Hash{common.HexToHash("a"), common.HexToHash("b"), common.HexToHash("c")}
	)
	frkhash := NewTester(nil, false)
	defer frkhash.Close()

	if tot := frkhash.Hashrate(); tot != 0 {
		t.Error("expect the result should be zero")
	}

	api := &API{frkhash}
	for i := 0; i < len(hashrate); i += 1 {
		if res := api.SubmitHashrate(hashrate[i], ids[i]); !res {
			t.Error("remote miner submit hashrate failed")
		}
		expect += uint64(hashrate[i])
	}
	if tot := frkhash.Hashrate(); tot != float64(expect) {
		t.Error("expect total hashrate should be same")
	}
}

func TestFrkhashClosedRemoteSealer(t *testing.T) {
	frkhash := NewTester(nil, false)
	time.Sleep(1 * time.Second) // ensure exit channel is listening
	frkhash.Close()

	api := &API{frkhash}
	if _, err := api.GetWork(); err != errEthashStopped {
		t.Error("expect to return an error to indicate frkhash is stopped")
	}

	if res := api.SubmitHashrate(hexutil.Uint64(100), common.HexToHash("a")); res {
		t.Error("expect to return false when submit hashrate to a stopped frkhash")
	}
}
