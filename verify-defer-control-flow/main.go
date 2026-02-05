package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := doThing2()
	if err != nil {
		log.Fatalf("Error running doThing: %v", err)
	}
}

func doThing() error {
	err := errors.New("Hello world!")
	if err != nil {
		return err
	}
	// This defer function won't run
	defer func() {
		fmt.Println("OMG defer didn't run.")
	}()
	return nil
}

func doThing2() error {
	err := errors.New("Hello world!")
	// This defer function would run
	defer func() {
		fmt.Println("OMG defer did run.")
	}()
	if err != nil {
		return err
	}
	return nil
}
