package kingdee

import "kingdee/client"

func New(api *client.API, login client.Login, options *client.Options) (*client.Client, error) {
	return client.NewClient(api, login, options)
}
