package main

import (
	"fmt"
)

func main() {
	// creating an empty slice
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1)
	fmt.Println(len(slice1))		// length of the slice
	fmt.Println(cap(slice1))		// capacity of the slice
	
	// appending elements to the slice
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(slice1)
	fmt.Println(len(slice1))	// length of the slice increased 
	fmt.Println(cap(slice1))	// capacity of the slice remains the same

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(slice1)
	fmt.Println(len(slice1))	// length of the slice increased
	fmt.Println(cap(slice1))	// capacity of the slice increased

	// pointing the slice to another variable
	slice2 := slice1
	fmt.Println(slice2)
	fmt.Println(slice1)

	// changes apply to the both the slices
	// pointing to the same array
	slice2[5] = 10
	fmt.Println(slice2)
	fmt.Println(slice1)

	// appending an new element
	// doesn't affect the slice1 elements
	slice2 = append(slice2, 15, 20)
	fmt.Println(slice2)
	fmt.Println(slice1)
}