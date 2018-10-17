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
	serviceBasePath = "/c/%s/e/%s/services"
)

type Service struct {
	Id            string `json:"_id,omitempty"`
	MembershipId  string `json:"membership_id,omitempty"`
	Service       string `json:"service,omitempty"`
	Name          string `json:"name,omitempty"`
	ServiceGUID   string `json:"service_guid,omitempty"`
	ServiceType   string `json:"service_type,omitempty"`
	State         string `json:"state,omitempty"`
	Revision      string `json:"_revision,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	EnvironmentId string `json:"environment_id,omitempty"`
	URLs          *struct {
		Http string `json:"http,omitempty"`
	} `json:"urls,omitempty"`
}

func NewService(service, name, membershipId string) Service {
	return Service{
		Name:         name,
		MembershipId: membershipId,
		Service:      service,
	}
}

func (c *KaleidoClient) ListServices(consortiumId, envId string, resultBox *[]Service) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath, consortiumId, envId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) DeployService(consortium, envId string, service *Service) (*resty.Response, error) {
	path := fmt.Sprintf(serviceBasePath, consortium, envId)
	return c.Client.R().SetResult(service).SetBody(service).Post(path)
}

func (c *KaleidoClient) GetService(consortium, envId, serviceId string, resultBox *[]Service) (*resty.Response, error) {
	return c.Client.R().SetQueryParam("_id", serviceId).SetResult(resultBox).Get("/services")
}
