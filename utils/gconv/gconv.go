/*
 * @Author       : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @Date         : 2022-10-11 17:47:59
 * @LastEditors  : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @FilePath     : \go-myexcel\utils\gconv\gconv.go
 * @Description  :
 * Copyright (c) 2022 by 刘元涛 snoopy_718@mails.ccnu.edu.cn, All Rights Reserved.
 */
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
