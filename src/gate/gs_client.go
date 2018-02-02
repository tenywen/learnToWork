package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"sync"
)

import (
	"cfg"
	. "logger"
	pb "protocol"
)

var gLock sync.RWMutex
var gsStreams map[int16]pb.Service_StreamClient

func init() {
	gsStreams = make(map[int16]pb.Service_StreamClient)
}

func gsDial() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	for _, config := range cfg.Config.GS {
		conn, err := grpc.Dial(config.Ip+config.Port, opts...)
		if err != nil {
			DEBUGF("dail gs %v", err)
			os.Exit(-1)
		}
		var callOpts []grpc.CallOption
		client := pb.NewServiceClient(conn)
		stream, err := client.Stream(context.Background(), callOpts...)
		if err != nil {
			DEBUGF("client stream %v", err)
			os.Exit(-1)
		}

		go func() {
			defer closeStream(config.Id)
			for {
				msg, err := stream.Recv()
				if err != nil {
					ERRORF("gs client %v", err)
					return
				}
				if msg.Err == pb.ERR_USER_LOGOUT {
					DEBUGF("%v kicked out", msg.Uid)
					dispatch(opClose, int(msg.Fd), nil)
				} else {
					dispatch(opWrite, int(msg.Fd), msg.Data)
				}
			}
		}()
		setStream(stream, config.Id)
		DEBUGF("dial gs %v succ!", config.Id)
	}
}

func getStream(server_id int16) pb.Service_StreamClient {
	gLock.RLock()
	defer gLock.RUnlock()
	if stream, ok := gsStreams[server_id]; ok {
		return stream
	}
	return nil
}

func setStream(stream pb.Service_StreamClient, server_id int16) {
	gLock.Lock()
	defer gLock.Unlock()
	gsStreams[server_id] = stream
}

func closeStream(server_id int16) {
	gLock.Lock()
	defer gLock.Unlock()
	if _, ok := gsStreams[server_id]; ok {
		delete(gsStreams, server_id)
	}
}
