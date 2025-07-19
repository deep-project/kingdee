package client

import (
	"errors"

	"github.com/deep-project/kingdee/consts"
	"github.com/deep-project/kingdee/models"
	"github.com/tidwall/gjson"
)

type Client struct {
	Handler *Handler
}

func NewClient(api *API, login Login, options *Options) (cli *Client, err error) {
	if api == nil {
		return nil, errors.New("API undefined")
	}
	if login == nil {
		return nil, errors.New("Login undefined")
	}
	if options == nil {
		return nil, errors.New("Options undefined")
	}
	cli = &Client{Handler: &Handler{API: api, Login: login, Options: options}}
	err = cli.Handler.RefreshAPISessionID()
	cli.Handler.RunRefreshAPISessionIdJob()
	return
}

// 查看
func (c *Client) View(formid string, data models.ViewParams) (*models.ViewResult, error) {
	b, err := c.Handler.Call(consts.View_API, map[string]any{"formId": formid, "data": data})
	if err != nil {
		return nil, err
	}
	return &models.ViewResult{
		ResponseStatus: getResultResponseStatus(b),
		Result:         gjson.Get(string(b), "Result.Result").String(),
		Raw:            b,
	}, nil
}
