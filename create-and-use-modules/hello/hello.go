package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Pujan")
	fmt.Println(message)
}
