// Copyright 2018 Kaleido, a ConsenSys business

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
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

func jsonPrintSingle(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
	return err
}

func jsonPrint(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			if err := jsonPrintSingle(val.Index(i)); err != nil {
				return err
			}
		}
	} else {
		return jsonPrintSingle(v)
	}
	return nil
}

type ethereumValue struct {
	address string
}

func (value *ethereumValue) String() string {
	return value.address
}

func (value *ethereumValue) Set(address string) error {
	if !eth.IsHexAddress(address) {
		return errors.New("must be a valid ethereum address")
	}
	value.address = address
	return nil
}

func (value *ethereumValue) Type() string {
	return "ethereum-address"
}
