package main

import (
	"context"
	. "httpext"
	. "logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"utils/com"
)

func TestDoHttpHandler(form *HttpFormAsync) interface{} {
	// 读http数据
	// 调用处理逻辑
	// 回写http数据
	serverid := form.GetString("serverId")
	LogInfo("test server1, serverid=%s", serverid)
	form.Resp <- true
	return nil
}

func TestDoHttpHandler2(form *HttpFormAsync) interface{} {
	// 读http数据
	// 调用处理逻辑
	// 回写http数据
	serverid := form.GetString("serverId")
	LogInfo("test server2, serverid=%s", serverid)
	form.Resp <- true
	return nil
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	com.GoRunArgs(DefaultLogger.LogStart, ctx, &wg)

	//httpServer := NewHttpServer()
	//httpServer.RegisterRouter("/test", TestDoHttpHandler)
	//httpServer.StartServer(":8080")

	httpServer := NewAsyncHttpServer()
	httpServer.RegisterRouter("/test1", TestDoHttpHandler)
	httpServer.RegisterRouter("/test2", TestDoHttpHandler2)
	httpServer.StartServer(":8080")

	<-sigs
	cancel()
	wg.Wait()
}
