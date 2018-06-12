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

// +build integration

package kaleido

import (
	"os"
	"testing"
)

func TestMembership(t *testing.T) {
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	consortium := NewConsortium("membershipTest", "members", "single-org")
	res, err := client.CreateConsortium(&consortium)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 201 {
		t.Fatalf("Failed to create consortium with status: %d", res.StatusCode())
	}
	defer client.DeleteConsortium(consortium.Id)

	membership := NewMembership("macdonalds")
	res, err = client.CreateMembership(consortium.Id, &membership)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Failed to create membership with status: %d", res.StatusCode())
	}

	var membership2 Membership
	res, err = client.GetMembership(consortium.Id, membership.Id, &membership2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Logf("%s", res.Request.URL)
		t.Fatalf("Failed to fetch membership with status: %d", res.StatusCode())
	}

	if membership.Id != membership2.Id {
		t.Fatalf("Fetched memberhsip id %s did not match %s.", membership2.Id, membership.Id)
	}

	res, err = client.DeleteMembership(consortium.Id, membership.Id)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 204 {
		t.Fatalf("Failed to delete membership %s with status: %d", membership.Id, res.StatusCode())
	}
}
