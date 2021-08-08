package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
	"sync/atomic"
)

var balance int64
// var mutex = &sync.Mutex{}

func credit(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		// mutex.Lock()
		// balance += 100
		// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		// fmt.Println("After Crediting, balance is ", balance)
		// mutex.Unlock()

		atomic.AddInt64(&balance, 100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func debit(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// mutex.Lock()
		// balance -= 100
		// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		// fmt.Println("After Debiting, balance is ", balance)
		// mutex.Unlock()

		atomic.AddInt64(&balance, -100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	balance = 1000
	fmt.Println("Initial Balance: ", balance)

	wg.Add(1)
	go credit(&wg)

	wg.Add(1)
	go debit(&wg)
	// fmt.Scanln()		// wait for user input 

	wg.Wait()
	fmt.Println(balance)
}