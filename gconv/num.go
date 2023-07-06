package gconv

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
