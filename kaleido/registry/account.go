package registry

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/kaleido-io/kaleido-sdk-go/contracts/directory"
)

// Account ...
type Account struct {
	Parent string `json:"parentName,omitempty"`
	Name   string `json:"name,omitempty"`
	Value  string `json:"-"`
}

// InvokeCreate create an account tied to an existing org or group
func (acct *Account) InvokeCreate(keystorePath string, signer string) error {
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	if account, err := Utils().getAccountForAddress(ks, signer); err == nil {
		client := Utils().getNodeClient()

		nonce, err := client.PendingNonceAt(context.Background(), account.Address)
		if err != nil {
			fmt.Printf("Error finding nonce.\n")
			fmt.Printf("Check your .kld.yaml file for correct node endpoint.\n")
			return err
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}

		auth := Utils().newKeyStoreTransactor(account, ks, nil) // TODO add chain id
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = gasPrice
		auth.GasLimit = uint64(5000000)
		auth.Value = big.NewInt(0)

		instance, err := directory.NewDirectory(common.HexToAddress(Utils().getDirectoryAddress()), client)
		if err != nil {
			return err
		}

		parentNodeID := Utils().GenerateNodeID(acct.Parent)

		var parent [32]byte
		parentBytes, _ := hexutil.Decode(parentNodeID)
		copy(parent[:], parentBytes)
		fmt.Println("Adding account to following parent node: ", parentNodeID)
		tx, err := instance.SetAccount(auth, parent, acct.Name, common.HexToAddress(acct.Value))
		if err != nil {
			return err
		}

		fmt.Println("tx sent:", tx.Hash().Hex())
		fmt.Println("waiting for tx to be mined (may take a few seconds)...")

		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}
		fmt.Printf("tx receipt: status=%v, gas used=%v\n", receipt.Status, receipt.CumulativeGasUsed)
	} else {
		return err
	}
	return nil
}

// InvokeGet retrieve an account
func (acct *Account) InvokeGet() error {
	client := Utils().getNodeClient()
	instance, err := directory.NewDirectory(common.HexToAddress(Utils().getDirectoryAddress()), client)
	if err != nil {
		return err
	}
	nodeID := Utils().GenerateNodeID(acct.Parent)
	var node [32]byte
	nodeBytes, _ := hexutil.Decode(nodeID)
	copy(node[:], nodeBytes)

	var parentID [32]byte
	var parentName string
	var name string
	var address common.Address
	var versionDescr string

	if acct.Name[:2] == "0x" {
		acctDecoded, _ := hexutil.Decode(acct.Name)
		var acctBytes [32]byte
		copy(acctBytes[:], acctDecoded)
		parentID, parentName, name, address, versionDescr, err = instance.GetLatestAccount(&bind.CallOpts{}, node, acctBytes)
	} else {
		parentID, parentName, name, address, versionDescr, err = instance.GetLatestAccountByName(&bind.CallOpts{}, node, acct.Name)
	}

	if err != nil {
		fmt.Println("Failed to find account in registry. Smart contract kicked back the function call.")
		return nil
	}
	if parentName == "" {
		fmt.Println("Failed to find account in registry.")
		return nil
	}
	fmt.Println("Successfully found account.")
	fmt.Println("**********************************************************")
	fmt.Printf("parent node     = 0x%x\n", parentID)
	fmt.Println("parent name     =", parentName)
	fmt.Println("account name    =", name)
	fmt.Printf("account address = 0x%x\n", address)
	fmt.Println("version descr.  =", versionDescr)
	return nil
}

// InvokeReverseLookup retrieve an account's details using Ethereum address
func (acct *Account) InvokeReverseLookup() error {
	client := Utils().getNodeClient()
	instance, err := directory.NewDirectory(common.HexToAddress(Utils().getDirectoryAddress()), client)
	if err != nil {
		return err
	}

	var parentID [32]byte
	var parentName string
	var name string
	var address common.Address
	var versionDescr string

	parentID, parentName, name, address, versionDescr, err = instance.AccountLookup(&bind.CallOpts{}, common.HexToAddress(acct.Value))

	if err != nil {
		fmt.Println("Failed to find account in registry. Smart contract kicked back the function call.")
		return nil
	}
	if parentName == "" {
		fmt.Println("Failed to find account in registry.")
		return nil
	}
	fmt.Println("Successfully found account.")
	fmt.Println("**********************************************************")
	fmt.Printf("parent node     = 0x%x\n", parentID)
	fmt.Println("parent name     =", parentName)
	fmt.Println("account name    =", name)
	fmt.Printf("account address = 0x%x\n", address)
	fmt.Println("version descr.  =", versionDescr)
	return nil
}
