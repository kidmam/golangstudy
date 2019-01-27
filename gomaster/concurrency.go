package main

import (
	"fmt"
	"time"
)

func main() {
	beep()
}

func beep() {
	for i := 0; ; i++ {
		fmt.Println("beep", i)
		time.Sleep(time.Second)
	}
}
