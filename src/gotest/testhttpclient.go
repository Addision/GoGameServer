package main

import (
	"fmt"
	. "httpext"
	"time"
)

func main() {
	StartHttpPostAsync()
	HttpPost.Request("test1", &JsonMap{
		"serverId": "1212",
	})

	HttpPostAsync.Request("test2", &JsonMap{
		"serverId": "3333"}, func(reqResult *RequestAsyncRet) {
		fmt.Println("HttpPostAsync.Request ok handler")
	}, func(err error) {
		fmt.Println("HttpPostAsync.Request error handler")
	})

	HttpPostAsync.HandleHttpResponse()
	time.Sleep(time.Second * 3)
	HttpPostAsync.Release()
}
