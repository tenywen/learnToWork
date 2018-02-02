package map_tbl

import (
	"fmt"
	"testing"
	. "types/gsmap"
)

func TestSave(t *testing.T) {
	tile := &Tile{
		Pos:   MapPos{1, 1},
		Level: 1,
	}

	Save("test1", tile)
}

func TestGetAll(t *testing.T) {
	tiles := GetAll("test1")
	for k := range tiles {
		fmt.Println(tiles[k])
	}
}

func TestGet(t *testing.T) {
	fmt.Println(Get("test1", MapPos{1, 1}))
}

func TestRemove(t *testing.T) {
	Remove("test1", MapPos{1, 1})
}
