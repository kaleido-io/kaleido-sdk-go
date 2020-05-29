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
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Provider      string        `json:"provider"`
	ConsensusType string        `json:"consensus_type"`
	ID            string        `json:"_id,omitempty"`
	State         string        `json:"state,omitempty"`
	ReleaseID     string        `json:"release_id,omitempty"`
	TestFeatures  *TestFeatures `json:"test_features,omitempty"`
	BlockPeriod   int           `json:"block_period"`
}

type TestFeatures struct {
	MultiRegion bool `json:"multi_region,omitempty"`
}

const (
	envBasePath = "/consortia/%s/environments"
)

func NewEnvironment(name, description, provider, consensus string, multiRegion bool, blockPeriod int) Environment {
	e := Environment{
		Name:          name,
		Description:   description,
		Provider:      provider,
		ConsensusType: consensus,
		BlockPeriod:   blockPeriod,
	}
	if multiRegion {
		e.TestFeatures = &TestFeatures{
			MultiRegion: multiRegion,
		}
	}
	return e
}

func (c *KaleidoClient) ListEnvironments(consortiumID string, resultBox *[]Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, consortiumID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) CreateEnvironment(consortiumID string, environment *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, consortiumID)
	return c.Client.R().SetResult(environment).SetBody(environment).Post(path)
}

func (c *KaleidoClient) DeleteEnvironment(consortiumID, environmentID string) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", consortiumID, environmentID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) GetEnvironment(consortiumID, environmentID string, resultBox *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", consortiumID, environmentID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
