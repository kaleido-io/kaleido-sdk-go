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

func TestListRegions(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/api/v1/regions").
		Reply(200)

	client := NewClient("http://example.com/api/v1", "KALEIDO_API_KEY")
	_, err := client.GetRegions(&Regions{})
	st.Expect(t, err, nil)
	st.Expect(t, gock.IsDone(), true)
}
