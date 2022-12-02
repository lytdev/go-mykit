package gconv

import (
	"strings"
)

// TagConvMap /** 结构体的tag(tag的多个kay-value是使用;隔开的)转map
func TagStrConvMap(tagStr string) map[string]string {
	resultData := make(map[string]string)
	splitData := strings.Split(tagStr, ";")
	for _, v := range splitData {
		key, val, ok := strings.Cut(v, ":")
		if ok {
			resultData[key] = val
		}
	}
	return resultData
}
