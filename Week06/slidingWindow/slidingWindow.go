package slidingWindow

import (
	"sync"
	"time"
)

type slidingWindow struct {
	sTime time.Time
	maxDelay  int64
	pool      []* item
	mu        sync.Mutex
}

func NewSlidingWindow(maxDelay int64) *slidingWindow {
	return &slidingWindow{
		maxDelay : maxDelay,
		pool:      make([] *item, 1),
	}
}

type item struct {
	sTime time.Time
	value int64
}

func (s *slidingWindow) Allow() bool {
	now := time.Now()
	during := now.Sub(s.sTime)
	return during < (time.Duration(s.maxDelay))
}

func (s *slidingWindow) AddItem(i int64) {
	if i == 0 {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pool = append(s.pool, &item{sTime: time.Now(), value: i})
	s.delOldItem()
	s.sTime = time.Now()
}

func (s *slidingWindow) delOldItem() {
	now := time.Now().Unix()
	for index, v := range s.pool {
		if now - v.sTime.Unix() > s.maxDelay{
			continue
		} else{
			s.pool = s.pool[index:]
			break
		}
	}
}