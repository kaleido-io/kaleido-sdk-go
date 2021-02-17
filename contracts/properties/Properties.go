// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package properties

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PropertiesABI is the input ABI used to generate the binding from.
const PropertiesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getByOwnerAndIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"}],\"name\":\"setWithVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"existsByKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getByKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"versionsByOwnerAndName\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"versionsByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"versionsByName\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"existsByName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getByName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getByOwnerAndName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"existsByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByOwnerAndName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"version\",\"type\":\"string\"}],\"name\":\"NewKeyValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NewKey\",\"type\":\"event\"}]"

// PropertiesBin is the compiled bytecode used for deploying new contracts.
const PropertiesBin = `60806040523480156200001157600080fd5b506040805190810160405280600581526020017f322e302e30000000000000000000000000000000000000000000000000000000815250600090805190602001906200005f92919062000066565b5062000115565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000a957805160ff1916838001178555620000da565b82800160010185558215620000da579182015b82811115620000d9578251825591602001919060010190620000bc565b5b509050620000e99190620000ed565b5090565b6200011291905b808211156200010e576000816000905550600101620000f4565b5090565b90565b61248080620001256000396000f30060806040526004361061011d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063095fc1f21461012257806319eccd22146102c05780631ae292e0146103b557806325a8e6cb146104f55780632d883a731461062b57806341d983bc146107a95780634a91da90146107f25780634efc8e80146109085780636df6087114610a605780638966e04714610afd5780638aa1043514610b62578063942bad2a14610bf2578063b26c502a14610c6f578063b336ad8314610cf0578063c883a1e014610e3e578063ce26640314610fac578063d493849214611003578063e942b5161461106c578063ea8f45131461111b578063ee1ce84114611293578063f5f62808146112be575b600080fd5b34801561012e57600080fd5b5061016d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506113de565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156101b557808201518184015260208101905061019a565b50505050905090810190601f1680156101e25780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561021b578082015181840152602081019050610200565b50505050905090810190601f1680156102485780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015610281578082015181840152602081019050610266565b50505050905090810190601f1680156102ae5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156102cc57600080fd5b506103b3600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611615565b005b3480156103c157600080fd5b5061040e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919080359060200190929190505050611900565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610452578082015181840152602081019050610437565b50505050905090810190601f16801561047f5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156104b857808201518184015260208101905061049d565b50505050905090810190601f1680156104e55780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561050157600080fd5b50610544600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190505050611c7c565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561058857808201518184015260208101905061056d565b50505050905090810190601f1680156105b55780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156105ee5780820151818401526020810190506105d3565b50505050905090810190601f16801561061b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561063757600080fd5b5061065660048036038101908080359060200190929190505050611cf2565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561069e578082015181840152602081019050610683565b50505050905090810190601f1680156106cb5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156107045780820151818401526020810190506106e9565b50505050905090810190601f1680156107315780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561076a57808201518184015260208101905061074f565b50505050905090810190601f1680156107975780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156107b557600080fd5b506107d86004803603810190808035600019169060200190929190505050611d0e565b604051808215151515815260200191505060405180910390f35b3480156107fe57600080fd5b506108216004803603810190808035600019169060200190929190505050611d21565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561086557808201518184015260208101905061084a565b50505050905090810190601f1680156108925780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156108cb5780820151818401526020810190506108b0565b50505050905090810190601f1680156108f85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561091457600080fd5b50610979600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611d95565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156109bd5780820151818401526020810190506109a2565b50505050905090810190601f1680156109ea5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610a23578082015181840152602081019050610a08565b50505050905090810190601f168015610a505780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b348015610a6c57600080fd5b50610ae7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611db5565b6040518082815260200191505060405180910390f35b348015610b0957600080fd5b50610b4c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190505050611dd1565b6040518082815260200191505060405180910390f35b348015610b6e57600080fd5b50610b77611e37565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610bb7578082015181840152602081019050610b9c565b50505050905090810190601f168015610be45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610bfe57600080fd5b50610c59600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611ed9565b6040518082815260200191505060405180910390f35b348015610c7b57600080fd5b50610cd6600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611eec565b604051808215151515815260200191505060405180910390f35b348015610cfc57600080fd5b50610d57600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611f06565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610d9b578082015181840152602081019050610d80565b50505050905090810190601f168015610dc85780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610e01578082015181840152602081019050610de6565b50505050905090810190601f168015610e2e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b348015610e4a57600080fd5b50610ec5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611f1c565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610f09578082015181840152602081019050610eee565b50505050905090810190601f168015610f365780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610f6f578082015181840152602081019050610f54565b50505050905090810190601f168015610f9c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b348015610fb857600080fd5b50610fed600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f3c565b6040518082815260200191505060405180910390f35b34801561100f57600080fd5b50611052600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190505050611f88565b604051808215151515815260200191505060405180910390f35b34801561107857600080fd5b50611119600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611ff0565b005b34801561112757600080fd5b506111ac600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050612034565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156111f05780820151818401526020810190506111d5565b50505050905090810190601f16801561121d5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561125657808201518184015260208101905061123b565b50505050905090810190601f1680156112835780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561129f57600080fd5b506112a8612056565b6040518082815260200191505060405180910390f35b3480156112ca57600080fd5b506112f7600480360381019080803560001916906020019092919080359060200190929190505050612066565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561133b578082015181840152602081019050611320565b50505050905090810190601f1680156113685780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156113a1578082015181840152602081019050611386565b50505050905090810190601f1680156113ce5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b60608060606000606080600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050871015156114a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f696e646578206f7574206f662072616e6765000000000000000000000000000081525060200191505060405180910390fd5b600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020878154811015156114ed57fe5b906000526020600020015492506115048884611c7c565b8093508192505050600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084600019166000191681526020019081526020016000206001018282828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115fe5780601f106115d3576101008083540402835291602001916115fe565b820191906000526020600020905b8154815290600101906020018083116115e157829003601f168201915b505050505092509550955095505050509250925092565b61161d612371565b60008084836000018190525083836020018190525061163b8661207f565b9150600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836000191660001916815260200190815260200160002060000160009054906101000a900460ff1615156116b8576116b7338388612196565b5b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836000191660001916815260200190815260200160002090508281600201600083600301548152602001908152602001600020600082015181600001908051906020019061174892919061238b565b50602082015181600101908051906020019061176592919061238b565b50905050806003016000815480929190600101919050555081600019167f90e19dc12d42f4ecaf99b1550dbdb58bf33ac733a22be994ebc288445120504587878760405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156117ee5780820151818401526020810190506117d3565b50505050905090810190601f16801561181b5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015611854578082015181840152602081019050611839565b50505050905090810190601f1680156118815780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156118ba57808201518184015260208101905061189f565b50505050905090810190601f1680156118e75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a2505050505050565b6060806000600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000866000191660001916815260200190815260200160002060000160009054906101000a900460ff1615156119e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4b657920646f6573206e6f742065786973742e0000000000000000000000000081525060200191505060405180910390fd5b60008410158015611a4d5750600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086600019166000191681526020019081526020016000206003015484105b1515611ac1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4b657920696e64657820646f6573206e6f742065786973742e0000000000000081525060200191505060405180910390fd5b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660001916600019168152602001908152602001600020600201600085815260200190815260200160002090508060010181600001818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611bcc5780601f10611ba157610100808354040283529160200191611bcc565b820191906000526020600020905b815481529060010190602001808311611baf57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c685780601f10611c3d57610100808354040283529160200191611c68565b820191906000526020600020905b815481529060010190602001808311611c4b57829003601f168201915b505050505090509250925050935093915050565b606080611ce7848460018060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600088600019166000191681526020019081526020016000206003015403611900565b915091509250929050565b6060806060611d0133856113de565b9250925092509193909250565b6000611d1a3383611f88565b9050919050565b606080611d8c338460018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600088600019166000191681526020019081526020016000206003015403611900565b91509150915091565b606080611daa611da48561207f565b84612066565b915091509250929050565b6000611dc983611dc48461207f565b611dd1565b905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836000191660001916815260200190815260200160002060030154905092915050565b606060008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611ecf5780601f10611ea457610100808354040283529160200191611ecf565b820191906000526020600020905b815481529060010190602001808311611eb257829003601f168201915b5050505050905090565b6000611ee53383611db5565b9050919050565b6000611eff611efa8361207f565b611d0e565b9050919050565b606080611f133384611f1c565b91509150915091565b606080611f3184611f2c8561207f565b611c7c565b915091509250929050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b600080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084600019166000191681526020019081526020016000206003015411905092915050565b61203082826040805190810160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250611615565b5050565b60608061204a856120448661207f565b85611900565b91509150935093915050565b600061206133611f3c565b905090565b606080612074338585611900565b915091509250929050565b60006002826040516020018082805190602001908083835b6020831015156120bc5780518252602082019150602081019050602083039250612097565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831015156121255780518252602082019150602081019050602083039250612100565b6001836020036101000a0380198251168184511680821785525050505050509050019150506020604051808303816000865af1158015612169573d6000803e3d6000fd5b5050506040513d602081101561217e57600080fd5b81019080805190602001909291905050509050919050565b61219e61240b565b818160200181905250600181600001901515908115158152505080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000856000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101908051906020019061224d92919061238b565b5060408201518160030155905050600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555082600019167fe2ec524f3bf279cdf6add99dd96825c9a48b1e7a5c892478646195040156e8bd836040518080602001828103825283818151815260200191508051906020019080838360005b83811015612331578082015181840152602081019050612316565b50505050905090810190601f16801561235e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b604080519081016040528060608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106123cc57805160ff19168380011785556123fa565b828001600101855582156123fa579182015b828111156123f95782518255916020019190600101906123de565b5b509050612407919061242f565b5090565b60606040519081016040528060001515815260200160608152602001600081525090565b61245191905b8082111561244d576000816000905550600101612435565b5090565b905600a165627a7a72305820cbefbaded30a7112caba364f8e86dff686e7551e310e2abe5597bd6a2bf665920029`

