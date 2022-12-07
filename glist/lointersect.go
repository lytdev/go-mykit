package glist

// Contains 判断切片是否包含指定的元素
// @param: collection 原切片数据
// @param: element 查找的数据
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// ContainsBy 和FindIndex类似,判断切片是否包含指定的元素(根据指定的函数)
// @param: collection 原切片数据
// @param: predicate 进行比较的函数
func ContainsBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every 判断一个切片是否完全包含另一个切片
// @param: collection 大切片
// @param: subset 小切片
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !Contains(collection, elem) {
			return false
		}
	}

	return true
}

// EveryBy 判断一个切片的元素进行函数计算全部返回true
// @param: collection 大切片
// @param: predicate 计算函数
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Some 判断一个切片是否至少包含另一个切片的一个元素(和子集的元素有重叠)
// @param: collection 大切片
// @param: subset 小切片
func Some[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return true
		}
	}

	return false
}

// SomeBy  切片元素经过函数计算,结果至少有一个true
// @param: collection 大切片
// @param: predicate 计算函数
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return true
		}
	}

	return false
}

// None 如果切片中没有包含子集的元素(和子集的元素没有重叠),或者子集为空,则返回true
func None[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return false
		}
	}

	return true
}

// NoneBy 切片元素经过函数计算结果都是false
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return false
		}
	}
	return true
}

// Intersect 和SliceIntersect一样,计算2个切片的交集
func Intersect[T comparable](list1 []T, list2 []T) []T {
	var result = make([]T, 0)
	seen := map[T]struct{}{}

	for _, elem := range list1 {
		seen[elem] = struct{}{}
	}

	for _, elem := range list2 {
		if _, ok := seen[elem]; ok {
			result = append(result, elem)
		}
	}

	return result
}

// Difference 和SliceDiff一样,计算2个集合的交集
func Difference[T comparable](list1 []T, list2 []T) ([]T, []T) {
	var left = make([]T, 0)
	var right = make([]T, 0)

	seenLeft := map[T]struct{}{}
	seenRight := map[T]struct{}{}

	for _, elem := range list1 {
		seenLeft[elem] = struct{}{}
	}

	for _, elem := range list2 {
		seenRight[elem] = struct{}{}
	}

	for _, elem := range list1 {
		if _, ok := seenRight[elem]; !ok {
			left = append(left, elem)
		}
	}

	for _, elem := range list2 {
		if _, ok := seenLeft[elem]; !ok {
			right = append(right, elem)
		}
	}

	return left, right
}

// Union 计算2个集合的并集
func Union[T comparable](lists ...[]T) []T {
	var result = make([]T, 0)
	seen := map[T]struct{}{}

	for _, list := range lists {
		for _, e := range list {
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				result = append(result, e)
			}
		}
	}

	return result
}

// Without 将给定的排除元素从切片移除
func Without[T comparable](collection []T, exclude ...T) []T {
	result := make([]T, 0, len(collection))
	for _, e := range collection {
		if !Contains(exclude, e) {
			result = append(result, e)
		}
	}
	return result
}

// WithoutEmpty 排除切片里面的空数据
func WithoutEmpty[T comparable](collection []T) []T {
	var empty T

	result := make([]T, 0, len(collection))
	for _, e := range collection {
		if e != empty {
			result = append(result, e)
		}
	}

	return result
}
