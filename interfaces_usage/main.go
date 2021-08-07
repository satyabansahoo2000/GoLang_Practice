package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
	name   string
}

type Square struct {
	length float64
	name   string
}

type Triangle struct {
	base   float64
	height float64
	name   string
}

type Shape interface {
	Area() float64
}

// Circle Implements Shape
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Square Implements Shape
func (s Square) Area() float64 {
	return math.Pow(s.length, 2)
}

// Triangle Implements Shape
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// Circle Stringer Function
func (c Circle) String() string {
	return fmt.Sprintf("Area of Circle : %v", c.Area())
}

// Square Stringer Function
func (s Square) String() string {
	return fmt.Sprintf("Area of Square : %v", s.Area())
}

// Triangle Stringer Function
func (t Triangle) String() string {
	return fmt.Sprintf("Area of Triangle : %v", t.Area())
}

// To calculate the Area and print result without using different lines
// func calculateArea(listOfShapes []Shape) {
// 	for _, shape := range listOfShapes {
// 		fmt.Println("Area of shape is : ", shape.Area())
// 	}
// }

func main() {
	c1 := Circle{radius: 5, name: "c1"}
	s1 := Square{length: 5, name: "s1"}
	t1 := Triangle{base: 5, height: 6, name: "t1"}

	// shapes := []Shape{c1, s1, t1}
	// calculateArea(shapes)

	fmt.Println(c1)
	fmt.Println(s1)
	fmt.Println(t1)
}
