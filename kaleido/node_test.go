package kaleido

import (
	"os"
	"testing"
	"time"
)

func TestNodeCreation(t *testing.T) {
	consensusType := "ibft"
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
	env := NewEnvironment("nodeCreate", "just create some nodes", "quorum", consensusType)

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

	var fetchedNode Node
	res, err = client.GetNode(consortium.Id, env.Id, node.Id, &fetchedNode)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Failed node fetch status code: %d", res.StatusCode())
	}

	if node.Id != fetchedNode.Id {
		t.Fatalf("Fetched node id %s does not match %s.", fetchedNode.Id, node.Id)
	}

	for fetchedNode.State != "started" {
		time.Sleep(time.Second)
		res, err = client.GetNode(consortium.Id, env.Id, node.Id, &fetchedNode)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode() != 200 {
			t.Fatalf("Failed node fetch status code: %d", res.StatusCode())
		}
		t.Logf("Node state is not started: %v", fetchedNode)
	}

	if fetchedNode.Urls.RPC == "" {
		t.Fatalf("Fetched node id %s missing RPC was '%s'.", fetchedNode.Id, fetchedNode.Urls.RPC)
	}

	if fetchedNode.Urls.WSS == "" {
		t.Fatalf("Fetched node id %s missing WSS, was '%s'.", fetchedNode.Id, fetchedNode.Urls.WSS)
	}

	if node.ConsensusType != consensusType {
		t.Fatalf("Fetched node %s has wrong consensusType %s", node.Id, node.ConsensusType)
	}

	if node.State == "" {
		t.Fatalf("Fetched node %s should have a state.", node.Id)
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
