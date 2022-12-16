package tester

import (
	"net"
)

func GetFreePort() uint32 {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	return uint32(l.Addr().(*net.TCPAddr).Port)
}
