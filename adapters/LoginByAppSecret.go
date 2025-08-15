package adapters

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/deep-project/kingdee/pkg/consts"
	"github.com/deep-project/kingdee/pkg/core"

	"github.com/tidwall/gjson"
)

type LoginByAppSecret struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	AppID      string // 应用ID
	AppSecret  string // 应用秘钥
	LanguageID string // 语言ID

	sessionData core.SessionData
}

func (login *LoginByAppSecret) GetSession() (*core.SessionData, error) {
	return &login.sessionData, nil
}

func (login *LoginByAppSecret) RefreshSession(c *core.Core) error {
	if c == nil {
		return errors.New("core undefined")
	}
	params := map[string]any{
		"parameters": []string{
			login.AccountID,
			login.Username,
			login.AppID,
			login.AppSecret,
			login.LanguageID,
		},
	}
	respBody, err := c.Request(consts.LoginByAppSecret_API, "", params)
	if err != nil {
		return err
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
		return fmt.Errorf("login failed 1: %s", Message)
	}
	if err := json.Unmarshal(respBody, &login.sessionData); err != nil {
		return fmt.Errorf("login failed 2: %s", err.Error())
	}
	return nil
}
