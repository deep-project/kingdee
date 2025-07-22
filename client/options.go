package client

import "time"

type LoginInterface interface {
	KDSVCSessionId(*Fetcher) (string, error)
}

type Options struct {
	BaseURL                  string            // api请求基地址
	UserAgent                string            // api请求的UserAgent
	RequestHeader            map[string]string // 额外的api请求头
	Login                    LoginInterface    // 登录接口
	RefreshSessionIdInterval time.Duration     // 定时刷新sessionID的时间间隔
	SessionExpiredRetryCount int               // session过期重试次数
}

func NewOptions(baseURL string, login LoginInterface) Options {
	return Options{
		BaseURL:                  baseURL,
		UserAgent:                "Kingdee/Golang WebApi SDK (author: https://github.com/deep-project/kingdee)",
		Login:                    login,
		RefreshSessionIdInterval: 5 * time.Minute,
		SessionExpiredRetryCount: 1,
	}
}
