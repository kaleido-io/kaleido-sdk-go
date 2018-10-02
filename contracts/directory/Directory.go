// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package directory

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

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\"}],\"name\":\"user\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"userKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"string\"},{\"name\":\"proof\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"profileProvider\",\"type\":\"address\"}],\"name\":\"setNodeDetailsEx\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\"}],\"name\":\"userClaims\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setUserDetails\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"usersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setNodeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"nodeDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"nodeChildrenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"profileProvider\",\"type\":\"address\"},{\"name\":\"claimsProvider\",\"type\":\"address\"}],\"name\":\"setUserDetailsEx\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\"}],\"name\":\"userEmail\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\"}],\"name\":\"userOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"nodeKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"bytes32\"}],\"name\":\"userProfile\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"nodeChild\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"nodeLabel\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"nodeUsersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"nodeUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"userLookup\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"string\"},{\"name\":\"proof\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setNodeDetails\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"nodeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"profileProvider\",\"type\":\"address\"},{\"name\":\"claimsProvider\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"userId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"orgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
const DirectoryBin = `608060405234801561001057600080fd5b506040516040806125a6833981018060405281019080805190602001909291908051906020019092919050505081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600160008060010260001916815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008060010260001916815260200190815260200160002060060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506123fc806101aa6000396000f300608060405260043610610128576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630dfb66151461012d5780631c92e7ac14610228578063259680b3146102745780634aae744e1461033957806361212f29146103aa5780636ba13a82146104415780637161ad191461046c578063766a5a25146104bd578063784042b0146105b757806382c84710146105fc57806386133592146106d35780638b52aa9d1461077d578063a963b057146107ee578063bcb284941461083a578063ca739ed5146108ab578063ce4d6d3c14610971578063e10031b114610a1b578063e6c875c914610a60578063f1a3c5b314610b68578063fcf3fb8514610b93578063ff67083c14610ca0578063ff75542b14610d45575b600080fd5b34801561013957600080fd5b5061015c6004803603810190808035600019169060200190929190505050610db6565b60405180856000191660001916815260200184600019166000191681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101ea5780820151818401526020810190506101cf565b50505050905090810190601f1680156102175780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561023457600080fd5b50610256600480360381019080803560ff169060200190929190505050610ee5565b60405180826000191660001916815260200191505060405180910390f35b34801561028057600080fd5b506103376004803603810190808035600019169060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f21565b005b34801561034557600080fd5b50610368600480360381019080803560001916906020019092919050505061140d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103b657600080fd5b5061043f6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611455565b005b34801561044d57600080fd5b50610456611525565b6040518082815260200191505060405180910390f35b34801561047857600080fd5b506104bb6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611532565b005b3480156104c957600080fd5b506104ec6004803603810190808035600019169060200190929190505050611675565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018560001916600019168152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b8381101561057857808201518184015260208101905061055d565b50505050905090810190601f1680156105a55780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156105c357600080fd5b506105e660048036038101908080356000191690602001909291905050506117ed565b6040518082815260200191505060405180910390f35b34801561060857600080fd5b506106d16004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611818565b005b3480156106df57600080fd5b506107026004803603810190808035600019169060200190929190505050611d83565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610742578082015181840152602081019050610727565b50505050905090810190601f16801561076f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561078957600080fd5b506107ac6004803603810190808035600019169060200190929190505050611e43565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107fa57600080fd5b5061081c600480360381019080803560ff169060200190929190505050611e8b565b60405180826000191660001916815260200191505060405180910390f35b34801561084657600080fd5b506108696004803603810190808035600019169060200190929190505050611ec6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156108b757600080fd5b506108e76004803603810190808035600019169060200190929190803560ff169060200190929190505050611f0e565b60405180836000191660001916815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561093557808201518184015260208101905061091a565b50505050905090810190601f1680156109625780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561097d57600080fd5b506109a06004803603810190808035600019169060200190929190505050612048565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109e05780820151818401526020810190506109c5565b50505050905090810190601f168015610a0d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610a2757600080fd5b50610a4a6004803603810190808035600019169060200190929190505050612108565b6040518082815260200191505060405180910390f35b348015610a6c57600080fd5b50610a9c6004803603810190808035600019169060200190929190803560ff169060200190929190505050612133565b60405180856000191660001916815260200184600019166000191681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b2a578082015181840152602081019050610b0f565b50505050905090810190601f168015610b575780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b348015610b7457600080fd5b50610b7d6121c4565b6040518082815260200191505060405180910390f35b348015610b9f57600080fd5b50610bd4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506121d0565b60405180856000191660001916815260200184600019166000191681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c62578082015181840152602081019050610c47565b50505050905090810190601f168015610c8f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b348015610cac57600080fd5b50610d436004803603810190808035600019169060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612234565b005b348015610d5157600080fd5b50610d7460048036038101908080356000191690602001909291905050506122e3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060006060846003600087600019166000191681526020019081526020016000206000015460036000886000191660001916815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360008960001916600019168152602001908152602001600020600401808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ecf5780601f10610ea457610100808354040283529160200191610ecf565b820191906000526020600020905b815481529060010190602001808311610eb257829003601f168201915b5050505050905093509350935093509193509193565b60006002805490508260ff16101515610efd57600080fd5b60028260ff16815481101515610f0f57fe5b90600052602060002001549050919050565b6000853373ffffffffffffffffffffffffffffffffffffffff1660016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610f9c57600080fd5b86866040516020018082805190602001908083835b602083101515610fd65780518252602082019150602081019050602083039250610fb1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310151561103f578051825260208201915060208101905060208303925061101a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206040516020018083600019166000191681526020018260001916600019168152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831015156110d757805182526020820191506020810190506020830392506110b2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150816000191687600019167f279b80e3bd3cce25eb67f32a0002db63b95e45f5fea0496c93aefae50cef229f8689604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111a557808201518184015260208101905061118a565b50505050905090810190601f1680156111d25780820380516001836020036101000a031916815260200191505b50935050505060405180910390a384600160008460001916600019168152602001908152602001600020600201816000191690555085600160008460001916600019168152602001908152602001600020600101908051906020019061123992919061232b565b508360016000846000191660001916815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660016000846000191660001916815260200190815260200160002060030181600019169055508260016000846000191660001916815260200190815260200160002060060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016000836000191660001916815260200190815260200160002060060160149054906101000a900460ff161515611404576001806000846000191660001916815260200190815260200160002060060160146101000a81548160ff02191690831515021790555060008290806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055506001600088600019166000191681526020019081526020016000206005018290806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055505b50505050505050565b600060036000836000191660001916815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b823373ffffffffffffffffffffffffffffffffffffffff1660016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156114ce57600080fd5b61151f848484600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611818565b50505050565b6000600280549050905090565b813373ffffffffffffffffffffffffffffffffffffffff1660016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156115ab57600080fd5b82600019167fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d26683604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a28160016000856000191660001916815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60006060600080600060016000876000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016000886000191660001916815260200190815260200160002060010160016000896000191660001916815260200190815260200160002060030154600160008a6000191660001916815260200190815260200160002060040180549050600160008b6000191660001916815260200190815260200160002060050180549050838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117d35780601f106117a8576101008083540402835291602001916117d3565b820191906000526020600020905b8154815290600101906020018083116117b657829003601f168201915b505050505093509450945094509450945091939590929450565b6000600160008360001916600019168152602001908152602001600020600501805490509050919050565b6000853373ffffffffffffffffffffffffffffffffffffffff1660016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561189357600080fd5b86866040516020018082805190602001908083835b6020831015156118cd57805182526020820191506020810190506020830392506118a8565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831015156119365780518252602082019150602081019050602083039250611911565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206040516020018083600019166000191681526020018260001916600019168152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831015156119ce57805182526020820191506020810190506020830392506119a9565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150866000191682600019167f38395024c961f08e23fb3be5b46bb73b51b3df560ed6c4a85009815e9d21ded4888860405180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015611a9c578082015181840152602081019050611a81565b50505050905090810190601f168015611ac95780820380516001836020036101000a031916815260200191505b50935050505060405180910390a38460036000846000191660001916815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866003600084600019166000191681526020019081526020016000206000018160001916905550856003600084600019166000191681526020019081526020016000206004019080519060200190611b8d92919061232b565b508360036000846000191660001916815260200190815260200160002060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260036000846000191660001916815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020816000191690555060036000836000191660001916815260200190815260200160002060050160009054906101000a900460ff161515611d7a57600160036000846000191660001916815260200190815260200160002060050160006101000a81548160ff02191690831515021790555060016000886000191660001916815260200190815260200160002060040182908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555060028290806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055505b50505050505050565b60606003600083600019166000191681526020019081526020016000206004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611e375780601f10611e0c57610100808354040283529160200191611e37565b820191906000526020600020905b815481529060010190602001808311611e1a57829003601f168201915b50505050509050919050565b600060036000836000191660001916815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600080805490508260ff16101515611ea257600080fd5b60008260ff16815481101515611eb457fe5b90600052602060002001549050919050565b600060036000836000191660001916815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600060606000600160008660001916600019168152602001908152602001600020600501805490508460ff16101515611f4657600080fd5b6001600086600019166000191681526020019081526020016000206005018460ff16815481101515611f7457fe5b9060005260206000200154905080600160008360001916600019168152602001908152602001600020600101808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156120355780601f1061200a57610100808354040283529160200191612035565b820191906000526020600020905b81548152906001019060200180831161201857829003601f168201915b5050505050905092509250509250929050565b60606001600083600019166000191681526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156120fc5780601f106120d1576101008083540402835291602001916120fc565b820191906000526020600020905b8154815290600101906020018083116120df57829003601f168201915b50505050509050919050565b6000600160008360001916600019168152602001908152602001600020600401805490509050919050565b600080600060606000600160008860001916600019168152602001908152602001600020600401805490508660ff1610151561216e57600080fd5b6001600088600019166000191681526020019081526020016000206004018660ff1681548110151561219c57fe5b906000526020600020015490506121b281610db6565b94509450945094505092959194509250565b60008080549050905090565b600080600060606000600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905061222481610db6565b9450945094509450509193509193565b833373ffffffffffffffffffffffffffffffffffffffff1660016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156122ad57600080fd5b6122dc85858585600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610f21565b5050505050565b600060016000836000191660001916815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061236c57805160ff191683800117855561239a565b8280016001018555821561239a579182015b8281111561239957825182559160200191906001019061237e565b5b5090506123a791906123ab565b5090565b6123cd91905b808211156123c95760008160009055506001016123b1565b5090565b905600a165627a7a72305820c68c4f637e807c324e96d21092bb980e67a6e640d2a65dcf61d2534fab60562b0029`

