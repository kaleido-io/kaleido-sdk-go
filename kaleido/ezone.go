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

	resty "gopkg.in/resty.v1"
)

const (
	ezoneBasePath = "/consortia/%s/environments/%s/zones"
)

type EZone struct {
	Name   string `json:"name"`
	Id     string `json:"_id,omitempty"`
	Region string `json:"region,omitempty"`
	Cloud  string `json:"cloud,omitempty"`
	Type   string `json:"type,omitempty"`
}

func NewEZone(name, region, cloud string) EZone {
	return EZone{
		Name:   name,
		Id:     "",
		Region: region,
		Cloud:  cloud,
		Type:   "kaleido",
	}
}

func (c *KaleidoClient) CreateEZone(consortium, envId string, ezone *EZone) (*resty.Response, error) {
	path := fmt.Sprintf(ezoneBasePath, consortium, envId)
	return c.Client.R().SetResult(ezone).SetBody(ezone).Post(path)
}

func (c *KaleidoClient) DeleteEZone(consortium, envId, ezoneId string) (*resty.Response, error) {
	path := fmt.Sprintf(ezoneBasePath+"/%s", consortium, envId, ezoneId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListEZones(consortium, envId string, resultBox *[]EZone) (*resty.Response, error) {
	path := fmt.Sprintf(ezoneBasePath, consortium, envId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetEZone(consortiumId, envId, ezoneId string, resultBox *EZone) (*resty.Response, error) {
	path := fmt.Sprintf(ezoneBasePath+"/%s", consortiumId, envId, ezoneId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
