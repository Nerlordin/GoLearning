package main

import (
	"fmt"

	"mystrings"
)

// main is the entry point of the program.
// It sets up the routes for the application and starts the server.
func main() {
	fmt.Println(mystrings.Reverse("hello world"))
}
