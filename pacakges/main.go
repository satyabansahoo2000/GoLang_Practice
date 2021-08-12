package main

import (
	"fmt"
)

func main() {
	prog1 := Program{Name: "GoLang", Year: "2009"}
	prog2 := Program{Name: "Python", Year: "1991"}
	
	fmt.Println(prog1.details())
	fmt.Println(prog2.details())
}