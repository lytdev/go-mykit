package grandom

import (
	"math/rand"
	"time"
)

// RandomInt 生成随机整数
func RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix()) //随机种子
	return rand.Intn(max-min) + min
}

// RandomFloat 生成随机浮点数
func RandomFloat(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	// 到这里确保 max>=min 并且二者一定是正数
	ret := min + rand.Float64()*(max-min)
	return ret
}
