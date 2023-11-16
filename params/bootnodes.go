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

import "github.com/expanse-org/go-expanse/common"

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
	//"enode://f6f0d6b9b7d02ec9e8e4a16e38675f3621ea5e69860c739a65c1597ca28aefb3cec7a6d84e471ac927d42a1b64c1cbdefad75e7ce8872d57548ddcece20afdd1@159.203.64.95:42786",
}

var RebirthBootnodes = []string{
	"enode://54d448a71d491eed10337caba6f593986e01bb09081d1a61d621e91748b872ff5da2161d58b7ca7d87c1dafb4ec59c28b28da1f41b317075296163fa3603c549@137.184.69.182:42786",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://a869b02cec167211fb4815a82941db2e7ed2936fd90e78619c53eb17753fcf0207463e3419c264e2a1dd8786de0df7e68cf99571ab8aeb7c4e51367ef186b1dd@51.15.116.226:30303",
	"enode://807b37ee4816ecf407e9112224494b74dd5933625f655962d892f2f0f02d7fbbb3e2a94cf87a96609526f30c998fd71e93e2f53015c558ffc8b03eceaf30ee33@51.15.119.157:30303",
	"enode://a59e33ccd2b3e52d578f1fbd70c6f9babda2650f0760d6ff3b37742fdcdfdb3defba5d56d315b40c46b70198c7621e63ffa3f987389c7118634b0fefbbdfa7fd@51.15.119.157:40303",
}

// YoloV3Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// YOLOv3 ephemeral test network.
// TODO: Set Yolov3 bootnodes
var YoloV3Bootnodes = []string{
	"enode://9e1096aa59862a6f164994cb5cb16f5124d6c992cdbf4535ff7dea43ea1512afe5448dca9df1b7ab0726129603f1a3336b631e4d7a1a44c94daddd03241587f9@3.9.20.133:30303",
}

var V5Bootnodes = []string{
	"enr:-J24QNp5DQKEMmOUMavO96CJPpD-15rKrkqmfPzH2JOD3hOzK6UFmQ9clpoOANcUriXqAnPLzKXwzoHmn4-b5Fx2eR-BiINldGjHxoQk81dAgIJpZIJ2NIJpcIRo7Ht2g2xlc8CJc2VjcDI1NmsxoQKXlt_hRqPXoXSjf06dakI8zI-0h0YlkoyBWw8c1xNdI4N0Y3CCpyKDdWRwgqci",
}

const dnsPrefix = "enrtree://AJJYVRBUC67IMBJFRJ4ZXOGNKHJK5BWVQF5746DSRG4BZ6CDE54CK@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "expanse"
	case RebirthGenesisHash:
		net = "rebirth"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".exp.blue"
}
