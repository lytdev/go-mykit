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

	"github.com/lytdev/go-myexcel/mapstructure"

	"github.com/xuri/excelize/v2"
)

type ExcelTest struct {
	Isbn      string  `json:"isbn" meg:"title:ISBN;index:0"`
	BookName  string  `json:"book_name" mapstructure:"book_name" meg:"title:书名;index:1"`
	Author    string  `json:"author" meg:"title:作者;index:2"`
	PubDate   string  `json:"pub_date" mapstructure:"pub_date" meg:"title:出版日期;index:3"`
	Price     float32 `json:"price" meg:"title:定价;index:4"`
	SuitObj   string  `json:"suit_obj" mapstructure:"suit_obj" meg:"title:适用对象;index:5"`
	MajorType string  `json:"major_type" mapstructure:"major_type" meg:"title:图书类目;index:6"`
	SubMajor  string  `json:"sub_major" mapstructure:"sub_major" meg:"title:细分类目;index:7"`
}

func Test_Read(t *testing.T) {
	filePath := "D:\\2022-2云展\\调整标题后\\人邮社本科教材.xlsx"
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
