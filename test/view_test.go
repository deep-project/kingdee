package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/models"
	"github.com/deep-project/kingdee/utils"
	"github.com/tidwall/gjson"
)

func TestView(t *testing.T) {
	cli, err := kingdee.New(client.NewOptions(os.Getenv("BASE_URL"), getLoginBySignAdapter()))
	if err != nil {
		t.Error(err)
		return
	}
	raw, err := cli.View("STK_InStock", models.ViewParams{Number: utils.Ptr("CGRK00019")})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(utils.GetResultResponseStatus(raw))
	t.Log("-----------")
	fmt.Printf("----------%v\n", gjson.Get(string(raw), "Result.Result").String())
}
