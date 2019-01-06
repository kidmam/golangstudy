package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)

	integer := make([]int, 2)
	fmt.Println("integer: ", integer)
	integer = nil
	fmt.Println(integer)
	integer = make([]int, 5)
	fmt.Println(integer)

	/*integer[0] = 1
	integer[1] = 2
	fmt.Println(integer)*/ //runtime error: index out of range

	anArray := [5]int{-1, -2, -3, -4, -5}
	refArray := anArray[:]
	fmt.Println(anArray)
	fmt.Println(refArray)
	anArray[4] = -100
	fmt.Println(refArray)

	s := make([]byte, 15)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}
	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value", y)
		}
		fmt.Println()
	}
}
