package gexcel

import (
	"errors"
	"fmt"
	"github.com/lytdev/go-mykit/gmap2struct"
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
	n.ConvertTypeErr = true
	return n
}

// NewExcelStruct /**StartRow开始行,索引从0开始;IndexMax索引最大行,如果结构体中的index大于配置的,那么使用结构体中的
func NewExcelStruct(StartRow, IndexMax int, fastErr bool) *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = StartRow
	n.IndexMax = IndexMax
	n.ConvertTypeErr = fastErr
	return n
}

// ReadFileToList /*读取本地文件
func ReadFileToList[T any](filePath string, sheetIndex int, ptr T) (resultData []T, err error) {
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("文件读取异常:" + err.Error())
		return nil, err
	}
	sheetName := xlsx.GetSheetName(sheetIndex)
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		fmt.Println("获取行数据异常:" + err.Error())
		return nil, err
	}
	err = NewExcelStructDefault().SetPointerStruct(&ptr).RowsAllProcess(rows, func(maps map[string]interface{}) error {
		// map转结构体
		if mapErr := gmap2struct.Decode(maps, &ptr); mapErr != nil {
			return mapErr
		}
		resultData = append(resultData, ptr)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resultData, nil
}

// WriteToFile /** 将二维数组数据写入到excel的sheet
func WriteToFile[T any](n string, dataList []T) (*excelize.File, error) {
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
		colName, convErr := toCharStrArr(field.Index)
		if convErr != nil {
			return nil, convErr
		}
		err := f.SetCellValue(n, colName+"1", field.Title)
		if err != nil {
			return nil, err
		}
	}
	//反射获取结构体的字段属性
	ref := reflect.ValueOf(ptr).Elem()
	for index, structData := range dataList {
		//获取单元格的值,此处获取的是内存值
		dataValue := reflect.ValueOf(structData)
		for i := 0; i < ref.NumField(); i++ {
			// 获取字段属性
			fieldInfo := ref.Type().Field(i)
			fieldName := fieldInfo.Name
			//获取之前初始化的字段属性元数据
			excelField := fm.Fields[fieldName]
			//获取属性具体的值
			tmpData := dataValue.Field(i).Interface()
			//初始化excel列的标题索引
			colTitle, convErr := toCharStrArr(excelField.Index)
			if convErr != nil {
				return nil, convErr
			}
			colTitle = colTitle + strconv.Itoa(index+2)
			//时间值的判断
			if excelField.FieldType == "time.Time" {
				cellTime := tmpData.(time.Time)
				if excelField.Format == "datetime" {
					err := f.SetCellValue(n, colTitle, cellTime.Format(DateTimePattern))
					if err != nil {
						return nil, err
					}
				} else {
					err := f.SetCellValue(n, colTitle, cellTime.Format(DatePattern))
					if err != nil {
						return nil, err
					}
				}

			} else {
				err := f.SetCellValue(n, colTitle, tmpData)
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
