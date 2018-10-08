package main

import (
	"encoding/binary"
	"syscall"
)

import (
	. "closure"
	. "logger"
	pb "protocol"
)

var (
	workerPool []*worker
)

type worker struct {
	ch       chan Closure
	sessions map[int]*session
}

//----------------------------------------------------------  handle loop
func (w *worker) loop() {
	defer PANIC("worker loop")
	for {
		select {
		case f := <-w.ch:
			f()
		}
	}
}

func startWorker(workNum int) {
	for i := 0; i <= workNum; i++ {
		w := worker{
			ch:       make(chan Closure, 1024),
			sessions: make(map[int]*session),
		}
		workerPool = append(workerPool, &w)
		go w.loop()
	}
}

func (w *worker) handleConn(fd int) []byte {
	var sess *session
	if _, ok := w.sessions[fd]; !ok {
		w.sessions[fd] = newsession(fd)
	}

	sess = w.sessions[fd]
	for {
		n, err := syscall.Read(sess.fd, sess.readerBuffer[sess.readerStart:sess.readerEnd])
		if n == 0 { // closed
			ERROR("fd closed ")
			w.close(fd)
			break
		}

		if err == syscall.EAGAIN { // eof
			break
		}

		if err != nil {
			ERROR("fd ", err)
			w.close(fd)
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
				result := handle(sess, data)
				if result != nil {
					w.write(fd, result)
				}

				if sess.flag&SESSION_KICKOUT != 0 {
					return w.close(fd)
				}
			}
			if len(sess.readerBuffer) < packetSize {
				sess.readerBuffer = make([]byte, dataMax)
			}
			sess.readerEnd = sess.readerStart + packetSize
		}
	}

	return nil
}

func (w *worker) close(fd int) []byte {
	syscall.Close(fd)
	epollEvent(fd, syscall.EPOLL_CTL_DEL, evRead)
	if sess, ok := w.sessions[fd]; ok {
		DEBUGF("fd %v stat cnt:%v avg:%v", sess.fd, sess.cnt, sess.delay/sess.cnt)

		stream := getStream(sess.serverId)
		if stream == nil {
			ERROR("gs not exist", sess.serverId)
			return nil
		}
		msg := pb.MSG{
			Fd:   int64(sess.fd),
			Data: nil,
			Api:  pb.PROTO_LOGOUT,
			Uid:  sess.uid,
		}
		DEBUGF("%v logout", sess.uid)
		stream.Send(&msg)
		delete(w.sessions, fd)
	}

	return nil
}

/*
func (w *worker) write(fd int, data []byte) []byte {
	if sess, ok := w.sessions[fd]; ok {
		for e := sess.sendList.Front(); e != nil; e = e.Next() {
			data, _ := e.Value.([]byte)
			length := len(data)
			n, err := syscall.Write(sess.fd, data)
			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK || err == syscall.EINTR {
				DEBUGF("write err:%v", err)
				epollEvent(sess.fd, syscall.EPOLL_CTL_MOD, evWrite)
				break
			}

			if err != nil {
				w.close(fd)
				break
			}

			if n != length {
				DEBUGF("write next time.res:%v", string(data))
				data = data[n:]
				e.Value = data
				epollEvent(sess.fd, syscall.EPOLL_CTL_MOD, evWrite)
				DEBUGF("write next time.res:%v", string(data))
				break
			}
			sess.sendList.Remove(e)
		}
	}
	return nil
}
*/

func (w *worker) write(fd int, data []byte) []byte {
	_, err := syscall.Write(fd, data)
	if err != nil {
		w.close(fd)
	}

	return nil
}
