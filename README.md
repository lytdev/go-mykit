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

结构体`TAG` 中

`gme->index`: 索序号(从`0`开始),如果不写默认和结构体的字段顺序一致(从`0`开始)

`gme->title`: 名称,如果不写则使用结构体的字段属性名

Excel 对应的结构体参考：

```golang
type ExcelBook struct {
 Isbn      string  `gme:"title:ISBN;index:0"`
 BookName  string  `gme:"title:书名;index:1"`
 Author    string  `gme:"title:作者;index:2"`
 PubDate   string  `gme:"title:出版日期;index:3"`
 Price     float32 `gme:"title:定价;index:4"`
 SuitObj   string  `gme:"title:适用对象;index:5"`
 MajorType string  `gme:"title:图书类目;index:6"`
 SubMajor  string  `gme:"title:细分类目;index:7"`
}
```

如果要执行案例,先引入案例依赖

```bin
go get github.com/xuri/excelize/v2
```

使用文档：<https://github.com/lytdev/go-myexcel/wiki>
