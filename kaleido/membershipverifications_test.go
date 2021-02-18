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
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

var mockMembershipVerificationCreatePlayload = map[string]interface{}{
	"test_certificate": true,
}

var mockMembershipReturnedFromVerify = map[string]interface{}{
	"_id":                     "zzc1dg3v9r",
	"org_id":                  "zzg3sqzl48",
	"org_name":                "Org3",
	"consortia_id":            "zzane7mkln",
	"verification_type":       "x509",
	"verification_proof":      "-----BEGIN CERTIFICATE-----...",
	"verification_selfsigned": true,
}

func TestMembershipVerificationCreate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia/con1/memberships/member1/verify").
		MatchType("json").
		JSON(mockMembershipVerificationCreatePlayload).
		Reply(200).
		JSON(mockMembershipReturnedFromVerify)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	verifyReq := NewMembershipVerification()
	_, err := client.CreateMembershipVerification("con1", "member1", &verifyReq)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestRegisterMembershipIdentity(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/idregistry/idreg1/identity").
		MatchType("json").
		Reply(200)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	_, err := client.RegisterMembershipIdentity("idreg1", "member1")

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
