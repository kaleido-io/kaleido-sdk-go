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

var mockDeployServicePayload = map[string]string{
	"service":       "idregistry",
	"name":          "blah",
	"membership_id": "member1",
}

var mockService = map[string]string{
	"name":           "blah",
	"membership_id":  "member1",
	"environment_id": "env1",
	"service":        "idregistry",
	"service_type":   "utility",
	"state":          "started",
	"consortium_id":  "cid",
}

var mockServices = []map[string]string{mockService}

func TestServiceDeploy(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/c/cid/e/env1/services").
		MatchType("json").
		JSON(mockDeployServicePayload).
		Reply(201).
		JSON(mockService)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	service := NewService("idregistry", "blah", "member1")
	res, err := client.DeployService("cid", "env1", &service)

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 201)
	st.Expect(t, gock.IsDone(), true)
}

func TestServiceGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/services").
		MatchParam("_id", "service-id").
		Reply(200).
		JSON(mockServices)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var services []Service
	res, err := client.GetService("cid", "env1", "service-id", &services)

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 200)
	st.Expect(t, gock.IsDone(), true)

}

func TestServiceList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/c/cid/e/env1/services").
		Reply(200).
		JSON(mockServices)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var services []Service
	res, err := client.ListServices("cid", "env1", &services)

	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 200)
	st.Expect(t, gock.IsDone(), true)
}
