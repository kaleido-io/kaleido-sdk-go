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
	"testing"

	"github.com/nbio/st"
	gock "gopkg.in/h2non/gock.v1"
)

var mockChannelCreatePayload = map[string]interface{}{
	"name":          "channel-name",
	"membership_id": "member1",
	"members":       []string{"member1", "member2"},
}

var mockChannel = Channel{
	ID:           "zzstcszriw",
	Name:         "channel-name",
	MembershipID: "member1",
	Members:      []string{"member1", "member2"},
	Contracts: map[string]Contract{
		"chaincode1": Contract{},
	},
}

var mockChannels = []Channel{mockChannel}

func TestChannelCreate(t *testing.T) {
	defer gock.Off()
	gock.Observe(gock.DumpRequest)

	gock.New("http://example.com").
		Post("/api/v1/consortia/c1/environments/env1/channels").
		MatchType("json").
		JSON(mockChannelCreatePayload).
		Reply(201).
		JSON(mockChannel)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	var channel = NewChannel("channel-name", "member1", []string{"member1", "member2"})
	_, err := client.CreateChannel("c1", "env1", &channel)
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestChannelGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/c1/environments/env1/channels/channel1").
		Reply(200).
		JSON(mockChannel)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.GetChannel("c1", "env1", "channel1", &Channel{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestChannelList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/c1/environments/env1/channels").
		Reply(200).
		JSON(mockChannels)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.ListChannel("c1", "env1", &[]Channel{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestChannelUpdate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Patch("/api/v1/consortia/c1/environments/env1/channels/channel1").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.UpdateChannel("c1", "env1", "channel1", &Channel{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
