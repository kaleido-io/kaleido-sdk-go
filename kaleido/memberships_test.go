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

var mockMembershipCreatePlayload = map[string]string{
	"org_name": "member1",
}

var mockMembership = map[string]string{
	"org_name":     "member1",
	"org_id":       "zzgl55vock",
	"state":        "active",
	"_id":          "zze8pz9jed",
	"consortia_id": "cid",
}

var mockMemberships = []map[string]string{mockMembership}

func TestMembershipCreate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia/cid/memberships").
		MatchType("json").
		JSON(mockMembershipCreatePlayload).
		Reply(201).
		JSON(mockMembership)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	membership := NewMembership("member1")
	_, err := client.CreateMembership("cid", &membership)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestMembershipGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/cid/memberships/zze8pz9jed").
		Reply(200).
		JSON(mockMembership)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var membership Membership
	_, err := client.GetMembership("cid", "zze8pz9jed", &membership)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestMembershipList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/cid/memberships").
		Reply(200).
		JSON(mockMemberships)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var memberships []Membership
	_, err := client.ListMemberships("cid", &memberships)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestMembershipDelete(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Delete("/api/v1/consortia/cid/memberships/zze8pz9jed").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	_, err := client.DeleteMembership("cid", "zze8pz9jed")

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestMembershipUpdate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Patch("/api/v1/consortia/cid/memberships/zze8pz9jed").
		Reply(200)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	_, err := client.UpdateMembership("cid", "zze8pz9jed", &Membership{OrgName: "new name"})

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
