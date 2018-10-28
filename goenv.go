package main

import (
	"fmt"
	"os"
)

func main() {
	// 모든 환경변수 출력
	for index, env := range os.Environ() {
		fmt.Println(index, env)
	}
}
