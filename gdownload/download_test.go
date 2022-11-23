package gdownload

import (
	"context"
	"fmt"
	"github.com/lytdev/go-mykit/helpers/progress"
	"testing"
	"time"
)

type Listener struct {
}

func (l Listener) ProgressChanged(event *hprogress.ProgressEvent) {
	fmt.Println(event)
}

func TestDownload(t *testing.T) {
	dl := Instance{
		Workers:  5,
		PartSize: 1024 * 1024 * 5,
		BufSize:  1024 * 500,
	}
	httpReader := HttpReader{Url: "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"}
	err := dl.Download(context.Background(), "../testdata/example1.mp4", &httpReader, &Listener{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestTimeout(t *testing.T) {
	dl := Instance{
		Workers:  5,
		PartSize: 1024 * 1024 * 5,
		BufSize:  1024 * 200,
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	httpReader := HttpReader{Url: "https://playback-tc.videocc.net/polyvlive/76490dba387702307790940685/f0.mp4"}
	err := dl.Download(ctx, "../testdata/example2.mp4", &httpReader, &Listener{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
