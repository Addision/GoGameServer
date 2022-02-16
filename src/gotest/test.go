package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// utils.UpdateSystemDate("2020-03-20 15:02:41")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("quit...")
				return
			default:
				time.Sleep(time.Second * 1)
				fmt.Println("time now:", time.Now().Second())
			}
			fmt.Println("test context...")
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()

	time.Sleep(3 * time.Second)
	fmt.Println("exit....")
}
