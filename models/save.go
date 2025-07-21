package models

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
