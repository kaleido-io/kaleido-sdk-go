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

func TestConsortiumCreationListDeletion(t *testing.T) {

	consortiaPostBody := map[string]string{
		"name":        "testConsortium",
		"description": "test description",
	}

	consortiaReplyBody := map[string]string{
		"_id":         "zzam3flatl",
		"description": "test description",
		"name":        "testConsortium",
		"owner":       "zzgl55vock",
		"state":       "setup",
	}

	mockConsortia := []map[string]string{consortiaReplyBody}

	// setup mock
	defer gock.Off()

	gock.New("http://example.com").
		Post("/api/v1/consortia").
		MatchType("json").
		JSON(consortiaPostBody).
		Reply(201).
		JSON(consortiaReplyBody)

	gock.New("http://example.com").
		Get("/api/v1/consortia/zzam3flatl").
		Reply(200).
		JSON(consortiaReplyBody)

	gock.New("http://example.com").
		Get("/api/v1/consortia").
		Reply(200).
		JSON(mockConsortia)

	gock.New("http://example.com").
		Delete("/api/v1/consortia/zzam3flatl").
		Reply(202)

	// TODO make the testing more useful by moving away from printing json the way we do today
	// for now, it is simply asserting that the request was called (and returned what we expected
	// but it is mocked so it will always return what we mocked, once the other commands are reworked
	// these tests can be used to validate what the function invocation returned (i.e. how we processed
	// a response as opposed to the response itself which is pretty much useless but left in place here
	// for now)
	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	consortium := NewConsortium("testConsortium", "test description")
	res, err := client.CreateConsortium(&consortium)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 201) // this will always be true,

	var respBody map[string]string
	if err := json.Unmarshal(res.Body(), &respBody); err != nil {
		panic(err)
	}
	st.Expect(t, respBody, consortiaReplyBody)

	var consortium2 Consortium
	res, err = client.GetConsortium(consortium.ID, &consortium2)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 200)

	if err := json.Unmarshal(res.Body(), &respBody); err != nil {
		panic(err)
	}
	st.Expect(t, respBody, consortiaReplyBody)

	var consortia []Consortium
	res, err = client.ListConsortium(&consortia)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode(), 200)

	var respArray []map[string]string
	if err := json.Unmarshal(res.Body(), &respArray); err != nil {
		panic(err)
	}
	st.Expect(t, respArray, mockConsortia)

	//Check for a newly created consortia and delete it.
	countNew := 0
	for _, x := range consortia {
		t.Logf("\n%v", x)
		if x.Name == "testConsortium" && (x.State != "deleted" && x.State != "delete_pending") {
			res, err = client.DeleteConsortium(x.ID)
			st.Expect(t, err, nil)
			st.Expect(t, res.StatusCode(), 202)
			countNew++
			t.Logf("\nNew Consortium: %v", x)
		}
	}

	st.Expect(t, gock.IsDone(), true)
}
