/*
 * @Author       : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @Date         : 2022-10-11 17:53:49
 * @LastEditors  : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @FilePath     : \go-myexcel\excel_test.go
 * @Description  :
 * Copyright (c) 2022 by 刘元涛 snoopy_718@mails.ccnu.edu.cn, All Rights Reserved.
 */
package gme

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/lytdev/go-myexcel/mapstructure"

	"github.com/xuri/excelize/v2"
)

type ExcelTest struct {
	Isbn      string    `gme:"title:ISBN;index:0"`
	BookName  string    `gme:"title:书名;index:1"`
	Author    string    `gme:"title:作者;index:2"`
	PubDate   time.Time `gme:"title:出版日期;index:3"`
	Price     float32   `gme:"title:定价;index:4"`
	SuitObj   string    `gme:"title:适用对象;index:5"`
	MajorType string    `gme:"title:图书类目;index:6"`
	SubMajor  string    `gme:"title:细分类目;index:7"`
}

func TestRead(t *testing.T) {
	filePath := "_doc/图书列表.xlsx"
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		t.Error("文件读取异常:", err)
		os.Exit(1)
	}
	// Get all the rows in a sheet.
	sheetName := xlsx.GetSheetName(0)
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		t.Error("获取行数据异常:", err)
		os.Exit(1)
	}
	var resultData []ExcelTest
	err = NewExcelStructDefault().SetPointerStruct(&ExcelTest{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
		var ptr ExcelTest
		// map 转 结构体
		if mapErr := mapstructure.Decode(maps, &ptr); mapErr != nil {
			return mapErr
		}
		resultData = append(resultData, ptr)
		return nil
	})
	if err != nil {
		t.Error("转换出现错误:", err)
		os.Exit(1)
	}
	for _, data := range resultData {
		fmt.Println(data)
		t.Log(data)
	}
}
