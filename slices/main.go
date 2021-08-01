package main

import (
	"fmt"
	"errors"
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

	// looping through the elements
	for index, value := range slice2 {
		fmt.Println("Index :", index, "Value: ", value)
	}

	// Copying a slice using copy()
	slice3 := make([]int, len(slice2))
	copy(slice3, slice2)
	fmt.Println(slice3)
	
	// Insert an element in the slice
	slice3, _ = insert(slice3, 4, 9)
	fmt.Println(slice3)

	// Remove an elements in the slice
	slice3, _ = remove(slice3, 3)
	fmt.Println(slice3)
}

// Function to insert an element in the slice
func insert(slice []int, index, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("index cannot be less than 0")
	}
	if index >= len(slice) {
		return append(slice, value), nil
	}

	slice = append(slice[:index+1], slice[index:]...)			// variadic function
	slice[index] = value
	return slice, nil
}

// Function to delete an element from the slice
func remove(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, errors.New("index out of the range")
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice, nil
}