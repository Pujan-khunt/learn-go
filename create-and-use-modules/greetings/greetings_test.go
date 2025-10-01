package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName() calls greeting.Hello with a name, checking for a valid return value.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHelloName() calls greeting.Hello with a name, checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Pujan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Pujan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
