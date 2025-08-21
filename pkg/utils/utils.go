package utils

import (
	"encoding/json"
	"maps"
)

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

// 查询迭代器
// 可以进行翻页查询
func IterateQuery[T map[string]any](limit int, f func(startRow, page int) ([]T, error)) ([]T, error) {
	var allData = []T{}
	var page = 0
	for {
		page++
		startRow := (page - 1) * limit
		list, err := f(startRow, page)
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
