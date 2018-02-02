package research

import (
	"timer"
)

type Research struct {
	Id        int16
	Type      int16
	Level     int16
	Status    int16
	Start     int64
	End       int64
	TimeEvent *timer.TimeEvent `bson:"-"`
}

type Manager struct {
	UserId     int64
	Researches []*Research
}
