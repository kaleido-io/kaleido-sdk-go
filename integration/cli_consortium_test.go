// intgration -- `go install` must be run before tests
package integration

import (
	"encoding/json"
	"os/exec"
	"strings"
	"testing"

	"github.com/nbio/st"

	"github.com/kaleido-io/kaleido-sdk-go/kaleido"
)

func TestKaleido(t *testing.T) {
	kld := exec.Command("../kld")

	output, err := kld.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(output), "Usage") {
		//t.Fatalf("missing usage on default execution: %s", string(output))
	}
}

func createConsortium(t *testing.T) *kaleido.Consortium {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "create", "consortium", "--name", "test-consortium"))

	var consortium kaleido.Consortium
	if err := json.Unmarshal(stdout.Bytes(), &consortium); err != nil {
		t.Fatal(err)
	}

	return &consortium
}

func deleteConsortium(t *testing.T, id string) {
	InvokeCLISuccess(t, exec.Command("../kld", "delete", "consortium", "--id", id))
}

func deleteNonExistingConsortium(t *testing.T) {
	InvokeCLIFailure(t, exec.Command("../kld", "delete", "consortium", "--id", "blahblahablaha"))
}

func listConsortia(t *testing.T) *[]kaleido.Consortium {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "list", "consortium"))

	var consortia []kaleido.Consortium
	if err := json.Unmarshal(stdout.Bytes(), &consortia); err != nil {
		t.Fatal(err)
	}

	return &consortia
}

func getConsortium(t *testing.T, id string) *kaleido.Consortium {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "get", "consortium", "--id", id))

	var consortium kaleido.Consortium
	if err := json.Unmarshal(stdout.Bytes(), &consortium); err != nil {
		t.Fatal(err)
	}

	return &consortium
}

func TestKaleidoConsortium_CreateAndDelete(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)
}

func TestKaleidoConsortium_DeleteNonExisting(t *testing.T) {
	t.Parallel()

	deleteNonExistingConsortium(t)
}

func TestKaleidoConsortium_List(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	newList := listConsortia(t)

	found := false
	for _, c := range *newList {
		if c.ID == consortium.ID {
			found = true
		}
	}
	st.Expect(t, found, true)
}

func TestKaleidoConsortium_Get(t *testing.T) {
	t.Parallel()

	newConsortium := createConsortium(t)
	defer deleteConsortium(t, newConsortium.ID)

	consortium := getConsortium(t, newConsortium.ID)
	st.Expect(t, consortium.ID, newConsortium.ID)
}
