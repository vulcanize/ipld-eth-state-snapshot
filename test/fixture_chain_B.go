package test

import (
	"math/big"

	"github.com/cerc-io/eth-testing/chains/premerge1"
	"github.com/cerc-io/plugeth-statediff/indexer/ipld"
	"github.com/cerc-io/plugeth-statediff/indexer/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

var ChainB = premerge1.ChainData

var ChainB_block1_Header = types.Header{
	ParentHash:  common.HexToHash("0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177"),
	UncleHash:   common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
	Coinbase:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
	Root:        common.HexToHash("0x53580584816f617295ea26c0e17641e0120cab2f0a8ffb53a866fd53aa8e8c2d"),
	TxHash:      common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
	ReceiptHash: common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
	Bloom:       types.Bloom{},
	Difficulty:  big.NewInt(+2),
	Number:      big.NewInt(+1),
	GasLimit:    4704588,
	GasUsed:     0,
	Time:        1492010458,
	Extra:       []byte{215, 131, 1, 6, 0, 132, 103, 101, 116, 104, 135, 103, 111, 49, 46, 55, 46, 51, 133, 108, 105, 110, 117, 120, 0, 0, 0, 0, 0, 0, 0, 0, 159, 30, 250, 30, 250, 114, 175, 19, 140, 145, 89, 102, 198, 57, 84, 74, 2, 85, 230, 40, 142, 24, 140, 34, 206, 145, 104, 193, 13, 190, 70, 218, 61, 136, 180, 170, 6, 89, 48, 17, 159, 184, 134, 33, 11, 240, 26, 8, 79, 222, 93, 59, 196, 141, 138, 163, 139, 202, 146, 228, 252, 197, 33, 81, 0},
	MixDigest:   common.Hash{},
	Nonce:       types.BlockNonce{},
	BaseFee:     nil,
}

var chainB_block1_stateNodeRLP = []byte{248, 113, 160, 147, 141, 92, 6, 119, 63, 191, 125, 121, 193, 230, 153, 223, 49, 102, 109, 236, 50, 44, 161, 215, 28, 224, 171, 111, 118, 230, 79, 99, 18, 99, 4, 160, 117, 126, 95, 187, 60, 115, 90, 36, 51, 167, 59, 86, 20, 175, 63, 118, 94, 230, 107, 202, 41, 253, 234, 165, 214, 221, 181, 45, 9, 202, 244, 148, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 160, 247, 170, 155, 102, 71, 245, 140, 90, 255, 89, 193, 131, 99, 31, 85, 161, 78, 90, 0, 204, 46, 253, 15, 71, 120, 19, 109, 123, 255, 0, 188, 27, 128}
var chainB_block1_stateNodeCID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(chainB_block1_stateNodeRLP))
var block_stateNodeLeafKey = "0x39fc293fc702e42b9c023f094826545db42fc0fdf2ba031bb522d5ef917a6edb"

var ChainB_block1_StateNodeIPLD = models.IPLDModel{
	BlockNumber: ChainB_block1_Header.Number.String(),
	Key:         chainB_block1_stateNodeCID.String(),
	Data:        chainB_block1_stateNodeRLP,
}

var ChainB_block1_EmptyRootNodeRLP, _ = rlp.EncodeToBytes([]byte{})

var ChainB_block1_StateNode0 = models.StateNodeModel{
	BlockNumber: ChainB_block1_Header.Number.String(),
	HeaderID:    ChainB_block1_Header.Hash().Hex(),
	CID:         chainB_block1_stateNodeCID.String(),
	Diff:        false,
	Balance:     "1000",
	Nonce:       1,
	CodeHash:    crypto.Keccak256Hash([]byte{}).Hex(),
	StorageRoot: crypto.Keccak256Hash(ChainB_block1_EmptyRootNodeRLP).Hex(),
	Removed:     false,
	StateKey:    block_stateNodeLeafKey,
}

