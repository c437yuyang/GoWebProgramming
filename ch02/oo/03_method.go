package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rect{2, 3}
	c1 := Circle{2}

	fmt.Println(r1.area(), c1.area())
}
