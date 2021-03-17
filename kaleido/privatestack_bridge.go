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

	resty "gopkg.in/resty.v1"
)

const (
	privateStackBridgeConfigPath = "/consortia/%s/environments/%s/services/%s/tunneler_config"
)

func (c *KaleidoClient) GetPrivateStackBridgeConfig(consortiumID, envID, serviceID string, resultBox *map[string]interface{}) (*resty.Response, error) {
	path := fmt.Sprintf(privateStackBridgeConfigPath, consortiumID, envID, serviceID)
	return c.Client.R().SetResult(resultBox).Get(path)
}
