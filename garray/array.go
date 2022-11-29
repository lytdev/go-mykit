package garray

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
