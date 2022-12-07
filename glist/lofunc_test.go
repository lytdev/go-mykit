package glist

import (
	"math"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5, 6}
	result := Filter(list, func(nbr int64, index int) bool {
		return nbr%2 == 0
	})
	// 过滤所有偶数的切片: [2 4 6]
	t.Logf("%v", result)
}

func TestMap(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5}

	result := Map(list, func(nbr int64, index int) string {
		return strconv.FormatInt(nbr*2, 10)
	})
	// 将原切片元素乘以2返回: [2 4 6 8 10]
	t.Logf("%v", result)
}

func TestFilterMap(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5}

	result := FilterMap(list, func(nbr int64, index int) (string, bool) {
		t.Logf("index:%d", index)
		return strconv.FormatInt(nbr*2, 10), nbr%2 == 0
	})

	// 只有是偶数才将计算的结果合并至新切片: [4 8]
	t.Logf("%v", result)
}

func TestFlatMap(t *testing.T) {
	list := []int64{1, 2, 3, 4}

	result := FlatMap(list, func(nbr int64, index int) []string {
		return []string{
			strconv.FormatInt(nbr, 10), // base 10
			strconv.FormatInt(nbr, 2),  // base 2
		}
	})
	//[1 1 2 10 3 11 4 100]
	t.Logf("%v", result)
}

func TestReduce(t *testing.T) {
	list := [][]int{{0, 1}, {2, 3}, {4, 5}}

	result := Reduce(list, func(agg, item []int, index int) []int {
		t.Logf("%v", agg)
		t.Logf("%v", item)
		t.Logf("%d", index)
		t.Log("-------------------")
		return append(agg, item...)
	}, []int{})
	//[0 1 2 3 4 5]
	t.Logf("%v", result)
}
func TestReduceRight(t *testing.T) {
	list := [][]int{{0, 1}, {2, 3}, {4, 5}}

	result := ReduceRight(list, func(agg []int, item []int, index int) []int {
		return append(agg, item...)
	}, []int{})
	//[4 5 2 3 0 1]
	t.Logf("%v", result)
}

func TestForEach(t *testing.T) {
	list := []int64{1, 2, 3, 4}
	ForEach(list, func(x int64, _ int) {
		t.Log(x)
	})
}

func TestTimes(t *testing.T) {
	result := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	t.Logf("%v", result)
}

func TestUniq(t *testing.T) {
	list := []int{1, 2, 2, 1}
	result := Uniq(list)
	// [1 2]
	t.Logf("%v", result)
}

func TestUniqBy(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7}

	result := UniqBy(list, func(i int) int {
		return i % 3
	})
	// [0 1 2]
	t.Logf("%v", result)
}

func TestGroupBy(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//按照余数分组
	result := GroupBy(list, func(i int) int {
		return i % 3
	})
	// [0 3 6 9]
	// [1 4 7]
	// [2 5 8]
	t.Logf("%v\n", result[0])
	t.Logf("%v\n", result[1])
	t.Logf("%v\n", result[2])
}

func TestChunk(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Chunk(list, 2)
	// [0 1]
	// [2 3]
	// [4 5]
	// [6 7]
	// [8 9]
	for _, item := range result {
		t.Logf("%v\n", item)
	}
}

func TestPartitionBy(t *testing.T) {
	list := []int{-2, -1, 0, 1, 2, 3, 4}

	result := PartitionBy(list, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	// [-2 -1]
	// [0 2 4]
	// [1 3]
	for _, item := range result {
		t.Logf("%v\n", item)
	}
}

func TestFlatten(t *testing.T) {
	list := [][]int{{0, 4, 2}, {3, 1, 5}}

	result := Flatten(list)
	// [0 4 2 3 1 5]
	t.Logf("%v", result)
}

func TestInterleave(t *testing.T) {
	list1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	list2 := [][]int{{1}, {2, 3, 4}, {5, 6}, {7, 8, 9, 10}}

	result1 := Interleave(list1...)
	result2 := Interleave(list2...)

	// [1 4 7 2 5 8 3 6 9]
	t.Logf("%v\n", result1)
	// [1 2 5 7 3 6 8 4 9 10]
	t.Logf("%v\n", result2)

}

func TestShuffle(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Shuffle(list)

	t.Logf("%v", result)
}

func TestReverse(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Reverse(list)

	t.Logf("%v", result)
	// Output: [5 4 3 2 1 0]
}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

func TestFill(t *testing.T) {
	list := []foo{{"a"}, {"a"}}

	result := Fill(list, foo{"b"})
	//  [{b} {b}]
	t.Logf("%v", result)
}

func TestRepeat(t *testing.T) {
	result := Repeat(2, foo{"a"})
	// [{a} {a}]
	t.Logf("%v", result)
}

func TestRepeatBy(t *testing.T) {
	result := RepeatBy(5, func(i int) string {
		return strconv.FormatInt(int64(math.Pow(float64(i), 2)), 10)
	})
	// [0 1 4 9 16]
	t.Logf("%v", result)
}

func TestKeyBy(t *testing.T) {
	list := []string{"a", "aa", "aaa"}

	result := KeyBy(list, func(str string) int {
		return len(str)
	})
	// map[1:a 2:aa 3:aaa]
	t.Logf("%v", result)
}

func TestAssociate(t *testing.T) {
	list := []string{"a", "aa", "aaa", "aa"}

	result := Associate(list, func(str string) (string, int) {
		return str, len(str)
	})
	// map[a:1 aa:2 aaa:3]
	t.Logf("%v", result)
}

func TestSliceToMap(t *testing.T) {
	list := []string{"a", "aa", "aaa", "aa"}

	result := SliceToMap(list, func(str string) (string, int) {
		return str, len(str)
	})
	// map[a:1 aa:2 aaa:3]
	t.Logf("%v", result)
}

func TestDrop(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Drop(list, 2)
	// [2 3 4 5]
	t.Logf("%v", result)

}

func TestDropRight(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}

	result := DropRight(list, 2)
	// [0 1 2 3]
	t.Logf("%v", result)
}

