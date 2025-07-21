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
