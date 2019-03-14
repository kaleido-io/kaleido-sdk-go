package registry

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strconv"

	"compress/gzip"

	"encoding/base64"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	profiles "github.com/kaleido-io/kaleido-sdk-go/contracts/properties"
)

// Property key-value details
type Property struct {
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
	Version string `json:"version,omitempty"`
}

// Profile object to interact with profile
type Profile struct {
	KeyStorePath string
	Signer       string
}

// SetProperty sets the key-value for the owner profile
func (p *Profile) SetProperty(key string, value string, revision string) error {
	ks := keystore.NewKeyStore(p.KeyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	if account, err := Utils().getAccountForAddress(ks, p.Signer); err == nil {
		client := Utils().getNodeClient()

		auth := Utils().newKeyStoreTransactor(account, ks, nil) // TODO add chain id
		auth.Value = big.NewInt(0)

		instance, err := profiles.NewProperties(common.HexToAddress(Utils().getProfilesAddress()), client)
		if err != nil {
			return err
		}

		// compress the data
		var buffer bytes.Buffer
		gz := gzip.NewWriter(&buffer)
		if _, err := gz.Write([]byte(value)); err != nil {
			return err
		}
		if err := gz.Flush(); err != nil {
			return err
		}
		if err := gz.Close(); err != nil {
			return err
		}

		var compressedValue = base64.StdEncoding.EncodeToString(buffer.Bytes())

		var tx *types.Transaction
		if revision != "" {
			tx, err = instance.SetWithVersion(auth, key, compressedValue, revision)
		} else {
			tx, err = instance.Set(auth, key, compressedValue)
		}
		if err != nil {
			return err
		}

		fmt.Println("tx sent:", tx.Hash().Hex())
		fmt.Println("waiting for tx to be mined (may take a few seconds)...")
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return err
		}
		fmt.Println("tx receipt, stats", receipt.Status, "gas used = ", receipt.CumulativeGasUsed)
	} else {
		return err
	}
	return nil
}

// GetProperty get a property's latest revision
func (p *Profile) GetProperty(owner string, key string) (*Property, error) {
	client := Utils().getProfilesClient()

	var property Property
	response, err := client.R().SetResult(&property).Get("/profiles/" + owner + "/" + key)
	if err := Utils().ValidateGetResponse(response, err, "profile key"); err != nil {
		return nil, err
	}
	return &property, nil
}

// GetPropertyByRevision as the name says
func (p *Profile) GetPropertyByRevision(owner string, key string, revisionIndex int64) (*Property, error) {
	client := Utils().getProfilesClient()

	var property Property
	response, err := client.R().SetResult(&property).Get("/profiles/" + owner + "/" + key + "/versions/" + strconv.FormatInt(revisionIndex, 10))
	if err := Utils().ValidateGetResponse(response, err, "key"); err != nil {
		return nil, err
	}
	return &property, nil
}

// GetPropertyAllVersions all version for the given property
func (p *Profile) GetPropertyAllVersions(owner string, key string) (*[]Property, error) {
	client := Utils().getProfilesClient()

	type responseBodyType struct {
		Count    string     `json:"count,omitempty"`
		Versions []Property `json:"versions,omitempty"`
	}

	var responseBody responseBodyType
	response, err := client.R().SetResult(&responseBody).Get("/profiles/" + owner + "/" + key + "/versions")
	if err := Utils().ValidateGetResponse(response, err, "key"); err != nil {
		return nil, err
	}
	return &responseBody.Versions, nil

}

// GetProperties get the latest revision of all properties associated with this owner
func (p *Profile) GetProperties(owner string) (*[]Property, error) {
	client := Utils().getProfilesClient()

	type keys struct {
		Count  string     `json:"count,omitempty"`
		Values []Property `json:"values,omitempty"`
	}

	type responseBodyType struct {
		Keys keys `json:"keys,omitempty"`
	}

	var responseBody responseBodyType
	response, err := client.R().SetResult(&responseBody).Get("/profiles/" + owner)
	if err := Utils().ValidateGetResponse(response, err, "profile"); err != nil {
		return nil, err
	}
	return &responseBody.Keys.Values, nil
}
