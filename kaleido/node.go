package kaleido

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type Node struct {
	Name         string `json:"name"`
	MembershipId string `json:"membership_id"`
	Id           string `json:"_id,omitempty"`
}

func NewNode(name, membershipId string) Node {
	return Node{
		Name:         name,
		MembershipId: membershipId,
		Id:           "",
	}
}

func (c *KaleidoClient) CreateNode(consortium, envId string, node *Node) (*resty.Response, error) {
	path := fmt.Sprintf("/consortia/%s/environments/%s/nodes", consortium, envId)
	return c.Client.R().SetResult(node).SetBody(node).Post(path)
}

func (c *KaleidoClient) DeleteNode(consortium, envId, nodeId string) (*resty.Response, error) {
	path := fmt.Sprintf("/consortia/%s/environments/%s/nodes/%s", consortium, envId, nodeId)
	return c.Client.R().Delete(path)
}

func (c *KaleidoClient) ListNodes(consortium, envId string, resultBox *[]Node) (*resty.Response, error) {
	path := fmt.Sprintf("/consortia/%s/environments/%s/nodes", consortium, envId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
