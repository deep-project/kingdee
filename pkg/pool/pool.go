package pool

import (
	"errors"

	"github.com/deep-project/kingdee/pkg/client"
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

// 通过回调函数批量创建
func NewBySize(size int, getClient func(i int) (*client.Client, error)) (*Pool, error) {
	if getClient == nil {
		return nil, errors.New("getClient function undefined")
	}
	var clients []*client.Client
	for i := range size {
		cli, err := getClient(i)
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
