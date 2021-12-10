package model

import "testing"

//TestUser...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "allan@example.org",
		Password: "password",
	}
}
