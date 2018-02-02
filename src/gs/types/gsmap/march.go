package gsmap

import (
	"timer"
)

type March struct {
	Id        int32
	TimeEvent *timer.TimeEvent `bson:"-"`
	Src       MapPos
	Dst       MapPos
	StartTime int64
	EndTime   int64
	UserId    int64
	Items     []GroupItem     `bson:",omitempty"`
	Troops    []Troop         `bson:",omitempty"`
	Captive   map[int64]int16 `bson:",omitempty"` // 抓捕的英雄
}

type MarchHero struct {
	Owner int64
	Id    int16
}

type Troop struct {
	UserId  int64
	Heroes  []int16
	Armyies []*Army
}

type GroupItem struct {
	Id  int16
	Cnt int32
}

type Army struct {
	Type int16
	Cnt  int64
}

func NewMarch() *March {
	return nil
}
