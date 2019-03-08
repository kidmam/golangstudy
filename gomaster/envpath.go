package main

import (
	"fmt"
	"os"
)

func main() {
	// Store the PATH environment variable in a variable
	path, exists := os.LookupEnv("PATH")

	if exists {
		// Print the value of the environment variable
		fmt.Print(path)
	}
}
