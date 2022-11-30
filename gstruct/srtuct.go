package gstruct

import (
	"github.com/lytdev/go-mykit/gconv"
	"reflect"
)

// Struct2Map 将结构体转换为map[string]interface{}
// struct{I int, S string}{I: 1, S: "a"} 将转换为 map[I:1 S:a].
func Struct2Map(obj interface{}) map[string]interface{} {
	// check params
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(obj)
	var m = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return m
}

// Struct2MapString 将结构体转换为map[string]string.
// struct{I int, S string}{I: 1, S: "a"} 将转换为 map[I:1 S:a].
func Struct2MapString(obj interface{}) map[string]string {
	// check params
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(obj)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = gconv.ToString(v.Field(i).Interface())
		}
	}
	return m
}
