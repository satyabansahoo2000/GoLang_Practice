// main package
package main

// Importing all the necessary packages
import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	// Without using external package
	fmt.Println("Hello World!!!")

	// using external package - rsc.io/quote
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
}