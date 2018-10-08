package ipc

import (
	. "closure"
	"gs/gsdb"
	"gs/ipc_service"
	. "gs/types"
)

func Send(sess *Session, serviceId int16, data []byte) {
	defer func() {
		if x := recover(); x != nil {

		}
	}()

	select {
	case sess.MQ <- func() []byte {
		return ipc_service.Services[serviceId](sess, data)
	}:
	default:
	}
}
