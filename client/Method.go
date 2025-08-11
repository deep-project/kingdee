package client

// TODO 封装的各种上层方法
type Method struct {
	client  *Client
	handler *Handler
}

func NewMethod(c *Client, h *Handler) *Method {
	return &Method{client: c, handler: h}
}

// 判断是否已经登录
// 目前先通过刷新sessionID的方式验证
// 尚不清楚获取当前用户登录信息的API接口，否则可以替换
func (m *Method) IsLogin() (bool, error) {
	if err := m.handler.RefreshSessionID(); err != nil {
		return false, err
	}
	return true, nil
}
