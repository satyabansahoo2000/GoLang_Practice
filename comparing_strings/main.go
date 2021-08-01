package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type Point struct {
	X float32
	Y float32
	Z float32
	Name []string
}

func main () {
	pt1 := Point{
		X: 5.6, 
		Y: 3.8, 
		Z: 6.9, 
		Name: []string{"pt1"},
	}
	
	pt2 := Point{
		X: 5.6, 
		Y: 3.8, 
		Z: 6.9, 
		Name: []string{"pt"},
	}
	
	pt3 := Point{
		X: 5.6, 
		Y: 3.8, 
		Z: 6.9,
		Name: []string{"pt"},
	}

	fmt.Println(cmp.Equal(pt1, pt2)) 
	fmt.Println(cmp.Equal(pt2, pt3))
}

// Custom-Error Function
func (p1 Point) Equal(p2 Point) bool {
	if p1.X == p1.X &&
	p2.Y == p1.Y &&
	p2.Z == p1.Z {
		return true
	}
	return false
}