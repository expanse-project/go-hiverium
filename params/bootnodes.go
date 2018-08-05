// Copyright 2015 The go-ethereum Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// EXP/DEV Go Bootnodes
	"enode://7f335a047654f3e70d6f91312a7cf89c39704011f1a584e2698250db3d63817e74b88e26b7854111e16b2c9d0c7173c05419aeee2d0321850227b126d8b1be3f@46.101.156.249:42786",
	"enode://df872f81e25f72356152b44cab662caf1f2e57c3a156ecd20e9ac9246272af68a2031b4239a0bc831f2c6ab34733a041464d46b3ea36dce88d6c11714446e06b@178.62.208.109:42786",
	"enode://96d3919b903e7f5ad59ac2f73c43be9172d9d27e2771355db03fd194732b795829a31fe2ea6de109d0804786c39a807e155f065b4b94c6fce167becd0ac02383@45.55.22.34:42786",
	"enode://5f6c625bf287e3c08aad568de42d868781e961cbda805c8397cfb7be97e229419bef9a5a25a75f97632787106bba8a7caf9060fab3887ad2cfbeb182ab0f433f@46.101.182.53:42786",
	"enode://d33a8d4c2c38a08971ed975b750f21d54c927c0bf7415931e214465a8d01651ecffe4401e1db913f398383381413c78105656d665d83f385244ab302d6138414@128.199.183.48:42786",
	"enode://df872f81e25f72356152b44cab662caf1f2e57c3a156ecd20e9ac9246272af68a2031b4239a0bc831f2c6ab34733a041464d46b3ea36dce88d6c11714446e06b@178.62.208.109:42786",
	"enode://f6f0d6b9b7d02ec9e8e4a16e38675f3621ea5e69860c739a65c1597ca28aefb3cec7a6d84e471ac927d42a1b64c1cbdefad75e7ce8872d57548ddcece20afdd1@159.203.64.95:42786",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://0540bc0d32aa96e3ffe9f7a1f46bf37ff4307a046d30eee13649b21348b8037a4ce9d16ab4f277be9a5d170ed0dcbda02a8d06869506d6df2079d45c10761d42@13.84.180.240:42786", // US-TX
	"enode://3fd6043ce992b2f8d1d1d098a2b97187f8fddd93fdead5de4f1204a8d438975ee474aa9f55470c5df2e9d902407cf6b0add2be4767ee6f8f565f55c100c82cca@165.227.99.29:42786",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