// DeployProperties deploys a new Ethereum contract, binding an instance of Properties to it.
func DeployProperties(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Properties, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PropertiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Properties{PropertiesCaller: PropertiesCaller{contract: contract}, PropertiesTransactor: PropertiesTransactor{contract: contract}, PropertiesFilterer: PropertiesFilterer{contract: contract}}, nil
}

// Properties is an auto generated Go binding around an Ethereum contract.
type Properties struct {
	PropertiesCaller     // Read-only binding to the contract
	PropertiesTransactor // Write-only binding to the contract
	PropertiesFilterer   // Log filterer for contract events
}

// PropertiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropertiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropertiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PropertiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropertiesSession struct {
	Contract     *Properties       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropertiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropertiesCallerSession struct {
	Contract *PropertiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PropertiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropertiesTransactorSession struct {
	Contract     *PropertiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PropertiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropertiesRaw struct {
	Contract *Properties // Generic contract binding to access the raw methods on
}

// PropertiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropertiesCallerRaw struct {
	Contract *PropertiesCaller // Generic read-only contract binding to access the raw methods on
}

// PropertiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropertiesTransactorRaw struct {
	Contract *PropertiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProperties creates a new instance of Properties, bound to a specific deployed contract.
func NewProperties(address common.Address, backend bind.ContractBackend) (*Properties, error) {
	contract, err := bindProperties(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Properties{PropertiesCaller: PropertiesCaller{contract: contract}, PropertiesTransactor: PropertiesTransactor{contract: contract}, PropertiesFilterer: PropertiesFilterer{contract: contract}}, nil
}

// NewPropertiesCaller creates a new read-only instance of Properties, bound to a specific deployed contract.
func NewPropertiesCaller(address common.Address, caller bind.ContractCaller) (*PropertiesCaller, error) {
	contract, err := bindProperties(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PropertiesCaller{contract: contract}, nil
}

// NewPropertiesTransactor creates a new write-only instance of Properties, bound to a specific deployed contract.
func NewPropertiesTransactor(address common.Address, transactor bind.ContractTransactor) (*PropertiesTransactor, error) {
	contract, err := bindProperties(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PropertiesTransactor{contract: contract}, nil
}

// NewPropertiesFilterer creates a new log filterer instance of Properties, bound to a specific deployed contract.
func NewPropertiesFilterer(address common.Address, filterer bind.ContractFilterer) (*PropertiesFilterer, error) {
	contract, err := bindProperties(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PropertiesFilterer{contract: contract}, nil
}

// bindProperties binds a generic wrapper to an already deployed contract.
func bindProperties(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PropertiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Properties *PropertiesRaw) Call(opts *bind.CallOpts, results *[]interface{}, method string, params ...interface{}) error {
	return _Properties.Contract.PropertiesCaller.contract.Call(opts, results, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Properties *PropertiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Properties.Contract.PropertiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Properties *PropertiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Properties.Contract.PropertiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Properties *PropertiesCallerRaw) Call(opts *bind.CallOpts, results *[]interface{}, method string, params ...interface{}) error {
	return _Properties.Contract.contract.Call(opts, results, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Properties *PropertiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Properties.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Properties *PropertiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Properties.Contract.contract.Transact(opts, method, params...)
}

// ExistsByKey is a free data retrieval call binding the contract method 0x41d983bc.
//
// Solidity: function existsByKey(key bytes32) constant returns(bool)
func (_Properties *PropertiesCaller) ExistsByKey(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "existsByKey", key)
	return *ret0, err
}

// ExistsByKey is a free data retrieval call binding the contract method 0x41d983bc.
//
// Solidity: function existsByKey(key bytes32) constant returns(bool)
func (_Properties *PropertiesSession) ExistsByKey(key [32]byte) (bool, error) {
	return _Properties.Contract.ExistsByKey(&_Properties.CallOpts, key)
}

// ExistsByKey is a free data retrieval call binding the contract method 0x41d983bc.
//
// Solidity: function existsByKey(key bytes32) constant returns(bool)
func (_Properties *PropertiesCallerSession) ExistsByKey(key [32]byte) (bool, error) {
	return _Properties.Contract.ExistsByKey(&_Properties.CallOpts, key)
}

// ExistsByName is a free data retrieval call binding the contract method 0xb26c502a.
//
// Solidity: function existsByName(name string) constant returns(bool)
func (_Properties *PropertiesCaller) ExistsByName(opts *bind.CallOpts, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "existsByName", name)
	return *ret0, err
}

// ExistsByName is a free data retrieval call binding the contract method 0xb26c502a.
//
// Solidity: function existsByName(name string) constant returns(bool)
func (_Properties *PropertiesSession) ExistsByName(name string) (bool, error) {
	return _Properties.Contract.ExistsByName(&_Properties.CallOpts, name)
}

// ExistsByName is a free data retrieval call binding the contract method 0xb26c502a.
//
// Solidity: function existsByName(name string) constant returns(bool)
func (_Properties *PropertiesCallerSession) ExistsByName(name string) (bool, error) {
	return _Properties.Contract.ExistsByName(&_Properties.CallOpts, name)
}

// ExistsByOwnerAndKey is a free data retrieval call binding the contract method 0xd4938492.
//
// Solidity: function existsByOwnerAndKey(owner address, key bytes32) constant returns(bool)
func (_Properties *PropertiesCaller) ExistsByOwnerAndKey(opts *bind.CallOpts, owner common.Address, key [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "existsByOwnerAndKey", owner, key)
	return *ret0, err
}

// ExistsByOwnerAndKey is a free data retrieval call binding the contract method 0xd4938492.
//
// Solidity: function existsByOwnerAndKey(owner address, key bytes32) constant returns(bool)
func (_Properties *PropertiesSession) ExistsByOwnerAndKey(owner common.Address, key [32]byte) (bool, error) {
	return _Properties.Contract.ExistsByOwnerAndKey(&_Properties.CallOpts, owner, key)
}

// ExistsByOwnerAndKey is a free data retrieval call binding the contract method 0xd4938492.
//
// Solidity: function existsByOwnerAndKey(owner address, key bytes32) constant returns(bool)
func (_Properties *PropertiesCallerSession) ExistsByOwnerAndKey(owner common.Address, key [32]byte) (bool, error) {
	return _Properties.Contract.ExistsByOwnerAndKey(&_Properties.CallOpts, owner, key)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(index uint256) constant returns(string, string, string)
func (_Properties *PropertiesCaller) GetByIndex(opts *bind.CallOpts, index *big.Int) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Properties.contract.Call(opts, out, "getByIndex", index)
	return *ret0, *ret1, *ret2, err
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(index uint256) constant returns(string, string, string)
func (_Properties *PropertiesSession) GetByIndex(index *big.Int) (string, string, string, error) {
	return _Properties.Contract.GetByIndex(&_Properties.CallOpts, index)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(index uint256) constant returns(string, string, string)
func (_Properties *PropertiesCallerSession) GetByIndex(index *big.Int) (string, string, string, error) {
	return _Properties.Contract.GetByIndex(&_Properties.CallOpts, index)
}

// GetByKey is a free data retrieval call binding the contract method 0x4a91da90.
//
// Solidity: function getByKey(key bytes32) constant returns(string, string)
func (_Properties *PropertiesCaller) GetByKey(opts *bind.CallOpts, key [32]byte) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getByKey", key)
	return *ret0, *ret1, err
}

// GetByKey is a free data retrieval call binding the contract method 0x4a91da90.
//
// Solidity: function getByKey(key bytes32) constant returns(string, string)
func (_Properties *PropertiesSession) GetByKey(key [32]byte) (string, string, error) {
	return _Properties.Contract.GetByKey(&_Properties.CallOpts, key)
}

// GetByKey is a free data retrieval call binding the contract method 0x4a91da90.
//
// Solidity: function getByKey(key bytes32) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetByKey(key [32]byte) (string, string, error) {
	return _Properties.Contract.GetByKey(&_Properties.CallOpts, key)
}

// GetByName is a free data retrieval call binding the contract method 0xb336ad83.
//
// Solidity: function getByName(name string) constant returns(string, string)
func (_Properties *PropertiesCaller) GetByName(opts *bind.CallOpts, name string) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getByName", name)
	return *ret0, *ret1, err
}

// GetByName is a free data retrieval call binding the contract method 0xb336ad83.
//
// Solidity: function getByName(name string) constant returns(string, string)
func (_Properties *PropertiesSession) GetByName(name string) (string, string, error) {
	return _Properties.Contract.GetByName(&_Properties.CallOpts, name)
}

// GetByName is a free data retrieval call binding the contract method 0xb336ad83.
//
// Solidity: function getByName(name string) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetByName(name string) (string, string, error) {
	return _Properties.Contract.GetByName(&_Properties.CallOpts, name)
}

// GetByOwnerAndIndex is a free data retrieval call binding the contract method 0x095fc1f2.
//
// Solidity: function getByOwnerAndIndex(owner address, index uint256) constant returns(string, string, string)
func (_Properties *PropertiesCaller) GetByOwnerAndIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Properties.contract.Call(opts, out, "getByOwnerAndIndex", owner, index)
	return *ret0, *ret1, *ret2, err
}

// GetByOwnerAndIndex is a free data retrieval call binding the contract method 0x095fc1f2.
//
// Solidity: function getByOwnerAndIndex(owner address, index uint256) constant returns(string, string, string)
func (_Properties *PropertiesSession) GetByOwnerAndIndex(owner common.Address, index *big.Int) (string, string, string, error) {
	return _Properties.Contract.GetByOwnerAndIndex(&_Properties.CallOpts, owner, index)
}

// GetByOwnerAndIndex is a free data retrieval call binding the contract method 0x095fc1f2.
//
// Solidity: function getByOwnerAndIndex(owner address, index uint256) constant returns(string, string, string)
func (_Properties *PropertiesCallerSession) GetByOwnerAndIndex(owner common.Address, index *big.Int) (string, string, string, error) {
	return _Properties.Contract.GetByOwnerAndIndex(&_Properties.CallOpts, owner, index)
}

// GetByOwnerAndKey is a free data retrieval call binding the contract method 0x25a8e6cb.
//
// Solidity: function getByOwnerAndKey(owner address, key bytes32) constant returns(string, string)
func (_Properties *PropertiesCaller) GetByOwnerAndKey(opts *bind.CallOpts, owner common.Address, key [32]byte) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getByOwnerAndKey", owner, key)
	return *ret0, *ret1, err
}

// GetByOwnerAndKey is a free data retrieval call binding the contract method 0x25a8e6cb.
//
// Solidity: function getByOwnerAndKey(owner address, key bytes32) constant returns(string, string)
func (_Properties *PropertiesSession) GetByOwnerAndKey(owner common.Address, key [32]byte) (string, string, error) {
	return _Properties.Contract.GetByOwnerAndKey(&_Properties.CallOpts, owner, key)
}

// GetByOwnerAndKey is a free data retrieval call binding the contract method 0x25a8e6cb.
//
// Solidity: function getByOwnerAndKey(owner address, key bytes32) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetByOwnerAndKey(owner common.Address, key [32]byte) (string, string, error) {
	return _Properties.Contract.GetByOwnerAndKey(&_Properties.CallOpts, owner, key)
}

