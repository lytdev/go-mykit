package glist

import (
	"github.com/lytdev/go-mykit/gfun"
	"math/rand"
)

//https://github.com/samber/lo/slice.go

// Filter 遍历集合的元素,返回满足条件的元素
// @param: collection原切片
// @param: predicate执行计算的函数,返回true才满足过滤的条件
func Filter[T any](collection []T, predicate func(item T, index int) bool) []T {
	var result = make([]T, 0)
	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}
	return result
}

// Map 遍历切片并根据给定的函数将原切片值进行转换后合并为一个新的切片
// @param: collection原切片
// @param: iteratee执行计算的函数,返回计算后的数据后合并至新的切片
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	var result = make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

// FilterMap 遍历切片并根据给定的函数计算的结果,将原切片数据计算结果为true,计算后的数据合并至新的切片
// @param: collection原切片
// @param: callback执行计算的函数,返true才将计算的结果合并至新的切片
func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R {
	result := make([]R, 0)

	for i, item := range collection {
		if r, ok := callback(item, i); ok {
			result = append(result, r)
		}
	}

	return result
}

// FlatMap 操作切片并将其转换和展平为另一种类型的切片
// 创建一个扁平化(注：同阶数组)的切片,这个切片的值来自collection(切片)中的每一个值经过iteratee(迭代函数)处理后返回的结果,并且扁平化合并.
// @param: collection原切片
// @param: iteratee执行计算的函数,函数计算的结果是一个新的切片
func FlatMap[T any, R any](collection []T, iteratee func(item T, index int) []R) []R {
	result := make([]R, 0)

	for i, item := range collection {
		result = append(result, iteratee(item, i)...)
	}

	return result
}

// Reduce 将切片缩减为一个值,该值是通过累加器运行切片中每个元素的累积结果,其中每次连续调用都提供前一次调用的返回值.
// 压缩 collection(切片)为一个值,通过 iteratee(迭代函数)遍历 collection(切片)中的每个元素,每次返回的值会作为下一次迭代使用(注：作为iteratee(迭代函数)的第一个参数使用).
// @param: collection切片
// @param: accumulator执行计算的函数,agg:每一次计算的初始值,item:切片的元素,index:切片的元素索引
// @param: initial初始值
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}

	return initial
}

// ReduceRight 从尾向头开始reduce操作
// Play: https://go.dev/play/p/Fq3W70l7wXF
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i := len(collection) - 1; i >= 0; i-- {
		initial = accumulator(initial, collection[i], i)
	}

	return initial
}

// ForEach 遍历集合的元素并为每个元素调用iteratee函数
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// Times 调用函数iteratee指定次数,返回每次调用结果的数组.
func Times[T any](count int, iteratee func(index int) T) []T {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		result[i] = iteratee(i)
	}

	return result
}

// Uniq 返回数组的无重复版本,其中只保留每个元素的第一次出现.
func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// UniqBy 返回数组的无重复版本,其中只保留每个元素的第一次出现.它接受数组中每个元素调用的iteratee,以生成计算惟一性的条件.
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key := iteratee(item)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}

// GroupBy 返回由通过iteratee运行集合的每个元素的结果生成的键组成的对象
// @param: collection切片
// @param: accumulator执行计算的函数,根据函数的结果进行分组,函数的结果作为分组后map的key,value是计算结果相同的元素的切片
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)
		result[key] = append(result[key], item)
	}
	return result
}

// Chunk 返回一个元素数组,将其分成长度为size的小切片.如果切片不能被平均分割,最后一个块将是剩余的元素.
// @param: collection 切片
// @param: size 每个切片分组里数据的大小
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	chunksNum := len(collection) / size
	if len(collection)%size != 0 {
		chunksNum += 1
	}

	result := make([][]T, 0, chunksNum)

	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}
		result = append(result, collection[i*size:last])
	}

	return result
}

// PartitionBy 和GroupBy类似,返回拆分为组的元素切片.分组是通过iteratee运行切片的每个元素的结果生成的.
func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) K) [][]T {
	var result = make([][]T, 0)
	seen := map[K]int{}

	for _, item := range collection {
		key := iteratee(item)

		resultIndex, ok := seen[key]
		if !ok {
			resultIndex = len(result)
			seen[key] = resultIndex
			result = append(result, []T{})
		}

		result[resultIndex] = append(result[resultIndex], item)
	}

	return result

	// unordered:
	// groups := GroupBy[T, K](collection, iteratee)
	// return Values[K, []T](groups)
}

