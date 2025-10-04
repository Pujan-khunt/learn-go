// Package example2
package main

import "fmt"

func main() {
	// No deadlock since channel is buffered with a size of 2.
	c := make(chan string, 2)
	c <- "Hello, "
	c <- "World!"
	// c <- "This line will create deadlocks upon running. Just uncomment this code to see it in action"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
