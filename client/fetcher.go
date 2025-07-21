package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"maps"
	"net/http"
	"strings"
)

type Fetcher struct {
	BaseURL        string
	UserAgent      string
	KDSVCSessionId string
}

func NewFetcher(baseURL, userAgent string) (*Fetcher, error) {
	if baseURL == "" {
		return nil, errors.New("BaseURL undefined")
	}
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}
	return &Fetcher{BaseURL: baseURL, UserAgent: userAgent}, nil
}

func (f *Fetcher) SetKDSVCSessionId(kdsvcSessionId string) {
	f.KDSVCSessionId = kdsvcSessionId
}

func (f *Fetcher) GetRequestURL(serviceName string) string {
	return f.BaseURL + serviceName
}

func (f *Fetcher) RequestHeader(more map[string]string) map[string]string {
	res := map[string]string{
		"Content-Type":        "application/json",
		"User-Agent":          f.UserAgent,
		"kdservice-sessionid": f.KDSVCSessionId,
	}
	maps.Copy(res, more)
	return res
}

func (f *Fetcher) Request(serviceName string, params any) (res []byte, err error) {
	return f.BaseRequest(serviceName, f.RequestHeader(nil), params)
}

func (f *Fetcher) BaseRequest(serviceName string, header map[string]string, params any) (res []byte, err error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", f.GetRequestURL(serviceName), bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
