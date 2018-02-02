package timer

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &Heap{}
	heap.Init(h)
	heap.Push(h, &TimeEvent{3, nil, 3})
	heap.Push(h, &TimeEvent{2, nil, 4})
	heap.Push(h, &TimeEvent{1, nil, 2})

	for h.Len() > 0 {
		e := (*h)[0]
		fmt.Println(e.Index, e)
		heap.Pop(h)
	}
}
