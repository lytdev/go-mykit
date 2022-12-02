package gconv

// ToBool 将接口数据类型转换为bool
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}
