package main

import (
	"fmt"
	vf "Functions/variadicFunctions"
	af "Functions/anonymousFunctions"
)

func main() {
	fmt.Println(vf.AddNums(1,2,3,4,5))
	fmt.Println(vf.AddNums(5,4,5))

	function := af.Anonymous()
	fmt.Println(function())
}