package test

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/pkg/client"
	"github.com/deep-project/kingdee/pkg/pool"
)

func TestPool(t *testing.T) {
	var baseURL = os.Getenv("BASE_URL")
	client_1, err := kingdee.New(baseURL, getLoginBySignAdapter())
	if err != nil {
		t.Error(err)
	}
	client_2, _ := kingdee.New(baseURL, getLoginBySignAdapter())
	client_3, _ := kingdee.New(baseURL, getLoginBySignAdapter())
	client_4, _ := kingdee.New(baseURL, getLoginBySignAdapter())
	client_5, _ := kingdee.New(baseURL, getLoginBySignAdapter())

	var p = pool.New([]*client.Client{client_1, client_2, client_3, client_4, client_5})
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := p.Get()
			defer p.Put(client)
			time.Sleep(2 * time.Second) // 延迟两秒，查看效果
			t.Log(client.Core.GetKDSVCSessionId())
		}()
	}
	wg.Wait()
}

func TestPoolNewBySize(t *testing.T) {
	p, err := pool.NewBySize(5, func(i int) (*client.Client, error) {
		return kingdee.New(os.Getenv("BASE_URL"), getLoginBySignAdapter())
	})
	if err != nil {
		t.Error(err)
		return
	}
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := p.Get()
			defer p.Put(client)
			time.Sleep(2 * time.Second) // 延迟两秒，查看效果
			t.Log(client.Core.GetKDSVCSessionId())
		}()
	}
	wg.Wait()
}
