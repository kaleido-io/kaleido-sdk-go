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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type ConsortiumDefinition struct {
	Consensus  string `json:"consensus"`
	Consortium struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Mode         string `json:"mode"`
		Environments []struct {
			Members []struct {
				Name  string `json:"name"`
				Nodes []struct {
					Name string `json:"name"`
				} `json:"nodes"`
			} `json:"members"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"environments"`
	} `json:"consortium"`
	Provider string `json:"provider"`
	Waitok   bool   `json:"waitok"`
}

type Identifier struct {
	ID string `json:"_id"`
}

type NodeState struct {
	State string `json:"state"`
}

type NodeStatus struct {
	BlockHeight int `json:"block_height"`
	Geth        struct {
		PublicAddress string   `json:"public_address"`
		Validators    []string `json:"validators"`
	} `json:"geth"`
	Quorum struct {
		PrivateAddress string `json:"private_address"`
		PublicAddress  string `json:"public_address"`
	} `json:"quorum"`
	ID   string `json:"id"`
	Urls struct {
		RPC string `json:"rpc"`
		Wss string `json:"wss"`
	} `json:"urls"`
	UserAccounts []string `json:"user_accounts"`
}

type AppcredsCreated struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Member struct {
	Appcreds AppcredsCreated `json:"appcreds"`
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Nodes    []NodeStatus    `json:"nodes"`
}

type EnvironmentCreated struct {
	ID      string   `json:"id"`
	Members []Member `json:"members"`
}

type ConsortiumOut struct {
	ConsortiumID string               `json:"consortium_id"`
	Environments []EnvironmentCreated `json:"environments"`
}

func (client *KaleidoClient) CreateConsortiumEnvironmentsMembersAndNodes(configPath string, waitForInitialization int) string {
	out := ConsortiumOut{}

	c := readProjectConfig(configPath)
	consortium := NewConsortium(c.Consortium.Name, c.Consortium.Description, c.Consortium.Mode)
	fmt.Printf("Creating consortium: %v ", c.Consortium.Name)
	consortiumRes, err := client.CreateConsortium(&consortium)
	ValidateCreationResponse(consortiumRes, err, "consortium")
	var cID Identifier
	unmarshallID(consortiumRes.Body(), &cID)
	out.ConsortiumID = cID.ID

	var environmentsCreated []EnvironmentCreated
	for _, env := range c.Consortium.Environments {
		fmt.Printf("Creating environment: %v ", env.Name)
		envCreated := EnvironmentCreated{}
		environment := NewEnvironment(env.Name, env.Description, c.Provider, c.Consensus)
		envRes, err := client.CreateEnvironment(cID.ID, &environment)
		ValidateCreationResponse(envRes, err, "environment")
		var eID Identifier
		unmarshallID(envRes.Body(), &eID)
		envCreated.ID = eID.ID

		var members []Member
		for _, m := range env.Members {
			memberCreated := Member{}
			membership := NewMembership(m.Name)
			fmt.Printf("Creating membership: %v ", m.Name)
			membershipRes, err := client.CreateMembership(cID.ID, &membership)
			ValidateCreationResponse(membershipRes, err, "membership")
			var mID Identifier
			unmarshallID(membershipRes.Body(), &mID)
			memberCreated.ID = mID.ID
			memberCreated.Name = m.Name

			nodesCreated := []NodeStatus{}
			for _, n := range m.Nodes {
				fmt.Printf("Creating Node: %v for member %v", n.Name, m.Name)
				node := NewNode(n.Name, mID.ID)
				resNode, err := client.CreateNode(cID.ID, eID.ID, &node)
				ValidateCreationResponse(resNode, err, "node")
				if c.Waitok {
					var nID Identifier
					unmarshallID(resNode.Body(), &nID)
					var node Node
					nodeInfoRes, err := client.GetNode(cID.ID, eID.ID, nID.ID, &node)
					ValidateGetResponse(nodeInfoRes, err, "node")
					var nState NodeState
					unmarshallNodeState(nodeInfoRes.Body(), &nState)
					start := time.Now()
					var elapsed time.Duration
					for initializing := true; initializing; initializing = (nState.State == "initializing" && elapsed < 1800000) {
						t := time.Now()
						elapsed = t.Sub(start)
						fmt.Printf("Elapsed: %v\n", elapsed)
						fmt.Printf("Waiting for node %v to start\n", nID.ID)
						time.Sleep(time.Duration(waitForInitialization) * time.Second)
						fmt.Printf("Getting status for node: %v\n", nID.ID)
						nodeInfoRes, err := client.GetNode(cID.ID, eID.ID, nID.ID, &node)
						ValidateGetResponse(nodeInfoRes, err, "node")
						unmarshallNodeState(nodeInfoRes.Body(), &nState)
					}
					resStatus, err := client.GetNodeStatus(cID.ID, eID.ID, nID.ID, &node)
					ValidateGetResponse(resStatus, err, "node")
					var nodeStatus NodeStatus
					err = json.Unmarshal(resStatus.Body(), &nodeStatus)
					if err != nil {
						msg := fmt.Sprintf("Unmarshal: %v", err)
						exit(msg)
					}
					nodesCreated = append(nodesCreated, nodeStatus)
				}
			}
			memberCreated.Nodes = nodesCreated
			appcreds := NewAppCreds(mID.ID)
			appcredsRes, err := client.CreateAppCreds(cID.ID, eID.ID, &appcreds)
			ValidateCreationResponse(appcredsRes, err, "appcreds")
			var appcredsCreated AppcredsCreated
			err = json.Unmarshal(appcredsRes.Body(), &appcredsCreated)
			if err != nil {
				msg := fmt.Sprintf("Unmarshal: %v", err)
				exit(msg)
			}
			memberCreated.Appcreds = appcredsCreated
			members = append(members, memberCreated)
		}
		envCreated.Members = members
		environmentsCreated = append(environmentsCreated, envCreated)
	}
	out.Environments = environmentsCreated
	consortiumOut, err := json.Marshal(out)
	return string(consortiumOut)
}

func unmarshallID(body []byte, data *Identifier) {
	err := json.Unmarshal(body, &data)
	if err != nil {
		msg := fmt.Sprintf("Unmarshal: %v", err)
		exit(msg)
	}
}

func unmarshallNodeState(body []byte, data *NodeState) {
	err := json.Unmarshal(body, &data)
	if err != nil {
		msg := fmt.Sprintf("Unmarshal: %v", err)
		exit(msg)
	}
}

func readProjectConfig(configPath string) ConsortiumDefinition {
	var c ConsortiumDefinition
	var yamlFile []byte
	var err error
	if len(configPath) != 0 {
		yamlFile, err = ioutil.ReadFile(configPath)
		if err != nil {
			msg := fmt.Sprintf("Error reading YAML config file  #%v ", err)
			exit(msg)
		}
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			msg := fmt.Sprintf("Unmarshal: %v", err)
			exit(msg)
		}
		return c
	}
	return ConsortiumDefinition{}
}

var exit = func(message string) {
	fmt.Println(message)
	os.Exit(1)
}
