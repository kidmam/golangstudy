package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{}
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)
}
