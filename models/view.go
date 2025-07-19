package models

type ViewParams struct {
	CreateOrgId string `json:",omitempty"` // 创建者组织内码（非必录）
	Number      string `json:",omitempty"` // 单据编码，字符串类型（使用编码时必录）
	Id          int    `json:",omitempty"` // 表单内码（使用内码时必录）
	IsSortBySeq bool   `json:",omitempty"` // 单据体是否按序号排序，默认false
}

type ViewResult struct {
	ResponseStatus *ResponseStatus // 响应状态
	Result         string          // 单据内容，json字符串
	Raw            []byte          // 原始响应内容
}

// 查询基础数据模型
// 可以根据基础模型扩展自己的业务模型
type ViewBaseData struct {
	Id             int    // 内码id
	FFormId        string // 业务对象表单Id
	BillNo         string // 单据编号
	DocumentStatus string // 单据状态
}
