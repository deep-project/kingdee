package core

import (
	"errors"
	"time"
)

type Core struct {
	Fetcher                  Fetcher       // 抓取器
	Session                  Session       // 会话
	RefreshSessionInterval   time.Duration // 定时刷新sessionID的时间间隔
	SessionExpiredRetryCount int           // session过期重试次数

	refreshSessionTicker *time.Ticker // 定时刷新sessionid的定时器
}

func New(f Fetcher, s Session) *Core {
	c := &Core{
		Fetcher:                  f,
		Session:                  s,
		RefreshSessionInterval:   15 * time.Minute,
		SessionExpiredRetryCount: 1,
	}
	c.RefreshSession()       // 初始化时刷新session
	c.RunRefreshSessionJob() // 初始化是启动定时刷新任务
	return c
}

// 设置session接口
func (c *Core) SetSession(val Session) *Core {
	c.Session = val
	return c
}

// 设置刷新session间隔时间
func (c *Core) SetRefreshSessionInterval(val time.Duration) *Core {
	c.RefreshSessionInterval = val
	c.RunRefreshSessionJob()
	return c
}

// 设置session过期重试次数
func (c *Core) SetSessionExpiredRetryCount(val int) *Core {
	c.SessionExpiredRetryCount = val
	return c
}

// 设置session接口并刷新
func (c *Core) SetSessionAndRefresh(login Session) error {
	c.SetSession(login)
	return c.RefreshSession()
}

// 获取session数据
func (c *Core) GetSessionData() (*SessionData, error) {
	if c.Session == nil {
		return nil, errors.New("session interface undefined")
	}
	return c.Session.GetSession()
}

// 获取KDSVCSessionId
func (c *Core) GetKDSVCSessionId() (string, error) {
	s, err := c.GetSessionData()
	if s == nil {
		return "", err
	}
	return s.KDSVCSessionId, nil
}

// 刷新session
func (c *Core) RefreshSession() (err error) {
	if c.Session == nil {
		return errors.New("session interface undefined")
	}
	return c.Session.RefreshSession(c)
}

// 启动定时刷新session任务
// 如果存在旧的定时器，会先停止旧的定时器并覆盖
func (c *Core) RunRefreshSessionJob() {
	go func() {
		if c.refreshSessionTicker != nil {
			c.refreshSessionTicker.Stop() // 先停止旧的
		}
		if c.RefreshSessionInterval <= 0 {
			return
		}
		c.refreshSessionTicker = time.NewTicker(c.RefreshSessionInterval)
		defer c.refreshSessionTicker.Stop()
		for range c.refreshSessionTicker.C {
			c.RefreshSession()
		}
	}()
}

func (c *Core) Call(serviceName string, params map[string]any) ([]byte, error) {
	return c.call(serviceName, params, 0)
}

func (c *Core) call(serviceName string, params map[string]any, count int) (res []byte, err error) {
	s, err := c.GetSessionData()
	if err != nil {
		return
	}
	if c.Fetcher == nil {
		return nil, errors.New("fetcher undefined")
	}
	res, err = c.Request(serviceName, s.KDSVCSessionId, params)
	if err != nil {
		return
	}
	// 如果sessionid过期，则刷新sessionid后再次请求
	if c.SessionExpiredRetryCount > 0 {
		if count >= c.SessionExpiredRetryCount {
			return
		}
		if IsSessionExpired(res) {
			if err = c.RefreshSession(); err != nil {
				return
			}
			return c.call(serviceName, params, count+1)
		}
	}
	return
}

func (c *Core) Request(serviceName, kdsvcSessionId string, params any) (res []byte, err error) {
	return c.Fetcher.Run(serviceName, kdsvcSessionId, params)
}
