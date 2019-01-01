package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("p1: ", p1)
	fmt.Println("*p2: ", *p2)
	fmt.Println("p2: ", p2)

	*p1 = 56373748499555
	fmt.Println("value: ", value)
	fmt.Println("*p2: ", *p2)

	*p1 = 563737
	fmt.Println("value: ", value)
	fmt.Println("*p2: ", *p2)
}
