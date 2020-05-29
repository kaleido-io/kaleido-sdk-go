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
	configurationsBasePath = "/consortia/%s/environments/%s/configurations"
)

type Configuration struct {
	ID           string                 `json:"_id,omitempty"`
	Name         string                 `json:"name,omitempty"`
	MembershipID string                 `json:"membership_id,omitempty"`
	Type         string                 `json:"type,omitempty"`
	Details      map[string]interface{} `json:"details"`
}

func NewConfiguration(name, membershipID, configType string, details map[string]interface{}) Configuration {
	return Configuration{
		Name:         name,
		MembershipID: membershipID,
		Type:         configType,
		Details:      details,
	}
}

func (c *KaleidoClient) CreateConfiguration(consortium, envID string, config *Configuration) (*resty.Response, error) {
	path := fmt.Sprintf(configurationsBasePath, consortium, envID)
	return c.Client.R().SetResult(config).SetBody(config).Post(path)
}

func (c *KaleidoClient) DeleteConfiguration(consortium, envID, configID string) (*resty.Response, error) {
	path := fmt.Sprintf(configurationsBasePath+"/%s", consortium, envID, configID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListConfigurations(consortium, envID string, resultBox *[]Configuration) (*resty.Response, error) {
	path := fmt.Sprintf(configurationsBasePath, consortium, envID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetConfiguration(consortiumID, envID, configID string, resultBox *Configuration) (*resty.Response, error) {
	path := fmt.Sprintf(configurationsBasePath+"/%s", consortiumID, envID, configID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
