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

var mockConfigurationCreatePayload = map[string]interface{}{
	"name":          "configName",
	"membership_id": "member1",
	"type":          "node_config",
	"details": map[string]interface{}{
		"config_specific": "details",
	},
}

var mockConfiguration = Configuration{
	ID:           "zzstcszriw",
	Name:         "configName",
	MembershipID: "member1",
	Type:         "node_config",
	Details: map[string]interface{}{
		"config_specific": "details",
	},
}

var mockConfigurations = []Configuration{mockConfiguration}

func TestConfigurationCreate(t *testing.T) {
	defer gock.Off()
	gock.Observe(gock.DumpRequest)

	gock.New("http://example.com").
		Post("/api/v1/consortia/c1/environments/env1/configurations").
		MatchType("json").
		JSON(mockConfigurationCreatePayload).
		Reply(201).
		JSON(mockConfiguration)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	var config = NewConfiguration("configName", "member1", "node_config", map[string]interface{}{
		"config_specific": "details",
	})
	_, err := client.CreateConfiguration("c1", "env1", &config)
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestConfigurationGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/c1/environments/env1/configurations/configuration1").
		Reply(200).
		JSON(mockConfiguration)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.GetConfiguration("c1", "env1", "configuration1", &Configuration{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestConfigurationList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/c1/environments/env1/configurations").
		Reply(200).
		JSON(mockConfigurations)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.ListConfigurations("c1", "env1", &[]Configuration{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestConfigurationDelete(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Delete("/api/v1/consortia/c1/environments/env1/configurations/configuration1").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.DeleteConfiguration("c1", "env1", "configuration1")
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
