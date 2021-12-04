package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() 	float64
}

type circle struct {
	radius 	float64
}

type rect struct {
	width	float64
	height	float64
}

func getArea(s shape) float64 {
	return s.area()
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5,7}
	shapes := []shape{&c1,&r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
	fmt.Printf("The input to AREA FUNCTION were circle '%v' and rectangle '%v'", shapes[0], shapes[1])
}