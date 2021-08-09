package main

import (
	"fmt"
	"math/rand"
)

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{}
	size := 10
	for i := 0; i < size; i++ {
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println("Numbers are :", nums)
	ch := make(chan int)

	// Splitting the numbers into parts and sum it up
	partSize := 2
	parts := size / partSize
	i := 0
	for i < parts {
		go sum(nums[i*partSize:(i+1)*partSize], ch)
		i += 1
	}

	i = 0
	total := 0
	for i < parts {
		total += <- ch
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()
}