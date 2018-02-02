package user

import (
	"timer"
)

type Buffer struct {
	Src       int64
	Type      int16
	Effect    float64
	End       int64
	TimeEvent *timer.TimeEvent `bson:"-"`
}

type Hero struct {
	Id int16
}

type Res struct {
	Type int8
	Cnt  int64
}

type User struct {
	Id           int64
	Name         string
	UUID         string
	Buff         []Buffer
	Level        int32
	Exp          int32
	RegisterTime int64
	LoginTime    int64

	//1 	Food
	//2  	Energy
	//3 	Oil
	//4 	Steel
	//5 	Cash
	//Res map[int8]float64
}

/*
func (u *User) CheckRes(list []Res) bool {
	for _, res := range list {
		if cnt, ok := u.Res[res.Type]; ok && cnt >= float64(res.Cnt) {
			continue
		}
		return false
	}
	return true
}

func (u *User) AddRes(list []Res) {
	for _, res := range list {
		u.Res[res.Type] += float64(res.Cnt)
	}
}

func (u *User) DelRes(list []Res) {
	for _, res := range list {
		u.Res[res.Type] -= float64(res.Cnt)
	}
}
*/
