package adapters

import (
	"fmt"

	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/consts"
	"github.com/tidwall/gjson"
)

type LoginByValidateUser struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	Password   string // 密码
	LanguageID string // 语言ID
}

func (login *LoginByValidateUser) KDSVCSessionId(api *client.API) (string, error) {
	params := map[string]any{
		"parameters": []string{
			login.AccountID,
			login.Username,
			login.Password,
			login.LanguageID,
		},
	}
	respBody, err := api.Request(consts.LoginByValidateUser_API, params)
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
