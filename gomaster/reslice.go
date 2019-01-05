package main

import "fmt"

func main() {
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)

	//s1 := make([]int, 5)
	s1 := make([]int, 5, 10)
	//s1[4] = 100
	reSlice := s1[2:5]
	fmt.Println(s1)
	fmt.Println(reSlice)

	s1 = append(s1, 200)
	fmt.Println(s1)

	reSlice[0] = -100
	reSlice[1] = 123456

	fmt.Println(s1)
	fmt.Println(reSlice)
}
