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
	Email  string `json:"email,omitempty"`
	Parent string `json:"parent,omitempty"`
	UserID string `json:"id,omitempty"`
	Owner  string `json:"owner,omitempty"`
}

// InvokeReverseLookup get a username from Ethereum account ID
func (u *User) InvokeReverseLookup(userAcct string) (*User, error) {
	client := utils().getNodeClient()

	var myCallOpts bind.CallOpts
	myCallOpts.Pending = true
	myCallOpts.From = common.HexToAddress(u.Owner)
	myCallOpts.Context = nil

	var user User

	instance, err := directory.NewDirectory(common.HexToAddress(utils().getDirectoryAddress()), client)
	if err != nil {
		return &user, err
	}
	_, _, owner, email, err := instance.UserLookup(&myCallOpts, common.HexToAddress(userAcct))
	//user.UserID = string(userID[:32])
	//user.Parent = string(org[:32])
	user.Email = email
	user.Owner = owner.Hex()
	if err != nil {
		return &user, err
	}

	return &user, nil
}

// InvokeGet get a user
func (u *User) InvokeGet() (*User, error) {
	client := utils().getDirectoryClient()

	url := "/users/" + utils().generateUserID(u.Parent, u.Email)

	var user User
	response, err := client.R().SetResult(&user).Get(url)

	utils().validateGetResponse(response, err, "user")
	return &user, nil
}

// InvokeList get a list of users
func (u *User) InvokeList() (*[]User, error) {
	type userSummary struct {
		UserID string `json:"userId,omitempty"`
		OrgID  string `json:"orgId,omitempty"`
		Owner  string `json:"owner,omitempty"`
		Email  string `json:"email,omitempty"`
	}

	type responseBodyType struct {
		Count int           `json:"count,omitempty"`
		Users []userSummary `json:"users,omitempty"`
	}
	client := utils().getDirectoryClient()

	url := "/orgs/" + utils().generateNodeID(u.Parent) + "/users"

	var responseBody responseBodyType
	response, err := client.R().SetResult(&responseBody).Get(url)
	utils().validateGetResponse(response, err, "users")

	var users []User

	for _, u := range responseBody.Users {
		var user User
		user.Email = u.Email
		user.Parent = u.OrgID
		user.UserID = u.UserID
		user.Owner = u.Owner
		users = append(users, user)
	}

	return &users, nil
}

// InvokeCreate create a user
func (u *User) InvokeCreate(keystorePath string, signer string) error {
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	if account, err := utils().getAccountForAddress(ks, signer); err == nil {
		client := utils().getNodeClient()

		nonce, err := client.PendingNonceAt(context.Background(), account.Address)
		if err != nil {
			return err
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}

		auth := utils().newKeyStoreTransactor(account, ks, nil) // TODO add chain id
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = gasPrice
		auth.GasLimit = uint64(300000)
		auth.Value = big.NewInt(0)

		instance, err := directory.NewDirectory(common.HexToAddress(utils().getDirectoryAddress()), client)
		if err != nil {
			return err
		}

		parentNodeID := utils().generateNodeID(u.Parent)

		var parent [32]byte
		parentBytes, _ := hexutil.Decode(parentNodeID)
		copy(parent[:], parentBytes)
		tx, err := instance.SetUserDetails(auth, parent, u.Email, common.HexToAddress(u.Owner))
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