var chainB_block1_storageNodeRLP = []byte{3, 111, 15, 5, 141, 92, 6, 120, 63, 191, 125, 121, 193, 230, 153, 7, 49, 102, 109, 236, 50, 44, 161, 215, 28, 224, 171, 111, 118, 230, 79, 99, 18, 99, 4, 160, 117, 126, 95, 187, 60, 115, 90, 36, 51, 167, 59, 86, 20, 175, 63, 118, 94, 2, 107, 202, 41, 253, 234, 165, 214, 221, 181, 45, 9, 202, 244, 148, 128, 128, 32, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 160, 247, 170, 155, 102, 245, 71, 140, 90, 255, 89, 131, 99, 99, 31, 85, 161, 78, 90, 0, 204, 46, 253, 15, 71, 120, 19, 109, 123, 255, 0, 188, 27, 128}
var chainB_block1_storageNodeCID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_block1_storageNodeRLP))

var ChainB_block1_StorageNodeIPLD = models.IPLDModel{
	BlockNumber: ChainB_block1_Header.Number.String(),
	Key:         chainB_block1_storageNodeCID.String(),
	Data:        chainB_block1_storageNodeRLP,
}

var ChainB_block1_StorageNode0 = models.StorageNodeModel{
	BlockNumber: ChainB_block1_Header.Number.String(),
	HeaderID:    ChainB_block1_Header.Hash().Hex(),
	StateKey:    block_stateNodeLeafKey,
	StorageKey:  "0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d",
	Removed:     false,
	CID:         chainB_block1_storageNodeCID.String(),
	Diff:        false,
	Value:       []byte{1},
}

// Header for last block at height 32
var ChainB_Block32_Header = types.Header{
	ParentHash:  common.HexToHash("0x6983c921c053d1f637449191379f61ba844013c71e5ebfacaff77f8a8bd97042"),
	UncleHash:   common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
	Coinbase:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
	Root:        common.HexToHash("0xeaa5866eb37e33fc3cfe1376b2ad7f465e7213c14e6834e1cfcef9552b2e5d5d"),
	TxHash:      common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
	ReceiptHash: common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
	Bloom:       types.Bloom{},
	Difficulty:  big.NewInt(2),
	Number:      big.NewInt(32),
	GasLimit:    8253773,
	GasUsed:     0,
	Time:        1658408469,
	Extra:       []byte{216, 131, 1, 10, 19, 132, 103, 101, 116, 104, 136, 103, 111, 49, 46, 49, 56, 46, 50, 133, 108, 105, 110, 117, 120, 0, 0, 0, 0, 0, 0, 0, 113, 250, 240, 25, 148, 32, 193, 94, 196, 10, 99, 63, 251, 130, 170, 0, 176, 201, 149, 55, 230, 58, 218, 112, 84, 153, 122, 83, 134, 52, 176, 99, 53, 54, 63, 12, 226, 81, 38, 176, 57, 117, 92, 205, 237, 81, 203, 232, 220, 228, 166, 254, 206, 136, 7, 253, 2, 61, 47, 217, 235, 24, 140, 92, 1},
	MixDigest:   common.Hash{},
	Nonce:       types.BlockNonce{},
	BaseFee:     nil,
}

