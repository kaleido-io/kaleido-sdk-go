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

// Environment fields
type Environment struct {
	Name              string                 `json:"name,omitempty"`
	Description       string                 `json:"description,omitempty"`
	Provider          string                 `json:"provider,omitempty"`
	ConsensusType     string                 `json:"consensus_type,omitempty"`
	ID                string                 `json:"_id,omitempty,omitempty"`
	State             string                 `json:"state,omitempty"`
	ReleaseID         string                 `json:"release_id,omitempty"`
	TestFeatures      map[string]interface{} `json:"test_features,omitempty"`
	BlockPeriod       int                    `json:"block_period,omitempty"`
	ChainID           uint                   `json:"chain_id,omitempty"`
	PrefundedAccounts map[string]interface{} `json:"prefunded_accounts,omitempty"`
}

// AccountBalance represents an account's balance
type AccountBalance struct {
	Balance string `json:"balance,omitempty"`
}

const (
	envBasePath = "/consortia/%s/environments"
)

// NewEnvironment creates a new environment
func NewEnvironment(name, description, provider, consensus string, multiRegion bool, blockPeriod int, prefundedAccounts map[string]string, chainID uint) Environment {
	accounts := map[string]interface{}{}
	for account, balance := range prefundedAccounts {
		accountBalance := &AccountBalance{}
		accountBalance.Balance = balance
		accounts[account] = accountBalance
	}
	e := Environment{
		Name:              name,
		Description:       description,
		Provider:          provider,
		ConsensusType:     consensus,
		BlockPeriod:       blockPeriod,
		PrefundedAccounts: accounts,
		TestFeatures:      make(map[string]interface{}),
	}
	if chainID > 0 {
		e.ChainID = chainID
	}
	if multiRegion {
		e.TestFeatures["multi_region"] = &multiRegion
	}
	return e
}

// ListEnvironments lists existing environment
func (c *KaleidoClient) ListEnvironments(consortiumID string, resultBox *[]Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, consortiumID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

// CreateEnvironment initiates request to create a new environment
func (c *KaleidoClient) CreateEnvironment(consortiumID string, environment *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, consortiumID)
	return c.Client.R().SetResult(environment).SetBody(environment).Post(path)
}

// UpdateEnvironment updates an environment
func (c *KaleidoClient) UpdateEnvironment(consortiumID, environmentID string, environment *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", consortiumID, environmentID)
	return c.Client.R().SetResult(environment).SetBody(environment).Patch(path)
}

// DeleteEnvironment deletes an environment
func (c *KaleidoClient) DeleteEnvironment(consortiumID, environmentID string) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", consortiumID, environmentID)
	return c.Client.R().Delete(path)
}

// GetEnvironment details
func (c *KaleidoClient) GetEnvironment(consortiumID, environmentID string, resultBox *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", consortiumID, environmentID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