// DeployDirectory deploys a new Ethereum contract, binding an instance of Directory to it.
func DeployDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, profileProvider common.Address, claimsProvider common.Address) (common.Address, *types.Transaction, *Directory, error) {
	parsed, err := abi.JSON(strings.NewReader(DirectoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DirectoryBin), backend, profileProvider, claimsProvider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Directory{DirectoryCaller: DirectoryCaller{contract: contract}, DirectoryTransactor: DirectoryTransactor{contract: contract}, DirectoryFilterer: DirectoryFilterer{contract: contract}}, nil
}

// Directory is an auto generated Go binding around an Ethereum contract.
type Directory struct {
	DirectoryCaller     // Read-only binding to the contract
	DirectoryTransactor // Write-only binding to the contract
	DirectoryFilterer   // Log filterer for contract events
}

// DirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DirectorySession struct {
	Contract     *Directory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DirectoryCallerSession struct {
	Contract *DirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DirectoryTransactorSession struct {
	Contract     *DirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DirectoryRaw struct {
	Contract *Directory // Generic contract binding to access the raw methods on
}

// DirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DirectoryCallerRaw struct {
	Contract *DirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// DirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DirectoryTransactorRaw struct {
	Contract *DirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDirectory creates a new instance of Directory, bound to a specific deployed contract.
func NewDirectory(address common.Address, backend bind.ContractBackend) (*Directory, error) {
	contract, err := bindDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Directory{DirectoryCaller: DirectoryCaller{contract: contract}, DirectoryTransactor: DirectoryTransactor{contract: contract}, DirectoryFilterer: DirectoryFilterer{contract: contract}}, nil
}

// NewDirectoryCaller creates a new read-only instance of Directory, bound to a specific deployed contract.
func NewDirectoryCaller(address common.Address, caller bind.ContractCaller) (*DirectoryCaller, error) {
	contract, err := bindDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DirectoryCaller{contract: contract}, nil
}

// NewDirectoryTransactor creates a new write-only instance of Directory, bound to a specific deployed contract.
func NewDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DirectoryTransactor, error) {
	contract, err := bindDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DirectoryTransactor{contract: contract}, nil
}

// NewDirectoryFilterer creates a new log filterer instance of Directory, bound to a specific deployed contract.
func NewDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DirectoryFilterer, error) {
	contract, err := bindDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DirectoryFilterer{contract: contract}, nil
}

// bindDirectory binds a generic wrapper to an already deployed contract.
func bindDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DirectoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directory *DirectoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Directory.Contract.DirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directory *DirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.Contract.DirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directory *DirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directory.Contract.DirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directory *DirectoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Directory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directory *DirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directory *DirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directory.Contract.contract.Transact(opts, method, params...)
}

