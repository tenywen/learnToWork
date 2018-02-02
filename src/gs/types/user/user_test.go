package user

import (
	"fmt"
	"testing"
)

func TestCheckRes(t *testing.T) {
	u := User{
		Res: make(map[int8]float64),
	}

	resList := []Res{Res{3, 10.0}}

	fmt.Println(u.CheckRes(resList))

	resList = append(resList, Res{2, 10.0})
	u.Res[2] = 11.0
	u.Res[3] = 1.0
	fmt.Println(u.CheckRes(resList))

	u.Res[3] = 12.0
	fmt.Println(u.CheckRes(resList))
}
