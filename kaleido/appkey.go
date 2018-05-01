package kaleido

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type AppKey struct {
	MembershipId string `json:"membership_id"`
	AuthType     string `json:"auth_type,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Id           string `json:"_id,omitempty"`
}

const (
	appKeyBasePath = "/consortia/%s/environments/%s/appcreds"
)

func NewAppKey(membershipId string) AppKey {
	return AppKey{
		MembershipId: membershipId,
	}
}

func (c *KaleidoClient) CreateAppKey(consortiumId, envId string, appKey *AppKey) (*resty.Response, error) {
	path := fmt.Sprintf(appKeyBasePath, consortiumId, envId)
	return c.Client.R().SetBody(appKey).SetResult(appKey).Post(path)
}

func (c *KaleidoClient) GetAppKey(consortiumId, envId, appKeyId string, resultBox *AppKey) (*resty.Response, error) {
	path := fmt.Sprintf(appKeyBasePath+"/%s", consortiumId, envId, appKeyId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) DeleteAppKey(consortiumId, envId, appKeyId string) (*resty.Response, error) {
	path := fmt.Sprintf(appKeyBasePath+"/%s", consortiumId, envId, appKeyId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListAppKeys(consortiumId, envId string, resultBox *[]AppKey) (*resty.Response, error) {
	path := fmt.Sprintf(appKeyBasePath, consortiumId, envId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
