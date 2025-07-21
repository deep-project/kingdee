package adapters

import (
	"fmt"

	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/consts"

	"github.com/tidwall/gjson"
)

type LoginByAppSecret struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	AppID      string // 应用ID
	AppSecret  string // 应用秘钥
	LanguageID string // 语言ID
}

func (login *LoginByAppSecret) KDSVCSessionId(f *client.Fetcher) (string, error) {
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
		return "", err
	}

	var (
		respBodyStr     = string(respBody)
		LoginResultType = gjson.Get(respBodyStr, "LoginResultType").Int()
		Message         = gjson.Get(respBodyStr, "Message").String()
		KDSVCSessionId  = gjson.Get(respBodyStr, "KDSVCSessionId").String()
	)

	if LoginResultType != 1 {
		if Message == "" {
			Message = string(respBody)
		}
		return "", fmt.Errorf("Login failed: %s", Message)
	}
	return KDSVCSessionId, nil
}
