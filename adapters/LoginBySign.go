package adapters

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/consts"
	"github.com/tidwall/gjson"
)

type LoginBySign struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	AppID      string // 应用ID
	AppSecret  string // 应用秘钥
	LanguageID string // 语言ID
}

func (login *LoginBySign) KDSVCSessionId(f *client.Fetcher) (string, error) {
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
	respBody, err := f.Request(consts.LoginBySign_API, params)
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

func (login *LoginBySign) getSign(timestamp string) string {
	arr := []string{login.AccountID, login.Username, login.AppID, login.AppSecret, timestamp}
	slices.Sort(arr) // 排序
	str := strings.Join(arr, "")
	hash := sha256.Sum256([]byte(str)) // 返回 [32]byte 数组
	return hex.EncodeToString(hash[:]) // 转成十六进制字符串
}