// State nodes for all paths at height 32
// Total 7
var ChainB_Block32_stateNode0RLP = []byte{248, 145, 128, 128, 128, 160, 151, 6, 152, 177, 246, 151, 39, 79, 71, 219, 192, 153, 253, 0, 46, 66, 56, 238, 116, 176, 237, 244, 79, 132, 49, 29, 30, 82, 108, 53, 191, 204, 128, 128, 160, 46, 224, 200, 157, 30, 24, 225, 92, 222, 131, 123, 169, 124, 86, 228, 124, 79, 136, 236, 83, 185, 22, 67, 136, 5, 73, 46, 110, 136, 138, 101, 63, 128, 128, 160, 104, 220, 31, 84, 240, 26, 100, 148, 110, 49, 52, 120, 81, 119, 30, 251, 196, 107, 11, 134, 124, 238, 93, 61, 109, 109, 181, 208, 10, 189, 17, 92, 128, 128, 160, 171, 149, 11, 254, 75, 39, 224, 164, 133, 151, 153, 47, 109, 134, 15, 169, 139, 206, 132, 93, 220, 210, 0, 225, 235, 118, 121, 247, 173, 12, 135, 133, 128, 128, 128, 128}
var ChainB_Block32_stateNode0CID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(ChainB_Block32_stateNode0RLP))
var ChainB_Block32_stateNode1RLP = []byte{248, 81, 128, 128, 128, 160, 209, 34, 171, 171, 30, 147, 168, 199, 137, 152, 249, 118, 14, 166, 1, 169, 116, 224, 82, 196, 237, 83, 255, 188, 228, 197, 7, 178, 144, 137, 77, 55, 128, 128, 128, 128, 128, 160, 135, 96, 108, 173, 177, 63, 201, 196, 26, 204, 72, 118, 17, 30, 76, 117, 155, 63, 68, 187, 4, 249, 78, 69, 161, 82, 178, 234, 164, 48, 158, 173, 128, 128, 128, 128, 128, 128, 128}
var ChainB_Block32_stateNode1CID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(ChainB_Block32_stateNode1RLP))
var ChainB_Block32_stateNode2RLP = []byte{248, 105, 160, 32, 21, 58, 188, 102, 126, 135, 59, 96, 54, 200, 164, 107, 221, 132, 126, 42, 222, 63, 137, 185, 51, 28, 120, 239, 37, 83, 254, 161, 148, 197, 13, 184, 70, 248, 68, 1, 128, 160, 168, 127, 48, 6, 204, 116, 51, 247, 216, 182, 191, 182, 185, 124, 223, 202, 239, 15, 67, 91, 253, 165, 42, 2, 54, 10, 211, 250, 242, 149, 205, 139, 160, 224, 22, 140, 8, 116, 27, 79, 113, 64, 185, 215, 180, 38, 38, 236, 164, 5, 87, 211, 15, 88, 153, 138, 185, 94, 186, 125, 137, 164, 198, 141, 192}
var ChainB_Block32_stateNode2CID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(ChainB_Block32_stateNode2RLP))
var ChainB_Block32_stateNode3RLP = []byte{248, 105, 160, 32, 252, 41, 63, 199, 2, 228, 43, 156, 2, 63, 9, 72, 38, 84, 93, 180, 47, 192, 253, 242, 186, 3, 27, 181, 34, 213, 239, 145, 122, 110, 219, 184, 70, 248, 68, 1, 128, 160, 25, 80, 158, 144, 166, 222, 32, 247, 189, 42, 34, 60, 40, 240, 56, 105, 251, 184, 132, 209, 219, 59, 60, 16, 221, 204, 228, 74, 76, 113, 37, 226, 160, 224, 22, 140, 8, 116, 27, 79, 113, 64, 185, 215, 180, 38, 38, 236, 164, 5, 87, 211, 15, 88, 153, 138, 185, 94, 186, 125, 137, 164, 198, 141, 192}
var ChainB_Block32_stateNode3CID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(ChainB_Block32_stateNode3RLP))
var ChainB_Block32_stateNode4RLP = []byte{248, 118, 160, 55, 171, 60, 13, 215, 117, 244, 72, 175, 127, 180, 18, 67, 65, 94, 214, 251, 151, 93, 21, 48, 162, 216, 40, 246, 155, 234, 115, 70, 35, 26, 215, 184, 83, 248, 81, 10, 141, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 160, 86, 232, 31, 23, 27, 204, 85, 166, 255, 131, 69, 230, 146, 192, 248, 110, 91, 72, 224, 27, 153, 108, 173, 192, 1, 98, 47, 181, 227, 99, 180, 33, 160, 197, 210, 70, 1, 134, 247, 35, 60, 146, 126, 125, 178, 220, 199, 3, 192, 229, 0, 182, 83, 202, 130, 39, 59, 123, 250, 216, 4, 93, 133, 164, 112}
var ChainB_Block32_stateNode4CID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(ChainB_Block32_stateNode4RLP))
var ChainB_Block32_stateNode5RLP = []byte{248, 105, 160, 51, 151, 227, 61, 237, 218, 71, 99, 174, 161, 67, 252, 97, 81, 235, 205, 154, 147, 246, 45, 183, 166, 165, 86, 212, 108, 88, 93, 130, 173, 42, 252, 184, 70, 248, 68, 1, 128, 160, 54, 174, 96, 33, 243, 186, 113, 120, 188, 222, 254, 210, 63, 40, 4, 130, 154, 156, 66, 247, 130, 93, 88, 113, 144, 78, 47, 252, 174, 140, 130, 45, 160, 29, 80, 58, 104, 206, 141, 36, 93, 124, 217, 67, 93, 183, 43, 71, 98, 114, 126, 124, 105, 229, 48, 218, 194, 109, 83, 20, 76, 13, 102, 156, 130}
var ChainB_Block32_stateNode5CID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(ChainB_Block32_stateNode5RLP))
var ChainB_Block32_stateNode6RLP = []byte{248, 105, 160, 58, 188, 94, 219, 48, 85, 131, 227, 63, 102, 50, 44, 238, 228, 48, 136, 170, 153, 39, 125, 167, 114, 254, 181, 5, 53, 18, 208, 58, 10, 112, 43, 184, 70, 248, 68, 1, 128, 160, 54, 174, 96, 33, 243, 186, 113, 120, 188, 222, 254, 210, 63, 40, 4, 130, 154, 156, 66, 247, 130, 93, 88, 113, 144, 78, 47, 252, 174, 140, 130, 45, 160, 29, 80, 58, 104, 206, 141, 36, 93, 124, 217, 67, 93, 183, 43, 71, 98, 114, 126, 124, 105, 229, 48, 218, 194, 109, 83, 20, 76, 13, 102, 156, 130}
var ChainB_Block32_stateNode6CID = ipld.Keccak256ToCid(ipld.MEthStateTrie, crypto.Keccak256(ChainB_Block32_stateNode6RLP))

