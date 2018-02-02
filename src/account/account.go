package account

type Account struct {
	ServerId int16  `bson:",omitempty"`
	UserName string `bson:",omitempty"`
	UUID     string `bson:",omitempty"`
	PWD      string `bson:",omitempty"`
	UID      int64  `bson:",omitempty"`
}
