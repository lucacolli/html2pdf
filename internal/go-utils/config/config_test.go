package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	assert.Equal(t, "0.0.0.0", Get("API_IP"), "the strings should match")
	os.Setenv("API_IP", "1.1.1.1")
	assert.Equal(t, "1.1.1.1", Get("API_IP"), "the strings should match")
	assert.Equal(t, "", Get("UNKNOWN_VARIABLE"), "the strings should match")
	os.Setenv("UNKNOWN_VARIABLE", "test")
	assert.Equal(t, "test", Get("UNKNOWN_VARIABLE"), "the strings should match")
}
