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
	"encoding/binary"

	"github.com/expanse-org/go-expanse/crypto"
)

const (
	datasetInitBytes   = 1 << 30 // Bytes in dataset at genesis
	datasetGrowthBytes = 1 << 23 // Dataset growth per epoch
	cacheInitBytes     = 1 << 24 // Bytes in cache at genesis
	cacheGrowthBytes   = 1 << 17 // Cache growth per epoch
	epochLength        = 30000   // Blocks per epoch
	mixBytes           = 128     // Width of mix
	hashBytes          = 64      // Hash length in bytes
	hashWords          = 16      // Number of 32 bit ints in a hash
	datasetParents     = 256     // Number of parents of each dataset element
	cacheRounds        = 3       // Number of rounds in cache production
	loopAccesses       = 64      // Number of accesses in hashimoto loop
)

func frankomoto(hash []byte, nonce uint64) ([]byte, []byte) {

	// Combine header+nonce into a 64 byte seed
	digest := make([]byte, 40)
	copy(digest, hash)
	binary.LittleEndian.PutUint64(digest[32:], nonce)

	digest = crypto.Keccak512(digest)

	// Here it would be best to change digest into 32 byte slice
	// we could use the first or second half, doesnt really matter
	// Because common.ByteToHash in consensus.go is lobbing off the first 32 bytes
	// We could do this here instead
	// d0 := digest[:32] // gives us the first 32 bytes
	d1 := digest[32:] // gives us the last 32 bytes
	// we could use d0 as the MixDigest and keccak_256(d1) < target
	// return d0, crypto.Keccak256(d1)
	// ORRRRRR We could just change Keccak512 to sha256 or sha3.New256() or sha512_256 or blake2b
	// because they would output a 32byte hash

	return d1, crypto.Keccak256(digest)
}
