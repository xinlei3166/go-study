package main

import "fmt"

const (
	a = iota // 默认为0
	_
	b
	c
)

const d = iota // 遇到const重置为0

var s string = "Hello GO"
var arr [3]int

func main() {
	fmt.Printf("a: %d, b: %d, c: %d, d: %d\n", a, b, c, d)
	fmt.Printf("s: %s\n", s)
	num := 20000
	fmt.Printf("num: %d\n", num)
	arr[0] = 42 // 数组下标是从0开始的
	arr[1] = 13
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{4, 5, 6}
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("arr first: %d\n", arr[0])        // 获取数据，返回42
	fmt.Printf("arr end: %d\n", arr[len(arr)-1]) // 返回未赋值的最后一个元素，默认返回0
	fmt.Printf("arr2 first: %d\n", arr2[0])      // 返回未赋值的最后一个元素，默认返回0
	fmt.Printf("arr3 first: %d\n", arr3[0])      // 返回未赋值的最后一个元素，默认返回0
	fmt.Print(doubleArray, "\n")
	fmt.Print(easyArray, "\n")
	slice := []string{"a", "b", "c", "d"}
	slice1 := slice[:]
	fmt.Print(slice, "\n")
	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", slice1)
	fruits := map[string]string{"name": "junxi", "age": "18"}
	fruit1 := fruits
	fmt.Print(fruits, "\n")
	fmt.Printf("%p\n", fruits)
	fmt.Printf("%p\n", fruit1)
}
