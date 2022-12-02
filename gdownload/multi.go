package gdownload

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

// MultiDownload 原生的分片下载
func (d *Downloader) MultiDownload(wc *WriteCounter, url, filename string, contentLen int) error {
	wc.total = contentLen
	partSize := contentLen / d.concurrency
	//开启线程组
	wg := sync.WaitGroup{}
	wg.Add(d.concurrency)
	rangeStart := 0
	for i := 0; i < d.concurrency; i++ {
		go func(i, rangeStart int) {
			defer wg.Done()
			rangeEnd := rangeStart + partSize
			// 最后一部分,总长度不能超过 ContentLength
			if i == d.concurrency-1 {
				rangeEnd = contentLen
			}
			downloaded := 0
			// 断点续传
			if d.resume {
				content, err := ioutil.ReadFile(d.getPartFilename(filename, i))
				if err == nil {
					downloaded = len(content)
					wc.Write(content)
				}
			}
			d.downloadPartial(wc, url, filename, rangeStart+downloaded, rangeEnd, i)
		}(i, rangeStart)
		rangeStart += partSize + 1
	}
	wg.Wait()
	return d.merge(filename)
}

// 下载分片数据
func (d *Downloader) downloadPartial(wc *WriteCounter, url, filename string, rangeStart, rangeEnd, i int) {
	if rangeStart >= rangeEnd {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	flags := os.O_CREATE | os.O_WRONLY
	if d.resume {
		flags |= os.O_APPEND
	}

	partFile, err := os.OpenFile(d.getPartFilename(filename, i), flags, 0666)
	if err != nil {
		return
	}
	defer func(partFile *os.File) {
		_ = partFile.Close()

	}(partFile)

	_, err = io.Copy(partFile, io.TeeReader(resp.Body, wc))
	if err != nil {
		return
	}
}

// merge 合并分片
func (d *Downloader) merge(filename string) error {
	destFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func(destFile *os.File) {
		_ = destFile.Close()

	}(destFile)

	for i := 0; i < d.concurrency; i++ {
		partFileName := d.getPartFilename(filename, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(destFile, partFile)
		if err != nil {
			return err
		}
		err = partFile.Close()
		if err != nil {
			return err
		}
		err = os.Remove(partFileName)
		if err != nil {
			return err
		}
	}
	return nil
}

//获取分片的名称
func (d *Downloader) getPartFilename(filename string, partNum int) string {
	return fmt.Sprintf("%s_%d", filename, partNum)
}
