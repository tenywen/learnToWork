package timer

import (
	"fmt"
	"testing"
	"time"
)

func f(t ...interface{}) func() {
	return func() {
		fmt.Println("now =", time.Now().Unix(), "parma:", t)
	}
}

func TestAdd(t *testing.T) {
	now := time.Now().Unix()
	NewTimer(now+10, f("TestAdd", now+10))
	ff := <-Ch
	ff()
}

func TestUpdate(t *testing.T) {
	now := time.Now().Unix()
	e := NewTimer(now+100, f(now+100))
	NewTimer(now+20, f(now+20))
	NewTimer(now+30, f(now+30))
	e.Update(now + 50)

	tick := time.NewTicker(250 * time.Second)
	for {
		select {
		case ff := <-Ch:
			ff()
		case <-tick.C:
		}
	}
}
