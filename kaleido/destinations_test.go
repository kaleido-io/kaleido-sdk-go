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

var mockDestinationCreatePlayload = map[string]interface{}{
	"name":            "a1",
	"kaleido_managed": true,
}

var mockDestination = map[string]interface{}{
	"uri":             "kld://app2app/z/dev2/m/zzf1diukb5/e/zzkdova20l/s/zzl30hmtfb/d/a1",
	"name":            "a1",
	"kaleido_managed": true,
	"setup_complete":  true,
}

var mockDestinations = []map[string]interface{}{mockDestination}

func TestDestinationCreate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Put("/api/v1/app2app/sid/destinations/a1").
		MatchType("json").
		JSON(mockDestinationCreatePlayload).
		Reply(201).
		JSON(mockDestination)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	destination := NewDestination("a1")
	_, err := client.CreateDestination("app2app", "sid", &destination)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestDestinationGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/app2app/sid/destinations/a1").
		Reply(200).
		JSON(mockDestination)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var destination Destination
	_, err := client.GetDestination("app2app", "sid", "a1", &destination)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestDestinationList(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/app2app/sid/destinations").
		Reply(200).
		JSON(mockDestinations)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var destinations []Destination
	_, err := client.ListDestinations("app2app", "sid", &destinations)

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestDestinationDelete(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Delete("/api/v1/app2app/sid/destinations/a1").
		Reply(202)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	_, err := client.DeleteDestination("app2app", "sid", "a1")

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}

func TestDestinationUpdate(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Put("/api/v1/app2app/sid/destinations/a1").
		Reply(200)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	_, err := client.UpdateDestination("app2app", "sid", "a1", &Destination{Name: "a1"})

	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
