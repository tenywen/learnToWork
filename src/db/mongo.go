package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

import (
	. "logger"
)

const (
	UID      = "uid"
	NEXT_VAL = "nextval"
)

var Session *mgo.Session

func StartDB(url string) {
	sess, err := mgo.Dial(url)
	if err != nil {
		ERROR(err)
		os.Exit(-1)
	}

	// TODO set mgo mode etc...
	Session = sess
}

func NextVal(dbname, collection, name string) int64 {
	mgoSession := Session.Copy()
	defer mgoSession.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"value": 1}},
		ReturnNew: true,
	}

	type Val struct {
		Name  string
		Value int64
	}

	val := &Val{}
	_, err := mgoSession.DB(dbname).C(collection).Find(bson.M{"name": name}).Apply(change, val)
	if err != nil {
		ERRORF("NextVal %v %v %v %v", dbname, collection, name, err)
		return -1
	}

	return val.Value
}