// GetByOwnerAndName is a free data retrieval call binding the contract method 0xc883a1e0.
//
// Solidity: function getByOwnerAndName(owner address, name string) constant returns(string, string)
func (_Properties *PropertiesCaller) GetByOwnerAndName(opts *bind.CallOpts, owner common.Address, name string) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getByOwnerAndName", owner, name)
	return *ret0, *ret1, err
}

// GetByOwnerAndName is a free data retrieval call binding the contract method 0xc883a1e0.
//
// Solidity: function getByOwnerAndName(owner address, name string) constant returns(string, string)
func (_Properties *PropertiesSession) GetByOwnerAndName(owner common.Address, name string) (string, string, error) {
	return _Properties.Contract.GetByOwnerAndName(&_Properties.CallOpts, owner, name)
}

// GetByOwnerAndName is a free data retrieval call binding the contract method 0xc883a1e0.
//
// Solidity: function getByOwnerAndName(owner address, name string) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetByOwnerAndName(owner common.Address, name string) (string, string, error) {
	return _Properties.Contract.GetByOwnerAndName(&_Properties.CallOpts, owner, name)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(string)
func (_Properties *PropertiesCaller) GetContractVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "getContractVersion")
	return *ret0, err
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(string)
func (_Properties *PropertiesSession) GetContractVersion() (string, error) {
	return _Properties.Contract.GetContractVersion(&_Properties.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() constant returns(string)
func (_Properties *PropertiesCallerSession) GetContractVersion() (string, error) {
	return _Properties.Contract.GetContractVersion(&_Properties.CallOpts)
}

// GetKeyCount is a free data retrieval call binding the contract method 0xee1ce841.
//
// Solidity: function getKeyCount() constant returns(uint256)
func (_Properties *PropertiesCaller) GetKeyCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "getKeyCount")
	return *ret0, err
}

