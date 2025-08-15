package session

type Session struct {
	KDSVCSessionId string
	Context        SessionContext
}

type SessionContext struct {
	UserLocale     string // 用户语言 zh-CN
	SessionId      string // sessionid
	UserId         int
	UserName       string
	CustomName     string
	DataCenterName string
	UserToken      string
}
