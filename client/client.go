package client

import "errors"

type Options struct {
}

type Client struct {
	API     *API
	Login   Login
	Options *Options
}

func NewClient(api *API, login Login, options *Options) (*Client, error) {
	if api == nil {
		return nil, errors.New("API undefined")
	}
	if login == nil {
		return nil, errors.New("Login undefined")
	}
	if err := login.Refresh(api); err != nil {
		return nil, err
	}
	return &Client{Options: options}, nil
}
