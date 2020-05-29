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
	serviceBasePath = "/consortia/%s/environments/%s/services"
)

type Service struct {
	Name         string                 `json:"name"`
	Service      string                 `json:"service"`
	ZoneID       string                 `json:"zone_id"`
	MembershipID string                 `json:"membership_id"`
	ID           string                 `json:"_id,omitempty"`
	State        string                 `json:"state,omitempty"`
	Role         string                 `json:"role,omitempty"`
	Urls         map[string]interface{} `json:"urls,omitempty"`
}

func NewService(name, service, membershipID string, zoneID string) Service {
	return Service{
		Name:         name,
		Service:      service,
		MembershipID: membershipID,
		ZoneID:       zoneID,
		ID:           "",
		State:        "",
	}
}

func (c *KaleidoClient) CreateService(consortium, envID string, service *Service) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath, consortium, envID)
	return c.Client.R().SetResult(service).SetBody(service).Post(path)
}

func (c *KaleidoClient) DeleteService(consortium, envID, serviceID string) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath+"/%s", consortium, envID, serviceID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListServices(consortium, envID string, resultBox *[]Service) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath, consortium, envID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetService(consortiumID, envID, serviceID string, resultBox *Service) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath+"/%s", consortiumID, envID, serviceID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
