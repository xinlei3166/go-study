package main

import "fmt"

// fib函数会返回一个返回int的函数
func fib() func() int {
	a := 0
	b := 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

