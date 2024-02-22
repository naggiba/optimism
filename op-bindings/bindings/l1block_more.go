// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1BlockStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"number\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1001,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"timestamp\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1002,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"basefee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"hash\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1004,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"sequenceNumber\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint64\"},{\"astId\":1005,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"blobBaseFeeScalar\",\"offset\":8,\"slot\":\"3\",\"type\":\"t_uint32\"},{\"astId\":1006,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"baseFeeScalar\",\"offset\":12,\"slot\":\"3\",\"type\":\"t_uint32\"},{\"astId\":1007,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeOverhead\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeScalar\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_uint256\"},{\"astId\":1010,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"blobBaseFee\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"interopSetSize\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_uint8\"},{\"astId\":1012,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"chainIds\",\"offset\":0,\"slot\":\"9\",\"type\":\"t_array(t_uint256)dyn_storage\"}],\"types\":{\"t_array(t_uint256)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint256[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint256\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1BlockStorageLayout = new(solc.StorageLayout)

var L1BlockDeployedBin = "0x608060405234801561001057600080fd5b50600436106101515760003560e01c80638381f58a116100cd578063e38bbc3211610081578063e81b2c6d11610066578063e81b2c6d14610323578063ee6ef90c1461032c578063f82061401461034b57600080fd5b8063e38bbc32146102c0578063e591b282146102e357600080fd5b80639e8c4966116100b25780639e8c496614610277578063b80777ea14610280578063c5985918146102a057600080fd5b80638381f58a1461025a5780638b239f731461026e57600080fd5b806357b7bc8b1161012457806364ca23ef1161010957806364ca23ef146101f457806368d5dca614610221578063760ee04d1461025257600080fd5b806357b7bc8b146101d85780635cf24969146101eb57600080fd5b806309bd5a601461015657806321d9309014610172578063440a5e201461018557806354fd4d501461018f575b600080fd5b61015f60025481565b6040519081526020015b60405180910390f35b61015f610180366004610699565b610354565b61018d610375565b005b6101cb6040518060400160405280600581526020017f312e322e3000000000000000000000000000000000000000000000000000000081525081565b60405161016991906106b2565b61018d6101e636600461078e565b6103ca565b61015f60015481565b6003546102089067ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610169565b60035461023d9068010000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610169565b61018d61054a565b6000546102089067ffffffffffffffff1681565b61015f60055481565b61015f60065481565b6000546102089068010000000000000000900467ffffffffffffffff1681565b60035461023d906c01000000000000000000000000900463ffffffff1681565b6102d36102ce366004610699565b6105d0565b6040519015158152602001610169565b6102fe73deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610169565b61015f60045481565b6008546103399060ff1681565b60405160ff9091168152602001610169565b61015f60075481565b6009818154811061036457600080fd5b600091825260209091200154905081565b3373deaddeaddeaddeaddeaddeaddeaddeaddead00011461039e57633cc50b456000526004601cfd5b60043560801c60035560143560801c600055602435600155604435600755606435600255608435600455565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000114610471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f4c31426c6f636b3a206f6e6c7920746865206465706f7369746f72206163636f60448201527f756e742063616e20736574204c3120626c6f636b2076616c7565730000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff8d81167fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921691909117680100000000000000008d8316021790915560018a90556002899055600380547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016918916919091179055600486905560058590556006849055600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff851617905561053c60098383610639565b505050505050505050505050565b3373deaddeaddeaddeaddeaddeaddeaddeaddead00011461057357633cc50b456000526004601cfd5b60a43560f81c60a5602082020136146105945763613457f26000526004601cfd5b806008556009600052602060002060005b828110156105cb5760098054600190810190915560a560208302013583830155016105a5565b505050565b60004682036105e157506001919050565b60005b60085460ff1681101561063057826009828154811061060557610605610852565b90600052602060002001540361061e5750600192915050565b8061062881610881565b9150506105e4565b50600092915050565b828054828255906000526020600020908101928215610674579160200282015b82811115610674578235825591602001919060010190610659565b50610680929150610684565b5090565b5b808211156106805760008155600101610685565b6000602082840312156106ab57600080fd5b5035919050565b600060208083528351808285015260005b818110156106df578581018301518582016040015282016106c3565b818111156106f1576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803567ffffffffffffffff8116811461073d57600080fd5b919050565b60008083601f84011261075457600080fd5b50813567ffffffffffffffff81111561076c57600080fd5b6020830191508360208260051b850101111561078757600080fd5b9250929050565b60008060008060008060008060008060006101408c8e0312156107b057600080fd5b6107b98c610725565b9a506107c760208d01610725565b995060408c0135985060608c013597506107e360808d01610725565b965060a08c0135955060c08c0135945060e08c013593506101008c013560ff8116811461080f57600080fd5b92506101208c013567ffffffffffffffff81111561082c57600080fd5b6108388e828f01610742565b915080935050809150509295989b509295989b9093969950565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036108d9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1BlockStorageLayoutJSON), L1BlockStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1Block"] = L1BlockStorageLayout
	deployedBytecodes["L1Block"] = L1BlockDeployedBin
	immutableReferences["L1Block"] = false
}
