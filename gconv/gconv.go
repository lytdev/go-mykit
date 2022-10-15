package gconv

import "strings"

// TagConvMap /** tag 转map
func TagConvMap(tagStr string) map[string]string {
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
