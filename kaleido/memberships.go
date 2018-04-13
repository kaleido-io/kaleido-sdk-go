package kaleido

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

const (
	memBasePath = "/consortia/%s/memberships"
)

type Membership struct {
	OrgName string `json:"org_name"`
	Id      string `json:"_id,omitempty"`
}

func NewMembership(orgName string) Membership {
	return Membership{orgName, ""}
}

func (c *KaleidoClient) ListMemberships(consortiaId string, resultBox *[]Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath, consortiaId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) CreateMembership(consortiaId string, membership *Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath, consortiaId)
	return c.Client.R().SetResult(membership).SetBody(membership).Post(path)
}

func (c *KaleidoClient) DeleteMembership(consortiaId, membershipId string) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath+"/%s", consortiaId, membershipId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) GetMembership(consortiaId, membershipId string, resultBox *Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath+"/%s", consortiaId, membershipId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
