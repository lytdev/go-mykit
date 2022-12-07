package gnum

import (
	"fmt"
	"github.com/lytdev/go-mykit/gstr"
	"strconv"
	"strings"
)

// DivWithInt 两个整数相除,保留小数
func DivWithInt(n1, n2, p int) float64 {
	r, _ := strconv.ParseFloat(fmt.Sprintf(DecimalPointFormat(p), float64(n1)/float64(n2)), 64)
	return r
}

// DivWithFloat 两个float相除,保留小数
func DivWithFloat(n1, n2, float, p int) float64 {
	r, _ := strconv.ParseFloat(fmt.Sprintf(DecimalPointFormat(p), float64(n1)/float64(n2)), 64)
	return r
}

// DecimalPointFormat 获取小数点的格式化
func DecimalPointFormat(p int) string {
	if p == 0 {
		p = 2
	}
	return gstr.StrConcat("%.", strconv.Itoa(p), "f")
}

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

// NumMulti 是否一个数字是否是另一个数字的整数倍
func NumMulti(n1, n2 int) bool {
	return (n1 % n2) == 0
}
