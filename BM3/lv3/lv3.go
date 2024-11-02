package main

import (
	"fmt"
	"math"
)

func main() {
	Circle1 := Circle{
		radius: 10,
	}
	Rectangle1 := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Println(Circle1.Area(), Rectangle1.Area())
}

type shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}
type Rectangle struct {
	length, width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
