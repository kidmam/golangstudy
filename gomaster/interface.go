package main

import "fmt"

type Interface interface {
	Method(string) string
}

type T struct{}

func (T) Method(str string) string {
	return str
}

func f1(i Interface) {
	fmt.Println(i.Method("Hello"))
}
func main() {
	f1(T{})
}
