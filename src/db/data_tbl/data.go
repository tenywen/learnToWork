package data_tbl

import (
	"gopkg.in/mgo.v2/bson"
)

import (
	"db"
	. "logger"
)

const (
	ResearchCollection = "research"
	BuildCollection    = "build"
	ItemCollection     = "item"
)

func Get(dbName, collection string, uid int64, obj interface{}) bool {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	err := mgoSession.DB(dbName).C(collection).Find(bson.M{"userid": uid}).One(obj)
	if err != nil {
		ERROR("get", err, uid, dbName, collection)
		return false
	}
	return true
}

func Set(dbName, collection string, uid int64, obj interface{}) {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	_, err := mgoSession.DB(dbName).C(collection).Upsert(bson.M{"userid": uid}, obj)
	if err != nil {
		ERROR("set", err, uid, dbName, collection)
	}
}
