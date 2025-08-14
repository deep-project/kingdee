package models

type MethodFileInfo struct {
	Size  int64  `json:"size"`  // 文件大小
	Name  string `json:"name"`  // 文件名
	Bytes []byte `json:"bytes"` // 内容
}