// NodeChild is a free data retrieval call binding the contract method 0xca739ed5.
//
// Solidity: function nodeChild(parentNode bytes32, index uint8) constant returns(bytes32, string)
func (_Directory *DirectoryCaller) NodeChild(opts *bind.CallOpts, parentNode [32]byte, index uint8) ([32]byte, string, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Directory.contract.Call(opts, out, "nodeChild", parentNode, index)
	return *ret0, *ret1, err
}

// NodeChild is a free data retrieval call binding the contract method 0xca739ed5.
//
// Solidity: function nodeChild(parentNode bytes32, index uint8) constant returns(bytes32, string)
func (_Directory *DirectorySession) NodeChild(parentNode [32]byte, index uint8) ([32]byte, string, error) {
	return _Directory.Contract.NodeChild(&_Directory.CallOpts, parentNode, index)
}

// NodeChild is a free data retrieval call binding the contract method 0xca739ed5.
//
// Solidity: function nodeChild(parentNode bytes32, index uint8) constant returns(bytes32, string)
func (_Directory *DirectoryCallerSession) NodeChild(parentNode [32]byte, index uint8) ([32]byte, string, error) {
	return _Directory.Contract.NodeChild(&_Directory.CallOpts, parentNode, index)
}

// NodeChildrenCount is a free data retrieval call binding the contract method 0x784042b0.
//
// Solidity: function nodeChildrenCount(node bytes32) constant returns(uint256)
func (_Directory *DirectoryCaller) NodeChildrenCount(opts *bind.CallOpts, node [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "nodeChildrenCount", node)
	return *ret0, err
}

// NodeChildrenCount is a free data retrieval call binding the contract method 0x784042b0.
//
// Solidity: function nodeChildrenCount(node bytes32) constant returns(uint256)
func (_Directory *DirectorySession) NodeChildrenCount(node [32]byte) (*big.Int, error) {
	return _Directory.Contract.NodeChildrenCount(&_Directory.CallOpts, node)
}

