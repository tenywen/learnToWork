package main

import (
	"io"
)

import (
	. "logger"
	pb "protocol"
)

type service struct{}

func (s service) Stream(stream pb.Service_StreamServer) error {
	die := make(chan struct{})
	in := make(chan *pb.MSG, 100000)

	go startAgent(stream, in)

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
