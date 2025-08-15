package client

import (
	"bytes"
	"encoding/base64"
	"errors"

	"github.com/deep-project/kingdee/pkg/client/models"
	"github.com/deep-project/kingdee/pkg/core"
	"github.com/tidwall/gjson"
)

// 封装的各种上层方法
type Methods struct {
	client *Client
	core   *core.Core
}

func NewMethods(cli *Client, c *core.Core) *Methods {
	return &Methods{client: cli, core: c}
}

// TODO
// 判断是否已经登录
func (m *Methods) IsLogin() (bool, error) {
	return true, nil
}

// 下载附件
// 不同于client的原始方法
// 此方法封装了轮询的下载逻辑，可以更便捷的下载
func (m *Methods) AttachmentDownLoad(fileId string) (*models.MethodsFileInfo, error) {
	var buf bytes.Buffer
	var StartIndex int64 = 0
	var FileSize int64 = 0
	var FileName = ""
	for {
		b, err := m.client.AttachmentDownLoad(map[string]any{"FileId": fileId, "StartIndex": StartIndex})
		if err != nil {
			return nil, err
		}
		str := string(b)
		IsSuccess := gjson.Get(str, "Result.ResponseStatus.IsSuccess").Bool()
		Errors := gjson.Get(str, "Result.ResponseStatus.Errors").String()
		IsLast := gjson.Get(str, "Result.IsLast").Bool()
		FilePart := gjson.Get(str, "Result.FilePart").String()
		StartIndex = gjson.Get(str, "Result.StartIndex").Int()
		FileSize = gjson.Get(str, "Result.FileSize").Int()
		FileName = gjson.Get(str, "Result.FileName").String()

		if !IsSuccess {
			return nil, errors.New(Errors)
		}
		decoded, err := base64.StdEncoding.DecodeString(FilePart)
		if err != nil {
			return nil, err
		}
		buf.Write(decoded)
		if IsLast {
			break
		}
	}

	return &models.MethodsFileInfo{Name: FileName, Size: FileSize, Bytes: buf.Bytes()}, nil
}
