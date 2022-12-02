package gdownload

import (
	"io"
)

type Downloader struct {
	Workers int
	//PartSize string
	PartSize int64
	//BufSize  string
	BufSize int64
}

const name = "download"

var instance *Downloader

func (d *Downloader) GetName() string {
	return name
}

func Get() *Downloader {
	return instance
}

type FileReader interface {
	GetFileSize() (int64, error)
	OpenRange(offset, size int64) (io.ReadCloser, error)
}
