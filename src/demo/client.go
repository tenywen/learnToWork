package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
	"os"
	"time"
)

import (
	//. "logger"
	pb "protocol"
)

func main() {
	conn, err := net.Dial("tcp", ":8082")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	login_info := pb.LOGIN_REQ{
		Uuid:     "test1",
		Account:  "test1",
		Pwd:      "test1",
		Serverid: 1,
	}

	b, _ := proto.Marshal(&login_info)
	conn.Write(pack(b, pb.PROTO_LOGIN))

	// recv
	go func() {
		header := make([]byte, 2)
		for {
			n, err := io.ReadFull(conn, header)
			if err != nil {
				return
			}
			size := binary.LittleEndian.Uint16(header[:n])
			data := make([]byte, int(size))
			n, err = io.ReadFull(conn, data)
			if err != nil {
				return
			}

			fmt.Println(string(data[:n]))
			time.Sleep(10000 * time.Millisecond)
		}
	}()
	for {
		req := pb.TEST_GATE_REQ{time.Now().UnixNano()}
		b, _ := proto.Marshal(&req)
		_, err := conn.Write(pack(b, pb.PROTO_TEST_GATE))
		if err != nil {
			break
		}
		time.Sleep(10 * time.Nanosecond)
	}
}

func pack(data []byte, api pb.PROTO) []byte {
	result := make([]byte, len(data)+4)
	binary.LittleEndian.PutUint16(result[:2], uint16(len(data))+2)
	binary.LittleEndian.PutUint16(result[2:4], uint16(api))
	copy(result[4:], data)
	return result
}
