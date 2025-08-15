package adapters

import (
	"encoding/json"
	"fmt"

	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/consts"
	"github.com/deep-project/kingdee/pkg/session"

	"github.com/tidwall/gjson"
)

type LoginByAppSecret struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	AppID      string // 应用ID
	AppSecret  string // 应用秘钥
	LanguageID string // 语言ID
}

func (login *LoginByAppSecret) GetSession(f *client.Fetcher) (*session.Session, error) {
	params := map[string]any{
		"parameters": []string{
			login.AccountID,
			login.Username,
			login.AppID,
			login.AppSecret,
			login.LanguageID,
		},
	}
	respBody, err := f.Request(consts.LoginByAppSecret_API, params)
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
