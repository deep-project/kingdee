package adapters

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/deep-project/kingdee/pkg/consts"
	"github.com/deep-project/kingdee/pkg/core"
	"github.com/tidwall/gjson"
)

type LoginByValidateUser struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	Password   string // 密码
	LanguageID string // 语言ID

	sessionData core.SessionData
}

func (login *LoginByValidateUser) GetSession() (*core.SessionData, error) {
	return &login.sessionData, nil
}

func (login *LoginByValidateUser) RefreshSession(c *core.Core) error {
	if c == nil {
		return errors.New("core undefined")
	}
	params := map[string]any{
		"parameters": []string{
			login.AccountID,
			login.Username,
			login.Password,
			login.LanguageID,
		},
	}
	respBody, err := c.Request(consts.LoginByValidateUser_API, "", params)
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
