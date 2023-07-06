package gexcel

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/lytdev/go-mykit/gconv"
)

type FieldsModel struct {
	Err    error                  //错误
	Fields map[string]ExcelFields //所有字段
	// index sort
	MapIndex map[int]string //按照 index 排序
	//max index
	IndexMax int // index 最大
}

// getStructInit 根据结构体获取excel的元数据信息
func getStructInit(ptr interface{}) *FieldsModel {
	fm := new(FieldsModel)
	// 获取入参的类型
	t := reflect.TypeOf(ptr)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fm.Err = fmt.Errorf("the argument should be a pointer to the structure")
		return fm
	}
	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()
	//Parsing fields
	// 解析字段
	for i := 0; i < v.NumField(); i++ {
		// 取tag
		fieldInfo := v.Type().Field(i)
		//
		tag := fieldInfo.Tag
		fields := ExcelFields{}
		// 解析tag
		tagStr := tag.Get(TagCustomKey)
		index := i
		fields.Name = fieldInfo.Name
		if tagStr == "" {
			fields.Title = fieldInfo.Name
		} else {
			tagMap := gconv.StructTagToMap(tagStr)
			if title, ok := tagMap[TagCustomTitleKey]; ok {
				fields.Title = title
			}
			if indexStr, ok := tagMap[TagCustomIndexKey]; ok {
				index, _ = strconv.Atoi(indexStr)
			}
			if format, ok := tagMap[TagCustomFormatKey]; ok {
				fields.Format = format
			}
		}
		//If the index is large, the value is assigned
		//如果索引大,那么赋值
		if fm.IndexMax < index {
			fm.IndexMax = index
		}
		fields.Index = index
		fields.FieldType = fieldInfo.Type.String()
		m := make(map[string]string)
		m[TagCustomTitleKey] = fields.Title
		m[TagCustomIndexKey] = strconv.Itoa(i)
		fields.Tags = m
		//
		if fm.Fields == nil {
			fm.Fields = make(map[string]ExcelFields)
		}
		fm.Fields[fields.Name] = fields
		if fm.MapIndex == nil {
			fm.MapIndex = make(map[int]string)
		}
		fm.MapIndex[index] = fields.Name
	}
	return fm
}
