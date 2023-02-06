package main

import (
	"fmt"
	"time"
)

func worker1(done chan int) {
	for i := range done {
		fmt.Printf("Worker 1: %v", i)
	}
}
func worker2(done chan int) {
	for i := range done {
		fmt.Printf("Worker 2: %v", i)
	}
}

func main() {
	done := make(chan int)
	go worker1(done)
	go worker2(done)

	done <- 1
	time.Sleep(1 * time.Second)
	done <- 2
	time.Sleep(1 * time.Second)
	done <- 3
	time.Sleep(1 * time.Second)
	done <- 4
	time.Sleep(1 * time.Second)
	done <- 5

	fmt.Println("main function")
}
