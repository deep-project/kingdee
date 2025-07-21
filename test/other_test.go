package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/client"
)

// 测试获取账套列表
func TestGetDataCenterList(t *testing.T) {
	cli, err := kingdee.New(client.NewOptions(os.Getenv("BASE_URL"), nil))
	if err != nil {
		t.Error(err)
		return
	}
	raw, err := cli.GetDataCenterList()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("----------%v\n", string(raw))
}