// NodeChildrenCount is a free data retrieval call binding the contract method 0x784042b0.
//
// Solidity: function nodeChildrenCount(node bytes32) constant returns(uint256)
func (_Directory *DirectoryCallerSession) NodeChildrenCount(node [32]byte) (*big.Int, error) {
	return _Directory.Contract.NodeChildrenCount(&_Directory.CallOpts, node)
}

// NodeDetails is a free data retrieval call binding the contract method 0x766a5a25.
//
// Solidity: function nodeDetails(node bytes32) constant returns(address, string, bytes32, uint256, uint256)
func (_Directory *DirectoryCaller) NodeDetails(opts *bind.CallOpts, node [32]byte) (common.Address, string, [32]byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Directory.contract.Call(opts, out, "nodeDetails", node)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// NodeDetails is a free data retrieval call binding the contract method 0x766a5a25.
//
// Solidity: function nodeDetails(node bytes32) constant returns(address, string, bytes32, uint256, uint256)
func (_Directory *DirectorySession) NodeDetails(node [32]byte) (common.Address, string, [32]byte, *big.Int, *big.Int, error) {
	return _Directory.Contract.NodeDetails(&_Directory.CallOpts, node)
}

// NodeDetails is a free data retrieval call binding the contract method 0x766a5a25.
//
// Solidity: function nodeDetails(node bytes32) constant returns(address, string, bytes32, uint256, uint256)
func (_Directory *DirectoryCallerSession) NodeDetails(node [32]byte) (common.Address, string, [32]byte, *big.Int, *big.Int, error) {
	return _Directory.Contract.NodeDetails(&_Directory.CallOpts, node)
}

// NodeKey is a free data retrieval call binding the contract method 0xa963b057.
//
// Solidity: function nodeKey(index uint8) constant returns(bytes32)
func (_Directory *DirectoryCaller) NodeKey(opts *bind.CallOpts, index uint8) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "nodeKey", index)
	return *ret0, err
}

// NodeKey is a free data retrieval call binding the contract method 0xa963b057.
//
// Solidity: function nodeKey(index uint8) constant returns(bytes32)
func (_Directory *DirectorySession) NodeKey(index uint8) ([32]byte, error) {
	return _Directory.Contract.NodeKey(&_Directory.CallOpts, index)
}

// NodeKey is a free data retrieval call binding the contract method 0xa963b057.
//
// Solidity: function nodeKey(index uint8) constant returns(bytes32)
func (_Directory *DirectoryCallerSession) NodeKey(index uint8) ([32]byte, error) {
	return _Directory.Contract.NodeKey(&_Directory.CallOpts, index)
}

// NodeLabel is a free data retrieval call binding the contract method 0xce4d6d3c.
//
// Solidity: function nodeLabel(node bytes32) constant returns(string)
func (_Directory *DirectoryCaller) NodeLabel(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "nodeLabel", node)
	return *ret0, err
}

// NodeLabel is a free data retrieval call binding the contract method 0xce4d6d3c.
//
// Solidity: function nodeLabel(node bytes32) constant returns(string)
func (_Directory *DirectorySession) NodeLabel(node [32]byte) (string, error) {
	return _Directory.Contract.NodeLabel(&_Directory.CallOpts, node)
}

// NodeLabel is a free data retrieval call binding the contract method 0xce4d6d3c.
//
// Solidity: function nodeLabel(node bytes32) constant returns(string)
func (_Directory *DirectoryCallerSession) NodeLabel(node [32]byte) (string, error) {
	return _Directory.Contract.NodeLabel(&_Directory.CallOpts, node)
}

// NodeOwner is a free data retrieval call binding the contract method 0xff75542b.
//
// Solidity: function nodeOwner(node bytes32) constant returns(address)
func (_Directory *DirectoryCaller) NodeOwner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "nodeOwner", node)
	return *ret0, err
}

// NodeOwner is a free data retrieval call binding the contract method 0xff75542b.
//
// Solidity: function nodeOwner(node bytes32) constant returns(address)
func (_Directory *DirectorySession) NodeOwner(node [32]byte) (common.Address, error) {
	return _Directory.Contract.NodeOwner(&_Directory.CallOpts, node)
}

// NodeOwner is a free data retrieval call binding the contract method 0xff75542b.
//
// Solidity: function nodeOwner(node bytes32) constant returns(address)
func (_Directory *DirectoryCallerSession) NodeOwner(node [32]byte) (common.Address, error) {
	return _Directory.Contract.NodeOwner(&_Directory.CallOpts, node)
}

// NodeUser is a free data retrieval call binding the contract method 0xe6c875c9.
//
// Solidity: function nodeUser(node bytes32, index uint8) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectoryCaller) NodeUser(opts *bind.CallOpts, node [32]byte, index uint8) ([32]byte, [32]byte, common.Address, string, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new(common.Address)
		ret3 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Directory.contract.Call(opts, out, "nodeUser", node, index)
	return *ret0, *ret1, *ret2, *ret3, err
}

// NodeUser is a free data retrieval call binding the contract method 0xe6c875c9.
//
// Solidity: function nodeUser(node bytes32, index uint8) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectorySession) NodeUser(node [32]byte, index uint8) ([32]byte, [32]byte, common.Address, string, error) {
	return _Directory.Contract.NodeUser(&_Directory.CallOpts, node, index)
}

