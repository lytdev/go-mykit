package gdownload

import (
	"io"
	"net/http"
	"os"
)

// SingleDownload 原生的简单下载
func (d *Downloader) SingleDownload(wc *WriteCounter, url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)
	wc.total = int(resp.ContentLength)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	_, err = io.Copy(f, io.TeeReader(resp.Body, wc))
	return err
}
