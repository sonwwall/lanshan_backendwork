package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r Rectangle) area() float64 {
	return r.length * r.width
}
func CalculateArea(shape Shape) float64 {
	return shape.area()

}

func main() {
	circle := Circle{10}
	rectangle := Rectangle{
		10,
		20,
	}
	fmt.Println(CalculateArea(circle))
	fmt.Println(CalculateArea(rectangle))
}