// Flatten 返回深度仅为一级的切片,即二维切片转换为一维切片
func Flatten[T any](collection [][]T) []T {
	totalLen := 0
	for i := range collection {
		totalLen += len(collection[i])
	}

	result := make([]T, 0, totalLen)
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

// Interleave 循环交替输入切片,并依次在索引处追加值到结果中
func Interleave[T any](collections ...[]T) []T {
	if len(collections) == 0 {
		return []T{}
	}

	maxSize := 0
	totalSize := 0
	for _, c := range collections {
		size := len(c)
		totalSize += size
		if size > maxSize {
			maxSize = size
		}
	}

	if maxSize == 0 {
		return []T{}
	}

	result := make([]T, totalSize)

	resultIdx := 0
	for i := 0; i < maxSize; i++ {
		for j := range collections {
			if len(collections[j])-1 < i {
				continue
			}

			result[resultIdx] = collections[j][i]
			resultIdx++
		}
	}

	return result
}

// Shuffle 创建一个被打乱值的集合.使用Fisher-Yates shuffle算法。
func Shuffle[T any](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}

// Reverse 翻转切片
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}

	return collection
}

// Fill 用"初始值"填充数组元素,initial需要实现Cloneable接口
// https://go.dev/play/p/VwR34GzqEub
func Fill[T Cloneable[T]](collection []T, initial T) []T {
	result := make([]T, 0, len(collection))

	for range collection {
		result = append(result, initial.Clone())
	}

	return result
}

// Repeat 构建包含N个T元素副本的切片
// Play: https://go.dev/play/p/g3uHXbmc3b6
func Repeat[T Cloneable[T]](count int, initial T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, initial.Clone())
	}

	return result
}

// RepeatBy 根据函数predicate构建包含N个副本的切片.
// Play: https://go.dev/play/p/ozZLCtX_hNU
func RepeatBy[T any](count int, predicate func(index int) T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, predicate(i))
	}

	return result
}

// KeyBy 将切片转换为map，key为函数处理的结果,value为切片元数据.
// Play: https://go.dev/play/p/mdaClUAT-zZ
func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V {
	result := make(map[K]V, len(collection))

	for _, v := range collection {
		k := iteratee(v)
		result[k] = v
	}

	return result
}

// Associate 类似KeyBy,将切片转换为map,key为切片原数据,value为函数处理的结果.如果存在重复的,则忽略掉.
// Play: https://go.dev/play/p/WHa2CfMO3Lr
func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	result := make(map[K]V)

	for _, t := range collection {
		k, v := transform(t)
		result[k] = v
	}

	return result
}

// SliceToMap Associate的别名版本
// Play: https://go.dev/play/p/WHa2CfMO3Lr
func SliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	return Associate(collection, transform)
}

// Drop 从开头丢弃n个元素
// Play: https://go.dev/play/p/JswS7vXRJP2
func Drop[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return make([]T, 0)
	}

	result := make([]T, 0, len(collection)-n)
	return append(result, collection[n:]...)
}

// DropRight 从结尾丢弃n个元素.
// Play: https://go.dev/play/p/GG0nXkSJJa3
func DropRight[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return []T{}
	}

	result := make([]T, 0, len(collection)-n)
	return append(result, collection[:len(collection)-n]...)
}

// DropWhile 类似于Filter,丢弃满足条件的索引前面的数据
// Play: https://go.dev/play/p/7gBPYw2IK16
func DropWhile[T any](collection []T, predicate func(item T) bool) []T {
	i := 0
	for ; i < len(collection); i++ {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, 0, len(collection)-i)
	return append(result, collection[i:]...)
}

// DropRightWhile DropWhile的尾部丢弃
// Play: https://go.dev/play/p/3-n71oEC0Hz
func DropRightWhile[T any](collection []T, predicate func(item T) bool) []T {
	i := len(collection) - 1
	for ; i >= 0; i-- {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, 0, i+1)
	return append(result, collection[:i+1]...)
}

