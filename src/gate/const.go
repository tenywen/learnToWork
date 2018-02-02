package main

import (
	"syscall"
)

const (
	evRead     = uint32((syscall.EPOLLET | syscall.EPOLLIN) & 0xFFFFFFFF)
	evListen   = uint32((syscall.EPOLLET | syscall.EPOLLIN) & 0xFFFFFFFF)
	evWrite    = uint32((syscall.EPOLLET | syscall.EPOLLOUT | syscall.EPOLLONESHOT) & 0xFFFFFFFF)
	evMax      = 65536
	dataMax    = 65536
	headerSize = 2

	SESSION_KICKOUT = 0x01
	SESSION_LOGIN   = 0x02

	opHandle = iota
	opWrite
	opClose
)
