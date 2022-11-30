package grandom

import (
	"math/rand"
	"time"
)

func init() {
	//随机种子
	rand.Seed(time.Now().UnixNano())
}

// RandInt 生成随机整数
func RandInt(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// RandomIntRange 生成一定范围内的随机整数
func RandomIntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomFloatRange 生成一定范围内的随机浮点数
func RandomFloatRange(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	// 到这里确保 max>=min 并且二者一定是正数
	ret := min + rand.Float64()*(max-min)
	return ret
}

// RandIntSlice 生成随机整数切片
func RandIntSlice(l int, max int) []int {
	slice := make([]int, l)
	for i := 0; i < l; i++ {
		slice[i] = RandInt(max)
	}
	return slice
}
