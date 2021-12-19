package main

import "fmt"

type Inner func(string) string

func Decorate(f Inner) Inner {
	a := " and a"
	return func(base string) string {
		ret := f(base)
		ret = ret + a + " and Tshirt"
		return ret
	}
}

func Dressing(cloth string) string {
	return "dressing " + cloth
}

func main() {
	f := Decorate(Dressing)
	fmt.Println(f("shoes"))
}
