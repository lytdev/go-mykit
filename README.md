# 说明
>excelize/v2的增版版本，支持读取excel并转换为map和struct；核心部分 没有使用任何 第三方包,引入第三方包都是测试和转换使用的

# 使用方式
>go get -u github.com/lytdev/go-myexcel


结构体`TAG` 中 `json`,`meg` 必须存在

`json`: 字段名称,映射成`map`

`meg-index`: 索序号

`meg-title`: 名称
