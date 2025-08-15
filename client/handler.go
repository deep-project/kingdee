package client

import (
	"errors"
	"time"

	"github.com/deep-project/kingdee/utils"
)

type Handler struct {
	options *Options
	fetcher *Fetcher
	// 定时刷新sessionid的定时器
	refreshSessionIdTicker *time.Ticker
}

func NewHandler(options *Options) (_ *Handler, err error) {
	fetcher, err := NewFetcher(options)
	if err != nil {
		return
	}
	h := &Handler{options: options, fetcher: fetcher}
	h.RefreshSessionID()
	h.RunRefreshSessionIdJob()
	return h, nil
}

// 设置登录接口
func (h *Handler) SetLogin(login LoginInterface) {
	h.options.SetLogin(login)
}

// 设置登录接口并刷新sessionid
func (h *Handler) SetLoginAndRefreshSessionID(login LoginInterface) error {
	h.SetLogin(login)
	return h.RefreshSessionID()
}

// 获取sessionid
func (h *Handler) GetKDSVCSessionId() string {
	if h.fetcher == nil {
		return ""
	}
	return h.fetcher.GetKDSVCSessionId()
}

// 刷新sessionid
func (h *Handler) RefreshSessionID() (err error) {
	if h.options.Login == nil {
		return errors.New("login undefined")
	}
	s, err := h.options.Login.GetSession(h.fetcher)
	if err != nil {
		return
	}
	h.fetcher.SetKDSVCSessionId(s.KDSVCSessionId)
	return
}

// 定时刷新session ID任务
// 如果存在旧的定时器，会先停止旧的定时器并覆盖
func (h *Handler) RunRefreshSessionIdJob() {
	go func() {
		if h.options.RefreshSessionIdInterval <= 0 {
			return
		}
		if h.refreshSessionIdTicker != nil {
			h.refreshSessionIdTicker.Stop() // 先停止旧的
		}
		h.refreshSessionIdTicker = time.NewTicker(h.options.RefreshSessionIdInterval)
		defer h.refreshSessionIdTicker.Stop()
		for range h.refreshSessionIdTicker.C {
			h.RefreshSessionID()
		}
	}()
}

func (h *Handler) Call(serviceName string, params map[string]any) ([]byte, error) {
	return h.call(serviceName, params, 0)
}

func (h *Handler) call(serviceName string, params map[string]any, count int) (res []byte, err error) {
	if h.fetcher == nil {
		return nil, errors.New("API undefined")
	}
	res, err = h.fetcher.Request(serviceName, params)
	if err != nil {
		return
	}
	// 如果sessionid过期，则刷新sessionid后再次请求
	if h.options.SessionExpiredRetryCount > 0 {
		if count >= h.options.SessionExpiredRetryCount {
			return
		}
		if utils.IsSessionExpired(res) {
			if err = h.RefreshSessionID(); err != nil {
				return
			}
			return h.call(serviceName, params, count+1)
		}
	}
	return
}
