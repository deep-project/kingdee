package adapters

import (
	"bytes"
	"encoding/json"
	"io"
	"maps"
	"net/http"
	"strings"
)

type FetcherHTTP struct {
	BaseURL           string            // api请求基地址
	UserAgent         string            // api请求的UserAgent
	APIClientIdentity string            // 客户端标识(如果系统设置了验证标识，则需要填写)
	RequestHeaders    map[string]string // 额外的api请求头
}

func NewFetcherHTTP(baseURL string) *FetcherHTTP {
	return &FetcherHTTP{BaseURL: baseURL}
}

func (f *FetcherHTTP) Run(serviceName, kdsvcSessionId string, params any) (res []byte, err error) {
	var headers = map[string]string{}
	if kdsvcSessionId != "" {
		headers["kdservice-sessionid"] = kdsvcSessionId
	}
	return f.request(serviceName, f.getRequestHeaders(headers), params)
}

func (f *FetcherHTTP) getRequestHeaders(more map[string]string) map[string]string {
	res := map[string]string{
		"Content-Type":        "application/json",
		"User-Agent":          f.UserAgent,
		"api-client-identity": f.APIClientIdentity,
	}
	maps.Copy(res, f.RequestHeaders)
	maps.Copy(res, more)
	return res
}

func (f *FetcherHTTP) getRequestURL(serviceName string) string {
	baseURL := f.BaseURL
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}
	return baseURL + serviceName
}

func (f *FetcherHTTP) request(serviceName string, header map[string]string, params any) (res []byte, err error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", f.getRequestURL(serviceName), bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	for k, v := range header {
		if v != "" {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
