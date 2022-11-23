package gdownload

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

type HttpReader struct {
	Url string
}

// GetFileSize 获取文件大小,单位B,1KB=1024B
func (h HttpReader) GetFileSize() (int64, error) {
	resp, err := http.Get(h.Url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	fileSize := resp.ContentLength
	if fileSize < 1 {
		return fileSize, errors.New("file size is error")
	}
	return fileSize, nil
}

// OpenRange 获取请求数据指定位置的指定大小
func (h HttpReader) OpenRange(offset, size int64) (io.ReadCloser, error) {
	request, err := http.NewRequest("GET", h.Url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(
		"Range",
		"bytes="+strconv.FormatInt(offset, 10)+"-"+strconv.FormatInt(offset+size-1, 10),
	)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
