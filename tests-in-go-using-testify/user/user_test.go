package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser_Testify(t *testing.T) {
	name := "Pujan"
	email := "pujankhunt2412@gmail.com"
	user, err := CreateUser(name, email)
	if err != nil {
		t.Errorf("Error creating user.")
	}

	assert.Equal(t, name, user.Name, "User name should match with the expected value")
	assert.Equal(t, email, user.Email, "User name should match with the expected value")
}

func TestCreateUser(t *testing.T) {
	name := "Pujan"
	email := "pujankhunt2412@gmail.com"
	user, err := CreateUser(name, email)
	if err != nil {
		t.Errorf("Error creating user.")
	}

	if user.Name != name {
		t.Errorf("Expected user name to be %q, but got %q", name, user.Name)
	}

	if user.Email != email {
		t.Errorf("Expected user email to be %q, but got %q", email, user.Email)
	}
}
