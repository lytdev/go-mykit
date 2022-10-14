package gconv

import "strings"

/**
 * @Description : tag 转map
 * @param        {string} tagStr
 * @return       {*}
 * @Date        : 2022-10-11 17:24:08
 */
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
