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
	options        *Options
	kdsvcSessionId string
}

func NewFetcher(opt *Options) (*Fetcher, error) {
	if opt == nil {
		return nil, errors.New("options undefined")
	}
	if opt.BaseURL == "" {
		return nil, errors.New("BaseURL undefined")
	}
	if !strings.HasSuffix(opt.BaseURL, "/") {
		opt.BaseURL = opt.BaseURL + "/"
	}
	return &Fetcher{options: opt}, nil
}

func (f *Fetcher) SetKDSVCSessionId(kdsvcSessionId string) {
	f.kdsvcSessionId = kdsvcSessionId
}
func (f *Fetcher) GetKDSVCSessionId() string {
	return f.kdsvcSessionId
}
func (f *Fetcher) GetRequestURL(serviceName string) string {
	return f.options.BaseURL + serviceName
}

func (f *Fetcher) RequestHeader(more map[string]string) map[string]string {
	res := map[string]string{
		"Content-Type":        "application/json",
		"User-Agent":          f.options.UserAgent,
		"kdservice-sessionid": f.kdsvcSessionId,
		"api-client-identity": f.options.APIClientIdentity,
	}
	maps.Copy(res, f.options.RequestHeader)
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
