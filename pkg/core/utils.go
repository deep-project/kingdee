package core

import (
	"github.com/tidwall/gjson"
)

// 判断是否是session过期
func IsSessionExpired(b []byte) bool {
	str := string(b)
	IsSuccess := gjson.Get(str, "Result.ResponseStatus.IsSuccess").Bool()
	MsgCode := gjson.Get(str, "Result.ResponseStatus.MsgCode").Int()
	return !IsSuccess && MsgCode == 1
}
