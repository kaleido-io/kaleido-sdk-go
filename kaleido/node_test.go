package kaleido

import (
	"os"
	"testing"
)

func TestNodeCreation(t *testing.T) {
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	consortium := NewConsortium("nodeTestConsortium", "node creation", "single-org")
	res, err := client.CreateConsortium(&consortium)
	t.Logf("%v", consortium)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Could not create consortium status code: %d.", res.StatusCode())
	}
	defer client.DeleteConsortium(consortium.Id)
	env := NewEnvironment("nodeCreate", "just create some nodes", "quorum", "raft")

	res, err = client.CreateEnvironment(consortium.Id, &env)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Could not create environment status code: %d, %s", res.StatusCode(), string(res.Body()))
	}

	var members []Membership
	res, err = client.ListMemberships(consortium.Id, &members)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not list memberships: %d", res.StatusCode())
	}

	if len(members) != 1 {
		t.Fatalf("Environment unexpected had %d members.", len(members))
	}
	t.Logf("%v", members)

	var nodes []Node
	res, err = client.ListNodes(consortium.Id, env.Id, &nodes)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not list nodes: %d", res.StatusCode())
	}

	t.Logf("Nodes: %v", nodes)

	node := NewNode("testNode", members[0].Id)

	res, err = client.CreateNode(consortium.Id, env.Id, &node)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Creating node failed status code: %d", res.StatusCode())
	}

	nodes = nil
	res, err = client.ListNodes(consortium.Id, env.Id, &nodes)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Could not read nodes status code: %d", res.StatusCode())
	}

	nodeCount := 2 //Monitor node and your own node.
	if len(nodes) != nodeCount {
		t.Fatalf("Found unexpected number of nodes: %d should be %d.", len(nodes), nodeCount)
	}

}
