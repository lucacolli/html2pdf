package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otelia/hotc/lib/iam"
)

func TestArrayer(t *testing.T) {
	// iam
	u1 := iam.User{Username: "Test1", Email: "test1@otelia.io", Password: "Password1"}
	u2 := iam.User{Username: "Test2", Email: "test2@otelia.io", Password: "Password2"}
	u := []iam.User{u1, u2}

	ui := make([]interface{}, len(u))
	for i, v := range u {
		ui[i] = v
	}

	// Create output 1
	uo1 := Arrayer(ui, []string{"username", "email"})
	ue1 := [][]string{
		[]string{"Test1", "test1@otelia.io"},
		[]string{"Test2", "test2@otelia.io"},
	}
	assert.Equal(t, ue1, uo1, "Arrays should match")
}
