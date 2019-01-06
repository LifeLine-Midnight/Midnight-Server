package main

import (
	"midgo/httpsvr"

	"midnightapisvr/controller"
)

const (
	IPADDR = "127.0.0.1"
	PORT   = 8848
)

func main() {
	svr := httpsvr.GetMidgoSvr()
	svr.AddController("/midnightapisvr/api", new(controller.UserController))
	svr.AddController("/midnightapisvr/api", new(controller.SessionController))
	svr.AddController("/midnightapisvr/api", new(controller.ActionController))
	svr.Run(IPADDR, PORT)
}
