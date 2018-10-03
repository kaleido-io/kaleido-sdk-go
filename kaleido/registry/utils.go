package registry

import (
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	resty "gopkg.in/resty.v1"
)

type utilsInterface interface {
	getRegistryURL() string
	getAPIClient() *resty.Client

	getDirectoryAddress() string
	getDirectoryClient() *resty.Client

	getProfilesAddress() string
	getProfilesClient() *resty.Client

	getNodeClient() *ethclient.Client
	validateGetResponse(res *resty.Response, err error, resourceName string) error
	validateCreateResponse(res *resty.Response, err error, resourceName string) error

	generateNodeID(path string) string
	generateUserID(path string, email string) string
	newKeyStoreTransactor(from *accounts.Account, keystore *keystore.KeyStore, chainID *big.Int) *bind.TransactOpts
	getAccountForAddress(ks *keystore.KeyStore, hexAddress string) (*accounts.Account, error)
}

var instance *utilsImpl

func utils() utilsInterface {
	if instance == nil {
		instance = &utilsImpl{}
		err := instance.initialize()
		if err != nil {
			fmt.Println(err.Error())
			panic(0) // hacky but will do instead of adding an err check on all utils() call
		}
	}
	return instance
}

type utilsImpl struct {
	registryURL string
	apiClient   *resty.Client

	directoryAddress string
	directoryClient  *resty.Client

	profilesAddress string
	profilesClient  *resty.Client

	nodeClient *ethclient.Client
}

func (u *utilsImpl) initAPIClient() {
	u.registryURL = viper.GetString("api.url") + "/idregistry/" + viper.GetString("registry.id")
	u.apiClient = resty.New().SetHostURL(u.registryURL).SetAuthToken(viper.GetString("api.key"))
	viper.SetDefault("api.debug", false)
	u.apiClient.SetDebug(viper.GetBool("api.debug"))
}

func (u *utilsImpl) initDirectoryClient() error {
	directory, profiles, err := u.fetchOnChainAddresses()
	if err != nil {
		return err
	}
	u.directoryAddress = directory
	u.profilesAddress = profiles

	directoryURL := utils().getRegistryURL() + "/directories/" + u.directoryAddress
	profilesURL := utils().getRegistryURL() + "/properties/" + u.profilesAddress

	u.directoryClient = resty.New().SetHostURL(directoryURL).SetAuthToken(viper.GetString("api.key"))
	u.profilesClient = resty.New().SetHostURL(profilesURL).SetAuthToken(viper.GetString("api.key"))

	return nil
}

func (u *utilsImpl) initEthClient() error {
	urlKey := "networks." + viper.GetString("profile") + ".url"
	if urlKey == "networks..url" {
		return errors.New("No active profile set")
	}
	url := viper.GetString(urlKey)
	if url == "" {
		return errors.New("Unable to retrieve node url from profile. Have you setup your profile properly?")
	}

	var err error
	if u.nodeClient, err = ethclient.Dial(url); err != nil {
		log.Fatalf("Unable to connect to node at %s", url)
		return err
	}

	return nil
}

func (u *utilsImpl) initialize() error {
	u.initAPIClient()

	err := u.initDirectoryClient()
	if err != nil {
		return err
	}

	err = u.initEthClient()
	if err != nil {
		return err
	}

	return nil
}

func (u *utilsImpl) getRegistryURL() string {
	return u.registryURL
}

var client *resty.Client

func (u *utilsImpl) getAPIClient() *resty.Client {
	return u.apiClient
}

func (u *utilsImpl) getDirectoryAddress() string {
	return u.directoryAddress
}

func (u *utilsImpl) getDirectoryClient() *resty.Client {
	return u.directoryClient
}

func (u *utilsImpl) getProfilesAddress() string {
	return u.profilesAddress
}

func (u *utilsImpl) getProfilesClient() *resty.Client {
	return u.profilesClient
}

func (u *utilsImpl) getNodeClient() *ethclient.Client {
	return u.nodeClient
}

func (u *utilsImpl) validateGetResponse(res *resty.Response, err error, resourceName string) error {
	if res.StatusCode() != 200 {
		if err != nil {
			fmt.Printf(err.Error())
		}
		return fmt.Errorf("could not retrieve %s. status code: %d", resourceName, res.StatusCode())
	}
	return nil
}

func (u *utilsImpl) validateCreateResponse(res *resty.Response, err error, resourceName string) error {
	if res.StatusCode() != 201 && res.StatusCode() != 200 {
		return fmt.Errorf("could not create %s. status code: %d", resourceName, res.StatusCode())
	}
	return nil
}

func (u *utilsImpl) fetchOnChainAddresses() (string, string, error) {
	type responseBody struct {
		ID          string `json:"_id,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Directory   string `json:"directory,omitempty"`
		Profiles    string `json:"profiles,omitempty"`
		Claims      string `json:"claims,omitempty"`
	}
	var directories []responseBody
	client := utils().getAPIClient()
	_, err := client.R().SetResult(&directories).Get("/directories")
	if err != nil {
		return "", "", err
	}
	if len(directories) < 1 {
		return "", "", errors.New("Unexpected error: no directories available")
	}

	return directories[0].Directory, directories[0].Profiles, nil
}

func (u *utilsImpl) generateNodeID(path string) string {
	// TODO calculate path hash
	return path
}

func (u *utilsImpl) generateUserID(path string, email string) string {
	// TODO calculate email hass
	return email
}

func (u *utilsImpl) newKeyStoreTransactor(from *accounts.Account, keystore *keystore.KeyStore, chainID *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: from.Address,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return keystore.SignTxWithPassphrase(*from, "test", tx, chainID)
		},
	}
}

func (u *utilsImpl) getAccountForAddress(ks *keystore.KeyStore, hexAddress string) (*accounts.Account, error) {
	signerAccount := common.HexToAddress(hexAddress)
	if ks.HasAddress(signerAccount) {
		for _, account := range ks.Accounts() {
			if account.Address == signerAccount {
				return &account, nil
			}
		}
	}
	return nil, fmt.Errorf("Account for address %s not found", hexAddress)
}
