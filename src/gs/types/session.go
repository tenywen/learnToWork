package types

import (
	"gs/types/build"
	"gs/types/item"
	"gs/types/research"
	"gs/types/user"
)

const (
	SESSION_KICKOUT = 0x1
	SESSION_LOGIN   = 0x2
	SESSION_LOGOUT  = 0x3
)

type Session struct {
	Id         int64
	Flag       int
	Fd         int64
	Idx        int64
	UpdateTime int64
	IP         string
	User       *user.User
	Builds     *build.Manager
	Researches *research.Manager
	Items      *item.Manager
}
