package main

import "fmt"

func printValue1(c chan int) {
	val := <-c
	fmt.Printf("Value received: %d", val)
}
func main() {
	ch := make(chan int, 1)
	ch <- 10
	ch <- 20
	printValue1(ch)
}
