package iamtk_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/iamtk"
)

func TestPertinentCapabilities(t *testing.T) {
	// run the Fib function b.N times
	uuid1, _ := uuid.NewRandom()
	dc := []string{"r", "w"}

	// All via wildcard
	ps := []iamtk.Permission{iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{"r", "w"}}}
	oe := []string{"r", "w"}
	o := iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")

	// Only w via wildcard
	ps = []iamtk.Permission{iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{"w"}}}
	oe = []string{"w"}
	o = iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")

	// None via wildcard
	ps = []iamtk.Permission{iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{}}}
	oe = []string{}
	o = iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")

	// All via specific
	ps = []iamtk.Permission{iamtk.Permission{Resource: uuid1, Capabilities: []string{"r", "w"}}}
	oe = []string{"r", "w"}
	o = iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")

	// Only w via specific
	ps = []iamtk.Permission{iamtk.Permission{Resource: uuid1, Capabilities: []string{"w"}}}
	oe = []string{"w"}
	o = iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")

	// None via specific
	ps = []iamtk.Permission{iamtk.Permission{Resource: uuid1, Capabilities: []string{}}}
	oe = []string{}
	o = iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")

	// Mix and Match
	ps = []iamtk.Permission{
		iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{"r"}},
		iamtk.Permission{Resource: uuid1, Capabilities: []string{"w"}},
	}
	oe = []string{"r", "w"}
	o = iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")

	// Overlapping
	ps = []iamtk.Permission{
		iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{"w"}},
		iamtk.Permission{Resource: uuid1, Capabilities: []string{"w"}},
	}
	oe = []string{"w"}
	o = iamtk.PertinentCapabilities(&ps, uuid1, dc)
	assert.Equal(t, oe, o, "should match")
}

func BenchmarkPertinentCapabilities(b *testing.B) {
	// run the Fib function b.N times
	uuid1, _ := uuid.NewRandom()
	uuid2, _ := uuid.NewRandom()
	uuid3, _ := uuid.NewRandom()
	uuid4, _ := uuid.NewRandom()
	uuid5, _ := uuid.NewRandom()
	ps := []iamtk.Permission{
		iamtk.Permission{
			Resource:     uuid1,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid2,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid3,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid4,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid5,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     iamtk.WildCard,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "write"},
		},
	}
	dc := []string{"read", "write"}
	for n := 0; n < b.N; n++ {
		iamtk.PertinentCapabilities(&ps, uuid5, dc)
	}
}

func BenchmarkHasCapability(b *testing.B) {
	// run the Fib function b.N times
	uuid1, _ := uuid.NewRandom()
	uuid2, _ := uuid.NewRandom()
	uuid3, _ := uuid.NewRandom()
	uuid4, _ := uuid.NewRandom()
	uuid5, _ := uuid.NewRandom()
	ps := []iamtk.Permission{
		iamtk.Permission{
			Resource:     uuid1,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid2,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid3,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid4,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
		iamtk.Permission{
			Resource:     uuid5,
			Capabilities: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z", "read"},
		},
	}
	dc := []string{"read", "write"}
	for n := 0; n < b.N; n++ {
		iamtk.HasCapability(&ps, uuid5, dc)
	}
}
