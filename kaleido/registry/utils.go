package registry

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	resty "gopkg.in/resty.v1"
)

type utilsInterface interface {
	getRegistryURL() string
	getClient() *resty.Client
	validateGetResponse(res *resty.Response, err error, resourceName string) error
	validateCreateResponse(res *resty.Response, err error, resourceName string) error
}

var instance *utilsImpl

func utils() utilsInterface {
	if instance == nil {
		var mutex = &sync.Mutex{}
		mutex.Lock()
		if instance == nil {
			registryURL := viper.GetString("api.url") + "/idregistry/" + viper.GetString("services.idregistry.id")
			client := resty.New().SetHostURL(registryURL).SetAuthToken(viper.GetString("api.key"))
			viper.SetDefault("api.debug", false)
			client.SetDebug(viper.GetBool("api.debug"))
			instance = &utilsImpl{
				registryURL: registryURL,
				client:      client,
			}
		}
		mutex.Unlock()
	}
	return instance
}

type utilsImpl struct {
	registryURL string
	client      *resty.Client
}

func (u *utilsImpl) getRegistryURL() string {
	return u.registryURL
}

var client *resty.Client

func (u *utilsImpl) getClient() *resty.Client {
	return u.client
}

func (u *utilsImpl) validateGetResponse(res *resty.Response, err error, resourceName string) error {
	if res.StatusCode() != 200 {
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
