package glist

import (
	"strconv"
	"testing"
)

func TestFindStrIndex(t *testing.T) {
	tmpArr := []string{"abc", "123", "xyz", "abc", "ABC"}
	index, ok := FindStrIndex(tmpArr, "xyz")
	if ok {
		t.Log("找到了,索引位置:" + strconv.Itoa(index))
	}
}

func TestRemoveItem(t *testing.T) {
	tmpArr := []string{"abc", "123", "xyz", "abc", "ABC"}
	newArr1 := RemoveFirstItem(tmpArr, func(item string) bool {
		return "abc" == item
	})
	for _, item := range newArr1 {
		t.Log(item)
	}
	t.Log("*****************************")
	newArr2 := RemoveAllItem(tmpArr, func(item string) bool {
		return "abc" == item
	})
	for _, item := range newArr2 {
		t.Log(item)
	}
}

func TestSliceIntersect(t *testing.T) {
	s1 := []string{"a", "b", "c", "d", "f"}
	s2 := []string{"c", "e", "f"}
	t.Log("************2个切片的交集*****************")
	newArr1 := SliceIntersect(s1, s2, func(v1, v2 string) bool {
		return v1 == v2
	})
	for _, item := range newArr1 {
		t.Log(item)
	}
	t.Log("************2个切片的差集*****************")
	newArr2 := SliceDiff(s2, s1, func(v1, v2 string) bool {
		return v1 == v2
	})
	for _, item := range newArr2 {
		t.Log(item)
	}
}
