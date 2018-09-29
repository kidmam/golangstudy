package main

import "fmt"

// Alias *int to intAddress for clarity
type intAddress *int

// addressOf takes an integer and returns the memory address
// of that integer.
func addressOf(i int) intAddress {
	return &i
}

// valueOf takes an intAddress, and returns the value
// at that address.
func valueOf(ia intAddress) int {
	return *ia
}

func main() {
	x := 4
	y := addressOf(x)
	z := valueOf(y)
	fmt.Printf("y: %+v\n", y)
	fmt.Printf("z: %d\n", z)
}
