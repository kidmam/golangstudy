package main

import "fmt"

func printValue(c chan int) {
	val := <-c
	fmt.Printf("Value received: %d", val)
}

func main() {
	ch := make(chan int)

	go printValue(ch)

	value := 20

	ch <- value

}