// GetKeyCount is a free data retrieval call binding the contract method 0xee1ce841.
//
// Solidity: function getKeyCount() constant returns(uint256)
func (_Properties *PropertiesSession) GetKeyCount() (*big.Int, error) {
	return _Properties.Contract.GetKeyCount(&_Properties.CallOpts)
}

// GetKeyCount is a free data retrieval call binding the contract method 0xee1ce841.
//
// Solidity: function getKeyCount() constant returns(uint256)
func (_Properties *PropertiesCallerSession) GetKeyCount() (*big.Int, error) {
	return _Properties.Contract.GetKeyCount(&_Properties.CallOpts)
}

// GetVersionByKey is a free data retrieval call binding the contract method 0xf5f62808.
//
// Solidity: function getVersionByKey(key bytes32, index uint256) constant returns(string, string)
func (_Properties *PropertiesCaller) GetVersionByKey(opts *bind.CallOpts, key [32]byte, index *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getVersionByKey", key, index)
	return *ret0, *ret1, err
}

// GetVersionByKey is a free data retrieval call binding the contract method 0xf5f62808.
//
// Solidity: function getVersionByKey(key bytes32, index uint256) constant returns(string, string)
func (_Properties *PropertiesSession) GetVersionByKey(key [32]byte, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByKey(&_Properties.CallOpts, key, index)
}

