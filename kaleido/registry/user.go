package registry

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"
	directory "github.com/kaleido-io/kaleido-sdk-go/contracts/directory"
)

// User represents a user
type User struct {
	Consortium  string `json:"consortia_id,omitempty"`
	Environment string `json:"environment_id,omitempty"`
	MemberID    string `json:"membership_id,omitempty"`
	Email       string `json:"email,omitempty"`
	Parent      string `json:"parent,omitempty"`
	UserID      string `json:"id,omitempty"`
	Owner       string `json:"owner,omitempty"`
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
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return err
		}
		fmt.Println("tx receipt, stats", receipt.Status, "gas used = ", receipt.CumulativeGasUsed)
	} else {
		return err
	}
	return nil
}
