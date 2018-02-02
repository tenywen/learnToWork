package timer

import (
	"container/heap"
	"sync"
	"time"
)

const (
	INTERVAL = 1 // s
)

var h *Heap
var lock sync.Mutex
var Ch chan func()

func init() {
	Ch = make(chan func(), 100000)
	h = &Heap{}
	heap.Init(h)
	go schedule()
}

func schedule() {
	// 可以调整定时器精度
	ticker := time.NewTicker(INTERVAL * time.Second)
	for {
		select {
		case t := <-ticker.C:
			lock.Lock()
			for h.Len() > 0 {
				// 随精度一起调整
				if (*h)[0].Timeout > t.Unix() {
					break
				}
				e := heap.Pop(h).(*TimeEvent)
				select {
				case Ch <- e.Func:
				default:
				}
			}
			lock.Unlock()
		}
	}
}
