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

// 查询迭代器
// 可以进行翻页查询
func (m *Method) IterateQuery(limit int, f func(startRow int) ([]map[string]any, error)) ([]map[string]any, error) {
	var allData = []map[string]any{}
	var page = 0
	for {
		page++
		startRow := (page - 1) * limit
		list, err := f(startRow)
		if err != nil {
			return allData, err
		}
		allData = append(allData, list...)
		if len(list) < limit {
			break // 最后一批
		}
	}
	return allData, nil
}
