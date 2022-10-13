/*
 * @Author       : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @Date         : 2022-10-11 19:00:26
 * @LastEditors  : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @FilePath     : \go-myexcel\model.go
 * @Description  :
 * Copyright (c) 2022 by 刘元涛 snoopy_718@mails.ccnu.edu.cn, All Rights Reserved.
 */
package gme

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/lytdev/go-myexcel/gconv"
	"github.com/lytdev/go-myexcel/gformt"
)

const (
	DATE_PATTERN         = "2006-01-02"
	DATE_TIME_PATTERN    = "2006-01-02 15:04:05"
	TAG_CUSTOM_KEY       = "gme"
	TAG_CUSTOM_TITLE_KEY = "title"
	TAG_CUSTOM_INDEX_KEY = "index"
)

/**
 * @Description : 字段
 * @return       {*}
 * @Date        : 2022-10-11 19:00:44
 */
type ExcelFields struct {
	//name
	Title string //名称
	//Index starts at 0
	Index int //索引  从0 开始
	//JSON field name
	Name string //json 字段名称
	//Field type
	FieldType string //字段类型
	//Save all tags
	Tags map[string]string // 保存所有tags
}

/**
 * @Description : 封装对象
 * @Description : init struct
 * @return       {*}
 * @Date        : 2022-10-11 19:00:50
 */
type ExcelStruct struct {
	// index sort
	MapIndex map[int]string //按照 index 排序
	//max index
	IndexMax int // index 最大
	//All fields
	Fields map[string]ExcelFields //所有字段
	//The first few lines start with specific data
	StartRow int //第几行开始为具体数据
	//error
	Err error //错误
	//During type conversion, whether to directly prompt to report an error when an error occurs
	ConvertTypeErr bool //类型转换时候,产生错误时是否直接提示报错
}

/**
 * @Description :定义回调函数
 * @param        {map[string]interface{}} maps
 * @return       {*}
 * @Date        : 2022-10-13 17:20:25
 */
type Callback func(maps map[string]interface{}) error

/**
 * @Description : set Struct pointer
 * @Description : 设置并构造转换的结构体
 * @param        {interface{}} ptr
 * @return       {*}
 * @Date        : 2022-10-13 17:17:10
 */
func (c *ExcelStruct) SetPointerStruct(ptr interface{}) *ExcelStruct {
	//Gets the type of the input parameter
	// 获取入参的类型
	t := reflect.TypeOf(ptr)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		//Argument should be a struct pointer
		c.Err = fmt.Errorf("the argument should be a pointer to the structure")
		return c
	}
	//Take the structure variable that the pointer points to
	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()
	//Parsing fields
	// 解析字段
	for i := 0; i < v.NumField(); i++ {
		//Take tag
		// 取tag
		fieldInfo := v.Type().Field(i)
		//
		tag := fieldInfo.Tag
		fields := ExcelFields{}
		//Parsing
		// 解析tag
		tagStr := tag.Get(TAG_CUSTOM_KEY)
		index := i
		fields.Name = fieldInfo.Name
		if tagStr == "" {
			fields.Title = fieldInfo.Name
		} else {
			tagMap := gconv.TagConvMap(tagStr)
			if title, ok := tagMap[TAG_CUSTOM_TITLE_KEY]; ok {
				fields.Title = title
			}
			if indexStr, ok := tagMap[TAG_CUSTOM_INDEX_KEY]; ok {
				index, _ = strconv.Atoi(indexStr)
			}
		}
		//If the index is large, the value is assigned
		//如果索引大,那么赋值
		if c.IndexMax < index {
			c.IndexMax = index
		}
		fields.Index = index
		fields.FieldType = fieldInfo.Type.String()
		m := make(map[string]string)
		m[TAG_CUSTOM_TITLE_KEY] = fields.Title
		m[TAG_CUSTOM_INDEX_KEY] = strconv.Itoa(i)
		fields.Tags = m
		//
		if c.Fields == nil {
			c.Fields = make(map[string]ExcelFields)
		}
		c.Fields[fields.Name] = fields
		if c.MapIndex == nil {
			c.MapIndex = make(map[int]string)
		}
		c.MapIndex[index] = fields.Name
	}
	return c
}

/**
 * @Description : 处理函数
 * @param        {[][]string} rows
 * @param        {Callback} callback
 * @return       {*}
 * @Date        : 2022-10-13 17:20:40
 */
