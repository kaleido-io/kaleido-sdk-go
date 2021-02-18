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
	memVerifyBasePath    = "/consortia/%s/memberships/%s/verify"
	registerIdentityPath = "/idregistry/%s/identity"
)

type MembershipVerification struct {
	TestCertificate bool   `json:"test_certificate,omitempty"`
	ProofID         string `json:"proof_id,omitempty"`
}

type MembershipIdentityRegistration struct {
	MembershipID string `json:"membership_id,omitempty"`
}

func NewMembershipVerification() MembershipVerification {
	return MembershipVerification{
		TestCertificate: true,
	}
}

func (c *KaleidoClient) CreateMembershipVerification(consortiaID, membershipID string, verification *MembershipVerification) (*resty.Response, error) {
	path := fmt.Sprintf(memVerifyBasePath, consortiaID, membershipID)
	var membership Membership
	return c.Client.R().SetResult(&membership).SetBody(verification).Post(path)
}

func (c *KaleidoClient) RegisterMembershipIdentity(idregistryID, membershipID string) (*resty.Response, error) {
	path := fmt.Sprintf(registerIdentityPath, idregistryID)
	return c.Client.R().SetBody(MembershipIdentityRegistration{MembershipID: membershipID}).Post(path)
}
