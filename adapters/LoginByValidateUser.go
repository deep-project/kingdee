package adapters

import (
	"encoding/json"
	"fmt"

	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/consts"
	"github.com/deep-project/kingdee/pkg/session"
	"github.com/tidwall/gjson"
)

type LoginByValidateUser struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	Password   string // 密码
	LanguageID string // 语言ID
}

func (login *LoginByValidateUser) GetSession(f *client.Fetcher) (*session.Session, error) {
	params := map[string]any{
		"parameters": []string{
			login.AccountID,
			login.Username,
			login.Password,
			login.LanguageID,
		},
	}
	respBody, err := f.Request(consts.LoginByValidateUser_API, params)
	if err != nil {
		return nil, err
	}
	var (
		respBodyStr     = string(respBody)
		LoginResultType = gjson.Get(respBodyStr, "LoginResultType").Int()
		Message         = gjson.Get(respBodyStr, "Message").String()
	)
	if LoginResultType != 1 {
		if Message == "" {
			Message = respBodyStr
		}
		return nil, fmt.Errorf("login failed 1: %s", Message)
	}

	var s session.Session
	if err := json.Unmarshal(respBody, &s); err != nil {
		return nil, fmt.Errorf("login failed 2: %s", err.Error())
	}

	return &s, nil
}
