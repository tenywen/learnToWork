package main

import (
	"syscall"
)

import (
	. "logger"
)

var (
	lis  int
	epfd int
)

func epollEvent(fd, op int, ev uint32) {
	syscall.EpollCtl(epfd, op, fd, &syscall.EpollEvent{Events: ev, Fd: int32(fd)})
}

func server(network string, addr string, workNum int, bufferMax int) {
	startWorker(workNum, bufferMax)
	// listen event
	sockAddr, domain := resolveTCPAddr(network, addr)
	lis, _ = syscall.Socket(domain, syscall.SOCK_STREAM|syscall.SOCK_NONBLOCK, syscall.IPPROTO_TCP)
	defer syscall.Close(lis)
	syscall.Bind(lis, sockAddr)
	syscall.SetNonblock(lis, true)
	syscall.Listen(lis, syscall.SOMAXCONN)
	epfd, _ = syscall.EpollCreate1(syscall.EPOLL_CLOEXEC)
	epollEvent(lis, syscall.EPOLL_CTL_ADD, evListen)
	DEBUG("GATE START!")
	// epoll event loop
	var events [evMax]syscall.EpollEvent
	for {
		n, err := syscall.EpollWait(epfd, events[:], -1) // -1 = block
		if err == syscall.EINTR {
			continue
		}
		if err != nil {
			break
		}

		for i := 0; i <= n; i++ {
			if events[i].Fd == int32(lis) {
				accept()
			} else if events[i].Events&syscall.EPOLLIN != 0 {
				dispatch(opHandle, int(events[i].Fd), nil)
			} else if events[i].Events&syscall.EPOLLOUT != 0 {
				dispatch(opWrite, int(events[i].Fd), nil)
			} else if events[i].Events&syscall.EPOLLRDHUP != 0 || events[i].Events&syscall.EPOLLHUP != 0 || events[i].Events&syscall.EPOLLERR != 0 {
				dispatch(opClose, int(events[i].Fd), nil)
			}
		}
	}
}

func accept() {
	for {
		fd, _, err := syscall.Accept(lis)
		if err != nil {
			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK { // eof
				break
			}
			if err == syscall.ECONNABORTED { // connect closed
				continue
			}
			break
		}
		syscall.SetNonblock(fd, true)
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, dataMax)
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_SNDBUF, dataMax)
		epollEvent(fd, syscall.EPOLL_CTL_ADD, evRead)
	}
}

func dispatch(op int8, fd int, data []byte) {
	t := &task{op, fd, data}
	w := getWorker(fd)
	select {
	case w.taskPool <- t:
	default:
	}
}

func getWorker(fd int) *worker {
	return workerPool[fd%len(workerPool)]
}
