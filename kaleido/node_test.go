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
	"testing"

	"github.com/nbio/st"
	gock "gopkg.in/h2non/gock.v1"
)

var mockNodeCreatePayload = map[string]string{
	"name":          "blah",
	"membership_id": "member1",
}

var mockNode = map[string]string{
	"name":           "blah",
	"membership_id":  "member1",
	"role":           "validator",
	"provider":       "quorum",
	"consensus_type": "raft",
	"_id":            "zzy7ww2963",
	"environment_id": "env1",
}

var mockNodes = []map[string]string{mockNode}

func TestNodeCreation(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia/cid/environments/env1/nodes").
		MatchType("json").
		JSON(mockNodeCreatePayload).
		Reply(201).
		JSON(mockNode)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	node := NewNode("blah", "member1")
	res, err := client.CreateNode("cid", "env1", &node)

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 201)

	var respBody map[string]string
	if err := json.Unmarshal(res.Body(), &respBody); err != nil {
		panic(err)
	}
	st.Expect(t, respBody, mockNode)

	st.Expect(t, gock.IsDone(), true)
}

func TestNodeGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/cid/environments/env1/nodes/zzy7ww2963").
		Reply(200).
		JSON(mockNode)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var node Node
	res, err := client.GetNode("cid", "env1", "zzy7ww2963", &node)

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 200)

	var respBody map[string]string
	if err := json.Unmarshal(res.Body(), &respBody); err != nil {
		panic(err)
	}
	st.Expect(t, respBody, mockNode)

	st.Expect(t, gock.IsDone(), true)

}

func TestNodeList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/cid/environments/env1/nodes").
		Reply(200).
		JSON(mockNodes)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var nodes []Node
	res, err := client.ListNodes("cid", "env1", &nodes)

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 200)

	var respBody []map[string]string
	if err := json.Unmarshal(res.Body(), &respBody); err != nil {
		panic(err)
	}
	st.Expect(t, respBody, mockNodes)

	st.Expect(t, gock.IsDone(), true)
}

func TestNodeDelete(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Delete("/api/v1/consortia/cid/environments/env1/nodes/zzy7ww2963").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	res, err := client.DeleteNode("cid", "env1", "zzy7ww2963")

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 202)
	st.Expect(t, gock.IsDone(), true)
}
