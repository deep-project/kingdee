package methods

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/tidwall/gjson"
)

// 分组信息查询配置
type QueryGroupInfoOptions struct {
	GroupFieldKey string   `json:"groupFieldKey"` // 分组字段Key，字符串类型（必录） 注（不填时取默认，无默认，取第一个分组）
	GroupPkIds    []string `json:"groupPkIds"`    // 分组内码 (使用分组内码时必录，分组内码和单据内码同时录时，分组内码优先)
	Ids           []string `json:"ids"`           // 单据内码集合
}

// 通用分组信息查询
func (m *Methods) QueryGroupInfo(formId string, opt QueryGroupInfoOptions) (_ []byte, err error) {
	_opt := map[string]any{"FormId": formId}
	if opt.GroupFieldKey != "" {
		_opt["GroupFieldKey"] = opt.GroupFieldKey
	}
	if len(opt.GroupPkIds) > 0 {
		_opt["GroupPkIds"] = strings.Join(opt.GroupPkIds, ",")
	}
	if len(opt.Ids) > 0 {
		_opt["Ids"] = strings.Join(opt.Ids, ",")
	}
	b, err := m.client.QueryGroupInfo(formId, _opt)
	if err != nil {
		return
	}
	var (
		str       = string(b)
		isSuccess = gjson.Get(str, "Result.ResponseStatus.IsSuccess").Bool()
		_errors   = gjson.Get(str, "Result.ResponseStatus.Errors").String()
		data      = gjson.Get(str, "Result.NeedReturnData").String()
	)
	if !isSuccess {
		return nil, errors.New(_errors)
	}
	if data == "" {
		return nil, errors.New("获取到的数据为空")
	}
	return []byte(data), nil
}

func QueryGroupInfo[T any](m *Methods, formId string, opt QueryGroupInfoOptions) (res []T, err error) {
	b, err := m.QueryGroupInfo(formId, opt)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &res)
	return
}