var ChainB_Block32_StateIPLDs = []models.IPLDModel{
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         ChainB_Block32_stateNode0CID.String(),
		Data:        ChainB_Block32_stateNode0RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         ChainB_Block32_stateNode1CID.String(),
		Data:        ChainB_Block32_stateNode1RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         ChainB_Block32_stateNode2CID.String(),
		Data:        ChainB_Block32_stateNode2RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         ChainB_Block32_stateNode3CID.String(),
		Data:        ChainB_Block32_stateNode3RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         ChainB_Block32_stateNode4CID.String(),
		Data:        ChainB_Block32_stateNode4RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         ChainB_Block32_stateNode5CID.String(),
		Data:        ChainB_Block32_stateNode5RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         ChainB_Block32_stateNode6CID.String(),
		Data:        ChainB_Block32_stateNode6RLP,
	},
}

var ChainB_Block32_StateNodes = []models.StateNodeModel{
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		CID:         ChainB_Block32_stateNode2CID.String(),
		Diff:        false,
		Balance:     "0",
		Nonce:       1,
		CodeHash:    common.HexToHash("0xe0168c08741b4f7140b9d7b42626eca40557d30f58998ab95eba7d89a4c68dc0").Hex(),
		StorageRoot: common.HexToHash("0xa87f3006cc7433f7d8b6bfb6b97cdfcaef0f435bfda52a02360ad3faf295cd8b").Hex(),
		Removed:     false,
		StateKey:    "0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d",
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		CID:         ChainB_Block32_stateNode3CID.String(),
		Diff:        false,
		Balance:     "1000",
		Nonce:       1,
		CodeHash:    crypto.Keccak256Hash([]byte{}).Hex(),
		StorageRoot: crypto.Keccak256Hash(ChainB_block1_EmptyRootNodeRLP).Hex(),
		Removed:     false,
		StateKey:    "0x39fc293fc702e42b9c023f094826545db42fc0fdf2ba031bb522d5ef917a6edb",
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		CID:         ChainB_Block32_stateNode4CID.String(),
		Diff:        false,
		Balance:     "1000",
		Nonce:       1,
		CodeHash:    crypto.Keccak256Hash([]byte{}).Hex(),
		StorageRoot: crypto.Keccak256Hash(ChainB_block1_EmptyRootNodeRLP).Hex(),
		Removed:     false,
		StateKey:    "0x67ab3c0dd775f448af7fb41243415ed6fb975d1530a2d828f69bea7346231ad7",
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		CID:         ChainB_Block32_stateNode5CID.String(),
		Diff:        false,
		Balance:     "1000",
		Nonce:       1,
		CodeHash:    crypto.Keccak256Hash([]byte{}).Hex(),
		StorageRoot: crypto.Keccak256Hash(ChainB_block1_EmptyRootNodeRLP).Hex(),
		Removed:     false,
		StateKey:    "0x9397e33dedda4763aea143fc6151ebcd9a93f62db7a6a556d46c585d82ad2afc",
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		CID:         ChainB_Block32_stateNode6CID.String(),
		Diff:        false,
		Balance:     "0",
		Nonce:       1,
		CodeHash:    common.HexToHash("0x1d503a68ce8d245d7cd9435db72b4762727e7c69e530dac26d53144c0d669c82").Hex(),
		StorageRoot: common.HexToHash("0x36ae6021f3ba7178bcdefed23f2804829a9c42f7825d5871904e2ffcae8c822d").Hex(),
		Removed:     false,
		StateKey:    "0xcabc5edb305583e33f66322ceee43088aa99277da772feb5053512d03a0a702b",
	},
}

