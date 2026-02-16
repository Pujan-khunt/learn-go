// Package user
package user

import "fmt"

type User struct {
	Name  string
	Email string
}

func CreateUser(Name string, Email string) (*User, error) {
	if Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	if Email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	return &User{
		Name:  Name,
		Email: Email,
	}, nil
}