func (c *ExcelStruct) RowsProcess(rows [][]string, callback Callback) error {
	return c.RowsAllProcess(rows, callback)
}

/**
 * @Description : process sheet rows
 * @Description : 处理sheet的rows
 * @param        {[][]string} rows
 * @param        {Callback} callback
 * @return       {*}
 * @Date        : 2022-10-13 17:03:00
 */
func (c *ExcelStruct) RowsAllProcess(rows [][]string, callback Callback) error {
	if c.Fields == nil {
		//Please fill in the structure pointer
		return fmt.Errorf("请填写结构体指针")
	}
	if c.Err != nil {
		return c.Err
	}
	//data := []interface{}{}
	for index, row := range rows {
		//If the index is less than the set start row, skip
		//如果 索引 小于 已设置的 开始行,那么跳过
		if index < c.StartRow {
			continue
		}
		//单行处理
		maps, err := c.Row(row)
		if err != nil {
			return err
		}
		err2 := callback(maps)
		if err2 != nil {
			return err2
		}
	}
	return nil
}

/**
* @Description : Process a row of data and convert the row data into a map according to index
* @Description : 处理一行的数据,将行数据根据index转换成map
* @param        {[]string} row
* @return       {*}
* @Date        : 2022-10-13 16:55:02
 */
func (c *ExcelStruct) Row(row []string) (map[string]interface{}, error) {
	if c.Fields == nil {
		//Please fill in the structure pointer
		return nil, fmt.Errorf("please fill in the structure pointer")
	}
	if c.Err != nil {
		return nil, c.Err
	}
	//行转map
	return c.row2Map(row)
}

/**
 * @Description :the core logic of row to map is to deal with some special formats
 * @Description :行转map的核心逻辑,主要是处理一些特殊的格式(日期、数字、时间)
 * @param        {[]string} row
 * @return       {*}
 * @Date        : 2022-10-13 17:08:43
 */
func (c *ExcelStruct) row2Map(row []string) (map[string]interface{}, error) {
	maps := make(map[string]interface{})
	for i, colCell := range row {
		//len should be used for string judgments
		//字符串判断应该使用len
		if colCell == "" || len(colCell) < 1 {
			continue
		}
		//check the key exists
		//判断键名是否存在
		if field, ok := c.MapIndex[i]; ok {
			maps[field] = ""
			//Type conversion
			//类型转换
			fields := c.Fields[field]
			//character
			//字符
			if fields.FieldType == "string" {
				maps[field] = colCell
				continue
			}
			//time
			//时间
			if fields.FieldType == "time.Time" && len(colCell) > 0 {
				//colCell的日志在excel可能存在多种形式,这里统一转换为yyyy-MM-DD的形式
				dateStr, err := gformt.GetFormatDateStr(colCell)
				if err == nil {
					colCell = dateStr
				}
				t, err := time.ParseInLocation(DATE_PATTERN, colCell, time.Local)
				if err == nil {
					maps[field] = t
				} else {
					//During type conversion, whether to directly prompt to report an error when an error occurs
					//类型转换时候,产生错误时是否直接提示报错
					if c.ConvertTypeErr {
						return nil, err
					}
				}
			} else {
				//other
				//其他类型
				switch fields.FieldType {
				case "bool":
					lower := strings.ToLower(colCell)
					if lower == "true" {
						maps[field] = true
					} else {
						maps[field] = false
					}
				case "int":
					int, err := strconv.Atoi(colCell)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "int8":
					int, err := strconv.ParseInt(colCell, 10, 8)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "int16":
					int, err := strconv.ParseInt(colCell, 10, 16)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "int32":
					int, err := strconv.ParseInt(colCell, 10, 32)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "int64":
					int, err := strconv.ParseInt(colCell, 10, 64)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
					//fmt.Println("int64=", int)
				case "uint":
					int, err := strconv.Atoi(colCell)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = uint(int)
					}
				case "uint8":
					int, err := strconv.ParseUint(colCell, 10, 8)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "uint16":
					int, err := strconv.ParseUint(colCell, 10, 16)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "uint32":
					int, err := strconv.ParseUint(colCell, 10, 32)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "uint64":
					int, err := strconv.ParseUint(colCell, 10, 64)
					if err != nil {
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "float32":
					int, err := strconv.ParseFloat(colCell, 32)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "float64":
					int, err := strconv.ParseFloat(colCell, 64)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = int
					}
				case "string":
					maps[field] = colCell
				}
			}
		}
	}
	return maps, nil
}
