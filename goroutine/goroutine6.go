package main

import (
	"fmt"
	"runtime"
	"time"
)

// 超时
func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(2 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<- o
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