// Reject 是Filter的相反切片,排除掉返回结果为true的元素
// Play: https://go.dev/play/p/YkLMODy1WEL
func Reject[V any](collection []V, predicate func(item V, index int) bool) []V {
	var result = make([]V, 0)

	for i, item := range collection {
		if !predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// Count 统计元素出现的次数
// Play: https://go.dev/play/p/Y3FlK54yveC
func Count[T comparable](collection []T, value T) (count int) {
	for _, item := range collection {
		if item == value {
			count++
		}
	}

	return count
}

// CountBy 统计满足条件元素出现的次数
// Play: https://go.dev/play/p/ByQbNYQQi4X
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	for _, item := range collection {
		if predicate(item) {
			count++
		}
	}

	return count
}

// CountValues 统计切片里元素出现的次数,返回的map,key为切片的元素,value为次数
// Play: https://go.dev/play/p/-p-PyLT4dfy
func CountValues[T comparable](collection []T) map[T]int {
	result := make(map[T]int)

	for _, item := range collection {
		result[item]++
	}

	return result
}

// CountValuesBy 统计切片里元素满足条件的次数,返回的map,key为函数的结果,value为满足的元素个数
// Play: https://go.dev/play/p/2U0dG1SnOmS
func CountValuesBy[T any, U comparable](collection []T, mapper func(item T) U) map[U]int {
	result := make(map[U]int)

	for _, item := range collection {
		result[mapper(item)]++
	}

	return result
}

// Subset 从指定位置截取切片的指定长度
// @param: collection 原切片数据
// @param: offset 开始位置,包含
// @param: length 截取长度
// Play: https://go.dev/play/p/tOQu1GhFcog
func Subset[T any](collection []T, offset int, length uint) []T {
	size := len(collection)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset > size {
		return []T{}
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	return collection[offset : offset+int(length)]
}

// Slice 从开始位置和结束位置截取切片
// @param: collection 原切片数据
// @param: start 开始位置,包含
// @param: end 结束位置,不包含
// Play: https://go.dev/play/p/8XWYhfMMA1h
func Slice[T any](collection []T, start int, end int) []T {
	size := len(collection)

	if start >= end {
		return []T{}
	}

	if start > size {
		start = size
	}
	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}
	if end < 0 {
		end = 0
	}

	return collection[start:end]
}

// Replace 替换切片里的数据
// @param: collection 原切片数据
// @param: old 待替换的旧数据
// @param: new 替换的新数据
// @param: n 替换的次数
// Play: https://go.dev/play/p/XfPzmf9gql6
func Replace[T comparable](collection []T, old T, new T, n int) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll 替换切片里所有匹配的数据
// @param: collection 原切片数据
// @param: old 待替换的旧数据
// @param: new 替换的新数据
// Play: https://go.dev/play/p/a9xZFUHfYcV
func ReplaceAll[T comparable](collection []T, old T, new T) []T {
	return Replace(collection, old, new, -1)
}

// Compact 清除切片空值
// Play: https://go.dev/play/p/tXiy-iK6PAc
func Compact[T comparable](collection []T) []T {
	var zero T

	result := make([]T, 0)

	for _, item := range collection {
		if item != zero {
			result = append(result, item)
		}
	}

	return result
}

// IsSorted 检查切片是否是已排序的
// Play: https://go.dev/play/p/mc3qR-t4mcx
func IsSorted[T Ordered](collection []T) bool {
	for i := 1; i < len(collection); i++ {
		if collection[i-1] > collection[i] {
			return false
		}
	}

	return true
}

// IsSortedByKey 检查一个切片是否按iteratee函数的结果排序.
// Play: https://go.dev/play/p/wiG6XyBBu49
func IsSortedByKey[T any, K Ordered](collection []T, iteratee func(item T) K) bool {
	size := len(collection)

	for i := 0; i < size-1; i++ {
		if iteratee(collection[i]) > iteratee(collection[i+1]) {
			return false
		}
	}

	return true
}

// Range 根据给定的长度创建切片
// Play: https://go.dev/play/p/0r6VimXAi9H
func Range(elementNum int) []int {
	length := gfun.If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]int, length)
	step := gfun.If(elementNum < 0, -1).Else(1)
	for i, j := 0, 0; i < length; i, j = i+1, j+step {
		result[i] = j
	}
	return result
}

// RangeFrom 根据给定的其实数值和长度创建切片
// Play: https://go.dev/play/p/0r6VimXAi9H
func RangeFrom[T Integer | Float](start T, elementNum int) []T {
	length := gfun.If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]T, length)
	step := gfun.If(elementNum < 0, -1).Else(1)
	for i, j := 0, start; i < length; i, j = i+1, j+T(step) {
		result[i] = j
	}
	return result
}

// RangeWithStep 给定开始值和结束值,按照步长生成切片
// @param: start 开始值
// @param: end 结束值
// @param: step 步长
func RangeWithStep[T Integer | Float](start, end, step T) []T {
	var result = make([]T, 0)
	if start == end || step == 0 {
		return result
	}
	if start < end {
		if step < 0 {
			return result
		}
		for i := start; i < end; i += step {
			result = append(result, i)
		}
		return result
	}
	if step > 0 {
		return result
	}
	for i := start; i > end; i += step {
		result = append(result, i)
	}
	return result
}

// Sum 对切片元素的值进行累计
// Play: https://go.dev/play/p/upfeJVqs4Bt
func Sum[T Float | Integer | Complex](collection []T) T {
	var sum T = 0
	for _, val := range collection {
		sum += val
	}
	return sum
}

// SumBy 对切片的元素进行计算,计算的结果进行累加
// Play: https://go.dev/play/p/Dz_a_7jN_ca
func SumBy[T any, R Float | Integer | Complex](collection []T, iteratee func(item T) R) R {
	var sum R = 0
	for _, item := range collection {
		sum = sum + iteratee(item)
	}
	return sum
}
