package registry

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"
	directory "github.com/kaleido-io/kaleido-sdk-go/contracts/directory"
)

// User represents a user
type User struct {
	Name     string `json:"name,omitempty"`
	ParentID string `json:"orgId,omitempty"`
	UserID   string `json:"userId,omitempty"`
	Owner    string `json:"owner,omitempty"`
}

//// InvokeReverseLookup get a username from Ethereum account ID
//func (u *User) InvokeReverseLookup(userAcct string) (*User, error) {
//client := Utils().getNodeClient()
//
//var myCallOpts bind.CallOpts
//myCallOpts.Pending = true
//myCallOpts.From = common.HexToAddress(u.Owner)
//myCallOpts.Context = nil
//
//var user User
//
//instance, err := directory.NewDirectory(common.HexToAddress(Utils().getDirectoryAddress()), client)
//if err != nil {
//return &user, err
//}
//_, _, owner, email, err := instance.UserLookup(&myCallOpts, common.HexToAddress(userAcct))
////user.UserID = string(userID[:32])
////user.Parent = string(org[:32])
//user.Email = email
//user.Owner = owner.Hex()
//if err != nil {
//return &user, err
//}
//
//return &user, nil
//}

// InvokeGet get a user
func (u *User) InvokeGet() (*User, error) {
	client := Utils().getDirectoryClient()

	url := "/users/" + Utils().generateUserID(u.ParentID, u.Name)

	var user User
	response, err := client.R().SetResult(&user).Get(url)

	Utils().ValidateGetResponse(response, err, "user")
	return &user, nil
}

// InvokeList get a list of users
func (u *User) InvokeList() (*[]User, error) {
	type userSummary struct {
		UserID   string `json:"userId,omitempty"`
		ParentID string `json:"orgId,omitempty"`
		Owner    string `json:"owner,omitempty"`
		Name     string `json:"name,omitempty"`
	}

	type responseBodyType struct {
		Count int           `json:"count,omitempty"`
		Users []userSummary `json:"users,omitempty"`
	}
	client := Utils().getDirectoryClient()

	url := "/orgs/" + Utils().GenerateNodeID(u.ParentID) + "/users"

	var responseBody responseBodyType
	response, err := client.R().SetResult(&responseBody).Get(url)
	Utils().ValidateGetResponse(response, err, "users")

	var users []User

	for _, u := range responseBody.Users {
		var user User
		user.Name = u.Name
		user.ParentID = u.ParentID
		user.UserID = u.UserID
		user.Owner = u.Owner
		users = append(users, user)
	}

	return &users, nil
}

// InvokeCreate create a user
func (u *User) InvokeCreate(keystorePath string, signer string) error {
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

		parentNodeID := Utils().GenerateNodeID(u.ParentID)

		var parent [32]byte
		parentBytes, _ := hexutil.Decode(parentNodeID)
		copy(parent[:], parentBytes)
		tx, err := instance.SetUserDetails(auth, parent, u.Name, common.HexToAddress(u.Owner))
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
