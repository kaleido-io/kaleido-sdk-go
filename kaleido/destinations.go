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

package kaleido

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

const (
	destBasePath = "/%s/%s/destinations"
)

type Destination struct {
	Name           string `json:"name,omitempty"`
	KaleidoManaged bool   `json:"kaleido_managed,omitempty"`
	URI            string `json:"uri,omitempty"`
	SetupComplete  bool   `json:"setup_complete,omitempty"`
}

func NewDestination(name string) Destination {
	return Destination{
		Name:           name,
		KaleidoManaged: true,
	}
}

func (c *KaleidoClient) CreateDestination(serviceType, serviceID string, destination *Destination) (*resty.Response, error) {
	path := fmt.Sprintf(destBasePath+"/%s", serviceType, serviceID, destination.Name)
	return c.Client.R().SetResult(destination).SetBody(destination).Put(path)
}

func (c *KaleidoClient) UpdateDestination(serviceType, serviceID, destName string, destination *Destination) (*resty.Response, error) {
	path := fmt.Sprintf(destBasePath+"/%s", serviceType, serviceID, destName)
	return c.Client.R().SetResult(destination).SetBody(destination).Put(path)
}

func (c *KaleidoClient) DeleteDestination(serviceType, serviceID, destName string) (*resty.Response, error) {
	path := fmt.Sprintf(destBasePath+"/%s", serviceType, serviceID, destName)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListDestinations(serviceType, serviceID string, resultBox *[]Destination) (*resty.Response, error) {
	path := fmt.Sprintf(destBasePath, serviceType, serviceID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetDestination(serviceType, serviceID, destName string, resultBox *Destination) (*resty.Response, error) {
	path := fmt.Sprintf(destBasePath+"/%s", serviceType, serviceID, destName)
	return c.Client.R().SetResult(resultBox).Get(path)
}