// GetVersionByKey is a free data retrieval call binding the contract method 0xf5f62808.
//
// Solidity: function getVersionByKey(key bytes32, index uint256) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetVersionByKey(key [32]byte, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByKey(&_Properties.CallOpts, key, index)
}

// GetVersionByName is a free data retrieval call binding the contract method 0x4efc8e80.
//
// Solidity: function getVersionByName(name string, index uint256) constant returns(string, string)
func (_Properties *PropertiesCaller) GetVersionByName(opts *bind.CallOpts, name string, index *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getVersionByName", name, index)
	return *ret0, *ret1, err
}

// GetVersionByName is a free data retrieval call binding the contract method 0x4efc8e80.
//
// Solidity: function getVersionByName(name string, index uint256) constant returns(string, string)
func (_Properties *PropertiesSession) GetVersionByName(name string, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByName(&_Properties.CallOpts, name, index)
}

// GetVersionByName is a free data retrieval call binding the contract method 0x4efc8e80.
//
// Solidity: function getVersionByName(name string, index uint256) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetVersionByName(name string, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByName(&_Properties.CallOpts, name, index)
}

// GetVersionByOwnerAndKey is a free data retrieval call binding the contract method 0x1ae292e0.
//
// Solidity: function getVersionByOwnerAndKey(owner address, key bytes32, index uint256) constant returns(string, string)
func (_Properties *PropertiesCaller) GetVersionByOwnerAndKey(opts *bind.CallOpts, owner common.Address, key [32]byte, index *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getVersionByOwnerAndKey", owner, key, index)
	return *ret0, *ret1, err
}

