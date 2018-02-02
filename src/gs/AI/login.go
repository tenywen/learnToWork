package AI

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

import (
	"db/data_tbl"
	"db/user_tbl"
	"gs/closure"
	"gs/gsdb"
	. "gs/types"
	"gs/types/build"
	"gs/types/item"
	//"gs/types/research"
	. "gs/types/user"
	. "logger"
	"misc"
	pb "protocol"
)

const (
	dbName = "gs"
)

//---------------------------------------------------------- 登录
func P_login_req(sess *Session, data []byte) ([]byte, pb.ERR) {
	if gsdb.IsExit(sess.Id) {
		// TODO 多次登录
		return nil, pb.ERR_OK
	}

	var user *User
	user = user_tbl.Get(dbName, sess.Id)
	if user == nil {
		// new
		user = &User{
			Id:    sess.Id,
			Name:  "new_" + fmt.Sprint(sess.Id),
			UUID:  "",
			Level: 1,
		}
		user_tbl.Save(dbName, user)
	}
	sess.User = user
	Login(sess)
	PM12(sess)
	gsdb.Online(sess)

	result := pb.LOGIN_ACK{
		Uid:   sess.User.Id,
		Level: sess.User.Level,
		Name:  sess.User.Name,
	}

	b, _ := proto.Marshal(&result)
	DEBUG(sess.User.Id, " login")
	// TODO ipc add item
	// test
	closure.IPCClosure([]int64{sess.User.Id}, nil, closure.IPC_add_item)
	return b, pb.ERR_OK
}

func Login(sess *Session) bool {
	// load build
	if !data_tbl.Get(dbName, data_tbl.BuildCollection, sess.User.Id, &sess.Builds) {
		sess.Builds = &build.Manager{
			UserId: sess.User.Id,
			Builds: make([]*build.Build, 0),
		}
		data_tbl.Set(dbName, data_tbl.BuildCollection, sess.User.Id, sess.Builds)
	}

	// load item
	if !data_tbl.Get(dbName, data_tbl.ItemCollection, sess.User.Id, &sess.Items) {
		sess.Items = &item.Manager{
			UserId: sess.User.Id,
			Items:  make([]*item.Item, 0),
		}
		data_tbl.Set(dbName, data_tbl.ItemCollection, sess.User.Id, sess.Items)
	}

	// load research
	/*
		if !data_tbl.Get(dbName, data_tbl.ResearchCollection, sess.User.Id, &sess.Researches) {
			sess.Researches = &research.Manager{
				UserId:     sess.User.Id,
				Researches: make([]*research.Research, 0),
			}
			data_tbl.Set(dbName, data_tbl.ResearchCollection, sess.User.Id, sess.Researches)
		}
	*/
	sess.Flag |= SESSION_LOGIN
	return true
}

func PM12(sess *Session) {
	now := time.Now().Unix()
	if sess.User.LoginTime > misc.PM12(now) { // today

	} else { // next day

	}

}
