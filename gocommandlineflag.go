package main

import (
	"flag"
	"fmt"
)

func main() {

	file := flag.String("file", "default.txt", "Input file")
	trials := flag.Int("maxtrial", 10, "Max Trial Count")
	isroot := flag.Bool("root", false, "Run as root")

	flag.Parse()

	// 포인터 변수이므로 앞에 * 를 붙어 deference 해야
	fmt.Println(*file, *trials, *isroot)
}