// GetVersionByOwnerAndKey is a free data retrieval call binding the contract method 0x1ae292e0.
//
// Solidity: function getVersionByOwnerAndKey(owner address, key bytes32, index uint256) constant returns(string, string)
func (_Properties *PropertiesSession) GetVersionByOwnerAndKey(owner common.Address, key [32]byte, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByOwnerAndKey(&_Properties.CallOpts, owner, key, index)
}

// GetVersionByOwnerAndKey is a free data retrieval call binding the contract method 0x1ae292e0.
//
// Solidity: function getVersionByOwnerAndKey(owner address, key bytes32, index uint256) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetVersionByOwnerAndKey(owner common.Address, key [32]byte, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByOwnerAndKey(&_Properties.CallOpts, owner, key, index)
}

// GetVersionByOwnerAndName is a free data retrieval call binding the contract method 0xea8f4513.
//
// Solidity: function getVersionByOwnerAndName(owner address, name string, index uint256) constant returns(string, string)
func (_Properties *PropertiesCaller) GetVersionByOwnerAndName(opts *bind.CallOpts, owner common.Address, name string, index *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Properties.contract.Call(opts, out, "getVersionByOwnerAndName", owner, name, index)
	return *ret0, *ret1, err
}

// GetVersionByOwnerAndName is a free data retrieval call binding the contract method 0xea8f4513.
//
// Solidity: function getVersionByOwnerAndName(owner address, name string, index uint256) constant returns(string, string)
func (_Properties *PropertiesSession) GetVersionByOwnerAndName(owner common.Address, name string, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByOwnerAndName(&_Properties.CallOpts, owner, name, index)
}

