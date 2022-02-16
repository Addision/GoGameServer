package main

import (
	"fmt"
	. "utils"
	// pool "pool/threadpool"
	// pool "pool/dispatcher"
	"time"
)

func tFunc(args ...interface{}) error {
	fmt.Println(time.Now())
	for _, arg := range args {
		fmt.Println(arg)
	}
	return nil
}

func main() {
	t := NewTask(tFunc, 1, "222")

	p := NewTaskPool(5, 1000)
	go p.Start()
	go func() {
		for i := 0; i < 100; i++ {
			p.Submit(t)
		}
		p.Stop()
	}()

	// p := pool.NewThreadPool(5, 5)
	// go p.StartPool()
	// time.Sleep(time.Second*1)

	// go func() {
	// 	for {
	// 		p.Submit(0, t)
	// 		p.Submit(1, t)
	// 		p.Submit(2, t)
	// 		p.Submit(3, t)
	// 		p.Submit(4, t)

	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// d := pool.NewDispatcher(5, 1000)
	// go d.Start()
	// time.Sleep(time.Second * 1)
	// go func() {

	// 	for i := 0; i < 5; i++ {
	// 		d.Submit(0, t)
	// 		d.Submit(1, t)
	// 		d.Submit(2, t)
	// 		d.Submit(3, t)
	// 		d.Submit(4, t)

	// 		time.Sleep(1 * time.Second)
	// 	}
	// 	d.Stop()
	// }()

}
