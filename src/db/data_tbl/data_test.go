package data_tbl

import (
	"fmt"
	"testing"
)

import (
	. "types"
	"types/build"
	"types/user"
)

func TestGet(t *testing.T) {
	sess := &Session{
		User: &user.User{
			Id: 3,
		},
		Builds: &build.Manager{},
	}
	dbName := "gs"
	Get(dbName, BuildCollection, sess.User.Id, &sess.Builds)
	fmt.Println(sess.Builds)
}
