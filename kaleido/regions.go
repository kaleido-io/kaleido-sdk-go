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
	resty "gopkg.in/resty.v1"
)

const (
	regionBasePath = "/regions"
)

type Region struct {
	U0 RegionInfo `json:"u0,omitempty"`
	E0 RegionInfo `json:"e0,omitempty"`
	A0 RegionInfo `json:"a0,omitempty"`
	K0 RegionInfo `json:"k0,omitempty"`
	U1 RegionInfo `json:"u1,omitempty"`
	E1 RegionInfo `json:"e1,omitempty"`
}

type RegionInfo struct {
	APIHost         string           `json:"api_host,omitempty"`
	DeploymentZones []DeploymentZone `json:"deployment_zones,omitempty"`
}

type DeploymentZone struct {
	Host                 string `json:"host,omitempty"`
	AcceptingDeployments bool   `json:"accepting_deployments,omitempty"`
}

func (c *KaleidoClient) ListRegions(resultBox *[]Region) (*resty.Response, error) {
	return c.Client.R().SetResult(resultBox).Get(regionBasePath)
}