// GetVersionByOwnerAndName is a free data retrieval call binding the contract method 0xea8f4513.
//
// Solidity: function getVersionByOwnerAndName(owner address, name string, index uint256) constant returns(string, string)
func (_Properties *PropertiesCallerSession) GetVersionByOwnerAndName(owner common.Address, name string, index *big.Int) (string, string, error) {
	return _Properties.Contract.GetVersionByOwnerAndName(&_Properties.CallOpts, owner, name, index)
}

// VersionsByName is a free data retrieval call binding the contract method 0x942bad2a.
//
// Solidity: function versionsByName(name string) constant returns(uint256)
func (_Properties *PropertiesCaller) VersionsByName(opts *bind.CallOpts, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "versionsByName", name)
	return *ret0, err
}

// VersionsByName is a free data retrieval call binding the contract method 0x942bad2a.
//
// Solidity: function versionsByName(name string) constant returns(uint256)
func (_Properties *PropertiesSession) VersionsByName(name string) (*big.Int, error) {
	return _Properties.Contract.VersionsByName(&_Properties.CallOpts, name)
}

// VersionsByName is a free data retrieval call binding the contract method 0x942bad2a.
//
// Solidity: function versionsByName(name string) constant returns(uint256)
func (_Properties *PropertiesCallerSession) VersionsByName(name string) (*big.Int, error) {
	return _Properties.Contract.VersionsByName(&_Properties.CallOpts, name)
}

// VersionsByOwnerAndKey is a free data retrieval call binding the contract method 0x8966e047.
//
// Solidity: function versionsByOwnerAndKey(owner address, key bytes32) constant returns(uint256)
func (_Properties *PropertiesCaller) VersionsByOwnerAndKey(opts *bind.CallOpts, owner common.Address, key [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "versionsByOwnerAndKey", owner, key)
	return *ret0, err
}

// VersionsByOwnerAndKey is a free data retrieval call binding the contract method 0x8966e047.
//
// Solidity: function versionsByOwnerAndKey(owner address, key bytes32) constant returns(uint256)
func (_Properties *PropertiesSession) VersionsByOwnerAndKey(owner common.Address, key [32]byte) (*big.Int, error) {
	return _Properties.Contract.VersionsByOwnerAndKey(&_Properties.CallOpts, owner, key)
}

// VersionsByOwnerAndKey is a free data retrieval call binding the contract method 0x8966e047.
//
// Solidity: function versionsByOwnerAndKey(owner address, key bytes32) constant returns(uint256)
func (_Properties *PropertiesCallerSession) VersionsByOwnerAndKey(owner common.Address, key [32]byte) (*big.Int, error) {
	return _Properties.Contract.VersionsByOwnerAndKey(&_Properties.CallOpts, owner, key)
}

// VersionsByOwnerAndName is a free data retrieval call binding the contract method 0x6df60871.
//
// Solidity: function versionsByOwnerAndName(owner address, name string) constant returns(uint256)
func (_Properties *PropertiesCaller) VersionsByOwnerAndName(opts *bind.CallOpts, owner common.Address, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
	}
	err := _Properties.contract.Call(opts, out, "versionsByOwnerAndName", owner, name)
	return *ret0, err
}

// VersionsByOwnerAndName is a free data retrieval call binding the contract method 0x6df60871.
//
// Solidity: function versionsByOwnerAndName(owner address, name string) constant returns(uint256)
func (_Properties *PropertiesSession) VersionsByOwnerAndName(owner common.Address, name string) (*big.Int, error) {
	return _Properties.Contract.VersionsByOwnerAndName(&_Properties.CallOpts, owner, name)
}

