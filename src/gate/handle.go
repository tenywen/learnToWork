package main

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
)

import (
	pb "protocol"
)

func handle(sess *session, data []byte) []byte {
	api := pb.PROTO(binary.LittleEndian.Uint16(data[:2]))

	switch api {
	case pb.PROTO_LOGIN:
		return login_req(sess, data[2:])
	case pb.PROTO_LOGOUT:
		return logout_req(sess, nil)
	case pb.PROTO_TEST_GATE:
		return test_gate(sess, data[2:])
	}

	return forward(sess, api, data[2:])
}

func selectServer(uuid string) int16 {
	return 1
}

func commonAck(api pb.PROTO, err pb.ERR) []byte {
	b, _ := proto.Marshal(&pb.COMMON_ACK{api, err})
	size := uint16(len(b))
	result := make([]byte, 2+size)
	binary.LittleEndian.PutUint16(result[:2], size)
	copy(result[2:], b)
	return result
}
