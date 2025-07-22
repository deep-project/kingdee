package client

import "time"

type LoginInterface interface {
	KDSVCSessionId(*Fetcher) (string, error)
}

type Options struct {
	BaseURL                  string            // api请求基地址
	UserAgent                string            // api请求的UserAgent
	APIClientIdentity        string            // 客户端标识(如果系统设置了验证标识，则需要填写)
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
