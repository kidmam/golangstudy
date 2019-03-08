package main

import (
	"fmt"
	"os"
)

func main() {
	// Set the USERNAME environment variable to "MattDaemon"
	os.Setenv("USERNAME", "MattDaemon")

	// Get the USERNAME environment variable
	username := os.Getenv("USERNAME")

	// Prints out username environment variable
	fmt.Println(username)
}
