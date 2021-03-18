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

	gock "gopkg.in/h2non/gock.v1"
)

func TestPrivateStackBridgeConfigGet(t *testing.T) {

	consortiumID := "cons1"
	envID := "env1"
	serviceID := "svc1"

	gock.New("http://example.com").
		Get("/api/v1/consortia/cons1/environments/env1/services/svc1/tunneler_config").
		Reply(200).
		JSON(map[string]interface{}{
			"some": "config",
		})

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")

	var fetchedConfig map[string]interface{}
	res, err := client.GetPrivateStackBridgeConfig(consortiumID, envID, serviceID, &fetchedConfig)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Failed service fetch bridge config. Code: %d", res.StatusCode())
	}

}
