package gdownload

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Listener struct {
}

func (l Listener) ProgressChanged(event *ProgressEvent) {
	fmt.Println(event)
}

func TestDownloadSingle(t *testing.T) {
	onWatch := func(current, total int, percentage float64) {
		fmt.Printf("\r当前已下载大小 %f MB, 下载进度：%.2f%%, 总大小 %f MB",
			float64(current)/1024/1024,
			percentage,
			float64(total)/1024/1024,
		)
	}
	url := "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"
	downloader := Downloader{}
	err := downloader.Download(url, "../testdata/example2.mp4", onWatch)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDownloadMulti(t *testing.T) {
	dl := Downloader{
		Workers:  5,
		PartSize: 1024 * 1024,
		BufSize:  1024 * 500,
	}
	httpReader := HttpReader{Url: "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"}
	err := dl.MultiDownload(context.Background(), "../testdata/example1.mp4", &httpReader, &Listener{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestTimeout(t *testing.T) {
	dl := Downloader{
		Workers:  5,
		PartSize: 1024 * 1024 * 5,
		BufSize:  1024 * 200,
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	httpReader := HttpReader{Url: "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"}
	err := dl.MultiDownload(ctx, "../testdata/example2.mp4", &httpReader, &Listener{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
