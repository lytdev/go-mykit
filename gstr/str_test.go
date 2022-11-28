package gstr

import "testing"

func TestStrToByteArr(t *testing.T) {
	key := "123456"
	data := StrToByteArr(key)
	rawStr := ByteArrToStr(data)
	t.Log(rawStr)
}
