package glist

import "fmt"

// FindIndex 查询泛型在切片里的位置
func FindIndex[T any](s []T, compare func(T) bool) (int, bool) {
	for i, e := range s {
		if compare(e) {
			return i, true
		}
	}
	return -1, false
}

// FindLastIndexOf 查询泛型在切片里最后一次的位置
func FindLastIndexOf[T any](collection []T, predicate func(item T) bool) (int, bool) {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return i, true
		}
	}
	return -1, false
}

// FindStrIndex 查询字符串在切片里的位置
func FindStrIndex(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// FindIntIndex 查询数字在切片里的位置
func FindIntIndex(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// RemoveFirstItem 从切片中移除第一个匹配元素
func RemoveFirstItem[T any](s []T, compare func(T) bool) []T {
	r := make([]T, 0)
	idx, ok := FindIndex(s, compare)
	if ok {
		if len(s[:idx]) > 0 {
			r = append(r, s[:idx]...)
		}
		if len(s[idx+1:]) > 0 {
			r = append(r, s[idx+1:]...)
		}
		return r
	}
	return s
}

// RemoveAllItem 从切片中移除所有匹配的元素
func RemoveAllItem[T any](s []T, compare func(T) bool) []T {
	r := s[:]
	for {
		idx, ok := FindIndex(r, compare)
		if ok {
			r = append(r[:idx], r[idx+1:]...)
		} else {
			return r
		}
	}
}

// SliceIntersect 2个切片的交集
func SliceIntersect[T any](s1 []T, s2 []T, compare func(T, T) bool) []T {
	s := make([]T, 0)
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if compare(v1, v2) {
				s = append(s, v1)
			}
		}
	}
	return s
}

// SliceDiff 2个切片的差集
func SliceDiff[T any](s1 []T, s2 []T, compare func(T, T) bool) []T {
	s := make([]T, 0)
	for _, v1 := range s1 {
		_, ok := FindIndex(s2, func(v2 T) bool {
			return compare(v1, v2)
		})
		if !ok {
			s = append(s, v1)
		}
	}
	return s
}

// DistinctStr 字符串数组去重
func DistinctStr(slice []string) []string {
	result := make([]string, 0)
	for _, temp := range slice {
		_, ok := FindStrIndex(result, temp)
		if !ok {
			result = append(result, temp)
		}
	}
	return result
}

// DistinctInt 整型数组去重
func DistinctInt(slice []int) []int {
	result := make([]int, 0)
	for _, temp := range slice {
		_, ok := FindIntIndex(result, temp)
		if !ok {
			result = append(result, temp)
		}
	}
	return result
}

// DistinctItem 数组去重
func DistinctItem[T any](slice []T, compare func(T, T) bool) []T {
	result := make([]T, 0)
	for _, temp := range slice {
		_, ok := FindIndex(result, func(v2 T) bool {
			return compare(temp, v2)
		})
		if !ok {
			result = append(result, temp)
		}
	}
	return result
}

// Min 返回切片的最小值
func Min[T Ordered](collection []T) T {
	var min T

	if len(collection) == 0 {
		return min
	}

	min = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item < min {
			min = item
		}
	}

	return min
}

// MinBy 根据给定的比较函数,返回切片的最小值
func MinBy[T any](collection []T, comparison func(a T, b T) bool) T {
	var min T

	if len(collection) == 0 {
		return min
	}

	min = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if comparison(item, min) {
			min = item
		}
	}

	return min
}

// Max 返回切片的最大值
func Max[T Ordered](collection []T) T {
	var max T

	if len(collection) == 0 {
		return max
	}

	max = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item > max {
			max = item
		}
	}

	return max
}

// MaxBy 根据给定的比较函数,返回切片的最大值
func MaxBy[T any](collection []T, comparison func(a T, b T) bool) T {
	var max T

	if len(collection) == 0 {
		return max
	}

	max = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if comparison(item, max) {
			max = item
		}
	}

	return max
}

// Last returns the last element of a collection or error if empty.
func Last[T any](collection []T) (T, error) {
	length := len(collection)

	if length == 0 {
		var t T
		return t, fmt.Errorf("last: cannot extract the last element of an empty slice")
	}

	return collection[length-1], nil
}

// Nth returns the element at index `nth` of collection. If `nth` is negative, the nth element
// from the end is returned. An error is returned when nth is out of slice bounds.
func Nth[T any, N Integer](collection []T, nth N) (T, error) {
	n := int(nth)
	l := len(collection)
	if n >= l || -n > l {
		var t T
		return t, fmt.Errorf("nth: %d out of slice bounds", n)
	}

	if n >= 0 {
		return collection[n], nil
	}
	return collection[l+n], nil
}
