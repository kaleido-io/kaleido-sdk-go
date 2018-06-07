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

	resty "gopkg.in/resty.v1"
)

const (
	nodeBasePath = "/consortia/%s/environments/%s/nodes"
)

type Node struct {
	Name          string `json:"name"`
	MembershipId  string `json:"membership_id"`
	Id            string `json:"_id,omitempty"`
	State         string `json:"state,omitempty"`
	Role          string `json:"role,omitempty"`
	Provider      string `json:"provider,omitempty"`
	ConsensusType string `json:"consensus_type,omitempty"`
	Urls          struct {
		RPC string `json:"rpc,omitempty"`
		WSS string `json:"wss, omitempty"`
	} `json:"urls,omitempty"`
}

func NewNode(name, membershipId string) Node {
	return Node{
		Name:          name,
		MembershipId:  membershipId,
		Id:            "",
		State:         "",
		Provider:      "",
		ConsensusType: "",
	}
}

func (c *KaleidoClient) CreateNode(consortium, envId string, node *Node) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath, consortium, envId)
	return c.Client.R().SetResult(node).SetBody(node).Post(path)
}

func (c *KaleidoClient) DeleteNode(consortium, envId, nodeId string) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath+"/%s", consortium, envId, nodeId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListNodes(consortium, envId string, resultBox *[]Node) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath, consortium, envId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetNode(consortiumId, envId, nodeId string, resultBox *Node) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath+"/%s", consortiumId, envId, nodeId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
