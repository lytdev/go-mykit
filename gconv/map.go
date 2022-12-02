package gconv

import "reflect"

// ToStrMapBool 将接口数据类型转换为map[string]bool
func ToStrMapBool(i interface{}) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

// ToStrMapInt 将接口数据类型转换为map[string]int
func ToStrMapInt(i interface{}) map[string]int {
	v, _ := ToStringMapIntE(i)
	return v
}

// ToStringMapInt64 将接口数据类型转换为map[string]int64
func ToStringMapInt64(i interface{}) map[string]int64 {
	v, _ := ToStringMapInt64E(i)
	return v
}

// ToStringMap 将接口数据类型转换为map[string]interface{}
func ToStringMap(i interface{}) map[string]interface{} {
	v, _ := ToStringMapE(i)
	return v
}

// ToStrMapStr 将接口数据类型转换为map[string]string
func ToStrMapStr(i interface{}) map[string]string {
	v, _ := ToStringMapStringE(i)
	return v
}

// ToStrMapStrSlice 将接口数据类型转换为map[string][]string
func ToStrMapStrSlice(i interface{}) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}

// StructToMap 利用反射将结构体转化为map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		mapTag := obj1.Field(i).Tag.Get("map")
		if mapTag != "" {
			data[mapTag] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}
