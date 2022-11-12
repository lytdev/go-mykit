package gexcel

import (
	"errors"
	"fmt"
	"github.com/lytdev/go-mykit/gmap2struct"
	"io"
	"reflect"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// NewExcelStructDefault
//  @Description: 创建读取的结构体 默认,从第一行开始,索引从 0开始
//  @return *ExcelStruct
//
func NewExcelStructDefault() *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = 1
	n.IndexMax = 10
	n.ConvertTypeErr = true
	return n
}

// NewExcelStruct
//  @Description:
//  @param StartRow 开始行,索引从0开始
//  @param IndexMax 索引最大行,如果结构体中的index大于配置的,那么使用结构体中的
//  @param fastErr
//  @return *ExcelStruct
//
func NewExcelStruct(StartRow, IndexMax int, fastErr bool) *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = StartRow
	n.IndexMax = IndexMax
	n.ConvertTypeErr = fastErr
	return n
}

// ReadFileToList
//  @Description:读取本地文件至切片
//  @param filePath 本地文件的路径
//  @param sheetIndex 需要读取第几个单元薄
//  @param ptr 读取后的切片对象
//  @return resultData 返回的切片
//  @return err 错误信息
//
func ReadFileToList[T any](filePath string, sheetIndex int, ptr T) (resultData []T, err error) {
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("文件读取异常:" + err.Error())
		return nil, err
	}
	return readCore(xlsx, sheetIndex, ptr)
}

// ReadFileStreamToList
//  @Description: 读取文件流
//  @param r 文件流
//  @param sheetIndex 需要读取第几个单元薄
//  @param ptr 读取后的切片对象
//  @return resultData 返回的切片
//  @return err 错误信息
//
func ReadFileStreamToList[T any](r io.Reader, sheetIndex int, ptr T) (resultData []T, err error) {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		fmt.Println("文件读取异常:" + err.Error())
		return nil, err
	}
	return readCore(xlsx, sheetIndex, ptr)
}

// readCore
//  @Description: 读取excel的核心方法
//  @param xlsx excelize.File对象
//  @param sheetIndex 需要读取第几个单元薄
//  @param ptr 读取后的切片对象
//  @return resultData 返回的切片
//  @return err 错误信息
//
func readCore[T any](xlsx *excelize.File, sheetIndex int, ptr T) (resultData []T, err error) {
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

// WriteToFile
//  @Description: 将二维数组数据写入到excel的sheet
//  @param n 单元薄的名称
//  @param dataList 数据切片
//  @return *excelize.File 返回的excelFile对象
//  @return error 错误信息
//
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
