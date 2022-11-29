package garray

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
