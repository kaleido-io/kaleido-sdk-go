// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
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

var mockApiKeyCreatePayload = map[string]string{
	"org_id": "u0vjobhsvw",
}

var mockApiKeyCreateWithNamePayload = map[string]string{
	"org_id": "u0vjobhsvw",
	"name":   "testApiKey",
}

var mockApiKey = map[string]string{
	"org_id": "u0vjobhsvw",
	"name":   "testApiKey",
	"_id":    "u0pggiazfk",
}

var mockApiKeys = []map[string]string{mockApiKey}

func TestApiKeyCreate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/apikeys").
		MatchType("json").
		JSON(mockApiKeyCreatePayload).
		Reply(201).
		JSON(mockApiKey)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	var apikey = NewApiKey("u0vjobhsvw")
	_, err := client.CreateApiKey(&apikey)
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestApiKeyCreateWithName(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/apikeys").
		MatchType("json").
		JSON(mockApiKeyCreateWithNamePayload).
		Reply(201).
		JSON(mockApiKey)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	var apikey = NewApiKeyWithName("u0vjobhsvw", "testApiKey")
	_, err := client.CreateApiKey(&apikey)
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestApiKeyGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/apikeys/u0pggiazfk").
		Reply(200).
		JSON(mockApiKey)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.GetApiKey("u0pggiazfk", &ApiKey{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestApiKeyList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/apikeys").
		Reply(200).
		JSON(mockApiKeys)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.ListApiKey(&[]ApiKey{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestApiKeyDelete(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Delete("/api/v1/apikeys/u0pggiazfk").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.DeleteApiKey("u0pggiazfk")
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestApiKeyUpdate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Patch("/api/v1/apikeys/u0pggiazfk").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.UpdateApiKey("u0pggiazfk", &ApiKey{Name: "test"})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
