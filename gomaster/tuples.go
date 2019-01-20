package main

import "fmt"

func retThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
}

func main() {
	fmt.Println(retThree(10))
	n1, n2, n3 := retThree(20)
	fmt.Println(n1, n2, n3)

	n1, n2 = n2, n1
	fmt.Println(n1, n2, n3)
}
