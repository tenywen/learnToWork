package build

import (
	"timer"
)

type Build struct {
	Id          int16
	Type        int16
	Pos         int16
	Level       int16
	Status      int16
	End         int64
	CollectTime int64            `bson:",omitempty"`
	Res         int64            `bson:",omitempty"`
	TimeEvent   *timer.TimeEvent `bson:"-"`
}

type Manager struct {
	UserId int64
	Builds []*Build
	NextId int16
}

func (m *Manager) Get(id int16) *Build {
	for _, b := range m.Builds {
		if b.Id == id {
			return b
		}
	}

	return nil
}

func (m *Manager) IsExist(typ, level int16) bool {
	for _, b := range m.Builds {
		if b.Type == typ && b.Level == level {
			return true
		}
	}
	return false
}

func (m *Manager) BuildingNum() (cnt int32) {
	for _, b := range m.Builds {
		if b.End > 0 {
			cnt++
		}
	}
	return cnt
}
