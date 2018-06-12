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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	resty "gopkg.in/resty.v1"
)

type APIResponse struct {
	responseBody string
	responseCode int
	APIPath      string
}

func TestConsortuim_ReadConfiguration(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "kld")
	defer os.Remove(file.Name())
	consortia := "---\n" +
		"provider: quorum\n" +
		"consensus: raft\n" +
		"waitok: true\n" +
		"mode: single-org\n" +
		"consortium:\n" +
		" name: test-consortium\n" +
		" description: A description\n" +
		" environments:\n" +
		" - name: dev\n" +
		"   description: environment description\n" +
		"   members:\n" +
		"   - name: org1\n" +
		"     nodes:\n" +
		"     - name: node1-org1\n" +
		"   - name: org2\n" +
		"     nodes:\n" +
		"     - name: node1-org2\n" +
		"     - name: node2-org2"

	file.WriteString(consortia)

	c := readProjectConfig(file.Name())

	if c.Consensus != "raft" {
		t.Errorf("Consensus algorithm should be Raft")
	}
	if c.Provider != "quorum" {
		t.Errorf("Blockchain should be quorum")
	}
	if c.Consortium.Name != "test-consortium" {
		t.Errorf("Consortium name should be test-consortuim")
	}
	if c.Consortium.Description != "A description" {
		t.Errorf("Consortium description should be \"A description\", but was %v", c.Consortium.Description)
	}
	if len(c.Consortium.Environments) != 1 {
		t.Errorf("Consortium should have 1 environment")
	}
	if len(c.Consortium.Environments[0].Members) != 2 {
		t.Errorf("Environment should have 2 members")
	}
	if len(c.Consortium.Environments[0].Members[0].Nodes) != 1 {
		t.Errorf("Member1 should have 1 node")
	}
	if len(c.Consortium.Environments[0].Members[1].Nodes) != 2 {
		t.Errorf("Member2 should have 2 nodes")
	}

}

