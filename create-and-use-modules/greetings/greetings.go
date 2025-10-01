// Package greetings provides an interface to greet a named person.
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for a person with a given name.
func Hello(name string) (string, error) {
	// If no name was provided, return an error message.
	if name == "" {
		return "", errors.New("no name provided")
	}

	// If name is provided, then return a value that embeds the name in a greeting message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Is your name %v?",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}
