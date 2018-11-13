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
	"os"
	"testing"
)

func TestInvitation(t *testing.T) {
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	consortium := NewConsortium("invitationTest", "invitations", "single-org")
	res, err := client.CreateConsortium(&consortium)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 201 {
		t.Fatalf("Failed to create consortium with status: %d", res.StatusCode())
	}
	defer client.DeleteConsortium(consortium.Id)

	invitation := NewInvitation("Test Organization","peter.broadhurst@consensys.net")
	res, err = client.CreateInvitation(consortium.Id, &invitation)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Failed to create invitation with status: %d", res.StatusCode())
	}

	var invitation2 Invitation
	res, err = client.GetInvitation(consortium.Id, invitation.Id, &invitation2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Logf("%s", res.Request.URL)
		t.Fatalf("Failed to fetch invitation with status: %d", res.StatusCode())
	}

	if invitation.Id != invitation2.Id {
		t.Fatalf("Fetched invitation id %s did not match %s.", invitation2.Id, invitation.Id)
	}

	res, err = client.DeleteInvitation(consortium.Id, invitation.Id)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 204 {
		t.Fatalf("Failed to delete invitation %s with status: %d", invitation.Id, res.StatusCode())
	}
}
