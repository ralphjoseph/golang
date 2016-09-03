// structs project main.go
package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height int
}

func (r rectangle) area() int {
	return r.width * r.height
}
func (r *rectangle) perimeter() int {
	return 2 * (r.width + r.height)
}

//Interfaces

type shape interface {
	area() float64
	//perimeter() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s *square) area() float64 {
	return s.length * s.length
}

func area2(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}
func main() {

	box := rectangle{height: 9, width: 10}

	fmt.Println(box.area(), box.perimeter())

	boxPtr := &box
	fmt.Println(boxPtr.area(), boxPtr.perimeter())

	c := circle{radius: 9.8}
	fmt.Println(c.area())

	s := &square{length: 5.5}
	fmt.Println(s.area())

	area2(c)
	area2(s)

}
