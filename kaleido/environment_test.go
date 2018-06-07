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

func TestEnvironmentCreationDeletion(t *testing.T) {
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	consortium := NewConsortium("envCreateTest", "creating an environment", "single-org")
	_, err := client.CreateConsortium(&consortium)
	if err != nil {
		t.Error(err)
	}
	defer client.DeleteConsortium(consortium.Id)
	var envs []Environment
	client.ListEnvironments(consortium.Id, &envs)
	t.Logf("Envs: %v", envs)
	if len(envs) != 0 {
		t.Fatalf("New consortium should be empty.")
	}

	env := NewEnvironment("testingEnvironment", "just test", "quorum", "raft")
	res, err := client.CreateEnvironment(consortium.Id, &env)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 201 {
		t.Fatalf("Could not create environment status code: %d", res.StatusCode())
	}
	t.Logf("Env: %v", env)

	var env2 Environment
	res, err = client.GetEnvironment(consortium.Id, env.Id, &env2)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not get environment %s. Status was %d", env.Id, res.StatusCode())
	}

	if env.Id != env2.Id {
		t.Fatalf("Id mismatch on GetEnvironment %s and %s", env.Id, env2.Id)
	}

	if env2.State == "" {
		t.Fatal("Fetched environment should have a state was empty.")
	}

	//Delete all testing environments
	res, err = client.ListEnvironments(consortium.Id, &envs)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not list environments status code: %d", res.StatusCode())
	}

	for _, v := range envs {
		res, err := client.DeleteEnvironment(consortium.Id, v.Id)
		if err != nil {
			t.Fatal(err)
		}
		if (res.StatusCode() != 202) && (res.StatusCode() != 204) {
			t.Fatalf("Could not delete environment %s status: %d", v.Id, res.StatusCode())
		}
	}
}
