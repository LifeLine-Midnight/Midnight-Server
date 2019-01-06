package main

import (
	"midgo/httpsvr"
)

const (
	IPADDR = "127.0.0.1"
	PORT   = 8848
)

func main() {
	svr := httpsvr.GetMidgoSvr()
	svr.Run(IPADDR, PORT)
}
