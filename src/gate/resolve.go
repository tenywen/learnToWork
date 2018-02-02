package main

import (
	"net"
	"os"
	"syscall"
)

import (
	. "logger"
)

func resolveTCPAddr(network string, addr string) (syscall.Sockaddr, int) {
	tcpAddr, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		ERROR(err)
		os.Exit(-1)
	}
	var domain int
	var sockAddr syscall.Sockaddr
	switch network {
	case "tcp4":
		sockAddr = &syscall.SockaddrInet4{Port: tcpAddr.Port}
		domain = syscall.AF_INET
		copy(sockAddr.(*syscall.SockaddrInet4).Addr[:], tcpAddr.IP.To4())
	case "tcp", "tcp6":
		sockAddr = &syscall.SockaddrInet6{Port: tcpAddr.Port}
		domain = syscall.AF_INET6
		copy(sockAddr.(*syscall.SockaddrInet6).Addr[:], tcpAddr.IP.To16())
	default:
		ERRORF("network error:%v", network)
		os.Exit(-1)
	}

	return sockAddr, domain
}
