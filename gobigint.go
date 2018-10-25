package main

import (
	"fmt"
	"strconv"
)

func main() {
	const intSize = 32 << (^uint(0) >> 63)
	fmt.Println(strconv.IntSize)
}
