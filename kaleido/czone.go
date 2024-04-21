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
	czoneBasePath = "/consortia/%s/zones"
)

type CZone struct {
	Name   string `json:"name,omitempty"`
	ID     string `json:"_id,omitempty"`
	Region string `json:"region,omitempty"`
	Cloud  string `json:"cloud,omitempty"`
	Type   string `json:"type,omitempty"`
}

func NewCZone(name, region, cloud string) CZone {
	return CZone{
		Name:   name,
		ID:     "",
		Region: region,
		Cloud:  cloud,
		Type:   "kaleido",
	}
}

func (c *KaleidoClient) CreateCZone(consortium string, ezone *CZone) (*resty.Response, error) {
	path := fmt.Sprintf(czoneBasePath, consortium)
	return c.Client.R().SetResult(ezone).SetBody(ezone).Post(path)
}

func (c *KaleidoClient) UpdateCZone(consortium, ezoneID string, ezone *CZone) (*resty.Response, error) {
	path := fmt.Sprintf(czoneBasePath+"/%s", consortium, ezoneID)
	return c.Client.R().SetResult(ezone).SetBody(ezone).Patch(path)
}

func (c *KaleidoClient) DeleteCZone(consortium, ezoneID string) (*resty.Response, error) {
	path := fmt.Sprintf(czoneBasePath+"/%s", consortium, ezoneID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListCZones(consortium string, resultBox *[]CZone) (*resty.Response, error) {
	path := fmt.Sprintf(czoneBasePath, consortium)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetCZone(consortiumID, ezoneID string, resultBox *CZone) (*resty.Response, error) {
	path := fmt.Sprintf(czoneBasePath+"/%s", consortiumID, ezoneID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
