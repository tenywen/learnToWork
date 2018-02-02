package timer

type Heap []*TimeEvent

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool { return h[i].Timeout < h[j].Timeout }

func (h Heap) Swap(i, j int) {
	h[i].Index = j
	h[j].Index = i
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	x.(*TimeEvent).Index = len(*h)
	*h = append(*h, x.(*TimeEvent))
}

func (h *Heap) Pop() interface{} {
	n := len(*h)
	e := (*h)[n-1]
	e.Index = -1
	*h = (*h)[:n-1]
	return e
}
