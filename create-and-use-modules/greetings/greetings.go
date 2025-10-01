// Package greetings provides an interface to greet a named person.
package greetings

import "fmt"

// Hello returns a greeting for a person with a given name.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
