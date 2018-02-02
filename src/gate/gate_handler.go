package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

import (
	"account"
	"db"
	"db/account_tbl"
	. "logger"
	pb "protocol"
)

func forward(sess *session, api pb.PROTO, data []byte) []byte {
	if sess.flag&SESSION_LOGIN == 0 {
		ERRORF("%v not login", sess.uid)
		return commonAck(api, pb.ERR_USER_NOT_LOGIN)
	}
	stream := getStream(sess.serverId)
	if stream == nil {
		ERRORF("gs %v not exist", sess.serverId)
		sess.flag |= SESSION_KICKOUT
		return commonAck(pb.PROTO_LOGIN, pb.ERR_SERVER_NOT_EXIST)
	}

	msg := pb.MSG{
		Fd:   int64(sess.fd),
		Uid:  sess.uid,
		Data: data,
		Api:  api,
	}

	err := stream.Send(&msg)
	if err != nil {
		ERROR("gs stream error")
		sess.flag &= SESSION_KICKOUT
	}
	DEBUGF("forward api:%v", pb.PROTO_name[int32(api)])
	return nil
}

func login_req(sess *session, data []byte) []byte {
	DEBUG("handle login_req")
	if sess.flag&SESSION_LOGIN != 0 {
		ERRORF("玩家%v 已经登录", sess.uid)
		return commonAck(pb.PROTO_LOGIN, pb.ERR_DATA)
	}
	tbl := pb.LOGIN_REQ{}
	err := proto.Unmarshal(data, &tbl)
	if err != nil {
		ERROR("login req ", err)
		return commonAck(pb.PROTO_LOGIN, pb.ERR_DATA)
	}
	user := account_tbl.GetByUUID(tbl.Uuid)
	if user == nil {
		uid := db.NextVal(db.NEXT_VAL, db.NEXT_VAL, db.UID)
		if uid == -1 {
			ERROR("NextVal uid err")
			return commonAck(pb.PROTO_LOGIN, pb.ERR_DATA)
		}
		user = &account.Account{
			UUID:     tbl.Uuid,
			UserName: tbl.Account,
			PWD:      tbl.Pwd,
			ServerId: int16(tbl.Serverid),
			UID:      uid,
		}
		account_tbl.Save(user)
	}

	sess.uid = user.UID
	sess.serverId = user.ServerId
	sess.flag |= SESSION_LOGIN
	DEBUGF("%v login gs:%v succ!", sess.uid, sess.serverId)
	return forward(sess, pb.PROTO_LOGIN, data)
}

func test_gate(sess *session, data []byte) []byte {
	tbl := pb.TEST_GATE_REQ{}
	err := proto.Unmarshal(data, &tbl)
	if err != nil {
		ERROR("test gate ", err)
		return nil
	}

	sess.cnt++
	sess.delay += (time.Now().UnixNano() - tbl.Start)

	b := []byte("this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test this is a test " + fmt.Sprint(sess.cnt))
	size := uint16(len(b))
	result := make([]byte, 2+size)
	binary.LittleEndian.PutUint16(result[:2], size)
	copy(result[2:], b)
	return result
}

func logout_req(sess *session, data []byte) []byte {
	DEBUGF("%v logout", sess.uid)
	sess.flag |= SESSION_KICKOUT
	return nil
}
