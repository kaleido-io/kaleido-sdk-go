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
	"gopkg.in/h2non/gock.v1"
)

var mockAppCredCreatePayload = map[string]string{
	"membership_id": "member1",
}

var mockAppCredCreateWithNamePayload = map[string]string{
	"membership_id": "member1",
	"name":          "testCred",
}

var mockAppCred = map[string]string{
	"membership_id": "zzzipxyjew",
	"name":          "testCred",
	"auth_type":     "basic_auth",
	"_id":           "zzstcszriw",
	"username":      "userid",
	"password":      "userid-password",
}

var mockAppCreds = []map[string]string{mockAppCred}

func TestAppCredCreate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia/c1/environments/env1/appcreds").
		MatchType("json").
		JSON(mockAppCredCreatePayload).
		Reply(201).
		JSON(mockAppCred)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	var appCreds = NewAppCreds("member1")
	_, err := client.CreateAppCreds("c1", "env1", &appCreds)
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestAppCredCreateWithName(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia/c1/environments/env1/appcreds").
		MatchType("json").
		JSON(mockAppCredCreateWithNamePayload).
		Reply(201).
		JSON(mockAppCred)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	var appCreds = NewAppCredsWithName("member1", "testCred")
	_, err := client.CreateAppCreds("c1", "env1", &appCreds)
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestAppCredGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/c1/environments/env1/appcreds/appcred1").
		Reply(200).
		JSON(mockAppCred)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.GetAppCreds("c1", "env1", "appcred1", &AppCreds{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestAppCredList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/consortia/c1/environments/env1/appcreds").
		Reply(200).
		JSON(mockAppCreds)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.ListAppCreds("c1", "env1", &[]AppCreds{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestAppCredDelete(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Delete("/api/v1/consortia/c1/environments/env1/appcreds/appcred1").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.DeleteAppCreds("c1", "env1", "appcred1")
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestAppCredUpdate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Patch("/api/v1/consortia/c1/environments/env1/appcreds/appcred1").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.UpdateAppCreds("c1", "env1", "appcred1", &AppCreds{Name: "test"})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestAppCredRegenerate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia/c1/environments/env1/appcreds/appcred1/regenerate").
		Reply(200).
		JSON(mockAppCred)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.RegenerateAppCreds("c1", "env1", "appcred1", &AppCreds{Name: "test"})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
