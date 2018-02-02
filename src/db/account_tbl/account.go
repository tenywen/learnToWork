package account_tbl

import (
	"gopkg.in/mgo.v2/bson"
)

import (
	. "account"
	"db"
	. "logger"
)

const (
	collection = "account"
	dbName     = "account"
)

//---------------------------------------------------------- get
func GetByUUID(uuid string) *Account {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	var account Account
	err := mgoSession.DB(dbName).C(collection).Find(bson.M{"uuid": uuid}).One(&account)
	if err != nil {
		ERROR("get acount by uuid ", err)
		return nil
	}
	return &account
}

//--------------------------------------------------------- save
func Save(account *Account) {
	mgoSession := db.Session.Copy()
	defer mgoSession.Close()

	_, err := mgoSession.DB(dbName).C(collection).Upsert(bson.M{"uuid": account.UUID}, account)
	if err != nil {
		ERROR("save account ", err, account)
	}
}
