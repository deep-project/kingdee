package adapters

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/consts"
	"github.com/deep-project/kingdee/pkg/session"
	"github.com/tidwall/gjson"
)

type LoginBySign struct {
	AccountID  string // 账套ID
	Username   string // 用户名
	AppID      string // 应用ID
	AppSecret  string // 应用秘钥
	LanguageID string // 语言ID
}

func (login *LoginBySign) GetSession(f *client.Fetcher) (*session.Session, error) {
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
		return nil, err
	}
	// TODO 登录接口增加获取当前登录用户信息
	// 在这里面取数据
	// fmt.Println(string(respBody))
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

func (login *LoginBySign) getSign(timestamp string) string {
	arr := []string{login.AccountID, login.Username, login.AppID, login.AppSecret, timestamp}
	slices.Sort(arr) // 排序
	str := strings.Join(arr, "")
	hash := sha256.Sum256([]byte(str)) // 返回 [32]byte 数组
	return hex.EncodeToString(hash[:]) // 转成十六进制字符串
}
