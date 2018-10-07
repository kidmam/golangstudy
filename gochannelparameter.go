package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 에러발생
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
