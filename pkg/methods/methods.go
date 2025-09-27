package methods

import (
	"github.com/deep-project/kingdee/pkg/client"
)

type Methods struct {
	client *client.Client
}

func New(client *client.Client) *Methods {
	return &Methods{client: client}
}

// TODO 判断是否已经登录
func (m *Methods) IsLogin() (bool, error) {
	return true, nil
}
