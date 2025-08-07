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
func (m *Method) IsLogin() (bool, error) {
	return true, nil
}
