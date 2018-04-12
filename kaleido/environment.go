package kaleido

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type Environment struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Provider      string `json:"provider"`
	ConsensusType string `json:"consensus_type"`
	Id            string `json:"_id,omitempty"`
}

const (
	BASE_PATH = "/consortia/%s/environments"
)

func NewEnvironment(name, description, provider, consensus string) Environment {
	return Environment{
		Name:          name,
		Description:   description,
		Provider:      provider,
		ConsensusType: consensus,
		Id:            "",
	}
}

func (c *KaleidoClient) ListEnvironments(consortiumId string, resultBox *[]Environment) (*resty.Response, error) {
	path := fmt.Sprintf(BASE_PATH, consortiumId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) CreateEnvironment(consortiumId string, environment *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(BASE_PATH, consortiumId)
	return c.Client.R().SetResult(environment).SetBody(environment).Post(path)
}

func (c *KaleidoClient) DeleteEnvironment(consortiumId, environmentId string) (*resty.Response, error) {
	path := fmt.Sprintf(BASE_PATH+"/%s", consortiumId, environmentId)
	return c.Client.R().Delete(path)
}
