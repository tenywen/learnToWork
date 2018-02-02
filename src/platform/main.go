package main

import (
	"google.golang.org/grpc"
	"net"
	"os"
)

import (
	"cfg"
	. "logger"
	pb "protocol"
)

func main() {
	listen, err := net.Listen("tcp", config.Port)
	if err != nil {
		ERROR(err)
		os.Exit(-1)
	}
	INFO("PLATFORM START")
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterServiceServer(server, service{})
	server.Serve(listen)
}
