package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func alignSize(nums []int) int {
	size := 0
	for _, n := range nums {
		if s := int(math.Log10(float64(n))) + 1; s > size {
			size = s
		}
	}

	return size
}

func main() {

	var e interface{} = 2.7182
	fmt.Printf("e = %v (%T)\n", e, e) // e = 2.7182 (float64)

	fmt.Printf("%10d\n", 353)    // will print "       353"
	fmt.Printf("%*d\n", 10, 353) // will print "       353"

	nums := []int{12, 237, 3878, 3}
	size := alignSize(nums)
	for i, n := range nums {
		fmt.Printf("%02d %*d\n", i, size, n)
	}

	fmt.Printf("The price of %[1]s was $%[2]d. $%[2]d! imagine that.\n", "carrot", 23)

	p := &Point{1, 2}
	fmt.Printf("%v %+v %#v \n", p, p, p)
}
