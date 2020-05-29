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

	gock "gopkg.in/h2non/gock.v1"
)

func TestInvitation(t *testing.T) {

	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia").
		MatchType("json").
		JSON(map[string]interface{}{
			"name":        "invitationTest",
			"description": "invitations",
		}).
		Reply(201).
		JSON(Consortium{
			ID:          "cons1",
			Name:        "invitationTest",
			Description: "invitations",
		})

	gock.New("http://example.com").
		Post("/api/v1/consortia/cons1/invitations").
		MatchType("json").
		JSON(map[string]interface{}{
			"org_name": "Test Organization",
			"email":    "someone@example.com",
		}).
		Reply(201).
		JSON(Invitation{
			ID:      "inv1",
			OrgName: "Test Organization",
			Email:   "someone@example.com",
		})

	gock.New("http://example.com").
		Get("/api/v1/consortia/cons1/invitations/inv1").
		Reply(200).
		JSON(Invitation{
			ID:      "inv1",
			OrgName: "Test Organization",
			Email:   "someone@example.com",
			State:   "pending",
		})

	gock.New("http://example.com").
		Delete("/api/v1/consortia/cons1/invitations/inv1").
		Reply(204)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	consortium := NewConsortium("invitationTest", "invitations")
	res, err := client.CreateConsortium(&consortium)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 201 {
		t.Fatalf("Failed to create consortium with status: %d", res.StatusCode())
	}
	defer client.DeleteConsortium(consortium.ID)

	invitation := NewInvitation("Test Organization", "someone@example.com")
	res, err = client.CreateInvitation(consortium.ID, &invitation)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Failed to create invitation with status: %d", res.StatusCode())
	}

	var invitation2 Invitation
	res, err = client.GetInvitation(consortium.ID, invitation.ID, &invitation2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Logf("%s", res.Request.URL)
		t.Fatalf("Failed to fetch invitation with status: %d", res.StatusCode())
	}

	if invitation.ID != invitation2.ID {
		t.Fatalf("Fetched invitation id %s did not match %s.", invitation2.ID, invitation.ID)
	}

	res, err = client.DeleteInvitation(consortium.ID, invitation.ID)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 204 {
		t.Fatalf("Failed to delete invitation %s with status: %d", invitation.ID, res.StatusCode())
	}
}
