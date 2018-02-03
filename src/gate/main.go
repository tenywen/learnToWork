package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

import (
	"cfg"
	"db"
	. "logger"
)

var id int16

func init() {
	id = int16(*(flag.Int("i", 1, "gate_id")))
}

func main() {
	flag.Parse()
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	gateConfig := cfg.GetGateConfig(id)
	if gateConfig != nil {
		StartLogger(gateConfig.LogName)
		db.StartDB(gateConfig.DB)
		gsDial()
		server("tcp4", gateConfig.Port, gateConfig.WorkNum, gateConfig.BufferMax)
		return
	}
	fmt.Println("GATE NOT START!")
}
