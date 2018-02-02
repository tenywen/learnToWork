package main

import (
	"encoding/binary"
	"syscall"
)

import (
	. "logger"
	pb "protocol"
)

var (
	handlers   map[int8]func(*session)
	workerPool []*worker
)

type task struct {
	op   int8
	fd   int
	data []byte
}

type worker struct {
	taskPool chan *task
	sessions map[int]*session
}

//----------------------------------------------------------  handle loop
func (w *worker) loop() {
	defer PANIC("worker loop")
	for {
		select {
		case t := <-w.taskPool:
			handler := handlers[t.op]
			if handler != nil {
				w.before_work(t)
				if sess, ok := w.sessions[t.fd]; ok {
					handler(sess)
					w.after_work(t)
				}
			}
		}
	}
}

func (w *worker) before_work(t *task) {
	if _, ok := w.sessions[t.fd]; !ok && t.op != opClose {
		w.sessions[t.fd] = newsession(t.fd)
	}

	switch t.op {
	case opWrite:
		if t.data != nil {
			if sess, ok := w.sessions[t.fd]; ok {
				sess.sendList.PushBack(t.data)
			}
		}
	}

}

func (w *worker) after_work(t *task) {
	switch t.op {
	case opClose:
		delete(w.sessions, t.fd)
	}
}

func startWorker(workNum int, bufferMax int) {
	handlers = make(map[int8]func(*session))
	handlers[opHandle] = handleConn
	handlers[opClose] = closed
	handlers[opWrite] = write

	for i := 0; i <= workNum; i++ {
		w := worker{
			taskPool: make(chan *task, bufferMax),
			sessions: make(map[int]*session),
		}
		workerPool = append(workerPool, &w)
		go w.loop()
	}
}

func handleConn(sess *session) {
	for {
		n, err := syscall.Read(sess.fd, sess.readerBuffer[sess.readerStart:sess.readerEnd])
		if n == 0 { // closed
			ERROR("fd closed ")
			dispatch(opClose, sess.fd, nil)
			break
		}

		if err == syscall.EAGAIN { // eof
			break
		}

		if err != nil {
			ERROR("fd ", err)
			dispatch(opClose, sess.fd, nil)
			break
		}

		sess.readerStart += n
		if sess.readerStart == sess.readerEnd {
			data := sess.readerBuffer[:sess.readerEnd]
			sess.readerBuffer = sess.readerBuffer[sess.readerEnd:]
			sess.readerStart = 0
			packetSize := int(headerSize)
			switch len(data) {
			case headerSize:
				packetSize = int(binary.LittleEndian.Uint16(data))
			default:
				handle(sess, data)
			}
			if len(sess.readerBuffer) < packetSize {
				sess.readerBuffer = make([]byte, dataMax)
			}
			sess.readerEnd = sess.readerStart + packetSize
		}
	}
}

func closed(sess *session) {
	syscall.Close(sess.fd)
	epollEvent(sess.fd, syscall.EPOLL_CTL_DEL, evRead)
	DEBUGF("fd %v stat cnt:%v avg:%v", sess.fd, sess.cnt, sess.delay/sess.cnt)

	stream := getStream(sess.serverId)
	if stream == nil {
		ERROR("gs not exist", sess.serverId)
		return
	}
	msg := pb.MSG{
		Fd:   int64(sess.fd),
		Data: nil,
		Api:  pb.PROTO_LOGOUT,
		Uid:  sess.uid,
	}
	DEBUGF("%v logout", sess.uid)
	stream.Send(&msg)
}

func write(sess *session) {
	for e := sess.sendList.Front(); e != nil; e = e.Next() {
		data, _ := e.Value.([]byte)
		length := len(data)
		n, err := syscall.Write(sess.fd, data)
		if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK || err == syscall.EINTR {
			DEBUGF("write err:%v", err)
			epollEvent(sess.fd, syscall.EPOLL_CTL_MOD, evWrite)
			return
		}

		if err != nil {
			dispatch(opClose, sess.fd, nil)
			return
		}

		if n != length {
			DEBUGF("write next time.res:%v", string(data))
			data = data[n:]
			e.Value = data
			epollEvent(sess.fd, syscall.EPOLL_CTL_MOD, evWrite)
			DEBUGF("write next time.res:%v", string(data))
			return
		}
		sess.sendList.Remove(e)
	}
}
