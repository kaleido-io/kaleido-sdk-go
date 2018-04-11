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
	_, err = client.CreateEnvironment(consortium.Id, &env)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Env: %v", env)

	//Delete all testing environments
	client.ListEnvironments(consortium.Id, &envs)
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
