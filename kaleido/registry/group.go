package registry

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"
	directory "github.com/kaleido-io/kaleido-sdk-go/contracts/directory"
)

// Group represents a group
type Group struct {
	NodeID      string
	Owner       string
	Name        string
	Parent      string
	NumUsers    string
	NumChildren string
	Profile     string
}

// InvokeGet get a user
func (g *Group) InvokeGet() error {
	client := Utils().getNodeClient()
	instance, err := directory.NewDirectory(common.HexToAddress(Utils().getDirectoryAddress()), client)
	if err != nil {
		return err
	}
	nodeID := Utils().GenerateNodeID(g.Name)

	var node [32]byte
	nodeBytes, _ := hexutil.Decode(nodeID)
	copy(node[:], nodeBytes)
	owner, label, parent, _, numUsers, numChildren, profile, err := instance.NodeDetails(&bind.CallOpts{}, node)
	if err != nil {
		return err
	}
	if label == "" {
		fmt.Println("Failed to find group in registry.")
		return nil
	}
	fmt.Println("Successfully found group.")
	fmt.Println("**********************************************************")
	fmt.Printf("owner           = 0x%x\n", owner)
	fmt.Println("label           =", label)
	fmt.Println("node ID         =", nodeID)
	fmt.Printf("parent          = 0x%x\n", parent)
	fmt.Println("numUsers        =", numUsers)
	fmt.Println("numChildren     =", numChildren)
	fmt.Printf("profile address = 0x%x\n", profile)
	return nil
}

// InvokeList get a list of groups
func (g *Group) InvokeList(node [32]byte) (*[]ContractOrganization, error) {
	var groups []ContractOrganization
	client := Utils().getNodeClient()
	instance, err := directory.NewDirectory(common.HexToAddress(Utils().getDirectoryAddress()), client)
	if err != nil {
		return &groups, err
	}

	count, err := instance.NodeChildrenCount(&bind.CallOpts{}, node)
	if err != nil {
		return &groups, err
	}
	countInt := count.Int64()
	var index int64
	fmt.Println("**********************************************************")
	fmt.Println("Number of groups  =", count)
	for index = 0; index < countInt; index++ {
		var org ContractOrganization
		nodeID, _, err := instance.NodeChild(&bind.CallOpts{}, node, uint8(index))
		if err != nil {
			return &groups, err
		}
		owner, label, parent, _, _, _, _, err := instance.NodeDetails(&bind.CallOpts{}, nodeID)
		if err != nil {
			return &groups, err
		}
		org.ID = "0x" + hex.EncodeToString(nodeID[:32])
		org.Name = label
		org.Owner = owner.String()
		org.ParentID = "0x" + hex.EncodeToString(parent[:32])
		groups = append(groups, org)
	}
	return &groups, nil
}

// InvokeCreate create a node
func (g *Group) InvokeCreate(keystorePath string, signer string) error {
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

		parentNodeID := Utils().GenerateNodeID(g.Parent)

		var parent [32]byte
		parentBytes, _ := hexutil.Decode(parentNodeID)
		copy(parent[:], parentBytes)
		fmt.Println("Adding group to following parent node: ", parentNodeID)
		tx, err := instance.SetNodeDetails(auth, parent, g.Name, "", common.HexToAddress(g.Owner))
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
