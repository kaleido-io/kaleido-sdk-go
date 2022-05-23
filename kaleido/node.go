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
	Name                 string                 `json:"name,omitempty"`
	MembershipID         string                 `json:"membership_id,omitempty"`
	ZoneID               string                 `json:"zone_id,omitempty"`
	ID                   string                 `json:"_id,omitempty"`
	State                string                 `json:"state,omitempty"`
	Role                 string                 `json:"role,omitempty"`
	Provider             string                 `json:"provider,omitempty"`
	ConsensusType        string                 `json:"consensus_type,omitempty"`
	Size                 string                 `json:"size,omitempty"`
	Urls                 map[string]interface{} `json:"urls,omitempty"`
	FirstUserAccount     string                 `json:"first_user_account,omitempty"`
	OpsmetricID          string                 `json:"opsmetric_id,omitempty"`
	NetworkingID         string                 `json:"networking_id,omitempty"`
	KmsID                string                 `json:"kms_id,omitempty"`
	BackupID             string                 `json:"backup_id,omitempty"`
	NodeConfigID         string                 `json:"node_config_id,omitempty"`
	BafID                string                 `json:"baf_id,omitempty"`
	HybridPortAllocation int64                  `json:"hybrid_port_allocation,omitempty"`
	NodeIdentity         string                 `json:"node_identity_data,omitempty"`
	DatabaseType         string                 `json:"database_type,omitempty"`
}

func NewNode(name, membershipID, ezoneID string) Node {
	return Node{
		Name:          name,
		MembershipID:  membershipID,
		ID:            "",
		State:         "",
		Provider:      "",
		ConsensusType: "",
		Size:          "",
		ZoneID:        ezoneID,
	}
}

func (c *KaleidoClient) CreateNode(consortium, envID string, node *Node) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath, consortium, envID)
	return c.Client.R().SetResult(node).SetBody(node).Post(path)
}

func (c *KaleidoClient) UpdateNode(consortium, envID, nodeID string, node *Node) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath+"/%s", consortium, envID, nodeID)
	return c.Client.R().SetResult(node).SetBody(node).Patch(path)
}

func (c *KaleidoClient) ResetNode(consortium, envID, nodeID string) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath+"/%s/reset", consortium, envID, nodeID)
	return c.Client.R().SetBody(map[string]string{}).Put(path)
}

func (c *KaleidoClient) DeleteNode(consortium, envID, nodeID string) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath+"/%s", consortium, envID, nodeID)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListNodes(consortium, envID string, resultBox *[]Node) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath, consortium, envID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetNode(consortiumID, envID, nodeID string, resultBox *Node) (*resty.Response, error) {
	path := fmt.Sprintf(nodeBasePath+"/%s", consortiumID, envID, nodeID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
