/*
 * @Author       : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @Date         : 2022-10-13 15:39:47
 * @LastEditors  : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @FilePath     : \go-myexcel\gformt\dateformat.go
 * @Description  :
 * Copyright (c) 2022 by 刘元涛 snoopy_718@mails.ccnu.edu.cn, All Rights Reserved.
 */
package gformt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**
 * @Description : 提取日期为统一格式
 * @param        {string} numStr
 * @return       {*}
 * @Date        : 2022-10-12 19:21:33
 */
func GetFormatDateStr(numStr string) (string, error) {
	if len(numStr) == 0 {
		return "", nil
	}
	formatDateStr := ""
	if matched, _ := regexp.MatchString("^(\\d{4}\u5e74{0,1})$", numStr); matched {
		numStr = strings.Replace(numStr, "年", "", 1)
		formatDateStr = numStr + "-01-01"
	} else if matched, _ := regexp.MatchString("^(\\d{4}[-/\\.\u5e74]{1}[0-9]{0,2}[-/\\.\u6708]{0,1}[0-9]{0,2}\u65e5{0,1})$", numStr); matched {
		tmp, err := handlerDateStr(numStr)
		if err != nil {
			return formatDateStr, err
		}
		formatDateStr = tmp
	} else if matched, _ := regexp.MatchString("^(\\d{6})$", numStr); matched {
		year := numStr[:4]
		month := numStr[4:]
		formatDateStr = year + "-" + month + "-01"
	} else if matched, _ := regexp.MatchString("^(\\d{8})$", numStr); matched {
		year := numStr[:4]
		month := numStr[4:6]
		day := numStr[6:]
		formatDateStr = year + "-" + month + "-" + day
	} else {
		return formatDateStr, fmt.Errorf("出版日期格式不正确:%s", numStr)
	}

	return formatDateStr, nil
}

func handlerDateStr(numStr string) (string, error) {
	extraArr := extraNumFromStr(numStr)
	if len(extraArr) == 0 {
		return "", nil
	}
	resultStr := ""
	switch len(extraArr) {
	case 1:
		resultStr = extraArr[0] + "-01-01"
	case 2:
		monthStr := extraArr[1]
		if month, _ := strconv.Atoi(monthStr); month > 12 {
			return "", errors.New("日期月份格式不正确:" + numStr)
		}
		resultStr = extraArr[0] + "-" + standByNum(monthStr) + "-01"
	case 3:
		monthStr := extraArr[1]
		if month, _ := strconv.Atoi(monthStr); month > 12 {
			return "", errors.New("日期月份格式不正确:" + numStr)
		}
		dayStr := extraArr[2]
		if day, _ := strconv.Atoi(dayStr); day > 12 {
			return "", errors.New("日期天份格式不正确:" + numStr)
		}
		resultStr = extraArr[0] + "-" + standByNum(monthStr) + "-" + standByNum(dayStr)
	default:
		fmt.Println("无效的输入！")
	}
	return resultStr, nil
}

/**
 * @Description : 为小于10的数字前面补0
 * @param        {string} numStr
 * @return       {*}
 * @Date        : 2022-10-12 18:52:16
 */
func standByNum(numStr string) string {
	if len(numStr) == 1 {
		num, _ := strconv.Atoi(numStr)
		if num < 10 {
			return "0" + numStr
		}
	}
	return numStr
}

/**
 * @Description : 从字符串中提取出日期数字
 * @param        {string} numStr
 * @return       {*}
 * @Date        : 2022-10-12 18:53:06
 */
func extraNumFromStr(numStr string) []string {
	reg, _ := regexp.Compile(`[^0-9]`)
	arr := make([]string, 0)
	for _, i := range strings.Split(reg.ReplaceAllString(numStr, ","), ",") {
		if i != "" {
			arr = append(arr, i)
		}
	}
	result := make([]string, len(arr))
	copy(result, arr)
	return result
}
