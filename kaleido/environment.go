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

type Environment struct {
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Provider      string       `json:"provider"`
	ConsensusType string       `json:"consensus_type"`
	Id            string       `json:"_id,omitempty"`
	State         string       `json:"state,omitempty"`
	ReleaseId     string       `json:"release_id,omitempty"`
	TestFeatures  TestFeatures `json:"test_features,omitempty"`
}

type TestFeatures struct {
	MultiRegion bool `json:"multi_region,omitempty"`
}

const (
	envBasePath = "/consortia/%s/environments"
)

func NewEnvironment(name, description, provider, consensus string, multiRegion bool) Environment {
	return Environment{
		Name:          name,
		Description:   description,
		Provider:      provider,
		ConsensusType: consensus,
		TestFeatures: TestFeatures{
			MultiRegion: multiRegion,
		},
	}
}

func (c *KaleidoClient) ListEnvironments(consortiumId string, resultBox *[]Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, consortiumId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) CreateEnvironment(consortiumId string, environment *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, consortiumId)
	return c.Client.R().SetResult(environment).SetBody(environment).Post(path)
}

func (c *KaleidoClient) DeleteEnvironment(consortiumId, environmentId string) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", consortiumId, environmentId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) GetEnvironment(consortiumId, environmentId string, resultBox *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", consortiumId, environmentId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
