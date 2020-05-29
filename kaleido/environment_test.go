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

	"github.com/nbio/st"
	gock "gopkg.in/h2non/gock.v1"
)

var mockEnvCreatePayload = map[string]interface{}{
	"name":           "testingEnvironment",
	"description":    "just test",
	"provider":       "quorum",
	"consensus_type": "raft",
	"block_period":   0,
	"test_features": map[string]interface{}{
		"multi_region": true,
	},
}

var mockEnv = map[string]string{
	"_id":          "envid",
	"name":         "testingEnvironment",
	"description":  "just test",
	"provider":     "quorum",
	"consensus":    "raft",
	"consortia_id": "cid",
}

var mockEnvs = []map[string]string{mockEnv}

func TestEnvironmentCreation(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia/cid/environments").
		MatchType("json").
		JSON(mockEnvCreatePayload).
		Reply(201).
		JSON(mockEnv)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	env := NewEnvironment("testingEnvironment", "just test", "quorum", "raft", true, 0)
	_, err := client.CreateEnvironment("cid", &env)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestEnvironmentDelete(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Delete("/api/v1/consortia/cid/environments/envid").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	_, err := client.DeleteEnvironment("cid", "envid")

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestEnvironmentGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/cid/environments/envid").
		Reply(200).
		JSON(mockEnv)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var env Environment
	_, err := client.GetEnvironment("cid", "envid", &env)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestEnvironmentList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/cid/environments").
		Reply(200).
		JSON(mockEnvs)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var envs []Environment
	_, err := client.ListEnvironments("cid", &envs)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
