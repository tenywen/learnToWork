package main

import (
	"time"
)

import (
	. "closure"
	"gs/gsdb"
	. "gs/types"
	"gs/types/build"
	"gs/types/item"
	"gs/types/research"
	. "logger"
	pb "protocol"
	"timer"
)

const (
	timeout = 600 // s
)

func startAgent(stream pb.Service_StreamServer, in chan *pb.MSG) {
	closure.TimerClosure(time.Now().Unix()+60, nil, nil, closure.TIMER_minHandle)
	mq := make(chan Closure, 1000)
	for {
		select {
		case msg := <-in: // from gate
			result := handle(gsdb.Get(msg.Uid), msg.Api, msg.Data)
			if result != nil {
				stream.Send(result)
			}
		case f := <-mq:
			result := f()
			if result != nil {

			}
		case msg := <-out:
			stream.Send(msg)
		case f := <-closure.Ch: // from internal
			f()
		case f := <-timer.Ch: // from timer
			f()
		}
	}
}

func handle(sess *Session, api pb.PROTO, data []byte) []byte {
	if api != pb.PROTO_LOGIN && sess == nil {
		return nil
	}
	handler := handlers[api]
	if handler == nil {
		ERROR("api:", api, "not exist")
		return
	}

	result, err := handler(sess, data)
	if result != nil {

	}
}
