package account_tbl

import (
	"fmt"
	"testing"
)

func TestGetByUUID(t *testing.T) {
	fmt.Println(GetByUUID("test1"))
}
