package closure

import (
	"time"
)

import (
	"gs/gsdb"
	"gs/writer"
	. "logger"
	pb "protocol"
)

// 内部异步消息
type ipchandle func([]int64, []byte)

func IPCClosure(ids []int64, data []byte, handle ipchandle) {
	f := func() {
		handle(ids, data)
	}

	select {
	case Ch <- f:
	default:
	}
}

func IPC_BroadCast(ids []int64, data []byte) {
	for _, id := range ids {
		sess := gsdb.Get(id)
		if sess != nil {
			w := writer.Get(sess.Idx)
			select {
			case w <- &pb.MSG{
				Fd:   sess.Fd,
				Idx:  sess.Idx,
				Data: data,
			}:
			default:
			}
			DEBUG(sess.User.Id, "broadcast!")
		}
	}
}

//--------------------------------------------------------- 添加道具 test
func IPC_add_item(ids []int64, data []byte) {
	for _, id := range ids {
		sess := gsdb.Get(id)
		if sess != nil {
			duration := int64(10)
			item := sess.Items.Add(1, 1, time.Now().Unix()+duration)
			if duration != 0 {
				// TODO   道具过期处理函数
				item.TimeEvent = TimerClosure(item.End, []int64{sess.User.Id}, nil, TIMER_item)
			}
		}
	}
}
