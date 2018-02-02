package closure

import (
	//	"github.com/vmihailenco/msgpack"
	"time"
)

import (
	"gs/gsdb"
	"gs/writer"
	. "logger"
	pb "protocol"
	"timer"
)

const (
	timeout = 600 // s
)

type timerhandle func([]int64, []byte)

func TimerClosure(end int64, ids []int64, data []byte, handle timerhandle) *timer.TimeEvent {
	f := func() {
		handle(ids, data)
	}
	DEBUG("TimerClosure")
	return timer.NewTimer(end, f)
}

//--------------------------------------------------------- 分钟定时处理函数
func TIMER_minHandle(ids []int64, data []byte) {
	now := time.Now().Unix()

	// 处理玩家超时
	list := gsdb.GetAll()
	for _, sess := range list {
		if sess.UpdateTime+timeout < now {
			writer := writer.Get(sess.Idx)
			DEBUGF("user:%v timeout", sess.Id)
			if writer != nil {
				select {
				case writer <- &pb.MSG{
					Uid: sess.Id,
					Fd:  sess.Fd,
					Err: pb.ERR_USER_LOGOUT,
				}:
				default:
				}
			}
		}
	}

	TimerClosure(now+60, nil, nil, TIMER_minHandle)
}

//--------------------------------------------------------- 道具过期处理函数 test
func TIMER_item(ids []int64, data []byte) {
	sess := gsdb.Get(ids[0])
	if sess != nil {
		now := time.Now().Unix()
		for _, item := range sess.Items.Items {
			if item.End <= now {
				DEBUG(sess.User.Id, item.Id, "过期 数量:", item.Num)
			}
		}
	}
}
