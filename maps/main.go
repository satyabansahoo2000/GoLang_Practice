package main

import (
	"fmt"
)

func main() {
	age1 := make(map[string]int)
	age1["Satyaban"] = 22
	age1["Suchita"] = 25
	age1["Sabyasachi"] = 22

	fmt.Println(age1)			// map[Sabyasachi:22 Satyaban:22 Suchita:25]

	// initializing a map with a map literal
	age2 := map[string]int{
		"Satyaban": 22,
		"Suchita": 25,
		"Sabyasachi": 22,
	}
	fmt.Println(age2)			// map[Sabyasachi:22 Satyaban:22 Suchita:25]

	// Checking for a key existance
	fmt.Println(age2["Hemraj"])	// 0

	if value, ok := age2["Hemraj"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key is not present")		// Key is not present
	}								

	// deleting a key
	if _, ok := age2["Suchita"]; ok {
		delete(age2, "Suchita")
	} else {
		fmt.Println("Key is not present")
	}
	fmt.Println(age2)				// map[Sabyasachi:22 Satyaban:22]
}