package main

import (
	. "global"
	. "logger"
	. "gfservice"
	"time"
)

func main() {
	service := ServiceFactory.NewNodeService(ST_SERVER, SERVER_TYPE_GAME)
	LogInfo("server start...")
	service.Start()
	for {
		time.Sleep(time.Second * 1)
	}
}
