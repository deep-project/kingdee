package pool

import (
	"github.com/deep-project/kingdee/client"
)

type Pool struct {
	size    int
	clients chan *client.Client
}

func New(clients []*client.Client) *Pool {
	size := len(clients)
	pool := &Pool{
		size:    size,
		clients: make(chan *client.Client, size),
	}
	for _, v := range clients {
		pool.Put(v)
	}
	return pool
}

// 通过统一配置创建
func NewBySize(size int, options client.Options) (*Pool, error) {
	var clients []*client.Client
	for range size {
		cli, err := client.NewClient(options)
		if err != nil {
			return nil, err
		}
		clients = append(clients, cli)
	}
	p := New(clients)
	return p, nil
}

func (cp *Pool) Get() *client.Client {
	return <-cp.clients
}

func (cp *Pool) Put(client *client.Client) {
	cp.clients <- client
}
