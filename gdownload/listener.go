package gdownload

import "sync"

// WriteCounter 下载监听器
type WriteCounter struct {
	current int
	total   int
	onWatch func(current, total int, percentage float64)
	sync.Mutex
}

// SetWatch 设置下载监听
func (w *WriteCounter) SetWatch(onWatch func(current, total int, percentage float64)) *WriteCounter {
	w.onWatch = onWatch
	return w
}

// getWritePercentage 获取进度
func (wc *WriteCounter) getWritePercentage() float64 {
	return float64(wc.current*10000/wc.total) / 100
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	if wc.onWatch != nil {
		wc.Lock()
		defer wc.Unlock()
		wc.current += n
		wc.onWatch(wc.current, wc.total, wc.getWritePercentage())
	}
	return n, nil
}
