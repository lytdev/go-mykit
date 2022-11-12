package gexcel

import (
	"fmt"
	"github.com/lytdev/go-mykit/gmap2struct"
	"os"
	"testing"
	"time"

	"github.com/xuri/excelize/v2"
)

type ExcelTest struct {
	Isbn      string    `gxlsx:"title:ISBN;index:1"`
	BookName  string    `gxlsx:"title:书名;index:0"`
	Author    string    `gxlsx:"title:作者;index:2"`
	PubDate   time.Time `gxlsx:"title:出版日期;index:4;format:datetime"`
	Price     float32   `gxlsx:"title:定价;index:3"`
	SuitObj   string    `gxlsx:"title:适用对象;index:5"`
	MajorType string    `gxlsx:"title:图书类目;index:6"`
	SubMajor  string    `gxlsx:"title:细分类目;index:7"`
}

func TestRead(t *testing.T) {
	filePath := "../_doc/图书列表.xlsx"
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
		// map转结构体
		if mapErr := gmap2struct.Decode(maps, &ptr); mapErr != nil {
			return mapErr
		}
		resultData = append(resultData, ptr)
		return nil
	})
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}
	for _, data := range resultData {
		fmt.Println(data)
		t.Log(data)
	}
}

//测试直接读取本地文件
func TestReadLocalFile(t *testing.T) {
	filePath := "../_doc/图书列表.xlsx"
	var ptr ExcelTest
	dataList, err := ReadFileToList(filePath, 0, ptr)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}
	for _, data := range dataList {
		fmt.Println(data)
		t.Log(data)
	}
}

//测试写入至本都文件
func TestWriteFile(t *testing.T) {
	pubDate, err := time.ParseInLocation(DatePattern, "2021-12-01", time.Local)
	if err != nil {
		fmt.Println(err)
	}
	var dataList []ExcelTest
	dataList = append(dataList, ExcelTest{
		Isbn:      "9787115375698",
		BookName:  "Excel 2013在会计与财务管理日常工作中的应用",
		Author:    "神龙工作室 编著",
		PubDate:   time.Now(),
		Price:     49.8,
		SuitObj:   "专科",
		MajorType: "计算机类",
		SubMajor:  "办公软件"})
	dataList = append(dataList, ExcelTest{
		Isbn:      "9787115383433",
		BookName:  "Excel 2013数据处理与分析",
		Author:    "金桥  周奎奎",
		PubDate:   pubDate,
		Price:     69,
		SuitObj:   "专科",
		MajorType: "计算机类",
		SubMajor:  "办公软件"})
	dataList = append(dataList, ExcelTest{
		Isbn:      "9787115512819",
		BookName:  "中文版Rhino 5.0实用教程",
		Author:    "邓宇燕",
		PubDate:   pubDate,
		Price:     59.9,
		SuitObj:   "专科",
		MajorType: "计算机类",
		SubMajor:  "三维设计"})
	sheetName := "Sheet1"
	f, err := WriteToFile(sheetName, dataList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 根据指定路径保存文件
	if err := f.SaveAs("../_doc/测试excel输出.xlsx"); err != nil {
		fmt.Println(err)
	}
}
