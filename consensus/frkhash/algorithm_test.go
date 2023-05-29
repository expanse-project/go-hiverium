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
	"bytes"
	"testing"

	"github.com/expanse-org/go-expanse/common/hexutil"
)

// Tests whether the hashimoto lookup works for both light as well as the full
// datasets.
func TestFrankomoto(t *testing.T) {
	// Create a block to verify
	hash := hexutil.MustDecode("0xc9149cc0386e689d789a1c2f3d5d169a61a6218ed30e74414dc736e442ef3d1f")
	nonce := uint64(0)

	wantDigest := hexutil.MustDecode("0x5c763662184d57157ef85c4672c3a68acd6fd2e35533f55abaa13c238023b506")
	wantResult := hexutil.MustDecode("0x74d692675960275b0523dc248bf3d5783f13e6ec2bc045a661dd2641e95ef2e2")

	digest, result := frankomoto(hash, nonce)
	if !bytes.Equal(digest, wantDigest) {
		t.Errorf("frankomoto digest mismatch: have %x, want %x", digest, wantDigest)
	}
	if !bytes.Equal(result, wantResult) {
		t.Errorf("frankomoto result mismatch: have %x, want %x", result, wantResult)
	}
}

// Benchmarks the light verification performance.
func BenchmarkFrankomoto(b *testing.B) {
	hash := hexutil.MustDecode("0xc9149cc0386e689d789a1c2f3d5d169a61a6218ed30e74414dc736e442ef3d1f")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		frankomoto(hash, 0)
	}
}
