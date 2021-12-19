package main

import (
	"fmt"
	"os"
)

func simpleFunc() {
	if x := 1; x > 1 {
		fmt.Printf("x is gt 1")
	} else if x == 1 {
		fmt.Printf("x is equal 1")
	} else {
		fmt.Printf("x is lt 1")
	}
}

func myFunc() {
	i := 0
here:
	fmt.Println(i)
	i++
	if i <= 100 {
		goto here
	}
}

func sumAdd() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
}

func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func breakContinue() {
	for i := 10; i > 0; i-- {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

func keyValue() {
	fruits := map[string]int{
		"apple":      5,
		"watermelon": 2,
	}
	for k, v := range fruits {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
}

func switchCase() {
	i := 1
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2:
		fmt.Println("i is equal to 2")
	default:
		fmt.Println("i i is an integer")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func args(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func pointer(a *int) int {
	/*	x := 1
		fmt.Println(x)
		x1 := pointer(&x)
		fmt.Println(x1)
		fmt.Println(x)*/
	*a += 1
	return *a
}

func ReadWrite() bool {
	// defer 采用后进先出模式
	file, err := os.Open("file.txt")
	defer file.Close()
	if err != nil {
		return false
	}
	return true
}

type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	/*	slice := []int {1, 2, 3, 4, 5, 7}
		fmt.Println("slice = ", slice)
		odd := filter(slice, isOdd) // 函数当做值来传递了
		fmt.Println("Odd elements of slice are: ", odd)
		even := filter(slice, isEven) // 函数当做值来传递了
		fmt.Println("Even elements of slice are: ", even)*/
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

var user = os.Getenv("USER1")

func init1() {
	if user == "" {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			println("模拟器recover")
			b = true
		}
	}()
	f() // 执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}

func main() {
	throwsPanic(init1)
}
