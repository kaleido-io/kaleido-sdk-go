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
	memBasePath = "/consortia/%s/memberships"
)

type Membership struct {
	OrgName           string `json:"org_name,omitempty"`
	ID                string `json:"_id,omitempty"`
	VerificationProof string `json:"verification_proof,omitempty"`
}

func NewMembership(orgName string) Membership {
	return Membership{orgName, "", ""}
}

func (c *KaleidoClient) ListMemberships(consortiaID string, resultBox *[]Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath, consortiaID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) CreateMembership(consortiaID string, membership *Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath, consortiaID)
	return c.Client.R().SetResult(membership).SetBody(membership).Post(path)
}

func (c *KaleidoClient) UpdateMembership(consortiaID, membershipID string, membership *Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath+"/%s", consortiaID, membershipID)
	return c.Client.R().SetResult(membership).SetBody(membership).Patch(path)
}

func (c *KaleidoClient) DeleteMembership(consortiaID, membershipID string) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath+"/%s", consortiaID, membershipID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) GetMembership(consortiaID, membershipID string, resultBox *Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath+"/%s", consortiaID, membershipID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
