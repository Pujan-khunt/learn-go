package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine which is sends a message to the c1 channel every 500ms
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Sending message every 500ms"
		}
	}()

	// Goroutine which is sends a message to the c2 channel every 2s
	go func() {
		for {
			time.Sleep(time.Second * 2)
			c1 <- "Sending message every 2s"
		}
	}()

	// Infinite loop
	for {
		// Whenever there is a channel ready to transmit data, execute the code below that case statement.
		// Not the fast Goroutine will not be blocked by the slow Goroutine.
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
