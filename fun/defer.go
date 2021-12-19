package main

import "fmt"

// 在 defer 函数中参数会提前求值
func main() {
	var i = 1
	defer fmt.Println("result: ", func() int { return i * 2 }())
	i++
}
