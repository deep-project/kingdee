package kingdee

import "github.com/deep-project/kingdee/client"

func New(options client.Options) (*client.Client, error) {
	return client.NewClient(options)
}

func NewOptions(baseURL string, login client.LoginInterface) client.Options {
	return client.NewOptions(baseURL, login)
}
