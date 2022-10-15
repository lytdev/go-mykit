package gme

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// NewExcelStructDefault 默认 从第一行开始,索引从 0开始
func NewExcelStructDefault() *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = 1
	n.IndexMax = 10
	return n
}

// NewExcelStruct StartRow starts row, index starts from 0
//StartRow 开始行,索引从 0开始
//IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
func NewExcelStruct(StartRow, IndexMax int) *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = StartRow
	n.IndexMax = IndexMax
	return n
}

// WriteFile /**
func WriteFile[T any](n string, dataList []T) (*excelize.File, error) {
	if len(n) == 0 {
		n = "Sheet1"
	}
	if len(dataList) == 0 {
		return nil, errors.New("数据不存在")
	}
	f := excelize.NewFile()
	// 创建一个工作表
	sheetIndex := f.NewSheet(n)
	//获取结构体的指针
	ptr := &dataList[0]
	//获取结构体的字段类型
	fm := getStructInit(ptr)
	for _, field := range fm.Fields {
		err := f.SetCellValue(n, toCharStrArr(field.Index)+"1", field.Title)
		if err != nil {
			return nil, err
		}
	}

	ref := reflect.ValueOf(ptr).Elem()
	for index, structData := range dataList {
		dataValue := reflect.ValueOf(structData)
		for i := 0; i < ref.NumField(); i++ {
			// 获取字段属性
			fieldInfo := ref.Type().Field(i)
			fieldName := fieldInfo.Name
			//获取之前初始化的字段属性元数据
			excelField := fm.Fields[fieldName]
			//获取具体的值
			tmpData := dataValue.Field(i).Interface()
			if excelField.FieldType == "time.Time" {
				cellTime := tmpData.(time.Time)
				if excelField.Format == "datetime" {
					err := f.SetCellValue(n, toCharStrArr(excelField.Index)+strconv.Itoa(index+2), cellTime.Format(DATE_TIME_PATTERN))
					if err != nil {
						return nil, err
					}
				} else {
					err := f.SetCellValue(n, toCharStrArr(excelField.Index)+strconv.Itoa(index+2), cellTime.Format(DATE_PATTERN))
					if err != nil {
						return nil, err
					}
				}

			} else {
				err := f.SetCellValue(n, toCharStrArr(excelField.Index)+strconv.Itoa(index+2), tmpData)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(sheetIndex)

	return f, nil
}
