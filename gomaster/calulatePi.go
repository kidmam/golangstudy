package main

import "math/big"

var precision uint = 0

func Pi(accuracy uint) *big.Float {
	k := 0
	pi := new(big.Float).SetPrec(precision).SetFloat64(0)

	for {
		if k > int(precision) {
			break
		}
		k++
	}
	return pi
}

func main() {

}
