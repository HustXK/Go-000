package slidingWindow

import (
	"context"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"testing"
	"time"
)

func Test_slidingWindow_Allow(t *testing.T) {
	l := NewSlidingWindow(10)
	t.Error("start")
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			run1(t, l)
			return nil
		})
	}
	_ = eg.Wait()

}

func run1(t *testing.T, l *slidingWindow) {
	var value int64 = 0
	for i := 0; i < 100; i++ {
		if l.Allow() {
			atomic.AddInt64(&value, 1)
			l.AddItem(value)
			t.Log("add value: ", value)
		}
		time.Sleep(time.Second)
	}
}
