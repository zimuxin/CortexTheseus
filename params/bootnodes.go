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
var MainnetBootnodes = []string{}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Cerebro test network.
var CerebroBootnodes = []string{
	// cortex-jp
	"enode://5091a81e18f722e6aa8fa7fd887836a4e030b6117ec0d11a7024871a4b9c5bfde35ee60063145f0b3bb2ec5def3fae941423fa7dc6d57f3fe23920242ea8d184@47.74.15.143:37566",
	// cortex-de
	"enode://0d6eedee4ed1c4a1b9f273de9ed28fb30581b3bad6b733609870b39ed5a716cc59c5067deb1a43570b15f5965ab22ac04c7067a79aef1e240667d78e8ffdf81d@47.254.135.53:37566",
	// cortex-sg
	"enode://c8119e731b29bed3222ee5e63ec514926fa683f0bc999c73ff171d54a3100aea4211d70640bf937d1d61d39fdf313470b131213ae11ddfea5ccb96669f603749@47.88.174.57:37566",
	// cortex-cn
	"enode://0778ac6fe1a572a2f5834452c1d4b6f77b57a2a4baa72e0e2582ad0dbfe88f41e4e66aa020a451a32936a86c29788d257ccaa19f58e5676943595b945e615b9e@139.224.132.6:37566",

	// cortex-ali
	"enode://a4d20d02a05c3674791f1a0cf9900f0db50485037f96c99395287e931a6845f624e16adbce9e84e7106f803278360979e6a6999b7ca8c25cf64b4445a6de75e6@47.52.39.170:37566",

	// cortex-hk-cpu, cerebro-cn
	"enode://927fa865522a4737e9d773c1ea4fd77ab199872ad8cf42f0790df9a138908c5b94b372d297413bb489765c7241322e4413e0b3444c440e8a4d30652fe5d74116@47.75.211.148:37566",
	// cortex-us-cpu
	"enode://411fe1332751ecfafb77990156d1b8b9573bdf2e3ac099379792a8f296f1f951328c462b1c01f366a6768f020950bd15cf89d66a5f90e8f44a663c49cebdff8c@54.183.146.247:37566",
	// cortex-uk-cpu
	"enode://767ea197d2e19a3f9ec63c5966f1fb625629791428941b58fd290b9ecf232dcb07eec210c8a2207ef3c0eecd4cd5ad515d967c9fedabbd3451db540c46e486c4@18.130.232.151:37566",
	// cortex-jp-cpu
	"enode://3c893c87cef40b090feaad65bae4f69ba8545fd1e0d2212852d85172af6042d07a3373a5187ba93ad2f3c85562a8bac4a3f8969efbfbf4269e0239266be6136b@52.194.191.24:37566",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
