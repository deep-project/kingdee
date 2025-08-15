package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/deep-project/kingdee"
)

// 测试获取账套列表
func TestGetDataCenterList(t *testing.T) {
	cli, err := kingdee.New(os.Getenv("BASE_URL"), nil)
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
