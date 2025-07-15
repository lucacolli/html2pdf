package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	assert.Equal(t, "https://api.otelia.io", GetConfig("API_URL"), "the strings should match")
	os.Setenv("API_URL", "http://localhost:5000")
	assert.Equal(t, "http://localhost:5000", GetConfig("API_URL"), "the strings should match")
	assert.Equal(t, "", GetConfig("UNKNOWN_VARIABLE"), "the strings should match")
	os.Setenv("UNKNOWN_VARIABLE", "test")
	assert.Equal(t, "test", GetConfig("UNKNOWN_VARIABLE"), "the strings should match")
}
