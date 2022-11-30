package gnum

//https://github.com/dablelv/go-huge-util/blob/master/abs.go

// Example 1: AbsInt8(-5)
// -5 code value as below
// original code	1000,0101
// inverse code		1111,1010
// complement code	1111,1011
// Negative numbers are represented by complement code in memory.
// shifted = n >> 7 = (1111,1011) >> 7 = 1111,1111 = -1(10-base) (负数右移，左补1)
//               1111,1011
// n xor shifted =  ----------- = 0000,0100 = 4(10-base)
//               1111,1111
// (n ^ shifted) - shifted = 4 - (-1) = 5
//
// Example 2: AbsInt8(5)
// 5 code value as below
// original code 0000,0101
// Positive numbers are represented by original code in memory,
// and the XOR operation between positive numbers and 0 is equal to itself.
// shifted = n >> 7 = 0
//               0000,0101
// n xor shifted =  ----------- = 0000,0101 = 5(10-base)
//               0000,0000
// (n ^ shifted) - shifted = 5 - 0 = 5

// AbsInt8 数字int8的绝对值
func AbsInt8(n int8) int8 {
	shifted := n >> 7
	return (n ^ shifted) - shifted
}

// AbsInt16 数字int16的绝对值
func AbsInt16(n int16) int16 {
	shifted := n >> 15
	return (n ^ shifted) - shifted
}

// AbsInt32 数字int32的绝对值
func AbsInt32(n int32) int32 {
	shifted := n >> 31
	return (n ^ shifted) - shifted
}

// AbsInt64 数字int64的绝对值
func AbsInt64(n int64) int64 {
	shifted := n >> 63
	return (n ^ shifted) - shifted
}
