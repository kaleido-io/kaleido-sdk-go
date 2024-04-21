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
	invBasePath = "/consortia/%s/invitations"
)

type Invitation struct {
	OrgName string `json:"org_name,omitempty"`
	Email   string `json:"email,omitempty"`
	ID      string `json:"_id,omitempty"`
	State   string `json:"state,omitempty"`
}

func NewInvitation(orgName, email string) Invitation {
	return Invitation{orgName, email, "", ""}
}

func (c *KaleidoClient) ListInvitations(consortiaID string, resultBox *[]Invitation) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath, consortiaID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) CreateInvitation(consortiaID string, invitation *Invitation) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath, consortiaID)
	return c.Client.R().SetResult(invitation).SetBody(invitation).Post(path)
}

func (c *KaleidoClient) UpdateInvitation(consortiaID, invitationID string, invitation *Invitation) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath+"/%s", consortiaID, invitationID)
	return c.Client.R().SetResult(invitation).SetBody(invitation).Patch(path)
}

func (c *KaleidoClient) DeleteInvitation(consortiaID, invitationID string) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath+"/%s", consortiaID, invitationID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) GetInvitation(consortiaID, invitationID string, resultBox *Invitation) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath+"/%s", consortiaID, invitationID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
