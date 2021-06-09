package app

import (
	"testing"
)

type Test1 struct {
	name  string `json:"name"`
	age   int    `json:"age"`
	email string `json:"email,omitempty"`
}

func TestSignup(t *testing.T) {
}
