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
	invBasePath = "/consortia/%s/invitations"
)

type Invitation struct {
	OrgName string `json:"org_name"`
	Email	string `json:"email"`
	Id      string `json:"_id,omitempty"`
	State	string `json:"state,omitempty"`
}

func NewInvitation(orgName, email string) Invitation {
	return Invitation{orgName, email, "", ""}
}

func (c *KaleidoClient) ListInvitations(consortiaId string, resultBox *[]Invitation) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath, consortiaId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) CreateInvitation(consortiaId string, invitation *Invitation) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath, consortiaId)
	return c.Client.R().SetResult(invitation).SetBody(invitation).Post(path)
}

func (c *KaleidoClient) DeleteInvitation(consortiaId, invitationId string) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath+"/%s", consortiaId, invitationId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) GetInvitation(consortiaId, invitationId string, resultBox *Invitation) (*resty.Response, error) {
	path := fmt.Sprintf(invBasePath+"/%s", consortiaId, invitationId)
	return c.Client.R().SetResult(resultBox).Get(path)
}