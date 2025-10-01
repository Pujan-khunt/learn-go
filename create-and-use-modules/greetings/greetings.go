// Package greetings provides an interface to greet a named person.
package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for a person with a given name.
func Hello(name string) (string, error) {
	// If no name was provided, return an error message.
	if name == "" {
		return "", errors.New("no name provided")
	}

	// If name is provided, then return a value that embeds the name in a greeting message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
