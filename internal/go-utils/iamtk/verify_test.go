package iamtk_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/iamtk"
)

func TestValidateField(t *testing.T) {
	// Test set
	tests := []struct {
		Data  []byte
		Error error
	}{
		{[]byte(`{"public": "test"}`), nil},
		{[]byte(`{"private": "test"}`), errors.New("Field private not writable")},
		{[]byte(`{"absent": "test"}`), errors.New("Field absent not supported")},
		{[]byte(`{"omitted": "test"}`), nil},
	}

	// Collateral data
	type Device struct {
		PublicField  string `json:"public" groups:"uP"`
		PrivateField string `json:"private" groups:"uR"`
		Omitted      string `json:"omitted,omitempty" groups:"uP"`
	}
	ga := []string{"uP"}

	// Test
	for _, test := range tests {
		e := iamtk.ValidateFields(test.Data, Device{}, ga)
		assert.Equal(t, test.Error, e, "they should be equal")
	}
}
