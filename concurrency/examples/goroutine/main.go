package main

import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("Hello from a Goroutine.")
}

func main() {
	go printMessage()
	fmt.Println("Hello from the main function.")
	// Wait for the Goroutine to finish.
	time.Sleep(time.Second)
}
