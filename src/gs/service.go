package main

import (
	"io"
)

import (
	"gs/writer"
	. "logger"
	pb "protocol"
)

type service struct{}

var in chan *pb.MSG

func init() {
	in = make(chan *pb.MSG, 100000)
}

func (s service) Stream(stream pb.Service_StreamServer) error {
	die := make(chan struct{})
	idx := writer.New(die, stream)
	DEBUG(idx, " stream create")
	defer func() {
		DEBUGF("%v stream end!", idx)
		close(die)
	}()
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			DEBUG("gs stream closed", err)
			return nil
		}

		if err != nil {
			ERROR("gs stream", err)
			return err
		}

		msg.Idx = idx
		in <- msg
	}
}
