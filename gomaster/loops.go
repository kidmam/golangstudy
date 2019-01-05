package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Println(i, " ")
	}
	fmt.Println()

	i := 10
	for { //while
		if i < 0 {
			break
		}
		fmt.Println(i, " ")
		i--
	}
	fmt.Println()

	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression { //do ..while
		if i > 10 {
			anExpression = false
		}
		fmt.Println(i, " ")
		i++
	}
	fmt.Println()

	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index: ", i, " value: ", value)
	}
}
