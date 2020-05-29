package integration

import (
	"encoding/json"
	"errors"
	"os/exec"
	"testing"

	"github.com/nbio/st"

	"github.com/kaleido-io/kaleido-sdk-go/kaleido"
)

func createAppCreds(t *testing.T, consortiumID string, environmentID string, membershipID string) *kaleido.AppCreds {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "create", "appcreds", "-c", consortiumID, "-e", environmentID, "-m", membershipID))

	var appcreds kaleido.AppCreds
	if err := json.Unmarshal(stdout.Bytes(), &appcreds); err != nil {
		t.Fatal(err)
	}

	return &appcreds
}

func getAppCreds(t *testing.T, consortiumID string, environmentID string, id string) *kaleido.AppCreds {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "get", "appcreds", "-c", consortiumID, "-e", environmentID, "-a", id))

	var appcreds kaleido.AppCreds
	if err := json.Unmarshal(stdout.Bytes(), &appcreds); err != nil {
		t.Fatal(err)
	}

	return &appcreds
}

func listAppCreds(t *testing.T, consortiumID string, environmentID string) *[]kaleido.AppCreds {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "list", "appcreds", "-c", consortiumID, "-e", environmentID))

	var appcreds []kaleido.AppCreds
	if err := json.Unmarshal(stdout.Bytes(), &appcreds); err != nil {
		t.Fatal(err)
	}

	return &appcreds
}

func deleteAppCreds(t *testing.T, consortiumID string, environmentID string, id string) {
	InvokeCLISuccess(t, exec.Command("../kld", "delete", "appcreds", "-c", consortiumID, "-e", environmentID, "-a", id))
}

func TestAppCredsCreate(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	environment := createEnvironment(t, consortium.ID)
	defer deleteEnvironment(t, consortium.ID, environment.ID)

	memberships := listMembership(t, consortium.ID)

	appCreds := createAppCreds(t, consortium.ID, environment.ID, (*memberships)[0].ID)
	defer deleteAppCreds(t, consortium.ID, environment.ID, appCreds.ID)

	if len(appCreds.Username) == 0 || len(appCreds.Password) == 0 {
		t.Fatal(errors.New("appcreds create failed, username or password is empty"))
	}
}

func TestAppCredsGet(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	environment := createEnvironment(t, consortium.ID)
	defer deleteEnvironment(t, consortium.ID, environment.ID)

	memberships := listMembership(t, consortium.ID)

	appCreds := createAppCreds(t, consortium.ID, environment.ID, (*memberships)[0].ID)
	defer deleteAppCreds(t, consortium.ID, environment.ID, appCreds.ID)

	appCredsSaved := getAppCreds(t, consortium.ID, environment.ID, appCreds.ID)
	st.Expect(t, appCreds.ID, appCredsSaved.ID)
}

func TestAppCredsList(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	environment := createEnvironment(t, consortium.ID)
	defer deleteEnvironment(t, consortium.ID, environment.ID)

	memberships := listMembership(t, consortium.ID)

	appCreds := createAppCreds(t, consortium.ID, environment.ID, (*memberships)[0].ID)
	defer deleteAppCreds(t, consortium.ID, environment.ID, appCreds.ID)

	newList := listAppCreds(t, consortium.ID, environment.ID)
	found := false
	for _, a := range *newList {
		if a.ID == appCreds.ID {
			found = true
		}
	}
	st.Expect(t, found, true)
}

func TestAppCredsDelete(t *testing.T) {
	// covered during create with the deferred delete
}
