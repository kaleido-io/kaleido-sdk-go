package integration

import (
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/nbio/st"

	"github.com/kaleido-io/kaleido-sdk-go/kaleido"
)

func createMembership(t *testing.T, consortiumID string) *kaleido.Membership {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "create", "membership", "--name", "test-member", "--consortium", consortiumID))

	var membership kaleido.Membership
	if err := json.Unmarshal(stdout.Bytes(), &membership); err != nil {
		t.Fatal(err)
	}

	return &membership
}

func getMembership(t *testing.T, consortiumID string, id string) *kaleido.Membership {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "get", "membership", "--id", id, "--consortium", consortiumID))

	var membership kaleido.Membership
	if err := json.Unmarshal(stdout.Bytes(), &membership); err != nil {
		t.Fatal(err)
	}

	return &membership
}

func listMembership(t *testing.T, consortiumID string) *[]kaleido.Membership {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "list", "membership", "--consortium", consortiumID))

	var memberships []kaleido.Membership
	if err := json.Unmarshal(stdout.Bytes(), &memberships); err != nil {
		t.Fatal(err)
	}

	return &memberships
}

func deleteMembership(t *testing.T, consortiumID string, id string) {
	InvokeCLISuccess(t, exec.Command("../kld", "delete", "membership", "--id", id, "--consortium", consortiumID))
}

func TestMembership_Create(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	membership := createMembership(t, consortium.ID)
	defer deleteMembership(t, consortium.ID, membership.ID)
}

func TestMembership_Get(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	membership := createMembership(t, consortium.ID)
	defer deleteMembership(t, consortium.ID, membership.ID)

	membershipSaved := getMembership(t, consortium.ID, membership.ID)
	st.Expect(t, membership.ID, membershipSaved.ID)
}

func TestMembership_List(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	origList := listMembership(t, consortium.ID)
	st.Expect(t, len(*origList), 1)

	membership := createMembership(t, consortium.ID)
	defer deleteMembership(t, consortium.ID, membership.ID)

	newList := listMembership(t, consortium.ID)
	st.Expect(t, len(*newList), 2)

}

func TestMembership_Delete(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	membership := createMembership(t, consortium.ID)
	defer deleteMembership(t, consortium.ID, membership.ID)
}

func TestMembership_Default(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	memberships := listMembership(t, consortium.ID)
	st.Expect(t, len(*memberships), 1)
	st.Expect(t, (*memberships)[0].OrgName, "Default Organization")

}
