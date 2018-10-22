package main

import (
	"fmt"
	"time"
)

// Listing 3
func producer(pipeline chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		pipeline <- i
	}
}

func consumer(pipeline <-chan int) {
	for value := range pipeline {
		// processing value . . .
		fmt.Printf("Value processed: %d\n", value)
	}
}

func main() {
	pipeline := make(chan int)
	go producer(pipeline)
	go consumer(pipeline)
	// waiting for producer and consumer to complete
	time.Sleep(16 * time.Second)
}
