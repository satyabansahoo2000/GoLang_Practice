package main

import (
	"fmt"
	"time"
)

func fib(n int, ch chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		ch <- a // blocked until the value is retrieved from channel
		a, b = b, a+b
		time.Sleep(1 * time.Second)
	}
	close(ch) // close the channel
}

func counter(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go fib(10, ch1)
	go counter(10, ch2)

	ch1Closed := false
	ch2Closed := false

	for {
		select {
		case n, ok := <-ch1:
			if !ok {
				ch1Closed = true
				if ch1Closed && ch2Closed {
					return
				}
			} else {
				fmt.Println("Fib() = ", n)
			}
		case m, ok := <-ch2:
			if !ok {
				ch2Closed = true
				if ch1Closed && ch2Closed {
					return
				}
			} else {
				fmt.Println("Counter() = ", m)
			}
		}
	}
}
