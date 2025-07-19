package client

import (
	"errors"
	"time"
)

type Options struct {
	// 定时刷新sessionID的时间间隔
	RefreshAPISessionIdInterval time.Duration
}

type Handler struct {
	API     *API
	Login   Login
	Options *Options

	// 定时刷新sessionid的定时器
	refreshAPISessionIdTicker *time.Ticker
}

func (c *Handler) RefreshAPISessionID() (err error) {
	sid, err := c.Login.KDSVCSessionId(c.API)
	if err != nil {
		return
	}
	c.API.SetKDSVCSessionId(sid)
	return nil
}

// 定时刷新api session ID任务
// 如果存在旧的定时器，会先停止旧的定时器并覆盖
func (c *Handler) RunRefreshAPISessionIdJob() {
	go func() {
		if c.Options.RefreshAPISessionIdInterval <= 0 {
			return
		}
		if c.refreshAPISessionIdTicker != nil {
			c.refreshAPISessionIdTicker.Stop() // 先停止旧的
		}
		c.refreshAPISessionIdTicker = time.NewTicker(c.Options.RefreshAPISessionIdInterval)
		defer c.refreshAPISessionIdTicker.Stop()
		for range c.refreshAPISessionIdTicker.C {
			c.RefreshAPISessionID()
		}
	}()
}

func (c *Handler) Call(serviceName string, params map[string]any) ([]byte, error) {
	if c.API == nil {
		return nil, errors.New("API undefined")
	}
	if c.API.KDSVCSessionId == "" {
		return nil, errors.New("KDSVCSessionId undefined")
	}
	return c.API.Request(serviceName, params)
}
