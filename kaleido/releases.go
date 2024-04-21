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
	resty "github.com/go-resty/resty/v2"
)

const (
	releaseBasePath = "/releases"
)

type Release struct {
	ID            string `json:"_id,omitempty"`
	Provider      string `json:"provider,omitempty"`
	Version       string `json:"version,omitempty"`
	VersionPadded string `json:"version_padded,omitempty"`
	ReleaseStatus string `json:"release_status,omitempty"`
}

func (c *KaleidoClient) ListReleases(resultBox *[]Release) (*resty.Response, error) {
	return c.Client.R().SetResult(resultBox).Get(releaseBasePath)
}

func (c *KaleidoClient) GetRelease(provider, version string, resultBox *[]Release) (*resty.Response, error) {
	return c.Client.R().SetResult(resultBox).
		SetQueryParam("version", version).
		SetQueryParam("provider", provider).
		Get(releaseBasePath)
}
