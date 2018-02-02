package cfg

import (
	"fmt"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	readConfig(gopath + "/config.yaml")
	fmt.Println(Config.GS)
	fmt.Println(Config.GATE)
	fmt.Println(Config.DB)
	fmt.Println(Config.LOG.Level)
	fmt.Println(Config.LOG.LocalTime)
}
