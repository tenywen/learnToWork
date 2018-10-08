package ipc_service

import (
	. "gs/types"
)

var Services map[int16]func(*Session, []byte) []byte = map[int16]func(*Session, []byte) []byte{
	1: IPC_Notify,
}

func IPC_Notify(sess *Session, data []byte) []byte {
	return data
}
