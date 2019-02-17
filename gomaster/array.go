package main

import "fmt"

func main() {
	var a [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println("array a => ", a)
	fmt.Println("elements => ", a[0], a[1], a[2])

	var b [3]int = [3]int{1, 2, 3}
	fmt.Println("array b => ", b)

	var c = [3]int{1, 2, 3}
	fmt.Println("array c => ", c)

	d := [3]int{1, 2, 3}
	fmt.Println("array d => ", d)

	greetings := [...]string{
		"Good morning!",
		"Good afternoon!",
		"Good evening!",
		"Good night!",
	}

	fmt.Println(greetings)
	fmt.Println(len(greetings))

	a1 := [3]int{1, 2, 3}
	b1 := [3]int{1, 3, 2}
	c1 := [3]int{1, 1, 1}
	d1 := [3]int{1, 2, 3}

	fmt.Println("a == b", a1 == b1)
	fmt.Println("a == c", a1 == c1)
	fmt.Println("a == d", a1 == d1)

	for index, value := range greetings {
		fmt.Printf("%d =>  %s", index, value)
	}
}
