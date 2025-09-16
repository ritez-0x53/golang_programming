package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	width, height float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func calculate(s Shape) {
	fmt.Println("area is", s.Area())
}

func main() {

	rec1 := Rectangle{width: 6, height: 12}
	circ := Circle{radius: 7.6}
	fmt.Println("Area of rectangle is", rec1.Area())
	fmt.Println("Area of circle is", circ.Area())

	calculate(rec1)
	calculate(circ)

}
