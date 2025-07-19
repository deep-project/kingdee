package client

import (
	"encoding/json"

	"github.com/deep-project/kingdee/models"
	"github.com/tidwall/gjson"
)

// 获取结果响应状态
func getResultResponseStatus(b []byte) *models.ResponseStatus {
	str := gjson.Get(string(b), "Result.ResponseStatus").String()
	var res models.ResponseStatus
	json.Unmarshal([]byte(str), &res)
	return &res
}
