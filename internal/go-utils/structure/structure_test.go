package structure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/structure"
)

func TestPresentOrDefault(t *testing.T) {
	type Example struct {
		Present string
		Absent  string
	}

	e := Example{
		Present: "a_string",
	}

	structure.PresentOrDefault(&e.Present, "b_string")
	structure.PresentOrDefault(&e.Absent, "b_string")

	c := Example{
		Present: "a_string",
		Absent:  "b_string",
	}
	assert.Equal(t, c, e, "the objects should be the same")
}
