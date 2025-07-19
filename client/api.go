package client

import (
	"bytes"
	"encoding/json"
	"io"
	"maps"
	"net/http"
)

type API struct {
	BaseURL        string
	KDSVCSessionId string
	UserAgent      string
}

func NewAPI(baseURL string) *API {
	return &API{BaseURL: baseURL}
}

func (api *API) SetKDSVCSessionId(kdsvcSessionId string) {
	api.KDSVCSessionId = kdsvcSessionId
}

func (api *API) GetRequestURL(serviceName string) string {
	return api.BaseURL + serviceName
}

func (api *API) Request(serviceName string, params any) (res []byte, err error) {
	return api.BaseRequest(serviceName, api.RequestHeader(nil), params)
}

func (api *API) RequestHeader(more map[string]string) map[string]string {
	res := map[string]string{
		"Content-Type":        "application/json",
		"User-Agent":          "Kingdee/Golang WebApi SDK (compatible: K3Cloud 7.x)",
		"kdservice-sessionid": api.KDSVCSessionId,
	}
	if api.UserAgent != "" {
		res["User-Agent"] = api.UserAgent
	}
	maps.Copy(res, more)
	return res
}

func (api *API) BaseRequest(serviceName string, header map[string]string, params any) (res []byte, err error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", api.GetRequestURL(serviceName), bytes.NewBuffer(jsonData))
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
