package main

import (
	"fmt"
	"math/rand"
)

func sum(nums []int, ch chan int) {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	ch <- sum
	fmt.Println("Done and Continue")
}

func main() {
	s := []int{}
	sliceSize := 10
	for i:= 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}

	ch := make(chan int, 5)
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], ch)
	}

	i = 0
	total := 0

	for i < parts {
		partialSum := <- ch
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()
}