package main

import (
	"container/list"
)

type session struct {
	readerBuffer []byte
	readerEnd    int
	readerStart  int
	fd           int
	packetTime   int64
	uid          int64
	flag         int16
	serverId     int16
	sendList     *list.List

	// stat
	cnt   int64
	delay int64
}

func newsession(fd int) *session {
	return &session{
		fd: fd,
		//sendList:     list.New(),
		readerBuffer: make([]byte, dataMax),
		readerEnd:    headerSize,
	}
}
