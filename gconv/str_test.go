package gconv

import "testing"

func TestToStr(t *testing.T) {
	var cityArray = [...]string{"北京", "上海", "深圳"}
	s1 := ArrToStr(cityArray)
	t.Log(s1)
	var numArray = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numSlice = numArray[1:5]
	s2 := ArrToStr(numArray)
	t.Log(s2)
	s3 := ArrToStr(numSlice)
	t.Log(s3)

}
