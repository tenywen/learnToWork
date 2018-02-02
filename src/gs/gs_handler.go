package main

import ()

import (
	. "gs/AI"
	. "gs/types"
	pb "protocol"
)

var gsHandlers map[pb.PROTO]func(*Session, []byte) ([]byte, pb.ERR) = map[pb.PROTO]func(*Session, []byte) ([]byte, pb.ERR){
	pb.PROTO_LOGIN:  P_login_req,
	pb.PROTO_LOGOUT: P_logout_req,
	pb.PROTO_HEART:  P_heart_req,
}
