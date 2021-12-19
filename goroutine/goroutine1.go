package main

import (
	"fmt"
	"time"
)

// goroutine
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main()  {
	go say("hello")	// 开一个新的Goroutines执行
	say("world")	// 当前Goroutines执行
}