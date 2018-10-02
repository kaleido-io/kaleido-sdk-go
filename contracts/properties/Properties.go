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
const PropertiesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getByOwnerAndIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"existsByKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getByKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"versionsByOwnerAndName\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"versionsByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"versionsByName\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"existsByName\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getByName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getByOwnerAndName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"existsByOwnerAndKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByOwnerAndName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVersionByKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"version\",\"type\":\"string\"}],\"name\":\"NewKeyValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NewKey\",\"type\":\"event\"}]"

// PropertiesBin is the compiled bytecode used for deploying new contracts.
const PropertiesBin = `608060405234801561001057600080fd5b5061233a806100206000396000f300608060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063095fc1f2146101175780631ae292e0146102b557806325a8e6cb146103f55780632d883a731461052b57806341d983bc146106a95780634a91da90146106f25780634efc8e80146108085780636df60871146109605780638966e047146109fd578063942bad2a14610a62578063b26c502a14610adf578063b336ad8314610b60578063c883a1e014610cae578063ce26640314610e1c578063d493849214610e73578063da465d7414610edc578063e942b51614610fd1578063ea8f451314611080578063ee1ce841146111f8578063f5f6280814611223575b600080fd5b34801561012357600080fd5b50610162600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611343565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156101aa57808201518184015260208101905061018f565b50505050905090810190601f1680156101d75780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156102105780820151818401526020810190506101f5565b50505050905090810190601f16801561023d5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561027657808201518184015260208101905061025b565b50505050905090810190601f1680156102a35780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156102c157600080fd5b5061030e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919080359060200190929190505050611579565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610352578082015181840152602081019050610337565b50505050905090810190601f16801561037f5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156103b857808201518184015260208101905061039d565b50505050905090810190601f1680156103e55780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561040157600080fd5b50610444600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080356000191690602001909291905050506118f2565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561048857808201518184015260208101905061046d565b50505050905090810190601f1680156104b55780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156104ee5780820151818401526020810190506104d3565b50505050905090810190601f16801561051b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561053757600080fd5b5061055660048036038101908080359060200190929190505050611968565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561059e578082015181840152602081019050610583565b50505050905090810190601f1680156105cb5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156106045780820151818401526020810190506105e9565b50505050905090810190601f1680156106315780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561066a57808201518184015260208101905061064f565b50505050905090810190601f1680156106975780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156106b557600080fd5b506106d86004803603810190808035600019169060200190929190505050611984565b604051808215151515815260200191505060405180910390f35b3480156106fe57600080fd5b506107216004803603810190808035600019169060200190929190505050611997565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561076557808201518184015260208101905061074a565b50505050905090810190601f1680156107925780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156107cb5780820151818401526020810190506107b0565b50505050905090810190601f1680156107f85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561081457600080fd5b50610879600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611a0b565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156108bd5780820151818401526020810190506108a2565b50505050905090810190601f1680156108ea5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610923578082015181840152602081019050610908565b50505050905090810190601f1680156109505780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561096c57600080fd5b506109e7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611a2b565b6040518082815260200191505060405180910390f35b348015610a0957600080fd5b50610a4c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190505050611a47565b6040518082815260200191505060405180910390f35b348015610a6e57600080fd5b50610ac9600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611aac565b6040518082815260200191505060405180910390f35b348015610aeb57600080fd5b50610b46600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611abf565b604051808215151515815260200191505060405180910390f35b348015610b6c57600080fd5b50610bc7600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611ad9565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610c0b578082015181840152602081019050610bf0565b50505050905090810190601f168015610c385780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610c71578082015181840152602081019050610c56565b50505050905090810190601f168015610c9e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b348015610cba57600080fd5b50610d35600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611aef565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610d79578082015181840152602081019050610d5e565b50505050905090810190601f168015610da65780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610ddf578082015181840152602081019050610dc4565b50505050905090810190601f168015610e0c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b348015610e2857600080fd5b50610e5d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b0f565b6040518082815260200191505060405180910390f35b348015610e7f57600080fd5b50610ec2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190505050611b5b565b604051808215151515815260200191505060405180910390f35b348015610ee857600080fd5b50610fcf600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611bc2565b005b348015610fdd57600080fd5b5061107e600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611eab565b005b34801561108c57600080fd5b50611111600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611eef565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561115557808201518184015260208101905061113a565b50505050905090810190601f1680156111825780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156111bb5780820151818401526020810190506111a0565b50505050905090810190601f1680156111e85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561120457600080fd5b5061120d611f11565b6040518082815260200191505060405180910390f35b34801561122f57600080fd5b5061125c600480360381019080803560001916906020019092919080359060200190929190505050611f21565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156112a0578082015181840152602081019050611285565b50505050905090810190601f1680156112cd5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156113065780820151818401526020810190506112eb565b50505050905090810190601f1680156113335780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b60608060606000606080600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905087101515611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f696e646578206f7574206f662072616e6765000000000000000000000000000081525060200191505060405180910390fd5b600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208781548110151561145257fe5b9060005260206000200154925061146988846118f2565b80935081925050506000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084600019166000191681526020019081526020016000206001018282828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115625780601f1061153757610100808354040283529160200191611562565b820191906000526020600020905b81548152906001019060200180831161154557829003601f168201915b505050505092509550955095505050509250925092565b60608060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000866000191660001916815260200190815260200160002060000160009054906101000a900460ff16151561165a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4b657920646f6573206e6f742065786973742e0000000000000000000000000081525060200191505060405180910390fd5b600084101580156116c457506000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086600019166000191681526020019081526020016000206003015484105b1515611738576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4b657920696e64657820646f6573206e6f742065786973742e0000000000000081525060200191505060405180910390fd5b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660001916600019168152602001908152602001600020600201600085815260200190815260200160002090508060010181600001818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118425780601f1061181757610100808354040283529160200191611842565b820191906000526020600020905b81548152906001019060200180831161182557829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118de5780601f106118b3576101008083540402835291602001916118de565b820191906000526020600020905b8154815290600101906020018083116118c157829003601f168201915b505050505090509250925050935093915050565b60608061195d848460016000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600088600019166000191681526020019081526020016000206003015403611579565b915091509250929050565b60608060606119773385611343565b9250925092509193909250565b60006119903383611b5b565b9050919050565b606080611a02338460016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600088600019166000191681526020019081526020016000206003015403611579565b91509150915091565b606080611a20611a1a85611f3a565b84611f21565b915091509250929050565b6000611a3f83611a3a84611f3a565b611a47565b905092915050565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836000191660001916815260200190815260200160002060030154905092915050565b6000611ab83383611a2b565b9050919050565b6000611ad2611acd83611f3a565b611984565b9050919050565b606080611ae63384611aef565b91509150915091565b606080611b0484611aff85611f3a565b6118f2565b915091509250929050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b6000806000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084600019166000191681526020019081526020016000206003015411905092915050565b611bca61222b565b600080848360000181905250838360200181905250611be886611f3a565b91506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000836000191660001916815260200190815260200160002060000160009054906101000a900460ff161515611c6457611c63338388612051565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008360001916600019168152602001908152602001600020905082816002016000836003015481526020019081526020016000206000820151816000019080519060200190611cf3929190612245565b506020820151816001019080519060200190611d10929190612245565b50905050806003016000815480929190600101919050555081600019167f90e19dc12d42f4ecaf99b1550dbdb58bf33ac733a22be994ebc288445120504587878760405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015611d99578082015181840152602081019050611d7e565b50505050905090810190601f168015611dc65780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015611dff578082015181840152602081019050611de4565b50505050905090810190601f168015611e2c5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015611e65578082015181840152602081019050611e4a565b50505050905090810190601f168015611e925780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a2505050505050565b611eeb82826040805190810160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250611bc2565b5050565b606080611f0585611eff86611f3a565b85611579565b91509150935093915050565b6000611f1c33611b0f565b905090565b606080611f2f338585611579565b915091509250929050565b60006002826040516020018082805190602001908083835b602083101515611f775780518252602082019150602081019050602083039250611f52565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083101515611fe05780518252602082019150602081019050602083039250611fbb565b6001836020036101000a0380198251168184511680821785525050505050509050019150506020604051808303816000865af1158015612024573d6000803e3d6000fd5b5050506040513d602081101561203957600080fd5b81019080805190602001909291905050509050919050565b6120596122c5565b8181602001819052506001816000019015159081151581525050806000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000856000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190612107929190612245565b5060408201518160030155905050600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555082600019167fe2ec524f3bf279cdf6add99dd96825c9a48b1e7a5c892478646195040156e8bd836040518080602001828103825283818151815260200191508051906020019080838360005b838110156121eb5780820151818401526020810190506121d0565b50505050905090810190601f1680156122185780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b604080519081016040528060608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061228657805160ff19168380011785556122b4565b828001600101855582156122b4579182015b828111156122b3578251825591602001919060010190612298565b5b5090506122c191906122e9565b5090565b60606040519081016040528060001515815260200160608152602001600081525090565b61230b91905b808211156123075760008160009055506001016122ef565b5090565b905600a165627a7a723058202cf9c90f138469bcd9b1d8fa6be2a595a7ffda3274c61b9c65df1a908ed448370029`

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
func (_Properties *PropertiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Properties.Contract.PropertiesCaller.contract.Call(opts, result, method, params...)
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
func (_Properties *PropertiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Properties.Contract.contract.Call(opts, result, method, params...)
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
	out := ret0
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
	out := ret0
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
	out := ret0
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

// GetKeyCount is a free data retrieval call binding the contract method 0xee1ce841.
//
// Solidity: function getKeyCount() constant returns(uint256)
func (_Properties *PropertiesCaller) GetKeyCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
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
	out := ret0
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
	out := ret0
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
	out := ret0
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
