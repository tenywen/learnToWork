package misc

import (
	"fmt"
	"testing"
	"time"
)

func TestPM12(t *testing.T) {
	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(PM12(now))
}
