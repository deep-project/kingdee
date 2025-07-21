package consts

const (
	GetDataCenterList_API   = "Kingdee.BOS.ServiceFacade.ServicesStub.Account.AccountService.GetDataCenterList.common.kdsvc" // 获取账套列表(获取数据中心列表)
	LoginByAppSecret_API    = "Kingdee.BOS.WebApi.ServicesStub.AuthService.LoginByAppSecret.common.kdsvc"                    // 通过appSecret登录
	LoginBySign_API         = "Kingdee.BOS.WebApi.ServicesStub.AuthService.LoginBySign.common.kdsvc"                         // 通过签名登录
	LoginByValidateUser_API = "Kingdee.BOS.WebApi.ServicesStub.AuthService.ValidateUser.common.kdsvc"                        // 通过用户名密码登录
	View_API                = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.View.common.kdsvc"                         // 查看
	Save_API                = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.save.common.kdsvc"                         // 保存
	BatchSave_API           = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.batchSave.common.kdsvc"                    // 批量保存
	FlexSave_API            = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.FlexSave.common.kdsvc"                     // 弹性域保存
	Draft_API               = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.draft.common.kdsvc"                        // 暂存
	Delete_API              = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.delete.common.kdsvc"                       // 删除
	Submit_API              = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Submit.common.kdsvc"                       // 提交
	Audit_API               = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Audit.common.kdsvc"                        // 审核
	UnAudit_API             = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.UnAudit.common.kdsvc"                      // 反审核
	Push_API                = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Push.common.kdsvc"                         // 下推
	CancelAssign_API        = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.CancelAssign.common.kdsvc"                 // 撤销（撤销服务）
	ExecuteBillQuery_API    = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.ExecuteBillQuery.common.kdsvc"             // 单据查询
	BillQuery_API           = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.BillQuery.common.kdsvc"                    // 单据查询(json)
	ExecuteOperation_API    = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.ExecuteOperation.common.kdsvc"             // 操作
	ExcuteOperation_API     = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.ExcuteOperation.common.kdsvc"              // 操作（疑似写错了，应该和 ExecuteOperation_API 作用相同）
	QueryBusinessInfo_API   = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.QueryBusinessInfo.common.kdsvc"            // 元数据查询（查询单据信息）
	WorkflowAudit_API       = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.WorkflowAudit.common.kdsvc"                // 工作流审批
	SwitchOrg_API           = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.SwitchOrg.common.kdsvc"                    // 切换组织
	AttachmentUpLoad_API    = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.AttachmentUpLoad.common.kdsvc"             // 上传附件
	AttachmentDownLoad_API  = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.AttachmentDownLoad.common.kdsvc"           // 下载附件
	Allocate_API            = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Allocate.common.kdsvc"                     // 分配
	CancelAllocate_API      = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.CancelAllocate.common.kdsvc"               // 取消分配
	Disassembly_API         = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Disassembly.common.kdsvc"                  // 拆单
	GroupSave_API           = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.GroupSave.common.kdsvc"                    // 分组保存
	QueryGroupInfo_API      = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.QueryGroupInfo.common.kdsvc"               // 分组信息查询
	GroupDelete_API         = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.GroupDelete.common.kdsvc"                  // 分组删除
	GetSysReportData_API    = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.GetSysReportData.common.kdsvc"             // 查询报表数据
	SendMsg_API             = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.SendMsg.common.kdsvc"                      // 发送消息
)
