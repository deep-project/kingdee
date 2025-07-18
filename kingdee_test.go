package kingdee

import (
	"kingdee/adapters"
	"kingdee/client"
	"testing"
)

func TestKingdee(t *testing.T) {

	cli, err := New(client.NewAPI(""), &adapters.LoginByAppSecret{}, &client.Options{})
	if err != nil {
		t.Error(err)
	}
	t.Log(cli.Login.KDSVCSessionId())
}
