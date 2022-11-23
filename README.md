# 说明

> 打造类似`hutool`的 go 工具箱

## 工具概览

### 支持excel的读取和写入

- `gexcel.ReadFileToList(filePath, 0, ptr)(resultData []T, err error)`读取本地excel文件至切片
- `gexcel.ReadFileStreamToList`读取excel文件流至切片
- `gexcel.WriteToFile(sheetName, dataList)(f *excelize.File,err error)`写入片切数据至excelize.File对象

### 日期格式化

- `gfrmt.GetFormatDateStr`常见的日期字符串格式转time

### 对象转换

- `gmap2struct.Decode`map转结构

### 文件下载
#### 大文件分片下载
```
import (
    "github.com/lytdev/go-mykit/gdownload"
    "github.com/lytdev/go-mykit/helpers/progress"
)

type Listener struct {
}

//实现监听接口
func (l Listener) ProgressChanged(event *hprogress.ProgressEvent) {
    fmt.Println(event)
}

download := gdownload.Instance{
    //5个线程进行下载
    Workers:  5,
    //每个分片5M
    PartSize: 1024 * 1024 * 5,
    //分片的缓存500KB 
    BufSize:  1024 * 500,
}
httpReader := gdownload.HttpReader{Url: "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"}
err := download.Download(context.Background(), "../testdata/example1.mp4", &httpReader, &Listener{})
if err != nil {
    fmt.Println(err)
    return
}
```
#### 小文件完成下载
