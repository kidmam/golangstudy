package main

import "fmt"

func main() {
	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)
}

func printIt(v interface{}) {
	fmt.Println(v) //Tom
}
