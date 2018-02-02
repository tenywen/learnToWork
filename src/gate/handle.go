package main

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
)

import (
	pb "protocol"
)

func handle(sess *session, data []byte) {
	if len(data) < 2 {
		return
	}

	var result []byte
	api := pb.PROTO(binary.LittleEndian.Uint16(data[:2]))

	switch api {
	case pb.PROTO_LOGIN:
		result = login_req(sess, data[2:])
	case pb.PROTO_LOGOUT:
		result = logout_req(sess, nil)
	case pb.PROTO_TEST_GATE:
		result = test_gate(sess, data[2:])
	default:
		result = forward(sess, api, data[2:])
	}

	if result != nil {
		dispatch(opWrite, sess.fd, result)
	}

	if sess.flag&SESSION_KICKOUT != 0 {
		dispatch(opClose, sess.fd, nil)
	}
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
