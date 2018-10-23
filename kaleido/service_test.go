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
	"os"
	"testing"
	"time"
)

func TestServiceCreation(t *testing.T) {
	serviceType := "idregistry"
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	consortium := NewConsortium("serviceTestConsortium", "service creation", "single-org")
	res, err := client.CreateConsortium(&consortium)
	t.Logf("%v", consortium)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Could not create consortium status code: %d.", res.StatusCode())
	}
	defer client.DeleteConsortium(consortium.Id)
	env := NewEnvironment("serviceCreate", "just create some services", "quorum", "raft")

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

	node := NewNode("testNode", members[0].Id)
	res, err = client.CreateNode(consortium.Id, env.Id, &node)
	if err != nil {
		t.Fatal(err)
	}

	for node.State != "started" {
		res, err = client.GetNode(consortium.Id, env.Id, node.Id, &node)
		if err != nil {
			t.Fatal(err)
			break
		}
		time.Sleep(5 * time.Second)
		t.Logf("Node %s state: %s", node.Id, node.State)
	}

	var services []Service
	res, err = client.ListServices(consortium.Id, env.Id, &services)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Could not list services: %d", res.StatusCode())
	}

	t.Logf("Services: %v", services)

	service := NewService("testService", serviceType, members[0].Id)

	res, err = client.CreateService(consortium.Id, env.Id, &service)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Creating service failed status code: %d", res.StatusCode())
	}

	var fetchedService Service
	res, err = client.GetService(consortium.Id, env.Id, service.Id, &fetchedService)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Failed service fetch status code: %d", res.StatusCode())
	}

	if service.Id != fetchedService.Id {
		t.Fatalf("Fetched service id %s does not match %s.", fetchedService.Id, service.Id)
	}

	for fetchedService.State != "started" {
		time.Sleep(time.Second)
		res, err = client.GetService(consortium.Id, env.Id, service.Id, &fetchedService)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode() != 200 {
			t.Fatalf("Failed service fetch status code: %d", res.StatusCode())
		}
		t.Logf("Service state is not started: %v", fetchedService)
	}

	if fetchedService.Urls["http"] == "" {
		t.Fatalf("Fetched service id %s missing 'http' was '%s'.", fetchedService.Id, fetchedService.Urls)
	}

	if service.Service != serviceType {
		t.Fatalf("Fetched service %s has wrong service type %s", service.Id, serviceType)
	}

	if service.State == "" {
		t.Fatalf("Fetched service %s should have a state.", service.Id)
	}

	services = nil
	res, err = client.ListServices(consortium.Id, env.Id, &services)

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

}