// Storage nodes for all paths at height 32
// Total 18
var chainB_Block32_storageNode0RLP = []byte{248, 145, 128, 128, 128, 128, 160, 46, 77, 227, 140, 57, 224, 108, 238, 40, 82, 145, 79, 210, 174, 54, 248, 0, 145, 137, 64, 229, 230, 148, 145, 250, 132, 89, 198, 8, 249, 245, 133, 128, 160, 146, 250, 117, 217, 106, 75, 51, 124, 196, 244, 29, 16, 47, 173, 5, 90, 86, 19, 15, 48, 179, 174, 60, 171, 112, 154, 92, 70, 232, 164, 141, 165, 128, 160, 107, 250, 27, 137, 190, 180, 7, 172, 62, 97, 13, 157, 215, 114, 55, 219, 14, 244, 163, 155, 192, 255, 34, 143, 154, 149, 33, 227, 166, 135, 164, 93, 128, 128, 128, 160, 173, 131, 221, 2, 30, 147, 11, 230, 58, 166, 18, 25, 90, 56, 198, 126, 196, 130, 131, 1, 213, 112, 129, 155, 96, 143, 121, 231, 218, 97, 216, 200, 128, 128, 128, 128}
var chainB_Block32_storageNode0CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode0RLP))
var chainB_Block32_storageNode1RLP = []byte{248, 81, 160, 167, 145, 134, 15, 219, 140, 96, 62, 101, 242, 176, 129, 164, 160, 200, 221, 13, 1, 246, 167, 156, 45, 205, 192, 88, 236, 235, 80, 105, 178, 123, 2, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 160, 18, 136, 22, 150, 26, 170, 67, 152, 182, 246, 95, 49, 193, 199, 219, 163, 97, 25, 243, 70, 126, 235, 163, 59, 44, 16, 37, 37, 247, 50, 229, 70, 128, 128}
var chainB_Block32_storageNode1CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode1RLP))
var chainB_Block32_storageNode2RLP = []byte{236, 160, 32, 87, 135, 250, 18, 168, 35, 224, 242, 183, 99, 28, 196, 27, 59, 168, 130, 139, 51, 33, 202, 129, 17, 17, 250, 117, 205, 58, 163, 187, 90, 206, 138, 137, 54, 53, 201, 173, 197, 222, 160, 0, 0}
var chainB_Block32_storageNode2CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode2RLP))
var chainB_Block32_storageNode3RLP = []byte{226, 160, 32, 44, 236, 111, 71, 132, 84, 126, 80, 66, 161, 99, 128, 134, 227, 24, 137, 41, 243, 79, 60, 0, 5, 248, 222, 195, 102, 201, 110, 129, 149, 172, 100}
var chainB_Block32_storageNode3CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode3RLP))
var chainB_Block32_storageNode4RLP = []byte{236, 160, 58, 160, 42, 17, 221, 77, 37, 151, 49, 139, 113, 212, 147, 177, 69, 221, 246, 174, 8, 23, 169, 211, 148, 127, 69, 213, 41, 166, 167, 95, 43, 239, 138, 137, 54, 53, 201, 173, 197, 222, 159, 255, 156}
var chainB_Block32_storageNode4CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode4RLP))
var chainB_Block32_storageNode5RLP = []byte{248, 67, 160, 58, 53, 172, 251, 193, 95, 248, 26, 57, 174, 125, 52, 79, 215, 9, 242, 142, 134, 0, 180, 170, 140, 101, 198, 182, 75, 254, 127, 227, 107, 209, 155, 161, 160, 71, 76, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6}
var chainB_Block32_storageNode5CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode5RLP))
var chainB_Block32_storageNode6RLP = []byte{248, 67, 160, 58, 53, 172, 251, 193, 95, 248, 26, 57, 174, 125, 52, 79, 215, 9, 242, 142, 134, 0, 180, 170, 140, 101, 198, 182, 75, 254, 127, 227, 107, 209, 155, 161, 160, 71, 76, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6}
var chainB_Block32_storageNode6CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode6RLP))
var chainB_Block32_storageNode7RLP = []byte{248, 67, 160, 50, 87, 90, 14, 158, 89, 60, 0, 249, 89, 248, 201, 47, 18, 219, 40, 105, 195, 57, 90, 59, 5, 2, 208, 94, 37, 22, 68, 111, 113, 248, 91, 161, 160, 71, 111, 108, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8}
var chainB_Block32_storageNode7CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode7RLP))
var chainB_Block32_storageNode8RLP = []byte{248, 67, 160, 50, 87, 90, 14, 158, 89, 60, 0, 249, 89, 248, 201, 47, 18, 219, 40, 105, 195, 57, 90, 59, 5, 2, 208, 94, 37, 22, 68, 111, 113, 248, 91, 161, 160, 71, 111, 108, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8}
var chainB_Block32_storageNode8CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode8RLP))
var chainB_Block32_storageNode9RLP = []byte{248, 145, 128, 128, 128, 128, 160, 145, 86, 15, 219, 52, 36, 164, 68, 160, 227, 156, 111, 1, 245, 112, 184, 187, 242, 26, 138, 8, 98, 129, 35, 57, 212, 165, 21, 204, 151, 229, 43, 128, 160, 250, 205, 84, 126, 141, 108, 126, 228, 162, 8, 238, 234, 141, 159, 232, 175, 70, 112, 207, 55, 165, 209, 107, 153, 54, 183, 60, 172, 194, 251, 66, 61, 128, 160, 107, 250, 27, 137, 190, 180, 7, 172, 62, 97, 13, 157, 215, 114, 55, 219, 14, 244, 163, 155, 192, 255, 34, 143, 154, 149, 33, 227, 166, 135, 164, 93, 128, 128, 128, 160, 173, 131, 221, 2, 30, 147, 11, 230, 58, 166, 18, 25, 90, 56, 198, 126, 196, 130, 131, 1, 213, 112, 129, 155, 96, 143, 121, 231, 218, 97, 216, 200, 128, 128, 128, 128}
var chainB_Block32_storageNode9CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode9RLP))
var chainB_Block32_storageNode10RLP = []byte{236, 160, 48, 87, 135, 250, 18, 168, 35, 224, 242, 183, 99, 28, 196, 27, 59, 168, 130, 139, 51, 33, 202, 129, 17, 17, 250, 117, 205, 58, 163, 187, 90, 206, 138, 137, 54, 53, 201, 173, 197, 222, 160, 0, 0}
var chainB_Block32_storageNode10CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode10RLP))
var chainB_Block32_storageNode11RLP = []byte{236, 160, 58, 160, 42, 17, 221, 77, 37, 151, 49, 139, 113, 212, 147, 177, 69, 221, 246, 174, 8, 23, 169, 211, 148, 127, 69, 213, 41, 166, 167, 95, 43, 239, 138, 137, 54, 53, 201, 173, 197, 222, 160, 0, 0}
var chainB_Block32_storageNode11CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode11RLP))
var chainB_Block32_storageNode12RLP = []byte{248, 81, 128, 128, 160, 79, 197, 241, 58, 178, 249, 186, 12, 45, 168, 139, 1, 81, 171, 14, 124, 244, 216, 93, 8, 204, 164, 92, 205, 146, 60, 106, 183, 99, 35, 235, 40, 128, 128, 128, 128, 128, 128, 128, 128, 160, 82, 154, 228, 80, 107, 126, 132, 72, 3, 170, 88, 197, 100, 216, 50, 21, 226, 183, 86, 42, 208, 239, 184, 183, 152, 93, 188, 113, 224, 234, 218, 43, 128, 128, 128, 128, 128}
var chainB_Block32_storageNode12CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode12RLP))
var chainB_Block32_storageNode13RLP = []byte{248, 81, 128, 128, 160, 79, 197, 241, 58, 178, 249, 186, 12, 45, 168, 139, 1, 81, 171, 14, 124, 244, 216, 93, 8, 204, 164, 92, 205, 146, 60, 106, 183, 99, 35, 235, 40, 128, 128, 128, 128, 128, 128, 128, 128, 160, 82, 154, 228, 80, 107, 126, 132, 72, 3, 170, 88, 197, 100, 216, 50, 21, 226, 183, 86, 42, 208, 239, 184, 183, 152, 93, 188, 113, 224, 234, 218, 43, 128, 128, 128, 128, 128}
var chainB_Block32_storageNode13CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode13RLP))
var chainB_Block32_storageNode14RLP = []byte{226, 160, 57, 13, 236, 217, 84, 139, 98, 168, 214, 3, 69, 169, 136, 56, 111, 200, 75, 166, 188, 149, 72, 64, 8, 246, 54, 47, 147, 22, 14, 243, 229, 99, 1}
var chainB_Block32_storageNode14CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode14RLP))
var chainB_Block32_storageNode15RLP = []byte{226, 160, 57, 13, 236, 217, 84, 139, 98, 168, 214, 3, 69, 169, 136, 56, 111, 200, 75, 166, 188, 149, 72, 64, 8, 246, 54, 47, 147, 22, 14, 243, 229, 99, 1}
var chainB_Block32_storageNode15CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode15RLP))
var chainB_Block32_storageNode16RLP = []byte{226, 160, 49, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246, 4}
var chainB_Block32_storageNode16CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode16RLP))
var chainB_Block32_storageNode17RLP = []byte{226, 160, 49, 14, 45, 82, 118, 18, 7, 59, 38, 238, 205, 253, 113, 126, 106, 50, 12, 244, 75, 74, 250, 194, 176, 115, 45, 159, 203, 226, 183, 250, 12, 246, 4}
var chainB_Block32_storageNode17CID = ipld.Keccak256ToCid(ipld.MEthStorageTrie, crypto.Keccak256(chainB_Block32_storageNode17RLP))

