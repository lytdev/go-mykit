package gdatetime

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// GetFormatDateStr /** 提取日期为统一格式
func GetFormatDateStr(numStr string) (string, error) {
	if len(numStr) == 0 {
		return "", nil
	}
	formatDateStr := ""
	if matched, _ := regexp.MatchString("^(\\d{4}\u5e74?)$", numStr); matched {
		//只有年的格式:2022(年)
		numStr = strings.Replace(numStr, "年", "", 1)
		formatDateStr = numStr + "-01-01 00:00:00"
	} else if matched, _ := regexp.MatchString("^(\\d{6})$", numStr); matched {
		//只有年月的格式:202206
		year := numStr[:4]
		month := numStr[4:]
		formatDateStr = year + "-" + month + "-01 00:00:00"
	} else if matched, _ := regexp.MatchString("^(\\d{8})$", numStr); matched {
		//只有年月日的格式:20220605
		year := numStr[:4]
		month := numStr[4:6]
		day := numStr[6:]
		formatDateStr = year + "-" + month + "-" + day + " 00:00:00"
	} else if matched, _ := regexp.MatchString("^(\\d{4}[-/.\u5e74]\\d{0,2}[-/.\u6708]?\\d{0,2}\u65e5?)$", numStr); matched {
		//处理年月日的格式yyyy-dd-MM
		tmp, err := handlerDateStr(numStr)
		if err != nil {
			return formatDateStr, err
		}
		formatDateStr = tmp
	} else if matched, _ := regexp.MatchString("^(\\d{4}[-/.\u5e74]\\d{1,2}[-/.\u6708]\\d{1,2}\u65e5?\\s*\\d{1,2}[:\u65f6]\\d{1,2}[:\u5206]?\\d{0,2}\u79d2?[ APM]{0,3})$", numStr); matched {
		//处理年月日时分秒的格式yyyy-dd-MM hh:mm:ss
		tmp, err := handlerDateStr(numStr)
		if err != nil {
			return formatDateStr, err
		}
		formatDateStr = tmp
	} else {
		return formatDateStr, fmt.Errorf("出版日期格式不正确:%s", numStr)
	}

	return formatDateStr, nil
}

/**
 * @Description : 处理日期时间格式
 * @param        {string} numStr
 * @return       {*}
 * @Date        : 2022-10-14 09:57:49
 */
func handlerDateStr(numStr string) (string, error) {
	extraArr := extraNumFromStr(numStr)
	if len(extraArr) == 0 {
		return "", nil
	}
	resultStr := ""
	switch len(extraArr) {
	case 1:
		resultStr = extraArr[0] + "-01-01 00:00:00"
	case 2:
		monthStr := extraArr[1]
		if month, _ := strconv.Atoi(monthStr); month > 12 {
			return "", errors.New("日期格式月份不正确:" + numStr)
		}
		resultStr = extraArr[0] + "-" + standByNum(monthStr) + "-01 00:00:00"
	case 3:
		monthStr := extraArr[1]
		if month, _ := strconv.Atoi(monthStr); month > 12 {
			return "", errors.New("日期格式月份不正确:" + numStr)
		}
		dayStr := extraArr[2]
		if day, _ := strconv.Atoi(dayStr); day > 31 {
			return "", errors.New("日期格式天份不正确:" + numStr)
		}
		resultStr = extraArr[0] + "-" + standByNum(monthStr) + "-" + standByNum(dayStr) + " 00:00:00"
	case 5:
		monthStr := extraArr[1]
		if month, _ := strconv.Atoi(monthStr); month > 12 {
			return "", errors.New("日期格式月份不正确:" + numStr)
		}
		dayStr := extraArr[2]
		if day, _ := strconv.Atoi(dayStr); day > 31 {
			return "", errors.New("日期格式天份不正确:" + numStr)
		}
		hourStr := extraArr[3]
		if hour, _ := strconv.Atoi(hourStr); hour > 24 {
			return "", errors.New("时间格式小时不正确:" + numStr)
		}
		minuStr := extraArr[4]
		if minu, _ := strconv.Atoi(minuStr); minu > 60 {
			return "", errors.New("时间格式分钟不正确:" + numStr)
		}
		resultStr = extraArr[0] + "-" + standByNum(monthStr) + "-" + standByNum(dayStr) + " " + standByNum(hourStr) + ":" + standByNum(minuStr) + ":00"
	case 6:
		monthStr := extraArr[1]
		if month, _ := strconv.Atoi(monthStr); month > 12 {
			return "", errors.New("日期月份格式不正确:" + numStr)
		}
		dayStr := extraArr[2]
		if day, _ := strconv.Atoi(dayStr); day > 31 {
			return "", errors.New("日期天份格式不正确:" + numStr)
		}
		hourStr := extraArr[3]
		if hour, _ := strconv.Atoi(hourStr); hour > 24 {
			return "", errors.New("时间格式小时不正确:" + numStr)
		}
		minuStr := extraArr[4]
		if minu, _ := strconv.Atoi(minuStr); minu > 60 {
			return "", errors.New("时间格式分钟不正确:" + numStr)
		}
		secStr := extraArr[5]
		if sec, _ := strconv.Atoi(secStr); sec > 60 {
			return "", errors.New("时间格式秒不正确:" + numStr)
		}
		resultStr = extraArr[0] + "-" + standByNum(monthStr) + "-" + standByNum(dayStr) + " " + standByNum(hourStr) + ":" + standByNum(minuStr) + ":" + standByNum(secStr)
	default:
		return "", errors.New("无效的时间格式:" + numStr)
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
