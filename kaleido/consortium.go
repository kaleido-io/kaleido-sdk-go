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

	"gopkg.in/resty.v1"
)

const (
	SingleOrg      = "single-org"
	MultiOrg       = "multi-org"
	DELETE_PENDING = "delete_pending"
	DELETED        = "deleted"
)

type KaleidoClient struct {
	Client *resty.Client
}

type Consortium struct {
	Id          string `json:"_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Mode        string `json:"mode"`
	DeletedAt   string `json:"deleted_at,omitempty"`
	State       string `json:"state,omitempty"`
}

func NewConsortium(name, description, mode string) Consortium {
	return Consortium{
		Id:          "",
		Name:        name,
		Description: description,
		Mode:        mode,
		DeletedAt:   "",
		State:       "",
	}
}

func NewClient(api string, apiKey string) KaleidoClient {
	r := resty.New().SetHostURL(api).SetAuthToken(apiKey)
	return KaleidoClient{r}
}

func (c *KaleidoClient) CreateConsortium(consortium *Consortium) (*resty.Response, error) {
	return c.Client.R().SetBody(consortium).SetResult(consortium).Post("/consortia")
}

func (c *KaleidoClient) ListConsortium(resultBox *[]Consortium) (*resty.Response, error) {
	return c.Client.R().SetResult(resultBox).Get("/consortia")
}

func (c *KaleidoClient) GetConsortium(id string, resultBox *Consortium) (*resty.Response, error) {
	path := fmt.Sprintf("/consortia/%s", id)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) DeleteConsortium(consortiumId string) (*resty.Response, error) {
	return c.Client.R().Delete(fmt.Sprintf("/consortia/%s", consortiumId))
}
