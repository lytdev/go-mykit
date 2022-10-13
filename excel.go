/*
 * @Author       : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @Date         : 2022-10-11 17:47:08
 * @LastEditors  : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @FilePath     : \go-myexcel\excel.go
 * @Description  :
 * Copyright (c) 2022 by 刘元涛 snoopy_718@mails.ccnu.edu.cn, All Rights Reserved.
 */
package gme

//By default, it starts from the first row and the index starts from 0
//默认 从第一行开始,索引从 0开始
func NewExcelStructDefault() *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = 1
	n.IndexMax = 10
	return n
}

//StartRow starts row, index starts from 0
//Indexmax indexes the maximum row. If the index in the structure is larger than the configured, the index in the structure is used
//StartRow 开始行,索引从 0开始
//IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
func NewExcelStruct(StartRow, IndexMax int) *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = StartRow
	n.IndexMax = IndexMax
	return n
}
