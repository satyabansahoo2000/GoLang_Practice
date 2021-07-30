package main

import (
	"fmt"
)

func main() {
	// declaration of arrays
	var arr1 [5]int			// array of 5 elements with int type
	fmt.Println(arr1)

	// initializing an array with fixed size
	arr2 := [5]int{1,2,3,4,5}
	fmt.Println(arr2)

	// array with no fixed size
	arr3 := [...]int{1,2,3,4,5,6,7,8}
	fmt.Println(arr3)

	// MULTI-DIMENSIONAL ARRAYS
	// 2-D Array
	var arr4 [3][3]int
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			arr4[row][col] = row + col
		}
	}
	fmt.Println(arr4)

	// 3-D Array
	var arr5 [2][2][2]int
	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			for depth := 0; depth < 2; depth ++ {
				arr5[row][col][depth] = row + col + depth
			}
		}
	}
	fmt.Println(arr5)
}