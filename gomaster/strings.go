package main

import "fmt"

func main() {
	const sLiternal = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiternal)
	fmt.Printf("x: %x\n", sLiternal)
	fmt.Printf("sLiternal length: %d\n", len(sLiternal))

	for i := 0; i < len(sLiternal); i++ {
		fmt.Printf("%x ", sLiternal[i])
	}
	fmt.Println()

	fmt.Printf("q: %q\n", sLiternal)
	fmt.Printf("+q: %+q\n", sLiternal)
	fmt.Printf("x: % x\n", sLiternal)

	fmt.Printf("s: As a string %s\n", sLiternal)
}
