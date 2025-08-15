package adapters

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/deep-project/kingdee/pkg/consts"
	"github.com/deep-project/kingdee/pkg/core"
	"github.com/tidwall/gjson"
)

type LoginBySign struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	AppID      string // 应用ID
	AppSecret  string // 应用秘钥
	LanguageID string // 语言ID

	sessionData core.SessionData
}

func (login *LoginBySign) GetSession() (*core.SessionData, error) {
	return &login.sessionData, nil
}

func (login *LoginBySign) RefreshSession(c *core.Core) error {
	if c == nil {
		return errors.New("core undefined")
	}
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	params := map[string]any{
		"parameters": []string{
			login.AccountID,
			login.Username,
			login.AppID,
			timestamp,
			login.getSign(timestamp),
			login.LanguageID,
		},
	}
	respBody, err := c.Request(consts.LoginBySign_API, "", params)
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

func (login *LoginBySign) getSign(timestamp string) string {
	arr := []string{login.AccountID, login.Username, login.AppID, login.AppSecret, timestamp}
	slices.Sort(arr) // 排序
	str := strings.Join(arr, "")
	hash := sha256.Sum256([]byte(str)) // 返回 [32]byte 数组
	return hex.EncodeToString(hash[:]) // 转成十六进制字符串
}
