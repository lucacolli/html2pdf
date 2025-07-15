package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/maps"
)

func TestGetKeys(t *testing.T) {
	a := map[string]interface{}{
		"k1": "key1",
		"k2": "key2",
	}

	assert.Equal(t, []string{"k1", "k2"}, maps.GetKeys(a), "the objects should be the same")
}
