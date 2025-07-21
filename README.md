# kingdee
kingdee sdk / 金蝶云星空SDK


## 使用 / Usage

```

import (
	"fmt"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/adapters"
	"github.com/deep-project/kingdee/utils"
)

func main() {

	cli, err := kingdee.New(kingdee.NewOptions("http://127.0.0.1:9010/K3Cloud/", &adapters.LoginBySign{
		AccountID:  "ACCOUNT_ID",
		Username:   "USER_NAME",
		AppID:      "APP_ID",
		AppSecret:  "APP_SECRET",
		LanguageID: "LANGUAGE_ID",
	}))

	raw, err := cli.View("STK_InStock", map[string]any{"Number": "CGRK00019"})
	if err != nil {
		return
	}
	status := utils.GetResultResponseStatus(raw)
	fmt.Println("status", status)
	fmt.Println("json", string(raw))
}
```

### 内置方法
```

// 获取账套列表(获取数据中心列表)
cli.GetDataCenterList() 
// 查看
cli.View(formid string, data any)
// 保存
cli.Save(formid string, data any)
// 批量保存
cli.BatchSave(formid string, data any)
// 弹性域保存
cli.FlexSave(formid string, data any)
// 暂存
cli.Draft(formid string, data any)
// 删除
cli.Delete(formid string, data any)
// 提交
cli.Submit(formid string, data any)
// 审核
cli.Audit(formid string, data any)
// 反审核
cli.UnAudit(formid string, data any)
// 下推
cli.Push(formid string, data any)
// 撤销
cli.CancelAssign(formid string, data any)
// 单据查询
cli.ExecuteBillQuery(data any)
// 单据查询(json)
cli.BillQuery(data any)
// 操作
cli.ExecuteOperation(formid string, operateNumber string, data any)
cli.ExcuteOperation(formid string, operateNumber string, data any)
// 元数据查询（查询单据信息）
cli.QueryBusinessInfo(formid string)
// 工作流审批
cli.WorkflowAudit(data any)
// 切换组织
cli.SwitchOrg(data any)
// 上传附件
cli.AttachmentUpLoad(data any)
// 下载附件
cli.AttachmentDownLoad(data any)
// 分配
cli.Allocate(formid string, data any)
// 取消分配
cli.CancelAllocate(formid string, data any)
// 拆单
cli.Disassembly(formid string, data any)
// 分组保存
cli.GroupSave(formid string, data any)
// 分组信息查询
cli.QueryGroupInfo(formid string, data any)
// 分组删除
cli.GroupDelete(formid string, data any)
// 查询报表数据
cli.GetSysReportData(formid string, data any)
// 发送消息
cli.SendMsg(data any)

```

### 支持三种登录方式

```
// 通过sign登录
&adapters.LoginBySign{
  AccountID:  "ACCOUNT_ID",
  Username:   "USER_NAME",
  AppID:      "APP_ID",
  AppSecret:  "APP_SECRET",
  LanguageID: "LANGUAGE_ID",
}

// 通过appSecret登录
&adapters.LoginByAppSecret{
  AccountID:  "ACCOUNT_ID",
  Username:   "USER_NAME",
  AppID:      "APP_ID",
  AppSecret:  "APP_SECRET",
  LanguageID: "LANGUAGE_ID",
}

// 通过账号密码登录
&adapters.LoginByValidateUser{
  AccountID:  "ACCOUNT_ID",
  Username:   "USER_NAME",
  Password:   "USER_PASSWORD",
  LanguageID: "LANGUAGE_ID",
}
```

## 依赖 / Dependencies
[tidwall/gjson](https://github.com/tidwall/gjson)