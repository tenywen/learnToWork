#!/bin/sh

$GOPATH/bin/gs -i 1& 
sleep 2
$GOPATH/bin/gate -i 1&
