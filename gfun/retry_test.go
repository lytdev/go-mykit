package gfun

import (
	"sync"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	i := 0
	var calls []int
	mu := sync.Mutex{}

	debounce, cancel := NewDebounce(time.Millisecond*5, func() {
		mu.Lock()
		defer mu.Unlock()
		calls = append(calls, i)
	})
	//5秒后执行函数
	debounce()
	i++
	//这个过程中执行了一次
	time.Sleep(7 * time.Millisecond)

	debounce()
	i++
	debounce()
	i++
	debounce()
	i++

	time.Sleep(6 * time.Millisecond)

	cancel()
	t.Logf("%v", calls)
}
