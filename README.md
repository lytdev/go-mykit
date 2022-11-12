# 说明

> 打造类似`hutool`的 go 工具箱

## 工具概览

### 支持excel的读取和写入

- `gexcel.ReadFileToList`读取本地excel文件至切片
- `gexcel.ReadFileStreamToList`读取excel文件流至切片
- `gexcel.WriteToFile`写入片切数据至excelize.File对象

### 日期格式化

- `gfrmt.GetFormatDateStr`常见的日期字符串格式转time

### 对象转换

- `gmap2struct.Decode`map转结构体