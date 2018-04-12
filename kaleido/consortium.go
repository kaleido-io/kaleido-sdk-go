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
