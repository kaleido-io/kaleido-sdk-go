package kaleido

import (
	"os"
	"testing"
)

func TestAppKey(t *testing.T) {
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

	var envs []Environment
	res, err = client.ListEnvironments(consortium.Id, &envs)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not create environment! Status: %d.", res.StatusCode())
	}
	t.Logf("Envs: %v", envs)
	if len(envs) != 1 {
		t.Fatalf("New consortium has unexpected number of envs %d.", len(envs))
	}

	env := envs[0]
	var members []Membership
	res, err = client.ListMemberships(consortium.Id, &members)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not list memberships.")
	}

	member := members[0]

	appKey := NewAppKey(member.Id)
	res, err = client.CreateAppKey(consortium.Id, env.Id, &appKey)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 201 {
		t.Fatalf("Could not create AppKey! Status: %d.", res.StatusCode())
	}

	if appKey.Password == "" {
		t.Fatalf("AppKey did not include a password! %v", appKey)
	}

	if appKey.Username == "" {
		t.Fatalf("AppKey did not include a username! %v", appKey)
	}

	var appKey2 AppKey
	res, err = client.GetAppKey(consortium.Id, env.Id, appKey.Id, &appKey2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Failed to fetch remote AppKey for id %s. Status: %d", appKey.Id, res.StatusCode())
	}
	if appKey.Id != appKey2.Id {
		t.Fatalf("Fetched AppKey %s id did not match original %s.", appKey.Id, appKey2.Id)
	}

	var appKeys []AppKey
	res, err = client.ListAppKeys(consortium.Id, env.Id, &appKeys)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Failed to list App Keys. Status: %d.", res.StatusCode())
	}

	if len(appKeys) != 1 {
		t.Fatalf("Expected 1 AppKey found %d.", len(appKeys))
	}

	res, err = client.DeleteAppKey(consortium.Id, env.Id, appKey.Id)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 204 {
		t.Fatalf("Could not delete AppKey %s. Status: %d", appKey.Id, res.StatusCode())
	}
}
