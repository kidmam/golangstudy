package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "adfadfasdf"
	C2 = "adfadfasdf"
	C3 = "adfadfasdf"
)

const (
	p2_0 Power2 = 1 << iota
	_
	p2_2
	_
	p2_4
	_
	p2_6
)

func main() {
	const s1 = 123
	const s2 float64 = 123
	var v1 float32 = s1 * 12
	//var v2 float32 = s2 * 12

	fmt.Println(v1)
	//fmt.Println(v2)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Two)

	fmt.Println(p2_6)
}
