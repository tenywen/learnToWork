package item

import (
	. "logger"
	"timer"
)

type Item struct {
	Id        int16
	Num       int32
	End       int64
	TimeEvent *timer.TimeEvent `bson:"-"`
}

type Manager struct {
	UserId int64
	Items  []*Item
}

func (m *Manager) Check(id int16, num int32, now int64) bool {
	for _, item := range m.Items {
		if item.Id == id && item.Num >= num {
			if item.End < now {
				return true
			}
		}
	}
	return false
}

func (m *Manager) Add(id int16, num int32, end int64) *Item {
	DEBUG(m.UserId, "add item ", id, num)
	for _, item := range m.Items {
		if item.Id == id && item.End == end {
			item.Num += num
			return item
		}
	}

	item := &Item{id, num, end, nil}
	m.Items = append(m.Items, item)
	return item
}
