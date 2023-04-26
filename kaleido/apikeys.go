// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
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

type ApiKey struct {
	OrgID string `json:"org_id,omitempty"`
	Name  string `json:"name,omitempty"`
	ID    string `json:"_id,omitempty"`
}

const (
	apikeyBasePath = "/apikeys"
)

func NewApiKey(orgID string) ApiKey {
	return ApiKey{
		OrgID: orgID,
	}
}

func NewApiKeyWithName(orgID, name string) ApiKey {
	return ApiKey{
		OrgID: orgID,
		Name:  name,
	}
}

func (c *KaleidoClient) CreateApiKey(apikey *ApiKey) (*resty.Response, error) {
	path := fmt.Sprintf(apikeyBasePath)
	return c.Client.R().SetBody(apikey).SetResult(apikey).Post(path)
}

func (c *KaleidoClient) UpdateApiKey(apikeyID string, apikey *ApiKey) (*resty.Response, error) {
	path := fmt.Sprintf(apikeyBasePath+"/%s", apikeyID)
	return c.Client.R().SetBody(apikey).SetResult(apikey).Patch(path)
}

func (c *KaleidoClient) GetApiKey(apikeyID string, resultBox *ApiKey) (*resty.Response, error) {
	path := fmt.Sprintf(apikeyBasePath+"/%s", apikeyID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) DeleteApiKey(apikeyID string) (*resty.Response, error) {
	path := fmt.Sprintf(apikeyBasePath+"/%s", apikeyID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListApiKey(resultBox *[]ApiKey) (*resty.Response, error) {
	path := fmt.Sprintf(apikeyBasePath)
	return c.Client.R().SetResult(resultBox).Get(path)
}
