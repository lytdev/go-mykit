package gconv

import "time"

//https://github.com/spf13/cast

// ToBool 将接口数据类型转换为bool
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToTime 将接口数据类型转换为time.Time.
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

func ToTimeInDefaultLocation(i interface{}, location *time.Location) time.Time {
	v, _ := ToTimeInDefaultLocationE(i, location)
	return v
}

// ToDuration 将接口数据类型转换为time.Duration.
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToFloat64 将接口数据类型转换为float64
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat32 将接口数据类型转换为float32
func ToFloat32(i interface{}) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToInt64 将接口数据类型转换为int64
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt32 将接口数据类型转换为int32
func ToInt32(i interface{}) int32 {
	v, _ := ToInt32E(i)
	return v
}

// ToInt16 将接口数据类型转换为int16
func ToInt16(i interface{}) int16 {
	v, _ := ToInt16E(i)
	return v
}

// ToInt8 将接口数据类型转换为int8
func ToInt8(i interface{}) int8 {
	v, _ := ToInt8E(i)
	return v
}

// ToInt 将接口数据类型转换为int
func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

// ToUint 将接口数据类型转换为uint
func ToUint(i interface{}) uint {
	v, _ := ToUintE(i)
	return v
}

// ToUint64 将接口数据类型转换为uint64
func ToUint64(i interface{}) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToUint32 将接口数据类型转换为uint32
func ToUint32(i interface{}) uint32 {
	v, _ := ToUint32E(i)
	return v
}

// ToUint16 将接口数据类型转换为uint16
func ToUint16(i interface{}) uint16 {
	v, _ := ToUint16E(i)
	return v
}

// ToUint8 将接口数据类型转换为uint8
func ToUint8(i interface{}) uint8 {
	v, _ := ToUint8E(i)
	return v
}

// ToString 将接口数据类型转换为string
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}

// ToStringMapString 将接口数据类型转换为map[string]string
func ToStringMapString(i interface{}) map[string]string {
	v, _ := ToStringMapStringE(i)
	return v
}

// ToStringMapStringSlice 将接口数据类型转换为map[string][]string
func ToStringMapStringSlice(i interface{}) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}

// ToStringMapBool 将接口数据类型转换为map[string]bool
func ToStringMapBool(i interface{}) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

// ToStringMapInt 将接口数据类型转换为map[string]int
func ToStringMapInt(i interface{}) map[string]int {
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

// ToSlice 将接口数据类型转换为[]interface{}
func ToSlice(i interface{}) []interface{} {
	v, _ := ToSliceE(i)
	return v
}

// ToBoolSlice 将接口数据类型转换为[]bool
func ToBoolSlice(i interface{}) []bool {
	v, _ := ToBoolSliceE(i)
	return v
}

// ToStringSlice 将接口数据类型转换为[]string
func ToStringSlice(i interface{}) []string {
	v, _ := ToStringSliceE(i)
	return v
}

// ToIntSlice 将接口数据类型转换为[]int
func ToIntSlice(i interface{}) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToDurationSlice 将接口数据类型转换为[]time.Duration
func ToDurationSlice(i interface{}) []time.Duration {
	v, _ := ToDurationSliceE(i)
	return v
}
