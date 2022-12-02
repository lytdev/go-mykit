package gconv

import "time"

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
