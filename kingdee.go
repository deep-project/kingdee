package kingdee

import (
	"github.com/deep-project/kingdee/adapters"
	"github.com/deep-project/kingdee/pkg/client"
	"github.com/deep-project/kingdee/pkg/core"
)

func New(baseURL string, s core.Session) (*client.Client, error) {
	c := core.New(adapters.NewFetcherHTTP(baseURL), s)
	return client.New(c)
}

func NewByFetcher(f core.Fetcher, s core.Session) (*client.Client, error) {
	c := core.New(f, s)
	return client.New(c)
}

func NewByCore(c *core.Core) (*client.Client, error) {
	return client.New(c)
}
