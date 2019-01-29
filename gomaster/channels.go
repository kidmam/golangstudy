package main

import (
	"fmt"
	"time"
)

func main() {
	beeps := make(chan string)
	go beep(beeps)
	fmt.Println("Listening for the beeps")
	for i := 0; i < 5; i++ {
		fmt.Printf("We received: %s\n", <-beeps)
	}
	fmt.Println("That Beep is anoying, quit listening!")
}

func beep(beeps chan string) {
	for i := 0; ; i++ {
		beeps <- fmt.Sprintf("beep %d", i)
		time.Sleep(time.Second)
	}
}
