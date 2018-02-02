package main

import (
	"flag"
	"google.golang.org/grpc"
	"net"
	"os"
)

import (
	"cfg"
	"db"
	. "logger"
	pb "protocol"
)

var id int16

func init() {
	id = int16(*(flag.Int("i", 1, "gs_id")))
}

func main() {
	flag.Parse()

	gsConfig := cfg.GetGsConfig(id)
	if gsConfig != nil {
		StartLogger(gsConfig.LogName)
		go startAgent()
		db.StartDB(gsConfig.DB)
		listen, err := net.Listen("tcp", gsConfig.Port)
		if err != nil {
			ERROR(err)
			os.Exit(-1)
		}
		INFO("GS START")
		var opts []grpc.ServerOption
		server := grpc.NewServer(opts...)
		pb.RegisterServiceServer(server, service{})
		server.Serve(listen)
	}
	INFO("GS STOP!")
}
