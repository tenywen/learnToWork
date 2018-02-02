package map_tbl

import (
	"fmt"
	"testing"
)

import (
	. "types/gsmap"
)

func TestSavemarch(t *testing.T) {
	march := March{
		Id:  1,
		Src: MapPos{1, 2},
		Dst: MapPos{215, 215},
	}
	SaveMarch("test1", &march)
}

func TestGetAllMarches(t *testing.T) {
	fmt.Println(GetAllMarches("test1"))
}

func TestGetMarch(t *testing.T) {
	fmt.Println(GetMarch("test1", 1))
}

func TestRemoveMarch(t *testing.T) {
	RemoveMarch("test", 1)
}
