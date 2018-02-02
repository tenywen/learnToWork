package AI

import (
	. "gs/types"
	. "logger"
	pb "protocol"
)

func P_heart_req(sess *Session, data []byte) ([]byte, pb.ERR) {
	DEBUG(sess.User.Id, "heart req")
	return nil, pb.ERR_OK
}
