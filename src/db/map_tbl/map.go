package map_tbl

import (
	"gopkg.in/mgo.v2/bson"
)

import (
	"db"
	. "logger"
	. "types/gsmap"
)

const (
	mapCollection = "map"
)

//---------------------------------------------------------- 获得所有tiles
func GetAllTiles(dbname string) []*Tile {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	var tiles []*Tile
	err := mgoSession.DB(dbname).C(mapCollection).Find(nil).All(&tiles)
	if err != nil {
		ERROR("get all tiles", err)
	}
	return tiles
}

//---------------------------------------------------------- 保存tile
func SaveTile(dbname string, t *Tile) {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	_, err := mgoSession.DB(dbname).C(mapCollection).Upsert(bson.M{"pos": t.Pos}, t)
	if err != nil {
		ERROR("save tile", err, t)
	}
}

//---------------------------------------------------------- 删除tile
func RemoveTile(dbname string, pos MapPos) {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	err := mgoSession.DB(dbname).C(mapCollection).Remove(bson.M{"pos": pos})
	if err != nil {
		ERROR("remove tile", err, pos)
	}
}

//---------------------------------------------------------- 获得tile
func GetTile(dbname string, pos MapPos) *Tile {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	var t Tile
	err := mgoSession.DB(dbname).C(mapCollection).Find(bson.M{"pos": pos}).One(&t)
	if err != nil {
		ERROR("get tile ", err, pos)
		return nil
	}
	return &t
}
