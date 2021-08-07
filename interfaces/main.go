package main

import (
	"fmt"
)

type DigitsCounter interface {
	EvenOddCount() (int, int)
}

type DigitsString string

func  (ds DigitsString) EvenOddCount() (int, int) {
	evenCounts, oddCounts := 0, 0
	for _, digit := range ds {
		if digit % 2 == 0 {
			evenCounts += 1
		} else {
			oddCounts += 1
		}
	}
	return evenCounts, oddCounts
}

func main() {
	num := DigitsString("012395104")
	fmt.Println(num.EvenOddCount())
}