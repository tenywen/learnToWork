package user_tbl

import (
	"gopkg.in/mgo.v2/bson"
)

import (
	"db"
	. "gs/types/user"
	. "logger"
)

const (
	collection = "users"
)

func GetAllUsers() {

}

func Save(dbname string, user *User) {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()
	_, err := mgoSession.DB(dbname).C(collection).Upsert(bson.M{"id": user.Id}, user)
	if err != nil {
		ERROR("save user ", err)
	}
}

func Get(dbname string, uid int64) *User {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()
	var user User
	err := mgoSession.DB(dbname).C(collection).Find(bson.M{"id": uid}).One(&user)
	if err != nil {
		ERROR("get user ", err)
		return nil
	}
	return &user
}
