package client

import "fmt"

type API struct {
	baseURL string
}

func NewAPI(domain string) *API {
	return &API{
		baseURL: fmt.Sprintf("https://%s/K3Cloud/", domain),
	}
}

func NewAPIByBaseURL(baseURL string) *API {
	return &API{
		baseURL: baseURL,
	}
}
