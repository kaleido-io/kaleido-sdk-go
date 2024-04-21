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
	serviceBasePath = "/consortia/%s/environments/%s/services"
)

type Service struct {
	Name                 string                 `json:"name,omitempty"`
	Service              string                 `json:"service,omitempty"`
	ServiceType          string                 `json:"service_type,omitempty"`
	ZoneID               string                 `json:"zone_id,omitempty"`
	MembershipID         string                 `json:"membership_id,omitempty"`
	ID                   string                 `json:"_id,omitempty"`
	Size                 string                 `json:"size,omitempty"`
	State                string                 `json:"state,omitempty"`
	HybridPortAllocation int64                  `json:"hybrid_port_allocation,omitempty"`
	Urls                 map[string]interface{} `json:"urls,omitempty"`
	Details              map[string]interface{} `json:"details,omitempty"`
}

func NewService(name, service, membershipID string, zoneID string, details map[string]interface{}) Service {
	return Service{
		Name:         name,
		Service:      service,
		MembershipID: membershipID,
		ZoneID:       zoneID,
		ID:           "",
		State:        "",
		Details:      details,
	}
}

func (c *KaleidoClient) CreateService(consortium, envID string, service *Service) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath, consortium, envID)
	return c.Client.R().SetResult(service).SetBody(service).Post(path)
}

func (c *KaleidoClient) UpdateService(consortium, envID, serviceID string, service *Service) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath+"/%s", consortium, envID, serviceID)
	return c.Client.R().SetResult(service).SetBody(service).Patch(path)
}

func (c *KaleidoClient) ResetService(consortium, envID, serviceID string) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath+"/%s/reset", consortium, envID, serviceID)
	return c.Client.R().SetBody(map[string]string{}).Put(path)
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