var ChainB_Block32_StorageIPLDs = []models.IPLDModel{
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode0CID.String(),
		Data:        chainB_Block32_storageNode0RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode1CID.String(),
		Data:        chainB_Block32_storageNode1RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode2CID.String(),
		Data:        chainB_Block32_storageNode2RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode3CID.String(),
		Data:        chainB_Block32_storageNode3RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode4CID.String(),
		Data:        chainB_Block32_storageNode4RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode5CID.String(),
		Data:        chainB_Block32_storageNode5RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode6CID.String(),
		Data:        chainB_Block32_storageNode6RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode7CID.String(),
		Data:        chainB_Block32_storageNode7RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode8CID.String(),
		Data:        chainB_Block32_storageNode8RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode9CID.String(),
		Data:        chainB_Block32_storageNode9RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode10CID.String(),
		Data:        chainB_Block32_storageNode10RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode11CID.String(),
		Data:        chainB_Block32_storageNode11RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode12CID.String(),
		Data:        chainB_Block32_storageNode12RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode13CID.String(),
		Data:        chainB_Block32_storageNode13RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode14CID.String(),
		Data:        chainB_Block32_storageNode14RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode15CID.String(),
		Data:        chainB_Block32_storageNode15RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode16CID.String(),
		Data:        chainB_Block32_storageNode16RLP,
	},
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		Key:         chainB_Block32_storageNode17CID.String(),
		Data:        chainB_Block32_storageNode17RLP,
	},
}