func TestDropWhile(t *testing.T) {
	list := []int{0, 8, 1, 0, 4, 5}

	result := DropWhile(list, func(val int) bool {
		return val < 2
	})
	// [8 1 0 4 5]
	t.Logf("%v", result)

}

func TestDropRightWhile(t *testing.T) {
	list := []int{0, 8, 1, 0, 4, 5}

	result := DropRightWhile(list, func(val int) bool {
		return val > 2
	})
	// [0 8 1 0]
	t.Logf("%v", result)
}

func TestReject(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Reject(list, func(x int, _ int) bool {
		return x%2 == 0
	})
	//[1 3 5]
	t.Logf("%v", result)
}

func TestCount(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 2}

	result := Count(list, 2)
	// 3
	t.Logf("%v", result)
}

func TestCountBy(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3}

	result := CountBy(list, func(i int) bool {
		return i < 3
	})
	// 6
	t.Logf("%v", result)
}

func TestCountValues(t *testing.T) {
	result1 := CountValues([]int{})
	result2 := CountValues([]int{1, 2})
	result3 := CountValues([]int{1, 2, 2})
	result4 := CountValues([]string{"foo", "bar", ""})
	result5 := CountValues([]string{"foo", "bar", "bar"})
	// map[]
	t.Logf("%v\n", result1)
	// map[1:1 2:1]
	t.Logf("%v\n", result2)
	// map[1:1 2:2]
	t.Logf("%v\n", result3)
	// map[:1 bar:1 foo:1]
	t.Logf("%v\n", result4)
	// map[bar:2 foo:1]
	t.Logf("%v\n", result5)

}

func TestCountValuesBy(t *testing.T) {
	isEven := func(v int) bool {
		return v%2 == 0
	}
	result1 := CountValuesBy([]int{}, isEven)
	result2 := CountValuesBy([]int{1, 2}, isEven)
	result3 := CountValuesBy([]int{1, 2, 2}, isEven)
	length := func(v string) int {
		return len(v)
	}
	result4 := CountValuesBy([]string{"foo", "bar", ""}, length)
	result5 := CountValuesBy([]string{"foo", "bar", "bar"}, length)
	//map[]
	t.Logf("%v\n", result1)
	// map[false:1 true:1]
	t.Logf("%v\n", result2)
	// map[false:1 true:2]
	t.Logf("%v\n", result3)
	// map[0:1 3:2]
	t.Logf("%v\n", result4)
	// map[3:3]
	t.Logf("%v\n", result5)
}

func TestSubset(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Subset(list, 2, 3)
	// [2 3 4]
	t.Logf("%v", result)
}

func TestSlice(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}

	result := Slice(list, 1, 4)
	// [1 2 3]
	t.Logf("%v\n", result)

	result = Slice(list, 4, 1)
	// []
	t.Logf("%v\n", result)

	result = Slice(list, 4, 5)
	// [4]
	t.Logf("%v\n", result)
}
func TestReplace(t *testing.T) {
	list := []int{0, 1, 0, 1, 2, 3, 0}

	result := Replace(list, 0, 42, 1)
	// [42 1 0 1 2 3 0]
	t.Logf("%v\n", result)

	result = Replace(list, -1, 42, 1)
	// [0 1 0 1 2 3 0]
	t.Logf("%v\n", result)

	result = Replace(list, 0, 42, 2)
	// [42 1 42 1 2 3 0]
	t.Logf("%v\n", result)

	result = Replace(list, 0, 42, -1)
	// [42 1 42 1 2 3 42]
	t.Logf("%v\n", result)

}

func TestReplaceAll(t *testing.T) {
	list := []string{",", "foo", ",", "bar", ","}

	result := ReplaceAll(list, ",", "-")
	// [- foo - bar -]
	t.Logf("%v", result)
}

func TestCompact(t *testing.T) {
	list := []string{"", "foo", "", "bar", ""}

	result := Compact(list)
	// [foo bar]
	t.Logf("%v", result)
}

func TestIsSorted(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := IsSorted(list)
	// true
	t.Logf("%v", result)

	// Output:
}

func TestIsSortedByKey(t *testing.T) {
	list := []string{"a", "bb", "ccc"}

	result := IsSortedByKey(list, func(s string) int {
		return len(s)
	})
	// true
	t.Logf("%v", result)
}
