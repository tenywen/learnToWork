package writer

import (
	"sync"
)

import (
	pb "protocol"
)

var writers map[int64]chan *pb.MSG
var gLock sync.RWMutex
var i int64

func init() {
	writers = make(map[int64]chan *pb.MSG)
}

func New(die chan struct{}, stream pb.Service_StreamServer) int64 {
	in := make(chan *pb.MSG, 1000)
	idx := set(in)
	go func() {
		defer closeWriter(idx)
		for {
			select {
			case msg := <-in: // to gate
				stream.Send(msg)
			case <-die:
				return
			}
		}
	}()
	return idx
}

func set(in chan *pb.MSG) int64 {
	gLock.Lock()
	defer gLock.Unlock()
	i++
	writers[i] = in
	return i
}

func Get(idx int64) chan *pb.MSG {
	gLock.RLock()
	defer gLock.RUnlock()
	if in, ok := writers[idx]; ok {
		return in
	}
	return nil
}

func closeWriter(idx int64) {
	gLock.Lock()
	defer gLock.Unlock()
	delete(writers, idx)

}
