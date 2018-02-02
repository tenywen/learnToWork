package logger

import (
	"fmt"
	"testing"
)

type Test struct {
	A int
	B int
}

func TestPanic(t *testing.T) {
	StartLogger("test.log")
	defer PANIC("test")

	var x *Test

	fmt.Println(x.A)
}
