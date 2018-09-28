package registry

import (
	"fmt"

	eth "github.com/ethereum/go-ethereum/common"
	resty "gopkg.in/resty.v1"
)

// Node struct
type Node struct {
	Path    string
	Address eth.Address
	Proof   string
}

func NewAppCreds(membershipId string) AppCreds {
	return AppCreds{
		MembershipId: membershipId,
	}
}

func (c *KaleidoClient) CreateAppCreds(consortiumId, envId string, appcreds *AppCreds) (*resty.Response, error) {
	path := fmt.Sprintf(appcredsBasePath, consortiumId, envId)
	return c.Client.R().SetBody(appcreds).SetResult(appcreds).Post(path)
}

func (c *KaleidoClient) GetAppCreds(consortiumId, envId, appcredsId string, resultBox *AppCreds) (*resty.Response, error) {
	path := fmt.Sprintf(appcredsBasePath+"/%s", consortiumId, envId, appcredsId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *KaleidoClient) DeleteAppCreds(consortiumId, envId, appcredsId string) (*resty.Response, error) {
	path := fmt.Sprintf(appcredsBasePath+"/%s", consortiumId, envId, appcredsId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListAppCreds(consortiumId, envId string, resultBox *[]AppCreds) (*resty.Response, error) {
	path := fmt.Sprintf(appcredsBasePath, consortiumId, envId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
