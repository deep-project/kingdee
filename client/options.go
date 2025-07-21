package client

import "time"

type LoginInterface interface {
	KDSVCSessionId(*Fetcher) (string, error)
}

type Options struct {
	BaseURL                  string         // api请求基地址
	UserAgent                string         // api请求的UserAgent
	Login                    LoginInterface // 登录接口
	RefreshSessionIdInterval time.Duration  // 定时刷新sessionID的时间间隔
	SessionExpiredRetryCount int            // session过期重试次数
}

func NewOptions(baseURL string, login LoginInterface) Options {
	return Options{
		BaseURL:                  baseURL,
		UserAgent:                "Kingdee/Golang WebApi SDK (compatible: K3Cloud 7.x)",
		Login:                    login,
		RefreshSessionIdInterval: 5 * time.Minute,
		SessionExpiredRetryCount: 1,
	}
}
