package shield_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/shield"
)

func TestValidateField(t *testing.T) {
	type Device struct {
		PublicField  string `json:"public" groups:"uP"`
		PrivateField string `json:"private" groups:"uR"`
	}
	jss := []byte(`{"public": "test"}`)
	jsf := []byte(`{"private": "test"}`)
	jsa := []byte(`{"absent": "test"}`)
	ga := []string{"uP"}
	d := Device{}
	es := shield.ValidateFields(jss, d, ga)
	ef := shield.ValidateFields(jsf, d, ga)
	ea := shield.ValidateFields(jsa, d, ga)
	assert.Equal(t, nil, es, "they should be equal")
	assert.Equal(t, errors.New("Field private not writable"), ef, "they should be equal")
	assert.Equal(t, errors.New("Field absent not supported"), ea, "they should be equal")
}
