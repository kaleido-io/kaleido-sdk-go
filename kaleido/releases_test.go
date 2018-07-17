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
	"fmt"
	"os"
	"testing"
)

func TestListReleases(t *testing.T) {
	client := NewClient(os.Getenv("KALEIDO_API"), os.Getenv("KALEIDO_API_KEY"))
	var releases []Release
	res, err := client.ListReleases(&releases)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatal(fmt.Errorf("Fetching releases returned %d code.", res.StatusCode()))
	}

	if len(releases) <= 0 {
		t.Fatalf("No releases were returned")
	}

	expectedRelease := releases[0]
	var release []Release
	res, err = client.GetRelease(expectedRelease.Provider, expectedRelease.Version, &release)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatal(fmt.Errorf("Could not fetch the individual release: status code %d.", res.StatusCode()))
	}

	actualRelease := release[0]
	if expectedRelease.Provider != actualRelease.Provider {
		t.Fatal(fmt.Errorf("Expected retrieved release provider to be %s, but got %s", expectedRelease.Provider, actualRelease.Provider))
	}

	if expectedRelease.Version != actualRelease.Version {
		t.Fatal(fmt.Errorf("Expected retrieved provider version to be %s, but got %s.", expectedRelease.Version, actualRelease.Version))
	}
}
