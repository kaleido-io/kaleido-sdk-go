// Copyright 2021 Kaleido

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

	"github.com/go-resty/resty/v2"
)

const (
	channelBasePath = "/consortia/%s/environments/%s/channels"
)

type Contract struct {
	Label        string `json:"label,omitempty"`
	Sequence     string `json:"sequence,omitempty"`
	InitRequired bool   `json:"init_required,omitempty"`
	ContractId   string `json:"contract_id,omitempty"`
}

type Channel struct {
	ID                string              `json:"_id,omitempty"`
	Name              string              `json:"name"`
	Description       string              `json:"description,omitempty"`
	MembershipID      string              `json:"membership_id,omitempty"`
	Members           []string            `json:"members"`
	State             string              `json:"state,omitempty"`
	ChannelMapVersion uint                `json:"channel_map_version,omitempty"`
	Contracts         map[string]Contract `json:"contracts,omitempty"`
}

func NewChannel(name, membershipID string, members []string) Channel {
	return Channel{
		Name:         name,
		MembershipID: membershipID,
		Members:      members,
	}
}

func (c *KaleidoClient) CreateChannel(consortium, envID string, channel *Channel) (*resty.Response, error) {
	path := fmt.Sprintf(channelBasePath, consortium, envID)
	return c.Client.R().SetBody(channel).SetResult(channel).Post(path)
}

func (c *KaleidoClient) UpdateChannel(consortium, envID, id string, channel *Channel) (*resty.Response, error) {
	path := fmt.Sprintf(channelBasePath+"/%s", consortium, envID, id)
	return c.Client.R().SetBody(channel).SetResult(channel).Patch(path)
}

func (c *KaleidoClient) ListChannel(consortium, envID string, resultBox *[]Channel) (*resty.Response, error) {
	path := fmt.Sprintf(channelBasePath, consortium, envID)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) GetChannel(consortium, envID, id string, resultBox *Channel) (*resty.Response, error) {
	path := fmt.Sprintf(channelBasePath+"/%s", consortium, envID, id)
	return c.Client.R().SetResult(resultBox).Get(path)
}