func TestConsortium_CreateConsortiumEnvironmentsMembersAndNodes(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "kld")
	defer os.Remove(file.Name())
	if err != nil {
		t.Errorf("Error writing temporary file for consortia definition")
	}
	result := `{"consortium_id":"u0m7e3itin","environments":[{"id":"u0xuyxonsj","members":[{"appcreds":{"id":"","password":` +
		`"uaL7VZi0j4VYvxEhzNr-A2T5QdaPFt8YltEU_RgB_ac","username":"u0k719wsku"},"id":"u0lxv8poisdp","name":"org1","nodes":` +
		`[{"block_height":0,"geth":{"public_address":"","validators":null},"quorum":{"private_address":"8Y++vcn+pzEboXQ4HgwrHh7+gMfYtLtwyn8E60kaLx4=",` +
		`"public_address":"0x44a7638e0d3ff8fe99e78f1c4f77928bf36978be"},"id":"7ebf87398f0d2ed155babf260c54e698aefb13a9c7d7472739d42f99984a0fe500b275100fd238e28e8465da876ad9312b50337c685f7bd3e765ca8e94619a36",` +
		`"urls":{"rpc":"https://e0czbf6cxq-e0l6kgey7g-rpc.eu-central-1.kaleido.io","wss":"wss://e0czbf6cxq-e0l6kgey7g-wss.eu-central-1.kaleido.io"},` +
		`"user_accounts":["0xD68762958aaA7F3dEaE2d963aDcbF2db7c670aE0"]}]},{"appcreds":{"id":"","password":"uaL7VZi0j4VYvxEhzNr-A2T5QdaPFt8YltEU_RgB_ac",` +
		`"username":"u0k719wsku"},"id":"u0lxv8hppe","name":"org2","nodes":[{"block_height":0,"geth":{"public_address":"",` +
		`"validators":null},"quorum":{"private_address":"cn+pzEboXQ4twyn8E60kaLx4=8Y++vHgwrHh7+gMfYtL","public_address":` +
		`"0x6978be44a7c4f77928bf3638e0d3ff8fe99e78f1"},"id":"fe500b275100fd238e28e8465da876ad93129c7d7472739d42f99984a0b50337c685f7bd3e765ca8e94619a367ebf87398f0d2ed155babf260c54e698aefb13a",` +
		`"urls":{"rpc":"https://e0l6kgey7g-cxqe0czbf6-rpc.eu-central-1.kaleido.io","wss":"wss://e0l6kgey7g-bf6cxqe0cz-wss.eu-central-1.kaleido.io"},` +
		`"user_accounts":["8aaA7F3dEaE2bF2db7c670aE00xD6876295d963aDc"]},{"block_height":0,"geth":{"public_address":"","validators":null},"quorum":` +
		`{"private_address":"cn+pzEboXQ4twyn8E60kaLx4=8Y++vHgwrHh7+gMfYtL","public_address":"0x6978be44a7c4f77928bf3638e0d3ff8fe99e78f1"},` +
		`"id":"fe500b275100fd238e28e8465da876ad93129c7d7472739d42f99984a0b50337c685f7bd3e765ca8e94619a367ebf87398f0d2ed155babf260c54e698aefb13a",` +
		`"urls":{"rpc":"https://e0l6kgey7g-cxqe0czbf6-rpc.eu-central-1.kaleido.io","wss":"wss://e0l6kgey7g-bf6cxqe0cz-wss.eu-central-1.kaleido.io"},` +
		`"user_accounts":["8aaA7F3dEaE2bF2db7c670aE00xD6876295d963aDc"]}]}]}]}`

	consortia := "---\n" +
		"provider: quorum\n" +
		"consensus: raft\n" +
		"waitok: true\n" +
		"mode: single-org\n" +
		"description: A description\n" +
		"consortium:\n" +
		" name: test-consortium\n" +
		" environments:\n" +
		" - name: dev\n" +
		"   description: environment description\n" +
		"   members:\n" +
		"   - name: org1\n" +
		"     nodes:\n" +
		"     - name: node1-org1\n" +
		"   - name: org2\n" +
		"     nodes:\n" +
		"     - name: node1-org2\n" +
		"     - name: node2-org2"

	file.WriteString(consortia)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	APIKey := "123"
	fakeAPIURL := "https://fake-console.kaleido.io/api/v1"
	consoritiumID := "u0m7e3itin"
	environmentID := "u0xuyxonsj"
	member1ID := "u0lxv8hppe"
	member2ID := "u0lxv8poisdp"
	node11ID := "u0awuc2pt1"
	node21ID := "u0awuc2pt3"
	node22ID := "u0awuc2pt4"

	createConsortiumIsCalled := false
	createEnvironmentIsCalled := false
	createMember1IsCalled := false
	createMember2IsCalled := false
	createNode11IsCalled := false
	createNode21IsCalled := false
	createNode22IsCalled := false
	createAppcredsIsCalled := false
	getNode11IsCalled := false
	getNode21IsCalled := false
	getNode22IsCalled := false

	cases := []struct {
		name                           string
		consortiumResponse             APIResponse
		environmentResponse            APIResponse
		membership1Response            APIResponse
		membership2Response            APIResponse
		node11Response                 APIResponse
		node12Response                 APIResponse
		node22Response                 APIResponse
		appCredsResponse               APIResponse
		node11IsIniitalizedGetResponse APIResponse
		node21IsIniitalizedGetResponse APIResponse
		node22IsIniitalizedGetResponse APIResponse
		node11GetStatusResponse        APIResponse
		node21GetStatusResponse        APIResponse
		node22GetStatusResponse        APIResponse
	}{
		{
			"create_consortium_wait_nodes_initialized",
			APIResponse{
				fmt.Sprintf(`{"name":"testConsortium234","description":"this is a test consortium",
					"mode":"single-org","owner":"e0lmvhq9zy","_id":"%v","state":"setup",
					"_revision":"0","created_at":"2018-05-31T08:31:10.065Z"}`, consoritiumID),
				201,
				"consortia",
			},
			APIResponse{
				fmt.Sprintf(`{"name":"dev","description":"","provider":"quorum","consensus_type":"raft",
					"autopause_init_delay":168,"autopause_idle_hours":24,"_id":"%v",
					"state":"initializing","enable_tether":false,"block_period":0,
					"chain_id":353546420,"node_list":[],"region":"us-east","release_id":"u0kugmj500",
					"_revision":"0","created_at":"2018-05-31T08:39:41.550Z"}`, environmentID),
				201,
				fmt.Sprintf("consortia/%v/environments", consoritiumID),
			},
			APIResponse{
				fmt.Sprintf(`{"org_name":"ame","org_id":"e0lmvhq9zy","state":"active",
					"_id":"%v","minimum_nodes":1,"maximum_nodes":1,"_revision":"0",
					"created_at":"2018-05-31T08:41:40.375Z"}`, member1ID),
				201,
				fmt.Sprintf("consortia/%v/memberships", consoritiumID),
			},
			APIResponse{
				fmt.Sprintf(`{"org_name":"ame","org_id":"e0lmvhq9zy","state":"active",
					"_id":"%v","minimum_nodes":1,"maximum_nodes":1,"_revision":"0",
					"created_at":"2018-05-31T08:41:40.375Z"}`, member2ID),
				201,
				fmt.Sprintf("consortia/%v/memberships", consoritiumID),
			},
			APIResponse{
				fmt.Sprintf(`{"name":"org1-node1","membership_id":"u0lxv8hppe",
					"urls":{"wss":""},"role":"validator","state":"initializing",
					"provider":"quorum","consensus_type":"raft","_id":"%v",
					"_revision":"0","created_at":"2018-05-31T08:44:47.846Z"}`, node11ID),
				201,
				fmt.Sprintf("consortia/%v/environments/%v/nodes", consoritiumID, environmentID),
			},
			APIResponse{
				fmt.Sprintf(`{"name":"org2-node1","membership_id":"u0lxv8hppe",
					"urls":{"wss":""},"role":"validator","state":"initializing",
					"provider":"quorum","consensus_type":"raft","_id":"%v",
					"_revision":"0","created_at":"2018-05-31T08:44:47.846Z"}`, node21ID),
				201,
				fmt.Sprintf("consortia/%v/environments/%v/nodes", consoritiumID, environmentID),
			},
			APIResponse{
				fmt.Sprintf(`{"name":"org2-node2","membership_id":"u0lxv8hppe",
					"urls":{"wss":""},"role":"validator","state":"initializing",
					"provider":"quorum","consensus_type":"raft","_id":"%v",
					"_revision":"0","created_at":"2018-05-31T08:44:47.846Z"}`, node22ID),
				201,
				fmt.Sprintf("consortia/%v/environments/%v/nodes", consoritiumID, environmentID),
			},
			APIResponse{
				`{"membership_id":"u0ypwvdcan","auth_type":"basic_auth","_id":"u0k719wsku",
					"_revision":"0","username":"u0k719wsku",
					"password":"uaL7VZi0j4VYvxEhzNr-A2T5QdaPFt8YltEU_RgB_ac"}`,
				201,
				fmt.Sprintf("consortia/%v/environments/%v/appcreds", consoritiumID, environmentID),
			},
			APIResponse{
				fmt.Sprintf(`{"name":"org1-node1","membership_id":"u0lxv8hppe",
					"urls":{"wss":""},"role":"validator","state":"ready",
					"provider":"quorum","consensus_type":"raft","_id":"%v",
					"_revision":"0","created_at":"2018-05-31T08:44:47.846Z"}`, node11ID),
				200,
				fmt.Sprintf("consortia/%v/environments/%v/nodes/%v", consoritiumID, environmentID, node11ID),
			},
			APIResponse{
				fmt.Sprintf(`{"name":"org2-node1","membership_id":"8hppeu0lxv",
					"urls":{"wss":""},"role":"validator","state":"ready",
					"provider":"quorum","consensus_type":"raft","_id":"%v",
					"_revision":"0","created_at":"2018-05-31T08:44:47.846Z"}`, node21ID),
				200,
				fmt.Sprintf("consortia/%v/environments/%v/nodes/%v", consoritiumID, environmentID, node21ID),
			},
			APIResponse{
				fmt.Sprintf(`{"name":"org2-node2","membership_id":"8hppeu0lxv",
					"urls":{"wss":""},"role":"validator","state":"ready",
					"provider":"quorum","consensus_type":"raft","_id":"%v",
					"_revision":"0","created_at":"2018-05-31T08:44:47.846Z"}`, node22ID),
				200,
				fmt.Sprintf("consortia/%v/environments/%v/nodes/%v", consoritiumID, environmentID, node22ID),
			},
			APIResponse{
				`{
					"id": "7ebf87398f0d2ed155babf260c54e698aefb13a9c7d7472739d42f99984a0fe500b275100fd238e28e8465da876ad9312b50337c685f7bd3e765ca8e94619a36",
					"quorum": {
						"private_address": "8Y++vcn+pzEboXQ4HgwrHh7+gMfYtLtwyn8E60kaLx4=",
						"public_address": "0x44a7638e0d3ff8fe99e78f1c4f77928bf36978be"
					},
					"user_accounts": ["0xD68762958aaA7F3dEaE2d963aDcbF2db7c670aE0"],
					"block_height": 0,
					"consensus_identity": "3",
					"urls": {
						"rpc": "https://e0czbf6cxq-e0l6kgey7g-rpc.eu-central-1.kaleido.io",
						"wss": "wss://e0czbf6cxq-e0l6kgey7g-wss.eu-central-1.kaleido.io"
					}
				}`,
				200,
				fmt.Sprintf("consortia/%v/environments/%v/nodes/%v/status", consoritiumID, environmentID, node11ID),
			},
			APIResponse{
				`{
					"id": "9c7d7472739d42f99984a0fe500b275100fd238e28e8465da876ad9312b50337c685f7bd3e765ca8e94619a367ebf87398f0d2ed155babf260c54e698aefb13a",
					"quorum": {
						"private_address": "twyn8E60kaLx4=8Y++vcn+pzEboXQ4HgwrHh7+gMfYtL",
						"public_address": "0xc4f77928bf36978be44a7638e0d3ff8fe99e78f1"
					},
					"user_accounts": ["bF2db7c670aE00xD68762958aaA7F3dEaE2d963aDc"],
					"block_height": 0,
					"consensus_identity": "3",
					"urls": {
						"rpc": "https://e0l6kgey7g-e0czbf6cxq-rpc.eu-central-1.kaleido.io",
						"wss": "wss://e0l6kgey7g-e0czbf6cxq-wss.eu-central-1.kaleido.io"
					}
				}`,
				200,
				fmt.Sprintf("consortia/%v/environments/%v/nodes/%v/status", consoritiumID, environmentID, node21ID),
			},
			APIResponse{
				`{
					"id": "fe500b275100fd238e28e8465da876ad93129c7d7472739d42f99984a0b50337c685f7bd3e765ca8e94619a367ebf87398f0d2ed155babf260c54e698aefb13a",
					"quorum": {
						"private_address": "cn+pzEboXQ4twyn8E60kaLx4=8Y++vHgwrHh7+gMfYtL",
						"public_address": "0x6978be44a7c4f77928bf3638e0d3ff8fe99e78f1"
					},
					"user_accounts": ["8aaA7F3dEaE2bF2db7c670aE00xD6876295d963aDc"],
					"block_height": 0,
					"consensus_identity": "3",
					"urls": {
						"rpc": "https://e0l6kgey7g-cxqe0czbf6-rpc.eu-central-1.kaleido.io",
						"wss": "wss://e0l6kgey7g-bf6cxqe0cz-wss.eu-central-1.kaleido.io"
					}
				}`,
				200,
				fmt.Sprintf("consortia/%v/environments/%v/nodes/%v/status", consoritiumID, environmentID, node22ID),
			},
		},
		{
			"create_consortium_fails",
			APIResponse{
				fmt.Sprintf(`{"name":"testConsortium234","description":"this is a test consortium",
					"mode":"single-org","owner":"e0lmvhq9zy","_id":"%v","state":"setup",
					"_revision":"0","created_at":"2018-05-31T08:31:10.065Z"}`, consoritiumID),
				404,
				"consortia",
			},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
			APIResponse{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.Reset()
			fakeConsortiumAPIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.consortiumResponse.APIPath)
			httpmock.RegisterResponder("POST", fakeConsortiumAPIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.consortiumResponse.responseCode, tc.consortiumResponse.responseBody)
				createConsortiumIsCalled = true
				return resp, nil
			},
			)

			fakeEnvironmentAPIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.environmentResponse.APIPath)
			httpmock.RegisterResponder("POST", fakeEnvironmentAPIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.environmentResponse.responseCode, tc.environmentResponse.responseBody)
				createEnvironmentIsCalled = true
				return resp, nil
			},
			)

			var requestMembershipData []byte
			fakeMembershipAPIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.membership1Response.APIPath)
			httpmock.RegisterResponder("POST", fakeMembershipAPIURL, func(req *http.Request) (*http.Response, error) {
				requestMembershipData, _ = ioutil.ReadAll(req.Body)
				if string(requestMembershipData) != `{"org_name":"org1"}` {
					resp := httpmock.NewStringResponse(tc.membership1Response.responseCode, tc.membership1Response.responseBody)
					createMember1IsCalled = true
					return resp, nil
				} else {
					resp := httpmock.NewStringResponse(tc.membership1Response.responseCode, tc.membership2Response.responseBody)
					createMember2IsCalled = true
					return resp, nil
				}
			},
			)

			fakeNodesAPIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.node11Response.APIPath)
			httpmock.RegisterResponder("POST", fakeNodesAPIURL, func(req *http.Request) (*http.Response, error) {
				requestNodeData, _ := ioutil.ReadAll(req.Body)
				if strings.Contains(string(requestNodeData), "node1-org1") {
					createNode11IsCalled = true
					resp := httpmock.NewStringResponse(tc.node11Response.responseCode, tc.node11Response.responseBody)
					return resp, nil
				} else if strings.Contains(string(requestNodeData), "node1-org2") {
					resp := httpmock.NewStringResponse(tc.node12Response.responseCode, tc.node12Response.responseBody)
					createNode21IsCalled = true
					return resp, nil
				} else {
					resp := httpmock.NewStringResponse(tc.node22Response.responseCode, tc.node22Response.responseBody)
					createNode22IsCalled = true
					return resp, nil
				}
			},
			)

			fakeAppcredsAPIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.appCredsResponse.APIPath)
			httpmock.RegisterResponder("POST", fakeAppcredsAPIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.appCredsResponse.responseCode, tc.appCredsResponse.responseBody)
				createAppcredsIsCalled = true
				return resp, nil
			},
			)

			fakeGetNode11APIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.node11IsIniitalizedGetResponse.APIPath)
			httpmock.RegisterResponder("GET", fakeGetNode11APIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.node11IsIniitalizedGetResponse.responseCode, tc.node11IsIniitalizedGetResponse.responseBody)
				getNode11IsCalled = true
				return resp, nil
			},
			)

			fakeGetNode21APIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.node21IsIniitalizedGetResponse.APIPath)
			httpmock.RegisterResponder("GET", fakeGetNode21APIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.node21IsIniitalizedGetResponse.responseCode, tc.node21IsIniitalizedGetResponse.responseBody)
				getNode21IsCalled = true
				return resp, nil
			},
			)

			fakeGetNode22APIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.node22IsIniitalizedGetResponse.APIPath)
			httpmock.RegisterResponder("GET", fakeGetNode22APIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.node21IsIniitalizedGetResponse.responseCode, tc.node21IsIniitalizedGetResponse.responseBody)
				getNode22IsCalled = true
				return resp, nil
			},
			)

			fakeGetNodeStatus11APIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.node11GetStatusResponse.APIPath)
			httpmock.RegisterResponder("GET", fakeGetNodeStatus11APIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.node11GetStatusResponse.responseCode, tc.node11GetStatusResponse.responseBody)
				getNode11IsCalled = true
				return resp, nil
			},
			)
			fakeGetNodeStatus21APIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.node21GetStatusResponse.APIPath)
			httpmock.RegisterResponder("GET", fakeGetNodeStatus21APIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.node21GetStatusResponse.responseCode, tc.node22GetStatusResponse.responseBody)
				getNode11IsCalled = true
				return resp, nil
			},
			)

			fakeGetNodeStatus22APIURL := fmt.Sprintf("%v/%v", fakeAPIURL, tc.node22GetStatusResponse.APIPath)
			httpmock.RegisterResponder("GET", fakeGetNodeStatus22APIURL, func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(tc.node22GetStatusResponse.responseCode, tc.node22GetStatusResponse.responseBody)
				getNode11IsCalled = true
				return resp, nil
			},
			)

			if tc.name == "create_consortium_wait_nodes_initialized" {
				r := resty.New().SetHostURL(fakeAPIURL).SetAuthToken(APIKey)
				c := KaleidoClient{r}
				consortiumCreated := c.CreateConsortiumEnvironmentsMembersAndNodes(file.Name(), 1)
				if consortiumCreated != result {
					t.Errorf("Failed when creating the consortium")
				}
			} else {
				if os.Getenv("EXIT_TEST") == "1" {
					r := resty.New().SetHostURL(fakeAPIURL).SetAuthToken(APIKey)
					c := KaleidoClient{r}
					c.CreateConsortiumEnvironmentsMembersAndNodes(file.Name(), 1)
				}
				cmd := exec.Command(os.Args[0], "-test.run=TestConsortium_CreateConsortium")
				cmd.Env = append(os.Environ(), "EXIT_TEST=1")
				out, err := cmd.Output()
				msg := strings.TrimSpace(string(out))
				if e, ok := err.(*exec.ExitError); ok && !e.Success() {
					if !strings.Contains(msg, "Could not create consortium. Status code: 404.") {
						t.Errorf("Expected to fail with error %q, but got %q",
							"Could not create consortium. Status code: 404.", msg)
					}
					return
				}
			}

			if tc.name == "create_consortium_wait_nodes_initialized" {
				if !createConsortiumIsCalled {
					t.Errorf("CreateConsortium should have been called")
				}

				if !createEnvironmentIsCalled {
					t.Errorf("CreateEnvironment should have been called")
				}
				if !createMember1IsCalled {
					t.Errorf("Member1 should have been created")
				}
				if !createMember2IsCalled {
					t.Errorf("Member2 should have been created")
				}
				if !createNode11IsCalled {
					t.Errorf("Node1 for member1 should have been created")
				}
				if !createNode21IsCalled {
					t.Errorf("Node1 for member2 should have been created")
				}
				if !createNode22IsCalled {
					t.Errorf("Node2 for member2 should have been created")
				}
				if !createAppcredsIsCalled {
					t.Errorf("CreateAppCreds should have been called")
				}
				if !getNode11IsCalled {
					t.Errorf("Get Node1 information for Memmber org1 should have been called ")
				}
				if !getNode21IsCalled {
					t.Errorf("Get Node1 information for Memmber org2 should have been called ")
				}
				if !getNode22IsCalled {
					t.Errorf("Get Node2 information for Memmber org2 should have been called ")
				}
			} else if tc.name == "create_consortium_fails" {
				if !createConsortiumIsCalled {
					t.Errorf("CreateConsortium should have been called")
				}

				if createEnvironmentIsCalled {
					t.Errorf("CreateEnvironment should have not been called")
				}
				if createMember1IsCalled {
					t.Errorf("Member1 should have not been created")
				}
				if createMember2IsCalled {
					t.Errorf("Member2 should have not been created")
				}
				if createNode11IsCalled {
					t.Errorf("Node1 for member1 should not have been created")
				}
				if createNode21IsCalled {
					t.Errorf("Node1 for member2 should not have been created")
				}
				if createNode22IsCalled {
					t.Errorf("Node2 for member2 should not have been created")
				}
				if createAppcredsIsCalled {
					t.Errorf("CreateAppCreds should not have been called")
				}
			}
		})
	}
}
