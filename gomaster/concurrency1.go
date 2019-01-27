package main

import (
	"fmt"
	"time"
)

func main() {
	go beep()
	fmt.Println("Listening for the beeps")
	time.Sleep(2 * time.Second)
	fmt.Println("That Beep is anoying, quit listening!")
}

func beep() {
	for i := 0; ; i++ {
		fmt.Println("beep", i)
		time.Sleep(time.Second)
	}
}
