package gmap

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestKeys(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2}

	result := Keys(kv)

	sort.StringSlice(result).Sort()
	t.Logf("%v", result)
	// Output: [bar foo]
}

func TestValues(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2}

	result := Values(kv)

	sort.IntSlice(result).Sort()
	t.Logf("%v", result)
	// Output: [1 2]
}

func TestPickBy(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := PickBy(kv, func(key string, value int) bool {
		return value%2 == 1
	})

	t.Logf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func TestPickByKeys(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := PickByKeys(kv, []string{"foo", "baz"})

	t.Logf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func TestPickByValues(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := PickByValues(kv, []int{1, 3})

	t.Logf("%v %v %v", len(result), result["foo"], result["baz"])
	// Output: 2 1 3
}

func TestOmitBy(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := OmitBy(kv, func(key string, value int) bool {
		return value%2 == 1
	})

	t.Logf("%v", result)
	// Output: map[bar:2]
}

func TestOmitByKeys(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := OmitByKeys(kv, []string{"foo", "baz"})

	t.Logf("%v", result)
	// Output: map[bar:2]
}

func TestOmitByValues(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := OmitByValues(kv, []int{1, 3})

	t.Logf("%v", result)
	// Output: map[bar:2]
}

func TestEntries(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := Entries(kv)

	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].Key, result[j].Key) < 0
	})
	t.Logf("%v", result)
	// Output: [{bar 2} {baz 3} {foo 1}]
}

func TestFromEntries(t *testing.T) {
	result := FromEntries([]Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
		{
			Key:   "baz",
			Value: 3,
		},
	})

	t.Logf("%v %v %v %v", len(result), result["foo"], result["bar"], result["baz"])
	// Output: 3 1 2 3
}

func TestInvert(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	result := Invert(kv)

	t.Logf("%v %v %v %v", len(result), result[1], result[2], result[3])
	// Output: 3 foo bar baz
}

func TestAssign(t *testing.T) {
	result := Assign(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)

	t.Logf("%v %v %v %v", len(result), result["a"], result["b"], result["c"])
	// Output: 3 1 3 4
}

func TestMapKeys(t *testing.T) {
	kv := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}

	result := MapKeys(kv, func(_ int, v int) string {
		return strconv.FormatInt(int64(v), 10)
	})

	t.Logf("%v %v %v %v %v", len(result), result["1"], result["2"], result["3"], result["4"])
	// Output: 4 1 2 3 4
}

func TestMapValues(t *testing.T) {
	kv := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}

	result := MapValues(kv, func(_ int, v int) string {
		return strconv.FormatInt(int64(v), 10)
	})

	t.Logf("%v %v %v %v %v", len(result), result[1], result[2], result[3], result[4])
	// Output: 4 1 2 3 4
}

func TestMapEntries(t *testing.T) {
	kv := map[string]int{"foo": 1, "bar": 2}

	result := MapEntries(kv, func(k string, v int) (int, string) {
		return v, k
	})

	t.Logf("%v\n", result)
	// Output: map[1:foo 2:bar]
}

func TestMapToSlice(t *testing.T) {
	kv := map[int]int64{1: 1, 2: 2, 3: 3, 4: 4}

	result := MapToSlice(kv, func(k int, v int64) string {
		return fmt.Sprintf("%d_%d", k, v)
	})

	sort.StringSlice(result).Sort()
	t.Logf("%v", result)
	// Output: [1_1 2_2 3_3 4_4]
}
