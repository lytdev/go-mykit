package gdownload

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/lytdev/go-mykit/gfile"
)

// Downloader 下载执行对象
type Downloader struct {
	// 是否开启并发下载 (若支持)
	multi bool
	// 并发协程数 (若支持,默认是cpu数)
	concurrency int
	// 断点续传 (若支持)
	resume bool
}

// NewWithSingle 创建简单下载对象
func NewWithSingle() *Downloader {
	return &Downloader{
		multi:       false,
		concurrency: 1,
		resume:      false,
	}
}

// NewWithMulti 创建分片下载对象
func NewWithMulti(runtimeNum int) *Downloader {
	d := &Downloader{
		multi:       true,
		concurrency: runtime.NumCPU(),
		resume:      true,
	}
	if runtimeNum > 0 {
		d.concurrency = runtimeNum
	}
	return d
}

// 设置是否多线程下载
func (d *Downloader) SetMulti(multi bool) *Downloader {
	d.multi = multi
	return d
}

// 设置下载的多线程数
func (d *Downloader) SetConcurrency(concurrency int) *Downloader {
	d.concurrency = concurrency
	return d
}

// 设置是否断点续传
func (d *Downloader) SetResume(resume bool) *Downloader {
	d.resume = resume
	return d
}

// 创建文件
func initFilePath(fp string) string {
	result := fp
	renameNum := 1
	mainName := gfile.MainName(fp)
	for {
		if gfile.IsExist(fp) {
			if len(gfile.ExtName(fp)) > 0 {
				fp = mainName + "_" + strconv.Itoa(renameNum) + "." + gfile.ExtName(fp)
			} else {
				fp = mainName + "_" + strconv.Itoa(renameNum)
			}
			renameNum += 1
		} else {
			break
		}
	}
	return result
}

// Download 下载文件
func (d *Downloader) Download(url, fp string, overwrite bool, onWatch func(current, total int, percentage float64)) error {
	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download fail: %s", resp.Status)
	}

	if len(fp) == 0 {
		return errors.New("file path is error")
	}
	if err := gfile.MkdirAll(fp); err != nil {
		return err
	}

	wc := new(WriteCounter)
	if onWatch != nil {
		wc.onWatch = onWatch
	}
	if !overwrite {
		fp = initFilePath(fp)
	}
	if d.multi && resp.Header.Get("Accept-Ranges") == "bytes" {
		// 支持分段下载
		return d.MultiDownload(wc, url, fp, int(resp.ContentLength))
	}

	return d.SingleDownload(wc, url, fp)
}