var ChainB_Block32_StorageNodes = []models.StorageNodeModel{
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace",
		CID:         chainB_Block32_storageNode2CID.String(),
		Value:       []byte{137, 54, 53, 201, 173, 197, 222, 160, 0, 0},
		StateKey:    "0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d",
	}, // 0
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x4e2cec6f4784547e5042a1638086e3188929f34f3c0005f8dec366c96e8195ac",
		CID:         chainB_Block32_storageNode3CID.String(),
		Value:       []byte{100},
		StateKey:    "0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d",
	}, // 1
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x6aa02a11dd4d2597318b71d493b145ddf6ae0817a9d3947f45d529a6a75f2bef",
		CID:         chainB_Block32_storageNode4CID.String(),
		Value:       []byte{137, 54, 53, 201, 173, 197, 222, 159, 255, 156},
		StateKey:    "0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d",
	}, // 2
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b",
		CID:         chainB_Block32_storageNode5CID.String(),
		Value:       []byte{},
		StateKey:    "0x39fc293fc702e42b9c023f094826545db42fc0fdf2ba031bb522d5ef917a6edb'",
	}, // 3
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b",
		CID:         chainB_Block32_storageNode6CID.String(),
		Value:       []byte{160, 71, 76, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6},
		StateKey:    "0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d",
	}, // 4
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0xc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b",
		CID:         chainB_Block32_storageNode7CID.String(),
		Value:       []byte{},
		StateKey:    "0x39fc293fc702e42b9c023f094826545db42fc0fdf2ba031bb522d5ef917a6edb'",
	}, // 5
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0xc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b",
		CID:         chainB_Block32_storageNode8CID.String(),
		Value:       []byte{160, 71, 111, 108, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8},
		StateKey:    "0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d",
	}, // 6
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace",
		CID:         chainB_Block32_storageNode10CID.String(),
		Value:       []byte{},
		StateKey:    "0x39fc293fc702e42b9c023f094826545db42fc0fdf2ba031bb522d5ef917a6edb'",
	}, // 7
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x6aa02a11dd4d2597318b71d493b145ddf6ae0817a9d3947f45d529a6a75f2bef",
		CID:         chainB_Block32_storageNode11CID.String(),
		Value:       []byte{},
		StateKey:    "0x39fc293fc702e42b9c023f094826545db42fc0fdf2ba031bb522d5ef917a6edb'",
	}, // 8
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563",
		CID:         chainB_Block32_storageNode14CID.String(),
		Value:       []byte{'\x01'},
		StateKey:    "0xcabc5edb305583e33f66322ceee43088aa99277da772feb5053512d03a0a702b",
	}, // 9
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563",
		CID:         chainB_Block32_storageNode15CID.String(),
		Value:       []byte{},
		StateKey:    "0x9397e33dedda4763aea143fc6151ebcd9a93f62db7a6a556d46c585d82ad2afc",
	}, // 10
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6",
		CID:         chainB_Block32_storageNode16CID.String(),
		Value:       []byte{'\x04'},
		StateKey:    "0xcabc5edb305583e33f66322ceee43088aa99277da772feb5053512d03a0a702b",
	}, // 11
	{
		BlockNumber: ChainB_Block32_Header.Number.String(),
		HeaderID:    ChainB_Block32_Header.Hash().Hex(),
		Diff:        false,
		Removed:     false,
		StorageKey:  "0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6",
		CID:         chainB_Block32_storageNode17CID.String(),
		Value:       []byte{},
		StateKey:    "0x9397e33dedda4763aea143fc6151ebcd9a93f62db7a6a556d46c585d82ad2afc",
	}, // 12
}

// Contracts used in chainB
/*
pragma solidity ^0.8.0;

contract Test {
    uint256 private count;
    uint256 private count2;

    event Increment(uint256 count);

    constructor() {
      count2 = 4;
    }

    function incrementCount() public returns (uint256) {
      count = count + 1;
      emit Increment(count);

      return count;
    }

    function destroy() public {
        selfdestruct(payable(msg.sender));
    }

    function deleteCount2() public {
        count2 = 0;
    }
}
*/

/*
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract GLDToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("Gold", "GLD") {
        _mint(msg.sender, initialSupply);
    }
}
*/