// NodeUser is a free data retrieval call binding the contract method 0xe6c875c9.
//
// Solidity: function nodeUser(node bytes32, index uint8) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectoryCallerSession) NodeUser(node [32]byte, index uint8) ([32]byte, [32]byte, common.Address, string, error) {
	return _Directory.Contract.NodeUser(&_Directory.CallOpts, node, index)
}

// NodeUsersCount is a free data retrieval call binding the contract method 0xe10031b1.
//
// Solidity: function nodeUsersCount(node bytes32) constant returns(uint256)
func (_Directory *DirectoryCaller) NodeUsersCount(opts *bind.CallOpts, node [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "nodeUsersCount", node)
	return *ret0, err
}

// NodeUsersCount is a free data retrieval call binding the contract method 0xe10031b1.
//
// Solidity: function nodeUsersCount(node bytes32) constant returns(uint256)
func (_Directory *DirectorySession) NodeUsersCount(node [32]byte) (*big.Int, error) {
	return _Directory.Contract.NodeUsersCount(&_Directory.CallOpts, node)
}

// NodeUsersCount is a free data retrieval call binding the contract method 0xe10031b1.
//
// Solidity: function nodeUsersCount(node bytes32) constant returns(uint256)
func (_Directory *DirectoryCallerSession) NodeUsersCount(node [32]byte) (*big.Int, error) {
	return _Directory.Contract.NodeUsersCount(&_Directory.CallOpts, node)
}

// NodesCount is a free data retrieval call binding the contract method 0xf1a3c5b3.
//
// Solidity: function nodesCount() constant returns(uint256)
func (_Directory *DirectoryCaller) NodesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "nodesCount")
	return *ret0, err
}

// NodesCount is a free data retrieval call binding the contract method 0xf1a3c5b3.
//
// Solidity: function nodesCount() constant returns(uint256)
func (_Directory *DirectorySession) NodesCount() (*big.Int, error) {
	return _Directory.Contract.NodesCount(&_Directory.CallOpts)
}

// NodesCount is a free data retrieval call binding the contract method 0xf1a3c5b3.
//
// Solidity: function nodesCount() constant returns(uint256)
func (_Directory *DirectoryCallerSession) NodesCount() (*big.Int, error) {
	return _Directory.Contract.NodesCount(&_Directory.CallOpts)
}

// User is a free data retrieval call binding the contract method 0x0dfb6615.
//
// Solidity: function user(userId bytes32) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectoryCaller) User(opts *bind.CallOpts, userId [32]byte) ([32]byte, [32]byte, common.Address, string, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new(common.Address)
		ret3 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Directory.contract.Call(opts, out, "user", userId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// User is a free data retrieval call binding the contract method 0x0dfb6615.
//
// Solidity: function user(userId bytes32) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectorySession) User(userId [32]byte) ([32]byte, [32]byte, common.Address, string, error) {
	return _Directory.Contract.User(&_Directory.CallOpts, userId)
}

// User is a free data retrieval call binding the contract method 0x0dfb6615.
//
// Solidity: function user(userId bytes32) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectoryCallerSession) User(userId [32]byte) ([32]byte, [32]byte, common.Address, string, error) {
	return _Directory.Contract.User(&_Directory.CallOpts, userId)
}

// UserClaims is a free data retrieval call binding the contract method 0x4aae744e.
//
// Solidity: function userClaims(userId bytes32) constant returns(address)
func (_Directory *DirectoryCaller) UserClaims(opts *bind.CallOpts, userId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "userClaims", userId)
	return *ret0, err
}

// UserClaims is a free data retrieval call binding the contract method 0x4aae744e.
//
// Solidity: function userClaims(userId bytes32) constant returns(address)
func (_Directory *DirectorySession) UserClaims(userId [32]byte) (common.Address, error) {
	return _Directory.Contract.UserClaims(&_Directory.CallOpts, userId)
}

// UserClaims is a free data retrieval call binding the contract method 0x4aae744e.
//
// Solidity: function userClaims(userId bytes32) constant returns(address)
func (_Directory *DirectoryCallerSession) UserClaims(userId [32]byte) (common.Address, error) {
	return _Directory.Contract.UserClaims(&_Directory.CallOpts, userId)
}

// UserEmail is a free data retrieval call binding the contract method 0x86133592.
//
// Solidity: function userEmail(userId bytes32) constant returns(string)
func (_Directory *DirectoryCaller) UserEmail(opts *bind.CallOpts, userId [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "userEmail", userId)
	return *ret0, err
}

// UserEmail is a free data retrieval call binding the contract method 0x86133592.
//
// Solidity: function userEmail(userId bytes32) constant returns(string)
func (_Directory *DirectorySession) UserEmail(userId [32]byte) (string, error) {
	return _Directory.Contract.UserEmail(&_Directory.CallOpts, userId)
}

// UserEmail is a free data retrieval call binding the contract method 0x86133592.
//
// Solidity: function userEmail(userId bytes32) constant returns(string)
func (_Directory *DirectoryCallerSession) UserEmail(userId [32]byte) (string, error) {
	return _Directory.Contract.UserEmail(&_Directory.CallOpts, userId)
}

