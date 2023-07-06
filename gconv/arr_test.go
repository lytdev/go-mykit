package gconv

import (
	"strings"
	"testing"
)

func TestToSlice(t *testing.T) {
	var s1 = "3,2154"
	//s1: [3.1415926,2154]
	d1 := strings.Split(s1, ",")
	t.Log("s1:", d1)
	d2 := ToIntSlice(d1)
	t.Log("d2:", d2)
}
