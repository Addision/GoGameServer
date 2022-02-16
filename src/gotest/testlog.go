package main

import (
	. "logger"
	"time"
)

func main() {
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)
	// var wg sync.WaitGroup
	//
	// ctx, cancel := context.WithCancel(context.Background())
	//
	// com.GoRunArgs(DefaultLogger.LogStart, ctx, &wg)
	LogError("some wrong with this program")
	LogInfo("some wrong with this program, %s", "xxxxxxxxxxxxxxxxx")
	time.Sleep(time.Second * 10)
	//
	// <-sigs
	// cancel()
	// wg.Wait()

}
