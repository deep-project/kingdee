package models

// 文件信息
type MethodsFileInfo struct {
	Size  int64  `json:"size"`  // 文件大小
	Name  string `json:"name"`  // 文件名
	Bytes []byte `json:"bytes"` // 内容
}

// 单据查询配置
type BillQueryOption struct {
	FormId          string        `json:"formId"`       // 业务对象表单Id（必录）
	FieldKeys       []string      `json:"fieldKeys"`    // 需查询的字段key集合
	FilterString    any           `json:"filterString"` // 过滤条件,数组或字符串格式
	Limit           int           `json:"limit"`        // 每轮查询的最大行数，最大10000
	QueryBeforeHook BillQueryHook // 查询前置钩子
	QueryAfterHook  BillQueryHook // 查询后置钩子
}

type BillQueryHook func(opt BillQueryOption, startRow, page int, current []map[string]any)
