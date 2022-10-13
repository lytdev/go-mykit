<!--
 * @Author       : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @Date         : 2022-10-11 18:47:44
 * @LastEditors  : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @FilePath     : \go-myexcel\README.md
 * @Description  :
 * Copyright (c) 2022 by 刘元涛 snoopy_718@mails.ccnu.edu.cn, All Rights Reserved.
-->

# 说明

> excelize/v2 的增版版本，支持读取 excel 并转换为 map 和 struct；核心部分 没有使用任何 第三方包,引入第三方包都是测试和转换使用的

# 使用方式

项目中执行引入包

> go get -u github.com/lytdev/go-myexcel

结构体`TAG` 中 `gme`(`index`) 必须存在,`json`可不写

`json`: 字段名称,映射成`map`

`gme->index`: 索序号(从 0 开始)

`gme->title`: 名称

Excel 对应的结构体参考：

```golang
type ExcelBook struct {
 Isbn      string  `json:"isbn" gme:"title:ISBN;index:0"`
 BookName  string  `json:"book_name" mapstructure:"book_name" gme:"title:书名;index:1"`
 Author    string  `json:"author" gme:"title:作者;index:2"`
 PubDate   string  `json:"pub_date" mapstructure:"pub_date" gme:"title:出版日期;index:3"`
 Price     float32 `json:"price" gme:"title:定价;index:4"`
 SuitObj   string  `json:"suit_obj" mapstructure:"suit_obj" gme:"title:适用对象;index:5"`
 MajorType string  `json:"major_type" mapstructure:"major_type" gme:"title:图书类目;index:6"`
 SubMajor  string  `json:"sub_major" mapstructure:"sub_major" gme:"title:细分类目;index:7"`
}
```

如果要执行案例,先引入案例依赖

```bin
go get github.com/xuri/excelize/v2
```

使用方式一：

```golang
package main

import (
 "fmt"
 "os"
 "testing"

 "github.com/lytdev/go-myexcel/mapstructure"

 "github.com/xuri/excelize/v2"
)

type ExcelTest struct {
 Isbn      string  `json:"isbn" gme:"title:ISBN;index:0"`
 BookName  string  `json:"book_name" mapstructure:"book_name" gme:"title:书名;index:1"`
 Author    string  `json:"author" gme:"title:作者;index:2"`
 PubDate   string  `json:"pub_date" mapstructure:"pub_date" gme:"title:出版日期;index:3"`
 Price     float32 `json:"price" gme:"title:定价;index:4"`
 SuitObj   string  `json:"suit_obj" mapstructure:"suit_obj" gme:"title:适用对象;index:5"`
 MajorType string  `json:"major_type" mapstructure:"major_type" gme:"title:图书类目;index:6"`
 SubMajor  string  `json:"sub_major" mapstructure:"sub_major" gme:"title:细分类目;index:7"`
}

func main() {
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

```
