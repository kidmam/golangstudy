package main

import "fmt"

func main() {
	//c := make(chan int)
	c := make(chan int, 1)
	c <- 101         //수신루틴이 없으므로 데드락
	fmt.Println(<-c) //코멘트해도 데르락 (별도의 Go루틴없기 때문)
}
