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
	MARCH_COLLECTION = "march"
)

//---------------------------------------------------------- 获取所有行军
func GetAllMarches(dbname string) []*March {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	var marches []*March
	err := mgoSession.DB(dbname).C(MARCH_COLLECTION).Find(nil).All(&marches)
	if err != nil {
		ERROR("get all marches", err)
	}
	return marches
}

//---------------------------------------------------------- 保存行军
func SaveMarch(dbname string, march *March) {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	_, err := mgoSession.DB(dbname).C(MARCH_COLLECTION).Upsert(bson.M{"id": march.Id}, march)
	if err != nil {
		ERROR("save march", err, march)
	}
}

//---------------------------------------------------------- 删除行军
func RemoveMarch(dbname string, id int32) {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	err := mgoSession.DB(dbname).C(MARCH_COLLECTION).Remove(bson.M{"id": id})
	if err != nil {
		ERROR("remove march", err, id)
	}
}

//---------------------------------------------------------- 获得行军
func GetMarch(dbname string, id int32) *March {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	var march *March
	err := mgoSession.DB(dbname).C(MARCH_COLLECTION).Find(bson.M{"id": id}).One(march)
	if err != nil {
		ERROR("get march", err, id)
	}
	return march
}
