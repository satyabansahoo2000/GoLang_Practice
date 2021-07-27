package main

import (
	"methods/readers"
	err "methods/customErrors"
)

func main() {
	err.PrintErrors()
	readers.PrintReader()
}