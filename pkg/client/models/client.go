package models

// 响应状态
type ResponseStatus struct {
	ErrorCode       int
	IsSuccess       bool
	Errors          []ResponseStatusError
	SuccessEntitys  []ResponseStatusSuccessEntitys
	SuccessMessages []ResponseStatusSuccessMessages

	// MsgCode的含义
	// 0：默认  1：上下文丢失 2：没有权限 3：操作标识为空 4：异常 5：单据标识为空 6：数据库操作失败 7：许可错误 8：参数错误
	// 9：指定字段/值不存在 10：未找到对应数据 11：验证失败 12：不可操作 13：网控冲突 14：调用限制 15：禁止管理员登录
	MsgCode int
}

// 响应状态错误信息
type ResponseStatusError struct {
	FieldName string
	Message   string
	DIndex    int // 数据索引
}

// 响应状态成功实体信息
type ResponseStatusSuccessEntitys struct {
	Id     int
	Number string
	DIndex int // 数据索引
}

// 响应状态成功消息
type ResponseStatusSuccessMessages struct {
	FieldName string
	Message   string
	DIndex    int // 数据索引
}

// 库存组织内码
// 当前业务操作所涉及的“库存组织”的唯一标识（内码）
// 用来承载库存相关的业务数据
// 例如：仓库库存、实收数量等
type StockOrgId struct {
	Id                int               // 内码id
	MultiLanguageText MultiLanguageText // 多语言文本
	Name              Name              // 名称
	Number            int               // 值
}

// 多语言文本
type MultiLanguageText struct {
	PkId     int    // 主键id
	LocaleId int    // 语言id
	Name     string // 名称
}

// 名称
type Name struct {
	Key   int    // 语言id
	Value string // 值
}

type FBillTypeID struct {
	Id                int               // 内码id
	MultiLanguageText MultiLanguageText // 多语言文本
	Name              Name              // 名称
	Number            string            // 值
}

type SaveParams struct {
	NeedUpDateFields      *[]string `json:",omitempty"` // 需要更新的字段，数组类型，格式：[key1,key2,...] （非必录）注（更新字段时Model数据包中必须设置内码，若更新单据体字段还需设置分录内码）
	NeedReturnFields      *[]string `json:",omitempty"` // 需返回结果的字段集合，数组类型，格式：[key,entitykey.key,...]（非必录） 注（返回单据体字段格式：entitykey.key）
	IsDeleteEntry         *bool     `json:",omitempty"` // 是否删除已存在的分录，布尔类型，默认true（非必录）
	SubSystemId           *string   `json:",omitempty"` // 表单所在的子系统内码，字符串类型（非必录）
	IsVerifyBaseDataField *bool     `json:",omitempty"` // 是否验证所有的基础资料有效性，布尔类，默认false（非必录）
	IsEntryBatchFill      *bool     `json:",omitempty"` // 是否批量填充分录，默认true（非必录）
	ValidateFlag          *bool     `json:",omitempty"` // 是否验证数据合法性标志，布尔类型，默认true（非必录）注（设为false时不对数据合法性进行校验）
	NumberSearch          *bool     `json:",omitempty"` // 是否用编码搜索基础资料，布尔类型，默认true（非必录）
	IsAutoAdjustField     *bool     `json:",omitempty"` // 是否自动调整JSON字段顺序，布尔类型，默认false（非必录）
	InterationFlags       *string   `json:",omitempty"` // 交互标志集合，字符串类型，分号分隔，格式："flag1;flag2;..."（非必录） 例如（允许负库存标识：STK_InvCheckResult）
	IgnoreInterationFlag  *bool     `json:",omitempty"` // 是否允许忽略交互，布尔类型，默认true（非必录）
	IsControlPrecision    *bool     `json:",omitempty"` // 是否控制精度，为true时对金额、单价和数量字段进行精度验证，默认false（非必录）
	ValidateRepeatJson    *bool     `json:",omitempty"` // 校验Json数据包是否重复传入，一旦重复传入，接口调用失败，默认false（非必录）
	Model                 any       `json:",omitempty"` // 表单数据包，JSON类型（必录）

	// 额外参数，用来补充未来新增查询条件（非必录）
	ExtraParams map[string]any `json:",omitempty"`
}

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
