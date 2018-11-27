package integration

import (
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/nbio/st"

	"github.com/kaleido-io/kaleido-sdk-go/kaleido"
)

func createMembership(t *testing.T, consortiumID string) *kaleido.Membership {
	stdout, _ := InvokeCLISuccess(t, exec.Command("kld", "create", "membership", "--name", "test-member", "--consortium", consortiumID))

	var membership kaleido.Membership
	if err := json.Unmarshal(stdout.Bytes(), &membership); err != nil {
		t.Fatal(err)
	}

	return &membership
}

func getMembership(t *testing.T, consortiumID string, id string) *kaleido.Membership {
	stdout, _ := InvokeCLISuccess(t, exec.Command("kld", "get", "membership", "--id", id, "--consortium", consortiumID))

	var membership kaleido.Membership
	if err := json.Unmarshal(stdout.Bytes(), &membership); err != nil {
		t.Fatal(err)
	}

	return &membership
}

func listMembership(t *testing.T, consortiumID string) *[]kaleido.Membership {
	stdout, _ := InvokeCLISuccess(t, exec.Command("kld", "list", "membership", "--consortium", consortiumID))

	var memberships []kaleido.Membership
	if err := json.Unmarshal(stdout.Bytes(), &memberships); err != nil {
		t.Fatal(err)
	}

	return &memberships
}

func deleteMembership(t *testing.T, consortiumID string, id string) {
	InvokeCLISuccess(t, exec.Command("kld", "delete", "membership", "--id", id, "--consortium", consortiumID))
}

func TestMembership_Create(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	membership := createMembership(t, consortium.Id)
	defer deleteMembership(t, consortium.Id, membership.Id)
}

func TestMembership_Get(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	membership := createMembership(t, consortium.Id)
	defer deleteMembership(t, consortium.Id, membership.Id)

	membershipSaved := getMembership(t, consortium.Id, membership.Id)
	st.Expect(t, membership.Id, membershipSaved.Id)
}

func TestMembership_List(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	origList := listMembership(t, consortium.Id)
	st.Expect(t, len(*origList), 1)

	membership := createMembership(t, consortium.Id)
	defer deleteMembership(t, consortium.Id, membership.Id)

	newList := listMembership(t, consortium.Id)
	st.Expect(t, len(*newList), 2)

}

func TestMembership_Delete(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	membership := createMembership(t, consortium.Id)
	defer deleteMembership(t, consortium.Id, membership.Id)
}

func TestMembership_Default(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.Id)

	memberships := listMembership(t, consortium.Id)
	st.Expect(t, len(*memberships), 1)
	st.Expect(t, (*memberships)[0].OrgName, "Default Organization")

}
