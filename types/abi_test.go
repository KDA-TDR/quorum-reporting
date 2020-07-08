package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventParsing(t *testing.T) {
	//Tests all the events parse correctly from ABIParsingContract.sol
	//The testTx is the tx from deploying the contract

	testTx := `{"hash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","status":true,"blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","index":0,"nonce":0,"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d","to":"0x0000000000000000000000000000000000000000","value":0,"gas":4700000,"gasPrice":0,"gasUsed":1891109,"cumulativeGasUsed":1891109,"createdContract":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","data":"","privateData":"0x","isPrivate":false,"timestamp":1594301001,"events":[{"index":0,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x35d7282d232693a58e27f2460e4aacced5663fd2e86b4645690bbab931e4772d"],"data":"0x123456789012345678901234567890123456789012345678901234567890123474000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":1,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x97db7e059ed2946cf7c8383948665db73baf23d369fc6b8eafb56fd7ca80b3cf"],"data":"0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":2,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0xc1865ec3e31fe00f17669214b0ac9bc55c1c34093a55b276a096eceb3598d8d1"],"data":"0xfffffffffffffffffffffffffffffffffffffffffffffffaa55ab2c71ad981160000000000000000000000000000000000000000000000000000000000003039ffffffffffffffffffffffffffffffffffffffffffffffffff642fe40c8aacfa000000000000000000000000000000000000000000000000ab54a98ceb1f0ad2","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":3,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x32021c01abdbd238344b0243e5256ab75d2e769dbea4b2b2fb56bd8738196c2c"],"data":"0x0000000000000000000000001932c48b2bf8102ba33b4a6b545c32236e342f340000000000000000000000009d13c6d3afe1721beef56b55d303b09e021e27ab","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":4,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x20cf3f101b104cb006fad7296551822b55d258a055008a81103d9ab3f233f644"],"data":"0x0000000000000000000000000000000000000000000000055aa54d38e5267eea0000000000000000000000000000000000000000000000000000000000003039000000000000000000000000000000000000000000000000009bd01bf3755306000000000000000000000000000000000000000000000000ab54a98ceb1f0ad2","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":5,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x256ec87dcd70c23fc01ff7a2f63fb7fb0073d152d0e0254ff578a46a631dd316"],"data":"0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000005736d616c6c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005a736f6d65207265616c6c79206c6172676520737472696e6720746861742077696c6c20676f206f76657220746865207468697274792d74776f2062797465206c696d697420666f7220612073696e676c65207661726961626c65000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":6,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0xe347e4329ff6519c831591d74f30a2748a4f38560924ee6970c881edc23ddd8d"],"data":"0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000064000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f60616263000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f6061626300000000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":7,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0xf62c419906de86aa15be7318a99c09b710c7e4d500fac892ce54ae145611728b"],"data":"0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000c8000000000000000000000000000000000000000000000000000000000000002c000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000000f4000000000000000000000000000000000000000000000000000000000000005800000000000000000000000000000000000000000000000000000000000000bc00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000084000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":8,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x9348b0e10a05565b31d4a3d3550bbd3c36fdbc9cb8eff07b960d7802910e6489"],"data":"0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000c8000000000000000000000000000000000000000000000000000000000000002c000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000000f4000000000000000000000000000000000000000000000000000000000000005800000000000000000000000000000000000000000000000000000000000000bc00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000084000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":9,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x06e13b4a8583b0c888cf9710d4e1e78378e52b39d0775c443acc85c8ccfe44ec"],"data":"0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000060123456789012345678901234567890123456789012345678901234567890123400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000005666972737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000604013487654674538507684738547847680974039786439857345674358096798000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000067365636f6e640000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":10,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0xe6f5c700a9879af36c5780605691c742a0bc6376e86f37727ffb72a27d0203fc"],"data":"0x00000000000000000000000000000000000000000000000000000000000000a043865789746478086504605430483574304038976310674530640434765847310000000000000000000000000000000000000000000000000000000000000d7f000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000060123456789012345678901234567890123456789012345678901234567890123400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000006717765727479000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000604013487654674538507684738547847680974039786439857345674358096798000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000096173646667686a6b6c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c737472696e672066696674680000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":11,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x5a5add0d273fac0599b39c4a0362e7f6c3aea02ddb3bec3a566aa057da3ffe4b"],"data":"0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000280000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000601234567890123456789012345678901234567890123456789012345678901234000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000067177657274790000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006012345678901234567890123456789012345678901234567890123456789012340000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000671776572747900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000601234567890123456789012345678901234567890123456789012345678901234000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000067177657274790000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":12,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x46f092a93f590e8af8ab57315116b41765e5fc7a88509f9aa66173774c0532a2"],"data":"0x0000000000000000000000000000000000000000000000000000000000003c3f0800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c126000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001},{"index":13,"address":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","topics":["0x5ef9e93d32c1a54a219a2e1083104d924cea95596bb1ba7c1a70f963a8976b39"],"data":"0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003c3f080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000003c3f080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000003c3f080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000003c3f08000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000003c3f080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000003c3f08000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001","blockNumber":1,"blockHash":"0xc7fd1915b4b8ac6344e750e4eaeacf9114d4e185f9c10b6b3bc7049511a96998","transactionHash":"0xbc77a72b3409ba3e098cb45bac1b7727b59dae9a05f37a0dbc61007949c8cede","transactionIndex":0,"timestamp":1594301001}],"internalCalls":null}`
	abiTx := `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"first","type":"address"},{"indexed":false,"internalType":"address","name":"second","type":"address"}],"name":"AddressFixed","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256[]","name":"first","type":"uint256[]"},{"indexed":false,"internalType":"bool[]","name":"second","type":"bool[]"}],"name":"ArrayDynamicSize","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256[10]","name":"first","type":"uint256[10]"},{"indexed":false,"internalType":"bool[6]","name":"second","type":"bool[6]"}],"name":"ArrayFixedSize","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bool","name":"first","type":"bool"},{"indexed":false,"internalType":"bool","name":"second","type":"bool"}],"name":"BoolFixed","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes","name":"first","type":"bytes"},{"indexed":false,"internalType":"bytes","name":"second","type":"bytes"}],"name":"BytesFixed","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"first","type":"bytes32"},{"indexed":false,"internalType":"bytes1","name":"second","type":"bytes1"},{"indexed":false,"internalType":"bytes1","name":"third","type":"bytes1"}],"name":"BytesFixedSize","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"int256","name":"first","type":"int256"},{"indexed":false,"internalType":"int16","name":"second","type":"int16"},{"indexed":false,"internalType":"int64","name":"third","type":"int64"},{"indexed":false,"internalType":"int256","name":"fourth","type":"int256"}],"name":"IntFixed","type":"event"},{"anonymous":false,"inputs":[{"components":[{"internalType":"string","name":"first","type":"string"},{"internalType":"bytes32","name":"second","type":"bytes32"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.Custom","name":"first","type":"tuple"},{"indexed":false,"internalType":"bytes32","name":"second","type":"bytes32"},{"indexed":false,"internalType":"int16","name":"third","type":"int16"},{"components":[{"internalType":"string","name":"first","type":"string"},{"internalType":"bytes32","name":"second","type":"bytes32"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.Custom","name":"fourth","type":"tuple"},{"indexed":false,"internalType":"string","name":"fifth","type":"string"}],"name":"Mixed","type":"event"},{"anonymous":false,"inputs":[{"components":[{"internalType":"uint64","name":"first","type":"uint64"},{"internalType":"bytes1","name":"second","type":"bytes1"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.StaticTuple","name":"first","type":"tuple"},{"components":[{"internalType":"uint64","name":"first","type":"uint64"},{"internalType":"bytes1","name":"second","type":"bytes1"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.StaticTuple","name":"second","type":"tuple"}],"name":"StaticTupleEventOne","type":"event"},{"anonymous":false,"inputs":[{"components":[{"internalType":"uint64","name":"first","type":"uint64"},{"internalType":"bytes1","name":"second","type":"bytes1"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.StaticTuple[5]","name":"first","type":"tuple[5]"},{"components":[{"internalType":"uint64","name":"first","type":"uint64"},{"internalType":"bytes1","name":"second","type":"bytes1"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.StaticTuple[]","name":"second","type":"tuple[]"}],"name":"StaticTupleEventTwo","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"first","type":"string"},{"indexed":false,"internalType":"string","name":"second","type":"string"}],"name":"StringFixed","type":"event"},{"anonymous":false,"inputs":[{"components":[{"internalType":"string","name":"first","type":"string"},{"internalType":"bytes32","name":"second","type":"bytes32"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.Custom[5]","name":"first","type":"tuple[5]"},{"components":[{"internalType":"string","name":"first","type":"string"},{"internalType":"bytes32","name":"second","type":"bytes32"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.Custom[]","name":"second","type":"tuple[]"}],"name":"StructArray","type":"event"},{"anonymous":false,"inputs":[{"components":[{"internalType":"string","name":"first","type":"string"},{"internalType":"bytes32","name":"second","type":"bytes32"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.Custom","name":"first","type":"tuple"},{"components":[{"internalType":"string","name":"first","type":"string"},{"internalType":"bytes32","name":"second","type":"bytes32"},{"internalType":"bool","name":"third","type":"bool"}],"indexed":false,"internalType":"struct ABIParsingContract.Custom","name":"second","type":"tuple"}],"name":"TupleDynamic","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"first","type":"uint256"},{"indexed":false,"internalType":"uint16","name":"second","type":"uint16"},{"indexed":false,"internalType":"uint64","name":"third","type":"uint64"},{"indexed":false,"internalType":"uint256","name":"fourth","type":"uint256"}],"name":"UintFixed","type":"event"}]`
	expectedEventResults := `{"AddressFixed":{"first":"0x1932c48b2bf8102ba33b4a6b545c32236e342f34","second":"0x9d13c6d3afe1721beef56b55d303b09e021e27ab"},"ArrayDynamicSize":{"first":[0,100,200,44,144,244,88,188,32,132],"second":[true,false,true,false,true,false,true,false,true,false,true,false,true,false,true,false,true,false,true,false]},"ArrayFixedSize":{"first":[0,100,200,44,144,244,88,188,32,132],"second":[true,false,true,false,true,false]},"BoolFixed":{"first":true,"second":false},"BytesFixed":{"first":["0x01","0x02","0x03","0x04","0x05","0x06","0x07","0x08","0x09","0x0a","0x0b","0x0c","0x0d","0x0e","0x0f","0x10","0x11","0x12","0x13","0x14","0x15","0x16","0x17","0x18","0x19","0x1a","0x1b","0x1c","0x1d","0x1e","0x1f","0x20","0x21","0x22","0x23","0x24","0x25","0x26","0x27","0x28","0x29","0x2a","0x2b","0x2c","0x2d","0x2e","0x2f","0x30","0x31","0x32","0x33","0x34","0x35","0x36","0x37","0x38","0x39","0x3a","0x3b","0x3c","0x3d","0x3e","0x3f","0x40","0x41","0x42","0x43","0x44","0x45","0x46","0x47","0x48","0x49","0x4a","0x4b","0x4c","0x4d","0x4e","0x4f","0x50","0x51","0x52","0x53","0x54","0x55","0x56","0x57","0x58","0x59","0x5a","0x5b","0x5c","0x5d","0x5e","0x5f","0x60","0x61","0x62","0x63","0x00"],"second":["0x01","0x02","0x03","0x04","0x05","0x06","0x07","0x08","0x09","0x0a","0x0b","0x0c","0x0d","0x0e","0x0f","0x10","0x11","0x12","0x13","0x14","0x15","0x16","0x17","0x18","0x19","0x1a","0x1b","0x1c","0x1d","0x1e","0x1f","0x20","0x21","0x22","0x23","0x24","0x25","0x26","0x27","0x28","0x29","0x2a","0x2b","0x2c","0x2d","0x2e","0x2f","0x30","0x31","0x32","0x33","0x34","0x35","0x36","0x37","0x38","0x39","0x3a","0x3b","0x3c","0x3d","0x3e","0x3f","0x40","0x41","0x42","0x43","0x44","0x45","0x46","0x47","0x48","0x49","0x4a","0x4b","0x4c","0x4d","0x4e","0x4f","0x50","0x51","0x52","0x53","0x54","0x55","0x56","0x57","0x58","0x59","0x5a","0x5b","0x5c","0x5d","0x5e","0x5f","0x60","0x61","0x62","0x63","0x00"]},"BytesFixedSize":{"first":"0x1234567890123456789012345678901234567890123456789012345678901234","second":"0x74","third":"0x12"},"IntFixed":{"first":-98765432109876543210,"fourth":12345678901234567890,"second":12345,"third":-43857439857398534},"Mixed":{"fifth":"string fifth","first":{"first":"qwerty","second":"0x1234567890123456789012345678901234567890123456789012345678901234","third":true},"fourth":{"first":"asdfghjkl","second":"0x4013487654674538507684738547847680974039786439857345674358096798","third":false},"second":"0x4386578974647808650460543048357430403897631067453064043476584731","third":3455},"StaticTupleEventOne":{"first":[15423,"0x08",true],"second":[193,"0x26",false]},"StaticTupleEventTwo":{"first":[[0,"0x00",false],[15423,"0x08",true],[15423,"0x08",true],[15423,"0x08",true],[15423,"0x08",true]],"second":[[15423,"0x08",true],[15423,"0x08",true]]},"StringFixed":{"first":"small","second":"some really large string that will go over the thirty-two byte limit for a single variable"},"StructArray":{"first":[{"first":"","second":"0x0000000000000000000000000000000000000000000000000000000000000000","third":false},{"first":"","second":"0x0000000000000000000000000000000000000000000000000000000000000000","third":false},{"first":"","second":"0x0000000000000000000000000000000000000000000000000000000000000000","third":false},{"first":"qwerty","second":"0x1234567890123456789012345678901234567890123456789012345678901234","third":true},{"first":"qwerty","second":"0x1234567890123456789012345678901234567890123456789012345678901234","third":true}],"second":[{"first":"","second":"0x0000000000000000000000000000000000000000000000000000000000000000","third":false},{"first":"qwerty","second":"0x1234567890123456789012345678901234567890123456789012345678901234","third":true}]},"TupleDynamic":{"first":{"first":"first","second":"0x1234567890123456789012345678901234567890123456789012345678901234","third":true},"second":{"first":"second","second":"0x4013487654674538507684738547847680974039786439857345674358096798","third":false}},"UintFixed":{"first":98765432109876543210,"fourth":12345678901234567890,"second":12345,"third":43857439857398534}}`

	var tx Transaction
	json.Unmarshal([]byte(testTx), &tx)

	var structure ABIStructure
	json.Unmarshal([]byte(abiTx), &structure)

	abi := structure.To()

	allParsedResults := make(map[string]interface{})
	for _, c := range tx.Events {
		for _, ev := range abi.Events {
			if ev.Signature() == c.Topics[0].String() {
				allParsedResults[ev.Name] = ev.Parse(c.Data)
			}
		}
	}

	asJson, _ := json.Marshal(allParsedResults)
	assert.JSONEq(t, expectedEventResults, string(asJson))
}
