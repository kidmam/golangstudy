package main

import (
	"fmt"
)

func main() {
	anArray := [4]int{1, 2, 3, -4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[3][3])
	fmt.Println("The length of", threeD, "is", len(threeD))
	fmt.Println("The first element of", threeD, "is", threeD[0][1][1])

	//일차원 배열
	for i := 0; i < len(anArray); i++ {
		fmt.Println("anArray", i, "=", anArray[i])
	}
	for i, v := range anArray {
		fmt.Println("anArray", i, "=", v)
	}

	//2차원 배열
	for i := 0; i < len(twoD); i++ {
		v := twoD[i]
		fmt.Println("v:", twoD[i])
		for j := 0; j < len(v); j++ {
			fmt.Println("Value:", twoD[i][j])
		}
	}
	for i, v := range twoD {
		fmt.Println("anArray", i, "=", v)
	}
	for _, v := range twoD {
		for _, m := range v {
			fmt.Println("anArray=", m)
		}
	}

	//3차원 배열
	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		fmt.Println("v:", threeD[i])
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Println("Value:", threeD[i][j][k])
			}
		}
	}
	for _, v := range threeD {
		for _, m := range v {
			for _, k := range m {
				fmt.Println("Value:", k)
			}
		}
	}
}
