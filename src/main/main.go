package main

import (
	"fmt"
	"go-worker/src/mypool"
	"time"
)

func main() {
	pool := mypool.NewPool(3)

	task := mypool.NewTask(func() {
		fmt.Println(time.Now())
	})

	go func() {
		for {
			pool.OutChan <- task
		}
	}()
	time.Sleep(1 * time.Second)
	pool.Run()
}
