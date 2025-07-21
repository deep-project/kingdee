package utils

import (
	"encoding/json"
	"maps"

	"github.com/deep-project/kingdee/models"
	"github.com/tidwall/gjson"
)

// 获取结果响应状态
func GetResultResponseStatus(raw []byte) models.ResponseStatus {
	str := gjson.Get(string(raw), "Result.ResponseStatus").String()
	var res models.ResponseStatus
	json.Unmarshal([]byte(str), &res)
	return res
}

// 获取值的指针
func Ptr[T any](v T) *T {
	return &v
}

// 结构体ExtraParams展开到根
// 用于处理入参的结构体模型
func StructWithExtraParamsToMap(input any) (map[string]any, error) {
	// Step 1: 将结构体转换为 map
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var baseMap map[string]any
	err = json.Unmarshal(jsonBytes, &baseMap)
	if err != nil {
		return nil, err
	}

	// Step 2: 检查 ExtraParams 是否存在并是 map[string]any
	if extra, ok := baseMap["ExtraParams"].(map[string]any); ok {
		// 合并 ExtraParams 到 baseMap
		maps.Copy(baseMap, extra)
		// 删除原来的 ExtraParams 字段
		delete(baseMap, "ExtraParams")
	}

	return baseMap, nil
}

// 判断是否是session过期
func IsSessionExpired(b []byte) bool {
	status := GetResultResponseStatus(b)
	return !status.IsSuccess && status.MsgCode == 1
}
