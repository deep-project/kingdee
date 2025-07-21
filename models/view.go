package models

type ViewParams struct {
	CreateOrgId *string `json:",omitempty"` // 创建者组织内码（非必录）
	Number      *string `json:",omitempty"` // 单据编码，字符串类型（使用编码时必录）
	Id          *int    `json:",omitempty"` // 表单内码（使用内码时必录）
	IsSortBySeq *bool   `json:",omitempty"` // 单据体是否按序号排序，默认false

	// 额外参数，用来补充未来新增查询条件（非必录）
	ExtraParams map[string]any `json:",omitempty"`
}

// 查询基础数据模型
// 可以根据基础模型扩展自己的业务模型
type ViewResultBaseData struct {
	Id              int         // 内码id
	FFormId         string      // 业务对象表单Id
	BillNo          string      // 单据编号
	DocumentStatus  string      // 单据状态
	StockOrgId_Id   int         // 库存组织内码id
	StockOrgId      StockOrgId  // 库存组织内码
	Date            string      // 日期
	FBillTypeID_Id  string      // 单据类型内码id
	FBillTypeID     FBillTypeID // 单据类型
	OwnerTypeIdHead string      // 业务对象头部所有者类型
	OwnerIdHead_Id  int         // 业务对象头部所有者内码id
}
