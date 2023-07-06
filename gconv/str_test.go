package gconv

import "testing"

func TestToNum(t *testing.T) {
	var s1 = "3.1415926"
	i1 := ToInt(s1)
	i2 := ToInt64(s1)
	t.Log("i1:", i1)
	t.Log("i2:", i2)
	t.Log("demo:", int64(ToFloat64(s1)))
}
