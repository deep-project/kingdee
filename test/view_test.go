package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/models"
)

func TestView(t *testing.T) {
	cli, err := kingdee.New(client.NewAPI(os.Getenv("BASE_URL")), getLoginBySignAdapter(), &client.Options{})
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := cli.View("STK_InStock", models.ViewParams{Number: "CGRK00019"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.ResponseStatus)
	t.Log("-----------")
	fmt.Printf("----------%v\n", resp.Result)
}
