package glist

// FindIndex 查询泛型在切片里的位置
func FindIndex[T any](s []T, compare func(t T) bool) (int, bool) {
	for i, e := range s {
		if compare(e) {
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
func RemoveFirstItem[T any](s []T, compare func(t T) bool) []T {
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
func RemoveAllItem[T any](s []T, compare func(t T) bool) []T {
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
func SliceIntersect[T any](s1 []T, s2 []T, compare func(t1, t2 T) bool) []T {
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
func SliceDiff[T any](s1 []T, s2 []T, compare func(t1, t2 T) bool) []T {
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
