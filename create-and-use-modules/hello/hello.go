package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	name := "Pujan"
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
