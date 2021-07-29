package main

import (
	"fmt"
	f "Functions/filter"
	vf "Functions/variadicFunctions"
	af "Functions/anonymousFunctions"
)

func main() {
	fmt.Println("Variadic Functions")
	fmt.Println(vf.AddNums(1,2,3,4,5))
	fmt.Println(vf.AddNums(5,4,5))

	fmt.Println("Anonymous Functions")
	function := af.Anonymous()
	fmt.Println(function())

	// Printing the even numbers using Filter() Function
	fmt.Println("Filter Function")
	arr := []int{0,1,2,3,4,5,6,7,8,9}
	evens := f.Filter(arr,
		func(val int) bool {
			return val % 2 == 0
		})
	fmt.Println(evens)
}