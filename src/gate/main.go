package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

import (
	"cfg"
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

	StartLogger(gateConfig.LogName)
	server("tcp4", gateConfig.Port)
}
