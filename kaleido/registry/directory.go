package registry

import (
	"errors"

	"github.com/spf13/viper"
	resty "gopkg.in/resty.v1"
)

// DirectoryInterface defines the interface for retrieving directory specific client
type DirectoryInterface interface {
	GetContractAddress() string
	GetClient() *resty.Client
}

// NewDirectory attaches to a new directory on the smart contract
func NewDirectory() (DirectoryInterface, error) {
	directory := directoryImpl{}
	err := directory.fetchContractAddress()
	if err != nil {
		return nil, err
	}
	directoryURL := utils().getRegistryURL() + "/" + directory.contractAddress
	directory.client = resty.New().SetHostURL(directoryURL).SetAuthToken(viper.GetString("api.key"))
	return &directory, nil
}

type directoryImpl struct {
	contractAddress string
	client          *resty.Client
}

func (d *directoryImpl) GetClient() *resty.Client {
	return d.client
}

func (d *directoryImpl) GetContractAddress() string {
	return d.contractAddress
}

func (d *directoryImpl) fetchContractAddress() error {
	type responseBody struct {
		_id         string
		name        string
		description string
		directory   string
		properties  string
		claims      string
	}
	var directories []responseBody
	client := utils().getAPIClient()
	_, err := client.R().SetResult(&directories).Get("/directories")
	if err != nil {
		return err
	}
	if len(directories) < 1 {
		return errors.New("Unexpected error: no directories available")
	}
	d.contractAddress = directories[0].directory
	return nil
}
