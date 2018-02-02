#!/usr/bin/env bash
python parse.py ../protocol/
go fmt api.go
go fmt proto.go
go fmt gate_handler.go
go fmt gs_handler.go
go fmt errcode.go
go fmt const.go
go fmt decode.go
mv api.go ../src/protocol/
mv proto.go ../src/protocol/
mv errcode.go ../src/protocol/
mv decode.go ../src/protocol/
mv const.go ../src/protocol/
mv gate_handler.go ../src/gate/
mv gs_handler.go ../src/gs/
