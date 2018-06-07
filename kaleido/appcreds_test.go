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

func TestAppCreds(t *testing.T) {
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	consortium := NewConsortium("apiKeyTest", "creating api key", "single-org")
	res, err := client.CreateConsortium(&consortium)

	if res.StatusCode() != 201 {
		t.Fatalf("Could not create consortium status code: %d.", res.StatusCode())
	}

	if err != nil {
		t.Fatal(err)
	}

	defer client.DeleteConsortium(consortium.Id)

	var env Environment
	client.CreateEnvironment(consortium.Id, &env)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 201 {
		t.Fatalf("Could not create environment! Status: %d.", res.StatusCode())
	}

	var members []Membership
	res, err = client.ListMemberships(consortium.Id, &members)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not list memberships.")
	}

	member := members[0]

	appcreds := NewAppCreds(member.Id)
	res, err = client.CreateAppCreds(consortium.Id, env.Id, &appcreds)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 201 {
		t.Fatalf("Could not create AppCreds! Status: %d.", res.StatusCode())
	}

	if appcreds.Password == "" {
		t.Fatalf("AppCreds did not include a password! %v", appcreds)
	}

	if appcreds.Username == "" {
		t.Fatalf("AppCreds did not include a username! %v", appcreds)
	}

	var appcreds2 AppCreds
	res, err = client.GetAppCreds(consortium.Id, env.Id, appcreds.Id, &appcreds2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Failed to fetch remote AppCreds for id %s. Status: %d", appcreds.Id, res.StatusCode())
	}
	if appcreds.Id != appcreds2.Id {
		t.Fatalf("Fetched AppCreds %s id did not match original %s.", appcreds.Id, appcreds2.Id)
	}

	var appcreds []AppCreds
	res, err = client.ListAppCreds(consortium.Id, env.Id, &appcreds)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Failed to list App Keys. Status: %d.", res.StatusCode())
	}

	if len(appcreds) != 1 {
		t.Fatalf("Expected 1 AppCreds found %d.", len(appcreds))
	}

	res, err = client.DeleteAppCreds(consortium.Id, env.Id, appcreds.Id)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 204 {
		t.Fatalf("Could not delete AppCreds %s. Status: %d", appcreds.Id, res.StatusCode())
	}
}
