# 说明


> 打造类似`hutool`的 go 工具箱

## 安装
go get -u github.com/lytdev/go-mykit

## 工具概览
>更多方法参考[官方文档](https://pkg.go.dev/github.com/lytdev/go-mykit)

### uuid
- `guid.FastUuid()` 使用Uuid4生成不带-的uuid
- `snowflake, _ := guid.NewSnowflake(0)` 创建雪花算法生成器

### datetime
> 日期时间推荐使用[carbon](https://github.com/golang-module/carbon)
- `gtime.FormatDateTimeToStr` 提取日期为统一格式 yyyy-mm-dd hh:mm:ss
- `gtime.FormatDurationToSecond` 将持续时间转为秒数 1:01:03
- `gtime.TimeToStrAsFormat` 获取时间字符串
- `gtime.TimeToTimeStampMill` 时间转毫秒级别时间戳
- `gtime.TimestampSecToTime` 秒级别时间戳转时间

### 文件操作
- `gfile.ReadWithLine` 按行读取文件的文本
- `gfile.CopyFile` 复制文件
- `gfile.FileDir` 获取文件所在的路径
- `gfile.MainName` 获取文件的名称,不带后缀

### 加密解密
> 加密解密推荐使用[dongle](https://github.com/golang-module/dongle)

### 支持excel的读取和写入

- `gexcel.ReadFileToList(filePath, 0, ptr)(resultData []T, err error)`读取本地excel文件至切片
- `gexcel.ReadFileStreamToList`读取excel文件流至切片
- `gexcel.WriteToFile(sheetName, dataList)(f *excelize.File,err error)`写入片切数据至excelize.File对象

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
#### 小文件下载
```
wc := new(WriteCounter)
wc.SetWatch(func(current, total int, percentage float64) {
	fmt.Printf("\r当前已下载大小 %f MB, 下载进度：%.2f%%, 总大小 %f MB",
		float64(current)/1024/1024,
		percentage,
		float64(total)/1024/1024,
	)
})
downloader := NewWithSingle()

err := downloader.SingleDownload(wc, downloadUrl, "../testdata/example2.mp4")
if err != nil {
	fmt.Println(err)
	return
}
```