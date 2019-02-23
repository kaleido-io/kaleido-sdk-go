package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	kaleido "github.com/kaleido-io/kaleido-sdk-go/common"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	resty "gopkg.in/resty.v1"
)

type utilsInterface interface {
	getServiceDefinition() (*serviceDefinitionType, error)

	getRegistryURL() string
	getAPIClient() *resty.Client
	getNetworkManagerClient() *resty.Client

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
	readPassword(envVarName string, prompt string) (string, error)
}

var instance *utilsImpl

func utils() utilsInterface {
	if instance == nil {
		instance = &utilsImpl{}
		err := instance.initialize()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1) // hacky but will do instead of adding an err check on all utils() call
		}
	}
	return instance
}

type utilsImpl struct {
	registryURL          string
	apiClient            *resty.Client
	networkManagerClient *resty.Client

	directoryAddress string
	directoryClient  *resty.Client

	profilesAddress string
	profilesClient  *resty.Client

	nodeClient *ethclient.Client

	serviceID string
}

func (u *utilsImpl) newClient(url, authToken string) *resty.Client {
	client := resty.New().SetHostURL(url).SetAuthToken(authToken)
	client.SetDebug(viper.GetBool("api.debug"))
	return client
}

func (u *utilsImpl) initAPIClient() error {
	u.registryURL = viper.GetString("api.url") + "/idregistry/" + u.serviceID
	u.apiClient = u.newClient(u.registryURL, viper.GetString("api.key"))
	return nil
}

func (u *utilsImpl) initNetworkManagerClient() error {
	u.networkManagerClient = u.newClient(viper.GetString("api.url"), viper.GetString("api.key"))
	return nil
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

	u.directoryClient = u.newClient(directoryURL, viper.GetString("api.key"))
	u.profilesClient = u.newClient(profilesURL, viper.GetString("api.key"))

	return nil
}

func (u *utilsImpl) initEthClient() error {
	url := viper.GetString("networks.url")
	if url == "" {
		return errors.New("Unable to load node url. Have you setup your configuration (~/.kld.yaml) properly?")
	}

	var err error
	if u.nodeClient, err = ethclient.Dial(url); err != nil {
		log.Fatalf("Unable to connect to node at %s", url)
		return err
	}

	return nil
}

func (u *utilsImpl) initialize() error {
	u.serviceID = viper.GetString("service.id")

	if err := u.initAPIClient(); err != nil {
		return err
	}

	if err := u.initNetworkManagerClient(); err != nil {
		return err
	}

	if err := u.initDirectoryClient(); err != nil {
		return err
	}

	if err := u.initEthClient(); err != nil {
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

func (u *utilsImpl) getNetworkManagerClient() *resty.Client {
	return u.networkManagerClient
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
		if res.StatusCode() >= 400 && res.StatusCode() < 500 {
			type photicError struct {
				ErrorMessage string `json:"errorMessage,omitempty"`
			}

			var e photicError
			unmarshalError := json.Unmarshal(res.Body(), &e)
			if unmarshalError != nil {
				fmt.Println(unmarshalError)
			} else {
				return errors.New(e.ErrorMessage)
			}
		}
		return fmt.Errorf("could not retrieve %s. status code: %d", resourceName, res.StatusCode())
	}
	return err
}

func (u *utilsImpl) validateCreateResponse(res *resty.Response, err error, resourceName string) error {
	if res.StatusCode() != 201 && res.StatusCode() != 200 {
		if res.StatusCode() >= 400 && res.StatusCode() < 500 {
			type photicError struct {
				ErrorMessage string `json:"errorMessage,omitempty"`
			}

			var e photicError
			unmarshalError := json.Unmarshal(res.Body(), &e)
			if unmarshalError != nil {
				fmt.Println(unmarshalError)
			} else {
				return errors.New(e.ErrorMessage)
			}
		}
		return fmt.Errorf("could not create %s. status code: %d", resourceName, res.StatusCode())
	}

	return err
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
	response, err := client.R().SetResult(&directories).Get("/directories")
	err = u.validateGetResponse(response, err, "directories")
	if err != nil {
		return "", "", err
	}

	if len(directories) < 1 {
		return "", "", errors.New("Unexpected error: no directories available")
	}

	return directories[0].Directory, directories[0].Profiles, nil
}

type serviceDefinitionType struct {
	Consortium  string `json:"consortia_id,omitempty"`
	Environment string `json:"environment_id,omitempty"`
	MemberID    string `json:"membership_id,omitempty"`
}

func (u *utilsImpl) getServiceDefinition() (*serviceDefinitionType, error) {
	client := u.newClient(viper.GetString("api.url"), viper.GetString("api.key"))
	url := fmt.Sprintf("/services?_id=%s", u.serviceID)
	var services []serviceDefinitionType
	response, err := client.R().SetResult(&services).Get(url)
	if err = u.validateGetResponse(response, err, "services"); err != nil {
		return nil, err
	}
	return &services[0], nil
}

func (u *utilsImpl) generateNodeID(path string) string {
	nodeID := path
	if path[:2] != "0x" {
		nodeID, _ = kaleido.PathHash(path)
	}
	return nodeID
}

func (u *utilsImpl) generateUserID(path string, email string) string {
	userID := email
	if userID[:2] != "0x" {
		nodeID := u.generateNodeID(path)
		userID = kaleido.ChildHash(nodeID, email)
	}
	return userID
}

func (u *utilsImpl) readPassword(envVarName string, prompt string) (string, error) {
	// read a passphrase from the enviroment or the terminal and sign with the passphrase
	var passphrase string
	if passphrase = os.Getenv(envVarName); passphrase == "" {
		fmt.Printf(prompt)
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n") // advance to the next line and don't screw someone else up
		if err != nil {
			return "", err
		}
		passphrase = string(bytePassword)
	}
	return passphrase, nil
}

func (u *utilsImpl) newKeyStoreTransactor(from *accounts.Account, keystore *keystore.KeyStore, chainID *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: from.Address,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			var txSigned *types.Transaction
			var err error
			// attempt to sign without a passphrase
			if txSigned, err = keystore.SignTx(*from, tx, chainID); err != nil {
				// if it fails with authentication needed
				if strings.HasPrefix(err.Error(), "authentication needed") {
					// read a passphrase from the enviroment or the terminal and sign with the passphrase
					var passphrase string
					if passphrase, err = u.readPassword("KLD_GETH_KEYSTORE_PASSWORD", "Password needed to unlock keystore: "); err != nil {
						return nil, err
					}
					return keystore.SignTxWithPassphrase(*from, passphrase, tx, chainID)
				}
				// failed for some other reason other than passphrase, just return whatever was returned to us
				return txSigned, err
			}
			// aha, we didn't fail, you should think about securing your keystore
			fmt.Println(">>>> warning: keystore is not secured with a passphrase <<<<")
			return txSigned, nil
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
