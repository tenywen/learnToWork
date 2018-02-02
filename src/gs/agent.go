package main

import (
	"time"
)

import (
	"gs/closure"
	"gs/gsdb"
	. "gs/types"
	"gs/types/build"
	"gs/types/item"
	"gs/types/research"
	"gs/writer"
	. "logger"
	pb "protocol"
	"timer"
)

const (
	timeout = 600 // s
)

func startAgent() {
	closure.TimerClosure(time.Now().Unix()+60, nil, nil, closure.TIMER_minHandle)
	for {
		select {
		case msg := <-in: // from gate
			handle(msg)
		case f := <-closure.Ch: // from internal
			f()
		case f := <-timer.Ch: // from timer
			f()
		}
	}
}

func handle(msg *pb.MSG) {
	now := time.Now().Unix()
	var sess *Session

	switch msg.Api {
	case pb.PROTO_LOGIN:
		sess = &Session{
			Fd:         msg.Fd,
			Idx:        msg.Idx,
			Id:         msg.Uid,
			Builds:     &build.Manager{},
			Items:      &item.Manager{},
			Researches: &research.Manager{},
		}
	default:
		sess = gsdb.Get(msg.Uid)
	}

	handler := gsHandlers[msg.Api]
	if handler == nil {
		ERROR("api:", msg.Api, "not exist")
		return
	}

	result, err := handler(sess, msg.Data)
	if result != nil {
		writer := writer.Get(sess.Idx)
		if writer != nil {
			select {
			case writer <- &pb.MSG{
				Fd:   sess.Fd,
				Data: result,
				Err:  err,
			}:
			default:
			}
		}
	}
	sess.UpdateTime = now
	return
}
