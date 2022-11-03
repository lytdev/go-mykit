package gxlsx

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/lytdev/go-mykit/gconv"
	"github.com/lytdev/go-mykit/gfrmt"
)

const (
	DatePattern        = "2006-01-02"
	DateTimePattern    = "2006-01-02 15:04:05"
	TagCustomKey       = "gxlsx"
	TagCustomTitleKey  = "title"
	TagCustomIndexKey  = "index"
	TagCustomFormatKey = "format"
)

var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
	"T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM",
	"AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ"}

func toCharStrArr(i int) (string, error) {
	if i > 52 {
		return "", errors.New("不支持excel超过52列")
	}
	if i < 0 {
		return "", errors.New("列索引设置错误,必须是正数")
	}
	return arr[i], nil

}

// ExcelFields /**Excel的字段属性
type ExcelFields struct {
	Title     string            //名称
	Index     int               //索引,从0开始
	Format    string            //所以的格式化类型(datetime),没有标注的话,默认是属性的类型
	Name      string            //json字段名称
	FieldType string            //字段类型
	Tags      map[string]string //保存所有tags
}

// ExcelStruct /** 封装对象
type ExcelStruct struct {
	MapIndex       map[int]string         //按照index排序的map
	IndexMax       int                    //列index最大值
	Fields         map[string]ExcelFields //所有字段
	StartRow       int                    //第几行开始为具体数据
	Err            error                  //错误
	ConvertTypeErr bool                   //类型转换时候,产生错误时是否直接提示报错
}

// Callback /** 定义回调函数
type Callback func(maps map[string]interface{}) error

// SetPointerStruct /** 设置并构造转换的结构体
func (c *ExcelStruct) SetPointerStruct(ptr interface{}) *ExcelStruct {
	// 获取入参的类型
	t := reflect.TypeOf(ptr)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		//参数必须是结构体
		c.Err = fmt.Errorf("the argument should be a pointer to the structure")
		return c
	}
	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()
	// 解析字段属性
	for i := 0; i < v.NumField(); i++ {
		// 取tag
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		fields := ExcelFields{}
		// 解析tag
		tagStr := tag.Get(TagCustomKey)
		index := i
		fields.Name = fieldInfo.Name
		if tagStr == "" {
			fields.Title = fieldInfo.Name
		} else {
			tagMap := gconv.TagConvMap(tagStr)
			if title, ok := tagMap[TagCustomTitleKey]; ok {
				fields.Title = title
			}
			if indexStr, ok := tagMap[TagCustomIndexKey]; ok {
				index, _ = strconv.Atoi(indexStr)
			}
		}
		//如果索引大,那么赋值
		if c.IndexMax < index {
			c.IndexMax = index
		}
		//列的索引值
		fields.Index = index
		//结构体属性的类型
		fields.FieldType = fieldInfo.Type.String()
		m := make(map[string]string)
		m[TagCustomTitleKey] = fields.Title
		m[TagCustomIndexKey] = strconv.Itoa(i)
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

// RowsProcess /** 处理函数
func (c *ExcelStruct) RowsProcess(rows [][]string, callback Callback) error {
	return c.RowsAllProcess(rows, callback)
}

// RowsAllProcess /** 处理sheet的rows
func (c *ExcelStruct) RowsAllProcess(rows [][]string, callback Callback) error {
	if c.Fields == nil {
		return fmt.Errorf("请填写结构体指针")
	}
	if c.Err != nil {
		return c.Err
	}
	for index, row := range rows {
		//如果索引小于已设置的开始行,那么跳过
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

// Row /** 处理一行的数据,将行数据根据index转换成map
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
				dateStr, formatErr := gfrmt.GetFormatDateStr(colCell)
				if formatErr != nil {
					//类型转换时候,产生错误时是否直接提示报错
					if c.ConvertTypeErr {
						return maps, formatErr
					}
				} else {
					colCell = dateStr
				}
				t, parseErr := time.ParseInLocation(DateTimePattern, colCell, time.Local)
				if parseErr != nil {
					//类型转换时候,产生错误时是否直接提示报错
					if c.ConvertTypeErr {
						return nil, parseErr
					}
				} else {
					maps[field] = t
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
					tmpInt, err := strconv.Atoi(colCell)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "int8":
					tmpInt, err := strconv.ParseInt(colCell, 10, 8)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "int16":
					tmpInt, err := strconv.ParseInt(colCell, 10, 16)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "int32":
					tmpInt, err := strconv.ParseInt(colCell, 10, 32)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "int64":
					tmpInt, err := strconv.ParseInt(colCell, 10, 64)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "uint":
					tmpInt, err := strconv.Atoi(colCell)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = uint(tmpInt)
					}
				case "uint8":
					tmpInt, err := strconv.ParseUint(colCell, 10, 8)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "uint16":
					tmpInt, err := strconv.ParseUint(colCell, 10, 16)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "uint32":
					tmpInt, err := strconv.ParseUint(colCell, 10, 32)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "uint64":
					tmpInt, err := strconv.ParseUint(colCell, 10, 64)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "float32":
					tmpInt, err := strconv.ParseFloat(colCell, 32)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "float64":
					tmpInt, err := strconv.ParseFloat(colCell, 64)
					if err != nil {
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, fmtErr(fields.FieldType, colCell)
						}
						maps[field] = 0
					} else {
						maps[field] = tmpInt
					}
				case "string":
					maps[field] = colCell
				}
			}
		}
	}
	return maps, nil
}

func fmtErr(typeName, cellVal string) error {
	return fmt.Errorf("类型转换出现错误，属性类型：%s，待转换的值：%s", typeName, cellVal)
}