// UserKey is a free data retrieval call binding the contract method 0x1c92e7ac.
//
// Solidity: function userKey(index uint8) constant returns(bytes32)
func (_Directory *DirectoryCaller) UserKey(opts *bind.CallOpts, index uint8) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "userKey", index)
	return *ret0, err
}

// UserKey is a free data retrieval call binding the contract method 0x1c92e7ac.
//
// Solidity: function userKey(index uint8) constant returns(bytes32)
func (_Directory *DirectorySession) UserKey(index uint8) ([32]byte, error) {
	return _Directory.Contract.UserKey(&_Directory.CallOpts, index)
}

// UserKey is a free data retrieval call binding the contract method 0x1c92e7ac.
//
// Solidity: function userKey(index uint8) constant returns(bytes32)
func (_Directory *DirectoryCallerSession) UserKey(index uint8) ([32]byte, error) {
	return _Directory.Contract.UserKey(&_Directory.CallOpts, index)
}

// UserLookup is a free data retrieval call binding the contract method 0xfcf3fb85.
//
// Solidity: function userLookup(owner address) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectoryCaller) UserLookup(opts *bind.CallOpts, owner common.Address) ([32]byte, [32]byte, common.Address, string, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new(common.Address)
		ret3 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Directory.contract.Call(opts, out, "userLookup", owner)
	return *ret0, *ret1, *ret2, *ret3, err
}

// UserLookup is a free data retrieval call binding the contract method 0xfcf3fb85.
//
// Solidity: function userLookup(owner address) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectorySession) UserLookup(owner common.Address) ([32]byte, [32]byte, common.Address, string, error) {
	return _Directory.Contract.UserLookup(&_Directory.CallOpts, owner)
}

// UserLookup is a free data retrieval call binding the contract method 0xfcf3fb85.
//
// Solidity: function userLookup(owner address) constant returns(bytes32, bytes32, address, string)
func (_Directory *DirectoryCallerSession) UserLookup(owner common.Address) ([32]byte, [32]byte, common.Address, string, error) {
	return _Directory.Contract.UserLookup(&_Directory.CallOpts, owner)
}

// UserOwner is a free data retrieval call binding the contract method 0x8b52aa9d.
//
// Solidity: function userOwner(userId bytes32) constant returns(address)
func (_Directory *DirectoryCaller) UserOwner(opts *bind.CallOpts, userId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "userOwner", userId)
	return *ret0, err
}

// UserOwner is a free data retrieval call binding the contract method 0x8b52aa9d.
//
// Solidity: function userOwner(userId bytes32) constant returns(address)
func (_Directory *DirectorySession) UserOwner(userId [32]byte) (common.Address, error) {
	return _Directory.Contract.UserOwner(&_Directory.CallOpts, userId)
}

// UserOwner is a free data retrieval call binding the contract method 0x8b52aa9d.
//
// Solidity: function userOwner(userId bytes32) constant returns(address)
func (_Directory *DirectoryCallerSession) UserOwner(userId [32]byte) (common.Address, error) {
	return _Directory.Contract.UserOwner(&_Directory.CallOpts, userId)
}

// UserProfile is a free data retrieval call binding the contract method 0xbcb28494.
//
// Solidity: function userProfile(userId bytes32) constant returns(address)
func (_Directory *DirectoryCaller) UserProfile(opts *bind.CallOpts, userId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "userProfile", userId)
	return *ret0, err
}

// UserProfile is a free data retrieval call binding the contract method 0xbcb28494.
//
// Solidity: function userProfile(userId bytes32) constant returns(address)
func (_Directory *DirectorySession) UserProfile(userId [32]byte) (common.Address, error) {
	return _Directory.Contract.UserProfile(&_Directory.CallOpts, userId)
}

// UserProfile is a free data retrieval call binding the contract method 0xbcb28494.
//
// Solidity: function userProfile(userId bytes32) constant returns(address)
func (_Directory *DirectoryCallerSession) UserProfile(userId [32]byte) (common.Address, error) {
	return _Directory.Contract.UserProfile(&_Directory.CallOpts, userId)
}

// UsersCount is a free data retrieval call binding the contract method 0x6ba13a82.
//
// Solidity: function usersCount() constant returns(uint256)
func (_Directory *DirectoryCaller) UsersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "usersCount")
	return *ret0, err
}

// UsersCount is a free data retrieval call binding the contract method 0x6ba13a82.
//
// Solidity: function usersCount() constant returns(uint256)
func (_Directory *DirectorySession) UsersCount() (*big.Int, error) {
	return _Directory.Contract.UsersCount(&_Directory.CallOpts)
}

// UsersCount is a free data retrieval call binding the contract method 0x6ba13a82.
//
// Solidity: function usersCount() constant returns(uint256)
func (_Directory *DirectoryCallerSession) UsersCount() (*big.Int, error) {
	return _Directory.Contract.UsersCount(&_Directory.CallOpts)
}