// VersionsByOwnerAndName is a free data retrieval call binding the contract method 0x6df60871.
//
// Solidity: function versionsByOwnerAndName(owner address, name string) constant returns(uint256)
func (_Properties *PropertiesCallerSession) VersionsByOwnerAndName(owner common.Address, name string) (*big.Int, error) {
	return _Properties.Contract.VersionsByOwnerAndName(&_Properties.CallOpts, owner, name)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(name string, _value string) returns()
func (_Properties *PropertiesTransactor) Set(opts *bind.TransactOpts, name string, _value string) (*types.Transaction, error) {
	return _Properties.contract.Transact(opts, "set", name, _value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(name string, _value string) returns()
func (_Properties *PropertiesSession) Set(name string, _value string) (*types.Transaction, error) {
	return _Properties.Contract.Set(&_Properties.TransactOpts, name, _value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(name string, _value string) returns()
func (_Properties *PropertiesTransactorSession) Set(name string, _value string) (*types.Transaction, error) {
	return _Properties.Contract.Set(&_Properties.TransactOpts, name, _value)
}

// SetWithVersion is a paid mutator transaction binding the contract method 0x19eccd22.
//
// Solidity: function setWithVersion(name string, value string, version string) returns()
func (_Properties *PropertiesTransactor) SetWithVersion(opts *bind.TransactOpts, name string, value string, version string) (*types.Transaction, error) {
	return _Properties.contract.Transact(opts, "setWithVersion", name, value, version)
}

// SetWithVersion is a paid mutator transaction binding the contract method 0x19eccd22.
//
// Solidity: function setWithVersion(name string, value string, version string) returns()
func (_Properties *PropertiesSession) SetWithVersion(name string, value string, version string) (*types.Transaction, error) {
	return _Properties.Contract.SetWithVersion(&_Properties.TransactOpts, name, value, version)
}

// SetWithVersion is a paid mutator transaction binding the contract method 0x19eccd22.
//
// Solidity: function setWithVersion(name string, value string, version string) returns()
func (_Properties *PropertiesTransactorSession) SetWithVersion(name string, value string, version string) (*types.Transaction, error) {
	return _Properties.Contract.SetWithVersion(&_Properties.TransactOpts, name, value, version)
}

// PropertiesNewKeyIterator is returned from FilterNewKey and is used to iterate over the raw logs and unpacked data for NewKey events raised by the Properties contract.
type PropertiesNewKeyIterator struct {
	Event *PropertiesNewKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PropertiesNewKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertiesNewKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PropertiesNewKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PropertiesNewKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertiesNewKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertiesNewKey represents a NewKey event raised by the Properties contract.
type PropertiesNewKey struct {
	Key  [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewKey is a free log retrieval operation binding the contract event 0xe2ec524f3bf279cdf6add99dd96825c9a48b1e7a5c892478646195040156e8bd.
//
// Solidity: e NewKey(key indexed bytes32, name string)
func (_Properties *PropertiesFilterer) FilterNewKey(opts *bind.FilterOpts, key [][32]byte) (*PropertiesNewKeyIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Properties.contract.FilterLogs(opts, "NewKey", keyRule)
	if err != nil {
		return nil, err
	}
	return &PropertiesNewKeyIterator{contract: _Properties.contract, event: "NewKey", logs: logs, sub: sub}, nil
}

// WatchNewKey is a free log subscription operation binding the contract event 0xe2ec524f3bf279cdf6add99dd96825c9a48b1e7a5c892478646195040156e8bd.
//
// Solidity: e NewKey(key indexed bytes32, name string)
func (_Properties *PropertiesFilterer) WatchNewKey(opts *bind.WatchOpts, sink chan<- *PropertiesNewKey, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Properties.contract.WatchLogs(opts, "NewKey", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertiesNewKey)
				if err := _Properties.contract.UnpackLog(event, "NewKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PropertiesNewKeyValueIterator is returned from FilterNewKeyValue and is used to iterate over the raw logs and unpacked data for NewKeyValue events raised by the Properties contract.
type PropertiesNewKeyValueIterator struct {
	Event *PropertiesNewKeyValue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PropertiesNewKeyValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertiesNewKeyValue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PropertiesNewKeyValue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PropertiesNewKeyValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertiesNewKeyValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertiesNewKeyValue represents a NewKeyValue event raised by the Properties contract.
type PropertiesNewKeyValue struct {
	Key     [32]byte
	Name    string
	Value   string
	Version string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewKeyValue is a free log retrieval operation binding the contract event 0x90e19dc12d42f4ecaf99b1550dbdb58bf33ac733a22be994ebc2884451205045.
//
// Solidity: e NewKeyValue(key indexed bytes32, name string, value string, version string)
func (_Properties *PropertiesFilterer) FilterNewKeyValue(opts *bind.FilterOpts, key [][32]byte) (*PropertiesNewKeyValueIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Properties.contract.FilterLogs(opts, "NewKeyValue", keyRule)
	if err != nil {
		return nil, err
	}
	return &PropertiesNewKeyValueIterator{contract: _Properties.contract, event: "NewKeyValue", logs: logs, sub: sub}, nil
}

// WatchNewKeyValue is a free log subscription operation binding the contract event 0x90e19dc12d42f4ecaf99b1550dbdb58bf33ac733a22be994ebc2884451205045.
//
// Solidity: e NewKeyValue(key indexed bytes32, name string, value string, version string)
func (_Properties *PropertiesFilterer) WatchNewKeyValue(opts *bind.WatchOpts, sink chan<- *PropertiesNewKeyValue, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Properties.contract.WatchLogs(opts, "NewKeyValue", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertiesNewKeyValue)
				if err := _Properties.contract.UnpackLog(event, "NewKeyValue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
