package gsdb

import (
	. "gs/types"
)

var onlines map[int64]*Session

func init() {
	onlines = make(map[int64]*Session)
}

func Online(sess *Session) {
	onlines[sess.Id] = sess
}

func Offline(uid int64) {
	delete(onlines, uid)
}

func Get(uid int64) *Session {
	if sess, ok := onlines[uid]; ok {
		return sess
	}
	return nil
}

func IsExit(uid int64) bool {
	if _, ok := onlines[uid]; ok {
		return true
	}
	return false
}

func GetAll() []*Session {
	result := make([]*Session, 0, len(onlines))
	for k := range onlines {
		result = append(result, onlines[k])
	}
	return result
}
