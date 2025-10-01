package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	fmt.Println("Running executable...")

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Pujan", "Narendra", "Smit"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	fmt.Println("Executable executed successfully.")
}
