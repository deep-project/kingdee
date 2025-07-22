package test

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/pool"
)

func TestPool(t *testing.T) {
	options := client.NewOptions(os.Getenv("BASE_URL"), getLoginBySignAdapter())
	client_1, err := kingdee.New(options)
	if err != nil {
		t.Error(err)
	}
	client_2, _ := kingdee.New(options)
	client_3, _ := kingdee.New(options)
	client_4, _ := kingdee.New(options)
	client_5, _ := kingdee.New(options)

	var p = pool.New([]*client.Client{client_1, client_2, client_3, client_4, client_5})
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := p.Get()
			defer p.Put(client)
			time.Sleep(2 * time.Second) // 延迟两秒，查看效果
			t.Log("KDSVCSessionId: ", client.Handler.GetKDSVCSessionId())
		}()
	}
	wg.Wait()
}