// SetNodeDetails is a paid mutator transaction binding the contract method 0xff67083c.
//
// Solidity: function setNodeDetails(parentNode bytes32, label string, proof bytes32, owner address) returns()
func (_Directory *DirectoryTransactor) SetNodeDetails(opts *bind.TransactOpts, parentNode [32]byte, label string, proof [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setNodeDetails", parentNode, label, proof, owner)
}

// SetNodeDetails is a paid mutator transaction binding the contract method 0xff67083c.
//
// Solidity: function setNodeDetails(parentNode bytes32, label string, proof bytes32, owner address) returns()
func (_Directory *DirectorySession) SetNodeDetails(parentNode [32]byte, label string, proof [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetNodeDetails(&_Directory.TransactOpts, parentNode, label, proof, owner)
}

// SetNodeDetails is a paid mutator transaction binding the contract method 0xff67083c.
//
// Solidity: function setNodeDetails(parentNode bytes32, label string, proof bytes32, owner address) returns()
func (_Directory *DirectoryTransactorSession) SetNodeDetails(parentNode [32]byte, label string, proof [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetNodeDetails(&_Directory.TransactOpts, parentNode, label, proof, owner)
}

// SetNodeDetailsEx is a paid mutator transaction binding the contract method 0x259680b3.
//
// Solidity: function setNodeDetailsEx(parentNode bytes32, label string, proof bytes32, owner address, profileProvider address) returns()
func (_Directory *DirectoryTransactor) SetNodeDetailsEx(opts *bind.TransactOpts, parentNode [32]byte, label string, proof [32]byte, owner common.Address, profileProvider common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setNodeDetailsEx", parentNode, label, proof, owner, profileProvider)
}

// SetNodeDetailsEx is a paid mutator transaction binding the contract method 0x259680b3.
//
// Solidity: function setNodeDetailsEx(parentNode bytes32, label string, proof bytes32, owner address, profileProvider address) returns()
func (_Directory *DirectorySession) SetNodeDetailsEx(parentNode [32]byte, label string, proof [32]byte, owner common.Address, profileProvider common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetNodeDetailsEx(&_Directory.TransactOpts, parentNode, label, proof, owner, profileProvider)
}

// SetNodeDetailsEx is a paid mutator transaction binding the contract method 0x259680b3.
//
// Solidity: function setNodeDetailsEx(parentNode bytes32, label string, proof bytes32, owner address, profileProvider address) returns()
func (_Directory *DirectoryTransactorSession) SetNodeDetailsEx(parentNode [32]byte, label string, proof [32]byte, owner common.Address, profileProvider common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetNodeDetailsEx(&_Directory.TransactOpts, parentNode, label, proof, owner, profileProvider)
}

// SetNodeOwner is a paid mutator transaction binding the contract method 0x7161ad19.
//
// Solidity: function setNodeOwner(node bytes32, owner address) returns()
func (_Directory *DirectoryTransactor) SetNodeOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setNodeOwner", node, owner)
}

// SetNodeOwner is a paid mutator transaction binding the contract method 0x7161ad19.
//
// Solidity: function setNodeOwner(node bytes32, owner address) returns()
func (_Directory *DirectorySession) SetNodeOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetNodeOwner(&_Directory.TransactOpts, node, owner)
}

// SetNodeOwner is a paid mutator transaction binding the contract method 0x7161ad19.
//
// Solidity: function setNodeOwner(node bytes32, owner address) returns()
func (_Directory *DirectoryTransactorSession) SetNodeOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetNodeOwner(&_Directory.TransactOpts, node, owner)
}

// SetUserDetails is a paid mutator transaction binding the contract method 0x61212f29.
//
// Solidity: function setUserDetails(node bytes32, email string, owner address) returns()
func (_Directory *DirectoryTransactor) SetUserDetails(opts *bind.TransactOpts, node [32]byte, email string, owner common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setUserDetails", node, email, owner)
}

// SetUserDetails is a paid mutator transaction binding the contract method 0x61212f29.
//
// Solidity: function setUserDetails(node bytes32, email string, owner address) returns()
func (_Directory *DirectorySession) SetUserDetails(node [32]byte, email string, owner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetUserDetails(&_Directory.TransactOpts, node, email, owner)
}

// SetUserDetails is a paid mutator transaction binding the contract method 0x61212f29.
//
// Solidity: function setUserDetails(node bytes32, email string, owner address) returns()
func (_Directory *DirectoryTransactorSession) SetUserDetails(node [32]byte, email string, owner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetUserDetails(&_Directory.TransactOpts, node, email, owner)
}

// SetUserDetailsEx is a paid mutator transaction binding the contract method 0x82c84710.
//
// Solidity: function setUserDetailsEx(node bytes32, email string, owner address, profileProvider address, claimsProvider address) returns()
func (_Directory *DirectoryTransactor) SetUserDetailsEx(opts *bind.TransactOpts, node [32]byte, email string, owner common.Address, profileProvider common.Address, claimsProvider common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setUserDetailsEx", node, email, owner, profileProvider, claimsProvider)
}

// SetUserDetailsEx is a paid mutator transaction binding the contract method 0x82c84710.
//
// Solidity: function setUserDetailsEx(node bytes32, email string, owner address, profileProvider address, claimsProvider address) returns()
func (_Directory *DirectorySession) SetUserDetailsEx(node [32]byte, email string, owner common.Address, profileProvider common.Address, claimsProvider common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetUserDetailsEx(&_Directory.TransactOpts, node, email, owner, profileProvider, claimsProvider)
}

