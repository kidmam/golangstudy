package main

import "fmt"

type Interface interface {
	Method(str string)
}
type T1 struct{}
type T2 struct{}

func (T1) Method(str string) {
	fmt.Println(str + " from T1 Method")
}
func (T2) Method(str string) {
	fmt.Println(str + " from T2 Method")
}
func main() {
	var i Interface
	i = T1{}
	i.Method("Hello") //invoke Method of T1
	i = T2{}
	i.Method("Hello") //invoke method of T2
}
