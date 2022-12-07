package gmap

import "github.com/lytdev/go-mykit/glist"

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Keys 通过map的key构造切片
// Play: https://go.dev/play/p/Uu11fHASqrU
func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

// Values 通过map的value构造切片
// Play: https://go.dev/play/p/nnRTQkzQfF6
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}

// PickBy 根据给定的条件过滤map
// Play: https://go.dev/play/p/kdg8GR_QMmf
func PickBy[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if predicate(k, v) {
			r[k] = v
		}
	}
	return r
}

// PickByKeys 过滤key在切片里的map数据
// Play: https://go.dev/play/p/R1imbuci9qU
func PickByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if glist.Contains(keys, k) {
			r[k] = v
		}
	}
	return r
}

// PickByValues 过滤value在切片里的map数据
// Play: https://go.dev/play/p/1zdzSvbfsJc
func PickByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if glist.Contains(values, v) {
			r[k] = v
		}
	}
	return r
}

// OmitBy  PickBy的相反结果,不满足给定条件的数据过滤map
// Play: https://go.dev/play/p/EtBsR43bdsd
func OmitBy[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !predicate(k, v) {
			r[k] = v
		}
	}
	return r
}

// OmitByKeys 过滤key不在切片里的map数据
// Play: https://go.dev/play/p/t1QjCrs-ysk
func OmitByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !glist.Contains(keys, k) {
			r[k] = v
		}
	}
	return r
}

// OmitByValues 过滤value不在切片里的map数据
// Play: https://go.dev/play/p/9UYZi-hrs8j
func OmitByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !glist.Contains(values, v) {
			r[k] = v
		}
	}
	return r
}

// Entries 将map转换为Entry类型的切片,类似java里面的Entry
func Entries[K comparable, V any](in map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(in))

	for k, v := range in {
		entries = append(entries, Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return entries
}

// ToPairs Entries的别名
// Play: https://go.dev/play/p/3Dhgx46gawJ
func ToPairs[K comparable, V any](in map[K]V) []Entry[K, V] {
	return Entries(in)
}

// FromEntries Entry类型的切片转换为map
// Play: https://go.dev/play/p/oIr5KHFGCEN
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	out := map[K]V{}

	for _, v := range entries {
		out[v.Key] = v.Value
	}

	return out
}

// FromPairs FromEntries的别名
// Play: https://go.dev/play/p/oIr5KHFGCEN
func FromPairs[K comparable, V any](entries []Entry[K, V]) map[K]V {
	return FromEntries(entries)
}

// Invert 调换map的key和value,如果value存在重复的,则后面遍历的会覆盖前面的
// Play: https://go.dev/play/p/rFQ4rak6iA1
func Invert[K comparable, V comparable](in map[K]V) map[V]K {
	out := map[V]K{}

	for k, v := range in {
		out[v] = k
	}

	return out
}

// Assign 合并多个map
// Play: https://go.dev/play/p/VhwfJOyxf5o
func Assign[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}

	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

// MapKeys 遍历map,并将map的key通过iteratee转换为另一个值作为新的map的key
// Play: https://go.dev/play/p/9_4WPIqOetJ
func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(value V, key K) R) map[R]V {
	result := map[R]V{}

	for k, v := range in {
		result[iteratee(v, k)] = v
	}

	return result
}

// MapValues 遍历map,并将map的value通过iteratee转换为另一个值作为新的map的value
// Play: https://go.dev/play/p/T_8xAfvcf0W
func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(value V, key K) R) map[K]R {
	result := map[K]R{}

	for k, v := range in {
		result[k] = iteratee(v, k)
	}

	return result
}

// MapEntries 遍历map,并将map的key和value通过iteratee转换为另一个值作为新的map的key和value
// Play: https://go.dev/play/p/VuvNQzxKimT
func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](in map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(in))

	for k1, v1 := range in {
		k2, v2 := iteratee(k1, v1)
		result[k2] = v2
	}

	return result
}

// MapToSlice 遍历map,并通过iteratee,将map的key和value进行转换作为新的切片的值
// Play: https://go.dev/play/p/ZuiCZpDt6LD
func MapToSlice[K comparable, V any, R any](in map[K]V, iteratee func(key K, value V) R) []R {
	result := make([]R, 0, len(in))

	for k, v := range in {
		result = append(result, iteratee(k, v))
	}

	return result
}
