package timer

import (
	"container/heap"
)

type TimeEvent struct {
	Timeout int64
	Func    func()
	Index   int
}

func NewTimer(timeout int64, f func()) *TimeEvent {
	lock.Lock()
	e := TimeEvent{timeout, f, -1}
	heap.Push(h, &e)
	lock.Unlock()
	return &e
}

func (e *TimeEvent) Cancel() {
	// 保证pop的event不会cancel
	if e.Index != -1 {
		lock.Lock()
		heap.Remove(h, e.Index)
		lock.Unlock()
	}
}

func (e *TimeEvent) Update(timeout int64) {
	lock.Lock()
	e.Timeout = timeout
	heap.Fix(h, e.Index)
	lock.Unlock()
}
