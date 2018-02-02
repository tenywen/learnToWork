package closure

import (
	//	"encoding/json"
	"fmt"
	//"github.com/vmihailenco/msgpack"
	"testing"
	"time"
)

import (
	"timer"
)

func TIMER_test(ids []int64, data []byte) {
	fmt.Println("TIMER_test")
}

func TestClosureTimer(t *testing.T) {
	f := func() {
		TIMER_test(nil, nil)
	}

	timer.NewTimer(time.Now().Unix()+10, f)

	f = <-timer.Ch
	f()
}