// SetUserDetailsEx is a paid mutator transaction binding the contract method 0x82c84710.
//
// Solidity: function setUserDetailsEx(node bytes32, email string, owner address, profileProvider address, claimsProvider address) returns()
func (_Directory *DirectoryTransactorSession) SetUserDetailsEx(node [32]byte, email string, owner common.Address, profileProvider common.Address, claimsProvider common.Address) (*types.Transaction, error) {
	return _Directory.Contract.SetUserDetailsEx(&_Directory.TransactOpts, node, email, owner, profileProvider, claimsProvider)
}

// DirectoryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Directory contract.
type DirectoryNewOwnerIterator struct {
	Event *DirectoryNewOwner // Event containing the contract specifics and raw log

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
func (it *DirectoryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectoryNewOwner)
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
		it.Event = new(DirectoryNewOwner)
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
func (it *DirectoryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectoryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectoryNewOwner represents a NewOwner event raised by the Directory contract.
type DirectoryNewOwner struct {
	ParentNode [32]byte
	Node       [32]byte
	Owner      common.Address
	Label      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x279b80e3bd3cce25eb67f32a0002db63b95e45f5fea0496c93aefae50cef229f.
//
// Solidity: e NewOwner(parentNode indexed bytes32, node indexed bytes32, owner address, label string)
func (_Directory *DirectoryFilterer) FilterNewOwner(opts *bind.FilterOpts, parentNode [][32]byte, node [][32]byte) (*DirectoryNewOwnerIterator, error) {

	var parentNodeRule []interface{}
	for _, parentNodeItem := range parentNode {
		parentNodeRule = append(parentNodeRule, parentNodeItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Directory.contract.FilterLogs(opts, "NewOwner", parentNodeRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &DirectoryNewOwnerIterator{contract: _Directory.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x279b80e3bd3cce25eb67f32a0002db63b95e45f5fea0496c93aefae50cef229f.
//
// Solidity: e NewOwner(parentNode indexed bytes32, node indexed bytes32, owner address, label string)
func (_Directory *DirectoryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *DirectoryNewOwner, parentNode [][32]byte, node [][32]byte) (event.Subscription, error) {

	var parentNodeRule []interface{}
	for _, parentNodeItem := range parentNode {
		parentNodeRule = append(parentNodeRule, parentNodeItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Directory.contract.WatchLogs(opts, "NewOwner", parentNodeRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectoryNewOwner)
				if err := _Directory.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// DirectoryNewUserIterator is returned from FilterNewUser and is used to iterate over the raw logs and unpacked data for NewUser events raised by the Directory contract.
type DirectoryNewUserIterator struct {
	Event *DirectoryNewUser // Event containing the contract specifics and raw log

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
func (it *DirectoryNewUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectoryNewUser)
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
		it.Event = new(DirectoryNewUser)
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
func (it *DirectoryNewUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectoryNewUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectoryNewUser represents a NewUser event raised by the Directory contract.
type DirectoryNewUser struct {
	UserId [32]byte
	OrgId  [32]byte
	Email  string
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewUser is a free log retrieval operation binding the contract event 0x38395024c961f08e23fb3be5b46bb73b51b3df560ed6c4a85009815e9d21ded4.
//
// Solidity: e NewUser(userId indexed bytes32, orgId indexed bytes32, email string, owner address)
func (_Directory *DirectoryFilterer) FilterNewUser(opts *bind.FilterOpts, userId [][32]byte, orgId [][32]byte) (*DirectoryNewUserIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Directory.contract.FilterLogs(opts, "NewUser", userIdRule, orgIdRule)
	if err != nil {
		return nil, err
	}
	return &DirectoryNewUserIterator{contract: _Directory.contract, event: "NewUser", logs: logs, sub: sub}, nil
}

// WatchNewUser is a free log subscription operation binding the contract event 0x38395024c961f08e23fb3be5b46bb73b51b3df560ed6c4a85009815e9d21ded4.
//
// Solidity: e NewUser(userId indexed bytes32, orgId indexed bytes32, email string, owner address)
func (_Directory *DirectoryFilterer) WatchNewUser(opts *bind.WatchOpts, sink chan<- *DirectoryNewUser, userId [][32]byte, orgId [][32]byte) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Directory.contract.WatchLogs(opts, "NewUser", userIdRule, orgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectoryNewUser)
				if err := _Directory.contract.UnpackLog(event, "NewUser", log); err != nil {
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

// DirectoryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Directory contract.
type DirectoryTransferIterator struct {
	Event *DirectoryTransfer // Event containing the contract specifics and raw log

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
func (it *DirectoryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectoryTransfer)
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
		it.Event = new(DirectoryTransfer)
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
func (it *DirectoryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectoryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectoryTransfer represents a Transfer event raised by the Directory contract.
type DirectoryTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: e Transfer(node indexed bytes32, owner address)
func (_Directory *DirectoryFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*DirectoryTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Directory.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DirectoryTransferIterator{contract: _Directory.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: e Transfer(node indexed bytes32, owner address)
func (_Directory *DirectoryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DirectoryTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Directory.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectoryTransfer)
				if err := _Directory.contract.UnpackLog(event, "Transfer", log); err != nil {
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
