package integration

import (
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/kaleido-io/kaleido-sdk-go/kaleido"
	"github.com/nbio/st"
)

func deployService(t *testing.T, consortium, environment, membership, serviceUniqueName string) *kaleido.Service {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "create", "service", "--name", "test-service", "-c", consortium, "-e", environment, "-m", membership, "-s", serviceUniqueName))

	var service kaleido.Service
	if err := json.Unmarshal(stdout.Bytes(), &service); err != nil {
		t.Fatal(err)
	}

	return &service
}

func listServices(t *testing.T, consortium, environment string) *[]kaleido.Service {
	stdout, _ := InvokeCLISuccess(t, exec.Command("../kld", "list", "services", "-c", consortium, "-e", environment))

	var services []kaleido.Service
	if err := json.Unmarshal(stdout.Bytes(), &services); err != nil {
		t.Fatal(err)
	}

	return &services
}

func TestService_Deploy(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	environment := createEnvironment(t, consortium.ID)
	defer deleteEnvironment(t, consortium.ID, environment.ID)

	memberships := listMembership(t, consortium.ID)
	membershipID := (*memberships)[0].ID

	createNode(t, consortium.ID, environment.ID, membershipID)
	// defer deleteNode, skipped to allow environment deletion to take care of the node

	idRegistry := deployService(t, consortium.ID, environment.ID, membershipID, "idregistry")
	hdWallet := deployService(t, consortium.ID, environment.ID, membershipID, "hdwallet")

	st.Expect(t, idRegistry.Service, "idregistry")
	st.Expect(t, hdWallet.Service, "hdwallet")
}

func TestService_List(t *testing.T) {
	t.Parallel()

	consortium := createConsortium(t)
	defer deleteConsortium(t, consortium.ID)

	environment := createEnvironment(t, consortium.ID)
	defer deleteEnvironment(t, consortium.ID, environment.ID)

	memberships := listMembership(t, consortium.ID)
	membershipID := (*memberships)[0].ID

	createNode(t, consortium.ID, environment.ID, membershipID)
	// defer deleteNode, skipped to allow environment deletion to take care of the node

	servicesOrig := listServices(t, consortium.ID, environment.ID)
	st.Expect(t, len(*servicesOrig), 0)

	deployService(t, consortium.ID, environment.ID, membershipID, "idregistry")
	deployService(t, consortium.ID, environment.ID, membershipID, "hdwallet")

	services := listServices(t, consortium.ID, environment.ID)
	st.Expect(t, len(*services), 2)
}
