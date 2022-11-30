package gnum

import (
	"strconv"
	"strings"
)

// NumFillZero 数字转字符串,位数不够的前面补0
func NumFillZero(n, l int) string {
	numStr := strconv.Itoa(n)
	nl := len(numStr)
	if nl >= l {
		return numStr
	}
	sb := strings.Builder{}
	for i := 0; i < (l - nl); i++ {
		sb.WriteString("0")
	}
	sb.WriteString(numStr)
	return sb.String()
}
