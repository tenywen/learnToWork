package closure

var Ch chan func()

func init() {
	Ch = make(chan func(), 100000)
}
