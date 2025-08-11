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
