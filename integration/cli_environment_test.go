package integration

import (
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/nbio/st"

	"github.com/kaleido-io/kaleido-sdk-go/kaleido"
)

func createEnvironment(t *testing.T, consortiumID string) *kaleido.Environment {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "create", "environment", "--name", "test-environment", "--consortium", consortiumID))

	var environment kaleido.Environment
	if err := json.Unmarshal(stdout.Bytes(), &environment); err != nil {
		t.Fatal(err)
	}

	return &environment
}

func deleteEnvironment(t *testing.T, consortiumID string, id string) {
	InvokeCLISuccess(t, exec.Command("../kld", "delete", "environment", "--consortium", consortiumID, "--id", id))
}

func getEnvironment(t *testing.T, consortiumID string, id string) *kaleido.Environment {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "get", "environment", "--consortium", consortiumID, "--id", id))

	var environment kaleido.Environment
	if err := json.Unmarshal(stdout.Bytes(), &environment); err != nil {
		t.Fatal(err)
	}

	return &environment
}

func listEnvironments(t *testing.T, consortiumID string) *[]kaleido.Environment {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "list", "environment", "--consortium", consortiumID))

	var environments []kaleido.Environment
	if err := json.Unmarshal(stdout.Bytes(), &environments); err != nil {
		t.Fatal(err)
	}

	return &environments
}

func TestEnvironment_Create(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	environment := createEnvironment(t, consortium.Id)
	defer deleteEnvironment(t, consortium.Id, environment.Id)
}

func TestEnvironment_Get(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	environment := createEnvironment(t, consortium.Id)
	defer deleteEnvironment(t, consortium.Id, environment.Id)

	environment2 := getEnvironment(t, consortium.Id, environment.Id)
	st.Expect(t, environment.Id, environment2.Id)
}

func TestEnvironment_List(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	origList := listEnvironments(t, consortium.Id)
	st.Expect(t, len(*origList), 0)

	environment := createEnvironment(t, consortium.Id)
	defer deleteEnvironment(t, consortium.Id, environment.Id)

	newList := listEnvironments(t, consortium.Id)
	st.Expect(t, len(*newList), 1)
}

func TestEnvironment_Delete(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	environment := createEnvironment(t, consortium.Id)
	defer deleteEnvironment(t, consortium.Id, environment.Id)
}
