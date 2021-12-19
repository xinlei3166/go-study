package main

func fib(n int) {
	a := 0
	b := 1
	for n > 0 {
		c := a
		a, b = b, a + b
		println(c)
		n--
	}
}

func main() {
	fib(40)
}
