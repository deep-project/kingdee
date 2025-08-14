# kingdee Golang SDK
金蝶云星空 Golang SDK,通过webapi实现。在金蝶云星空8.x版本测试通过。

## 项目亮点 / Features
+ 内置30+操作方法
+ 可直接根据服务名请求
+ 支持三种登录方式
+ 支持定时刷新sessionid
+ 支持被动刷新sessionid
+ 可后置登录接口
+ 可使用池并发执行

## 使用 / Usage
```go
go get github.com/deep-project/kingdee
```
```go

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

### 内置基础方法
> 基本上内置了大部分快捷方法，如果有缺失，可以提pull,也可以告诉我加上。也可以看下面的“直接根据服务名称调用”
```go
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
cli.QueryBusinessInfo(data any)
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
// 获取即时库存(官方自定义版本)
cli.GetInventoryData(data any)

```
### 上层封装方法
```go
// 根据附件ID下载附件
// 内部实现了轮询逻辑
cli.Method.AttachmentDownLoad(fileId string)
```

#### 直接根据服务名称调用
>如果内置方法无法满足需求，亦可以直接通过服务名请求结果。
```go
cli.Handler.Call("Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.CancelAllocate.common.kdsvc" , map[string]any{"data": data})
```

### 支持三种登录方式
>官方推荐的是sign登录方式，更安全。
```go
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
### 支持定时刷新sessionid
>可设置每隔多久刷新一次sessionid，防止过期。默认15分钟（因为金蝶服务端默认20分钟过期），可以在配置设置。根据金蝶内部规则，如果20分钟内请求不闲置的话，session是会自动续期的。但以防万一，启用定时刷新更保险一些。
```go
options := kingdee.NewOptions("http://127.0.0.1:9010/K3Cloud/", &adapters.LoginBySign{....})
options.SetRefreshSessionIdInterval(30 * time.Minute) // 每30分钟刷新一次sessionid
cli, err := kingdee.New(options)

```


### 支持被动刷新sessionid
>如果访问接口发现已经过期，则刷新sessionid再次请求。默认重试1次，可以在配置设置。
```go
options := kingdee.NewOptions("http://127.0.0.1:9010/K3Cloud/", &adapters.LoginBySign{....})
options.SessionExpiredRetryCount(3) // 过期时重试3次
cli, err := kingdee.New(options)

```
### 全部配置项
```go
type Options struct {
	BaseURL                  string            // api请求基地址
	UserAgent                string            // api请求的UserAgent
	APIClientIdentity        string            // 客户端标识(如果系统设置了验证标识，则需要填写)
	RequestHeader            map[string]string // 额外的api请求头
	Login                    LoginInterface    // 登录接口
	RefreshSessionIdInterval time.Duration     // 定时刷新sessionID的时间间隔
	SessionExpiredRetryCount int               // session过期重试次数
}
```

### 可后置登录接口
>如果初始需要通过免登录的GetDataCenterList()接口获取账套列表，可以先不设置登录接口，拿到账套信息后再设置
```go
cli, err := kingdee.New(kingdee.NewOptions("http://127.0.0.1:9010/K3Cloud/", nil))

// 先获取账套信息
AccountInfo,err :=cli.GetDataCenterList()

// 再用完善的信息设置登录接口
cli.Handler.SetLogin(&adapters.LoginByValidateUser{
  AccountID:  "ACCOUNT_ID",
  Username:   "USER_NAME",
  Password:   "USER_PASSWORD",
  LanguageID: "LANGUAGE_ID",
})

// 进行其他操作
cli.View(...)
```

### 使用池并发执行
> 因为金蝶云星空的请求默认是同步模式，也就是说，一个session多次访问时，是同步请求，所以要实现并发执行，就要创建多个client携带不同的sessionid去请求。

##### 创建不同的客户端池
```go
import (
	"fmt"
	"os"
	"sync"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/pool"
)

func main() {

	// 创建5个客户端
	// 金蝶建议的并发是3-5个
	client_1, _ := kingdee.New(...)
	client_2, _ := kingdee.New(...)
	client_3, _ := kingdee.New(...)
	client_4, _ := kingdee.New(...)
	client_5, _ := kingdee.New(...)

	var p = pool.New([]*client.Client{client_1, client_2, client_3, client_4, client_5})
	var wg sync.WaitGroup

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 只有存在空闲client时，才会被Get到
			// 否则会停在这里等待
			client := p.Get()
			// 用完必须归还，否则资源耗尽程序卡死
			defer p.Put(client)
			// 调用金蝶保存接口
			client.Save(...)
		}()
	}
	wg.Wait()
}

```

##### 通过统一配置创建池
```go
import (
	"fmt"
	"os"
	"sync"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/client"
	"github.com/deep-project/kingdee/pool"
)

func main() {

	options := kingdee.NewOptions("http://127.0.0.1:9010/K3Cloud/", &adapters.LoginBySign{....})
	// 一次性创建5个配置相同的client池
	var p = pool.NewBySize(5, options)
	var wg sync.WaitGroup

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 只有存在空闲client时，才会被Get到
			// 否则会停在这里等待
			client := p.Get()
			// 用完必须归还，否则资源耗尽程序卡死
			defer p.Put(client)
			// 调用金蝶保存接口
			client.Save(...)
		}()
	}
	wg.Wait()
}

```

## 依赖 / Dependencies
[tidwall/gjson](https://github.com/tidwall/gjson)