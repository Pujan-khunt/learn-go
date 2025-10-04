package main

import (
	"fmt"
	"time"
)

func printMessage(message chan string) {
	time.Sleep(time.Second * 2)
	message <- "Hello from Goroutine"
}

func main() {
	message := make(chan string)
	go printMessage(message)
	fmt.Println("Hello from main function.")
	fmt.Println(<-message)
}
