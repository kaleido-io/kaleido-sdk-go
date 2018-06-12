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
	"os"
	"testing"
)

func TestConsortiumCreationListDeletion(t *testing.T) {
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	consortium := NewConsortium("testConsortium", "test description", "single-org")
	res, err := client.CreateConsortium(&consortium)
	if res.StatusCode() != 201 {
		t.Fatalf("Could not create consortium status code: %d.", res.StatusCode())
	}
	if err != nil {
		t.Fatal(err)
	}

	var consortium2 Consortium
	res, err = client.GetConsortium(consortium.Id, &consortium2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Unable to fetch consortium %s response was: %d.", consortium.Id, res.StatusCode())
	}

	if consortium.Id != consortium2.Id {
		t.Fatalf("Fetched consortium id mismatch: expected %s found %s", consortium.Id, consortium2.Id)
	}

	var consortia []Consortium
	_, err = client.ListConsortium(&consortia)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	//Check for a newly created consortia and delete it.
	countNew := 0
	for _, x := range consortia {
		t.Logf("\n%v", x)
		if x.Name == "testConsortium" && (x.State != DELETED && x.State != DELETE_PENDING) {
			res, err = client.DeleteConsortium(x.Id)
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode() != 202 {
				t.Errorf("Consortium Deletion Failed Status %d.", res.StatusCode())
			}
			countNew += 1
			t.Logf("\nNew Consortium: %v", x)
		}
	}
}
