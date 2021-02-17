// Copyright 2018 Kaleido, a ConsenSys business

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kaleido

import (
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"
)

func TestServiceCreation(t *testing.T) {

	gock.New("http://example.com").
		Post("/api/v1/consortia").
		MatchType("json").
		JSON(map[string]interface{}{
			"name":        "serviceTestConsortium",
			"description": "service creation",
		}).
		Reply(201).
		JSON(Consortium{
			ID:          "cons1",
			Name:        "serviceTestConsortium",
			Description: "service creation",
		})

	gock.New("http://example.com").
		Post("/api/v1/consortia/cons1/environments").
		MatchType("json").
		JSON(map[string]interface{}{
			"name":           "serviceCreate",
			"description":    "just create some services",
			"provider":       "quorum",
			"consensus_type": "ibft",
			"block_period":   5,
		}).
		Reply(201).
		JSON(Environment{
			ID:            "env1",
			Name:          "serviceCreate",
			Description:   "just create some services",
			Provider:      "quorum",
			ConsensusType: "ibft",
			BlockPeriod:   5,
		})

	gock.New("http://example.com").
		Get("/api/v1/consortia/cons1/memberships").
		Reply(200).
		JSON([]Membership{
			Membership{
				ID:      "member1",
				OrgName: "Org 1",
			},
		})

	gock.New("http://example.com").
		Post("/api/v1/consortia/cons1/environments/env1/zones").
		MatchType("json").
		JSON(map[string]interface{}{
			"name":   "one zone",
			"region": "us-east-2",
			"cloud":  "aws",
			"type":   "kaleido",
		}).
		Reply(201).
		JSON(EZone{
			ID:     "zone1",
			Name:   "one zone",
			Region: "us-east-2",
			Cloud:  "aws",
			Type:   "kaleido",
		})

	gock.New("http://example.com").
		Post("/api/v1/consortia/cons1/environments/env1/nodes").
		MatchType("json").
		JSON(map[string]interface{}{
			"name":          "testNode",
			"membership_id": "member1",
			"zone_id":       "zone1",
		}).
		Reply(201).
		JSON(Node{
			ID:           "node1",
			Name:         "testNodee",
			MembershipID: "member1",
			ZoneID:       "zone1",
			State:        "initializing",
		})

	gock.New("http://example.com").
		Get("/api/v1/consortia/cons1/environments/env1/nodes/node1").
		Reply(200).
		JSON(Node{
			ID:           "node1",
			Name:         "testNodee",
			MembershipID: "member1",
			ZoneID:       "zone1",
			State:        "started",
		})

	gock.New("http://example.com").
		Get("/api/v1/consortia/cons1/environments/env1/services").
		Reply(200).
		JSON([]Service{})

	gock.New("http://example.com").
		Post("/api/v1/consortia/cons1/environments/env1/services").
		MatchType("json").
		JSON(map[string]interface{}{
			"name":          "testService",
			"service":       "idregistry",
			"zone_id":       "zone1",
			"membership_id": "member1",
		}).
		Reply(201).
		JSON(Service{
			ID:           "svc1",
			Service:      "idregistry",
			Name:         "testNodee",
			MembershipID: "member1",
			ZoneID:       "zone1",
			State:        "provisioning",
		})

	gock.New("http://example.com").
		Get("/api/v1/consortia/cons1/environments/env1/services/svc1").
		Reply(200).
		JSON(Service{
			ID:           "svc1",
			Service:      "idregistry",
			Name:         "testNodee",
			MembershipID: "member1",
			ZoneID:       "zone1",
			State:        "started",
		})

	gock.New("http://example.com").
		Get("/api/v1/consortia/cons1/environments/env1/services").
		Reply(200).
		JSON([]Service{Service{
			ID:           "svc1",
			Service:      "idregistry",
			Name:         "testNodee",
			MembershipID: "member1",
			ZoneID:       "zone1",
			State:        "started",
		}})

	gock.New("http://example.com").
		Delete("/api/v1/consortia/cons1/environments/env1/services/svc1").
		Reply(204)

	gock.New("http://example.com").
		Patch("/api/v1/consortia/cons1/environments/env1/services/svc1").
		Reply(200)

	serviceType := "idregistry"
	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	consortium := NewConsortium("serviceTestConsortium", "service creation")
	res, err := client.CreateConsortium(&consortium)
	t.Logf("%v", consortium)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Could not create consortium status code: %d.", res.StatusCode())
	}
	defer client.DeleteConsortium(consortium.ID)
	env := NewEnvironment("serviceCreate", "just create some services", "quorum", "ibft", false, 5, map[string]string{})

	res, err = client.CreateEnvironment(consortium.ID, &env)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Could not create environment status code: %d, %s", res.StatusCode(), string(res.Body()))
	}

	var members []Membership
	res, err = client.ListMemberships(consortium.ID, &members)

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

	zone := NewEZone("one zone", "us-east-2", "aws")
	res, err = client.CreateEZone(consortium.ID, env.ID, &zone)
	if err != nil {
		t.Fatal(err)
	}

	node := NewNode("testNode", members[0].ID, zone.ID)
	res, err = client.CreateNode(consortium.ID, env.ID, &node)
	if err != nil {
		t.Fatal(err)
	}

	for node.State != "started" {
		res, err = client.GetNode(consortium.ID, env.ID, node.ID, &node)
		if err != nil {
			t.Fatal(err)
			break
		}
		if node.State != "started" {
			time.Sleep(5 * time.Second)
		}
		t.Logf("Node %s state: %s", node.ID, node.State)
	}

	var services []Service
	res, err = client.ListServices(consortium.ID, env.ID, &services)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not list services: %d", res.StatusCode())
	}

	t.Logf("Services: %v", services)

	service := NewService("testService", serviceType, members[0].ID, zone.ID, nil)

	res, err = client.CreateService(consortium.ID, env.ID, &service)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Creating service failed status code: %d", res.StatusCode())
	}

	res, err = client.UpdateService(consortium.ID, env.ID, service.ID, &Service{Name: "New name"})

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Updating service failed status code: %d", res.StatusCode())
	}

	var fetchedService Service
	res, err = client.GetService(consortium.ID, env.ID, service.ID, &fetchedService)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Failed service fetch status code: %d", res.StatusCode())
	}

	if service.ID != fetchedService.ID {
		t.Fatalf("Fetched service id %s does not match %s.", fetchedService.ID, service.ID)
	}

	for fetchedService.State != "started" {
		res, err = client.GetService(consortium.ID, env.ID, service.ID, &fetchedService)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode() != 200 {
			t.Fatalf("Failed service fetch status code: %d", res.StatusCode())
		}
		t.Logf("Service state is not started: %v", fetchedService)
		if fetchedService.State != "started" {
			time.Sleep(5 * time.Second)
		}
	}

	if fetchedService.Urls["http"] == "" {
		t.Fatalf("Fetched service id %s missing 'http' was '%s'.", fetchedService.ID, fetchedService.Urls)
	}

	if service.Service != serviceType {
		t.Fatalf("Fetched service %s has wrong service type %s", service.ID, serviceType)
	}

	if service.State == "" {
		t.Fatalf("Fetched service %s should have a state.", service.ID)
	}

	services = nil
	res, err = client.ListServices(consortium.ID, env.ID, &services)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Could not read services status code: %d", res.StatusCode())
	}

	serviceCount := 1
	if len(services) != serviceCount {
		t.Fatalf("Found unexpected number of services: %d should be %d.", len(services), serviceCount)
	}

	res, err = client.DeleteService(consortium.ID, env.ID, service.ID)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 204 {
		t.Fatalf("Deleting service failed status code: %d", res.StatusCode())
	}

}
