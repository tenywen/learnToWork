package AI

import (
	"db/data_tbl"
	"db/user_tbl"
	"gs/gsdb"
	. "gs/types"
	. "logger"
	pb "protocol"
)

func P_logout_req(sess *Session, data []byte) ([]byte, pb.ERR) {
	if sess.Flag&SESSION_LOGIN == 0 {
		return nil, pb.ERR_USER_NOT_LOGIN
	}

	user_tbl.Save(dbName, sess.User)
	data_tbl.Set(dbName, data_tbl.BuildCollection, sess.User.Id, sess.Builds)
	data_tbl.Set(dbName, data_tbl.ItemCollection, sess.User.Id, sess.Items)
	//data_tbl.Set(dbName, data_tbl.ResearchCollection, sess.User.Id, sess.Items)
	gsdb.Offline(sess.Id)
	DEBUG(sess.User.Id, "logout")
	return nil, pb.ERR_OK
}
