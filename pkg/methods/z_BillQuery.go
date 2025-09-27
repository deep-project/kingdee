package methods

import (
	"encoding/json"
	"strings"

	"github.com/deep-project/kingdee/pkg/utils"
)

// 单据查询配置
type BillQueryOptions struct {
	FormId          string        `json:"formId"`       // 业务对象表单Id（必录）
	FieldKeys       []string      `json:"fieldKeys"`    // 需查询的字段key集合
	FilterString    any           `json:"filterString"` // 过滤条件,数组或字符串格式
	Limit           int           `json:"limit"`        // 每轮查询的最大行数，最大10000
	QueryBeforeHook BillQueryHook // 查询前置钩子
	QueryAfterHook  BillQueryHook // 查询后置钩子
}

type BillQueryHook func(opt BillQueryOptions, startRow, page int, current []map[string]any)

// 通用单据列表查询
// 内部封装了翻页逻辑
func (m *Methods) BillQuery(opt BillQueryOptions) (_ []byte, err error) {
	list, err := utils.IterateQuery(opt.Limit, func(startRow, page int) (current []map[string]any, _err error) {
		if opt.QueryBeforeHook != nil {
			opt.QueryBeforeHook(opt, startRow, page, current)
		}
		b, _err := m.client.BillQuery(map[string]any{
			"FormId":       opt.FormId,
			"FieldKeys":    strings.Join(opt.FieldKeys, ","),
			"FilterString": opt.FilterString,
			"Limit":        opt.Limit,
			"StartRow":     startRow,
		})
		if _err != nil {
			return
		}
		if _err = json.Unmarshal(b, &current); err != nil { // 响应内容转成数组
			return
		}
		if opt.QueryAfterHook != nil {
			opt.QueryAfterHook(opt, startRow, page, current)
		}
		return
	})
	if err != nil {
		return
	}
	return json.Marshal(list)
}

func BillQuery[T any](m *Methods, opt BillQueryOptions) (res []T, err error) {
	b, err := m.BillQuery(opt)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &res)
	return
}
