package gdownload

import (
	"io"
	"net/http"
	"os"
	"sync"
)

type writeCounter struct {
	current    int
	total      int
	percentage float64
	onWatch    func(current, total int, percentage float64)
	sync.Mutex
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	if wc.onWatch != nil {
		wc.Lock()
		defer wc.Unlock()
		wc.current += n
		wc.onWatch(wc.current, wc.total, float64(wc.current*10000/wc.total)/100)
	}
	return n, nil
}

// Download 进行简单下载
func (d *Downloader) Download(url, filename string, onWatch func(current, total int, percentage float64)) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	wc := new(writeCounter)
	if onWatch != nil {
		wc.onWatch = onWatch
	}

	wc.total = int(resp.ContentLength)

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, io.TeeReader(resp.Body, wc))
	return err
}
