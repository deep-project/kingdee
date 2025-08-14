package client

import (
	"github.com/deep-project/kingdee/consts"
)

type Client struct {
	Handler *Handler
	Methods *Methods
}

func NewClient(options Options) (cli *Client, err error) {
	handler, err := NewHandler(&options)
	if err != nil {
		return
	}
	cli = &Client{Handler: handler}
	cli.Methods = NewMethods(cli, handler)
	return
}

// 获取账套列表(获取数据中心列表)
func (c *Client) GetDataCenterList() (raw []byte, err error) {
	return c.Handler.Call(consts.GetDataCenterList_API, map[string]any{})
}

// TODO
func (c *Client) GetDataCenterListParsed() (raw []byte, err error) {
	return c.Handler.Call(consts.GetDataCenterList_API, map[string]any{})
}

// 查看
func (c *Client) View(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.View_API, map[string]any{"formId": formid, "data": data})
}

// 保存
func (c *Client) Save(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Save_API, map[string]any{"formId": formid, "data": data})
}

// 批量保存
func (c *Client) BatchSave(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.BatchSave_API, map[string]any{"formId": formid, "data": data})
}

// 弹性域保存
func (c *Client) FlexSave(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.FlexSave_API, map[string]any{"formId": formid, "data": data})
}

// 暂存
func (c *Client) Draft(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Draft_API, map[string]any{"formId": formid, "data": data})
}

// 删除
func (c *Client) Delete(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Delete_API, map[string]any{"formId": formid, "data": data})
}

// 提交
func (c *Client) Submit(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Submit_API, map[string]any{"formId": formid, "data": data})
}

// 审核
func (c *Client) Audit(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Audit_API, map[string]any{"formId": formid, "data": data})
}

// 反审核
func (c *Client) UnAudit(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.UnAudit_API, map[string]any{"formId": formid, "data": data})
}

// 下推
func (c *Client) Push(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Push_API, map[string]any{"formId": formid, "data": data})
}

// 撤销
func (c *Client) CancelAssign(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.CancelAssign_API, map[string]any{"formId": formid, "data": data})
}

// 单据查询
func (c *Client) ExecuteBillQuery(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.ExecuteBillQuery_API, map[string]any{"data": data})
}

// 单据查询(json)
func (c *Client) BillQuery(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.BillQuery_API, map[string]any{"data": data})
}

// 操作
// operateNumber = Cancel: 作废
// operateNumber = UnCancel: 反作废
// operateNumber = BillClose: 整单关闭
// operateNumber = BillUnClose: 整单反关闭
// operateNumber = ReqToRefundBill: 申请单退款
// operateNumber = StatusConvert: 开票结束
// operateNumber = PayableClose: 应付关闭
// operateNumber = PayableUnClose: 应付反关闭
// operateNumber = BackServiceUnAudit: 后台反审核
// operateNumber = Freeze: 业务冻结
// operateNumber = UnFreeze: 反业务冻结
// operateNumber = Terminate: 业务终止
// operateNumber = UnTerminate: 反业务终止
// operateNumber = Confirm: 确认
// operateNumber = UnComfirm: 反认
// operateNumber = StatusEnable: 状态启用
// operateNumber = StatusDisable: 状态停用
// operateNumber = Forbid: 禁用
// operateNumber = Enable: 反禁用
func (c *Client) ExecuteOperation(formid string, operateNumber string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.ExecuteOperation_API, map[string]any{"formId": formid, "opNumber": operateNumber, "data": data})
}

// 操作（别名）
// 文档中很多地方都是用了Excute这个单词，疑似是拼写错了
func (c *Client) ExcuteOperation(formid string, operateNumber string, data any) (raw []byte, err error) {
	return c.ExecuteOperation(formid, operateNumber, data)
}

// 元数据查询（查询单据信息）
// data示例：map[string]any{"FormId": "PUR_PurchaseOrder"}
func (c *Client) QueryBusinessInfo(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.QueryBusinessInfo_API, map[string]any{"data": data})
}

// 工作流审批
func (c *Client) WorkflowAudit(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.WorkflowAudit_API, map[string]any{"data": data})
}

// 切换组织
func (c *Client) SwitchOrg(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.SwitchOrg_API, map[string]any{"data": data})
}

// 上传附件
func (c *Client) AttachmentUpLoad(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.AttachmentUpLoad_API, map[string]any{"data": data})
}

// 下载附件
func (c *Client) AttachmentDownLoad(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.AttachmentDownLoad_API, map[string]any{"data": data})
}

// 分配
func (c *Client) Allocate(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Allocate_API, map[string]any{"formId": formid, "data": data})
}

// 取消分配
func (c *Client) CancelAllocate(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.CancelAllocate_API, map[string]any{"formId": formid, "data": data})
}

// 拆单
func (c *Client) Disassembly(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.Disassembly_API, map[string]any{"formId": formid, "data": data})
}

// 分组保存
func (c *Client) GroupSave(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.GroupSave_API, map[string]any{"formId": formid, "data": data})
}

// 分组信息查询
func (c *Client) QueryGroupInfo(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.QueryGroupInfo_API, map[string]any{"formId": formid, "data": data})
}

// 分组删除
func (c *Client) GroupDelete(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.GroupDelete_API, map[string]any{"formId": formid, "data": data})
}

// 查询报表数据
func (c *Client) GetSysReportData(formid string, data any) (raw []byte, err error) {
	return c.Handler.Call(consts.GetSysReportData_API, map[string]any{"formId": formid, "data": data})
}

// 发送消息
func (c *Client) SendMsg(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.SendMsg_API, map[string]any{"data": data})
}

// 获取即时库存(官方自定义版本)
// 支持版本： 补丁号：PT-146848 （即2020年8月13号之后的补丁）
// 位置在API文档的 供应量->库存管理->自定义API->即时库存（https://openapi.open.kingdee.com/ApiDoc）
// 此接口的说明地址 https://vip.kingdee.com/link/s/MU4cj
func (c *Client) GetInventoryData(data any) (raw []byte, err error) {
	return c.Handler.Call(consts.GetInventoryData_API, map[string]any{"data": data})
}
