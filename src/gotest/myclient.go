package main

import (
	. "global"
	. "logger"
	. "gfservice"
	. "network"
	"time"
)

func onStart(args ...interface{}) {
	LogInfo("onStart function...")
}

func main() {
	service := ServiceFactory.NewNodeService(ST_CLIENT, SERVER_TYPE_GAME)
	nodeClient := service.(*NodeClient)
	nodeClient.AddConnServer()
	nodeClient.SetOnStart(onStart)
	service.Start()
	for {
		time.Sleep(time.Second * 1)
	}
	service.Stop()
}
