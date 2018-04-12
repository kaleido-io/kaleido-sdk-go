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
	if len(envs) != 1 {
		t.Fatalf("New consortiums have a single auto generated environment.")
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
