package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{12}
	c2 := Circle{9}
	fmt.Println("Area of r1: ", r1.area())
	fmt.Println("Area of r2: ", r2.area())
	fmt.Println("Area of c1: ", c1.area())
	fmt.Println("Area of c2: ", c2.area())
}
