package iamtk_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/iamtk"
)

//func Marshal(ps *[]Permission, items interface{}, desiredCapabilities []string, field string) (interface{}, error) {
func TestMarshalSingleItems(t *testing.T) {
	type Device struct {
		ID      uuid.UUID `json:"id" groups:"DeviceList"`
		Public  string    `json:"public" groups:"DeviceList"`
		Private string    `json:"private" groups:"DevicePrivateList"`
	}

	uuid1, _ := uuid.NewRandom()
	item := Device{
		ID:      uuid1,
		Public:  "test",
		Private: "test",
	}
	var o interface{}
	var err error

	// Wildcard Permissions (all)
	userPermissions := []iamtk.Permission{
		iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{"DeviceList", "DevicePrivateList"}},
	}
	expected := map[string]interface{}{"id": uuid1, "public": "test", "private": "test"}
	o, err = iamtk.Marshal(&userPermissions, item, []string{"DeviceList", "DevicePrivateList"}, "ID", 10)
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, nil, err, "should match")

	// Wildcard Permissions (some)
	userPermissions = []iamtk.Permission{
		iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{"DeviceList"}},
	}
	expected = map[string]interface{}{"id": uuid1, "public": "test"}
	o, err = iamtk.Marshal(&userPermissions, item, []string{"DeviceList", "DevicePrivateList"}, "ID", 10)
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, nil, err, "should match")

	// Wildcard Permissions (none)
	userPermissions = []iamtk.Permission{
		iamtk.Permission{Resource: iamtk.WildCard, Capabilities: []string{}},
	}
	expected = map[string]interface{}{}
	o, err = iamtk.Marshal(&userPermissions, item, []string{"DeviceList", "DevicePrivateList"}, "ID", 10)
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, nil, err, "should match")

	// Specific Permissions (all)
	userPermissions = []iamtk.Permission{
		iamtk.Permission{Resource: uuid1, Capabilities: []string{"DeviceList", "DevicePrivateList"}},
	}
	expected = map[string]interface{}{"id": uuid1, "public": "test", "private": "test"}
	o, err = iamtk.Marshal(&userPermissions, item, []string{"DeviceList", "DevicePrivateList"}, "ID", 10)
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, nil, err, "should match")

	// Specific Permissions (some)
	userPermissions = []iamtk.Permission{
		iamtk.Permission{Resource: uuid1, Capabilities: []string{"DeviceList"}},
	}
	expected = map[string]interface{}{"id": uuid1, "public": "test"}
	o, err = iamtk.Marshal(&userPermissions, item, []string{"DeviceList", "DevicePrivateList"}, "ID", 10)
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, nil, err, "should match")

	// Specific Permissions (none)
	userPermissions = []iamtk.Permission{
		iamtk.Permission{Resource: uuid1, Capabilities: []string{}},
	}
	expected = map[string]interface{}{}
	o, err = iamtk.Marshal(&userPermissions, item, []string{"DeviceList", "DevicePrivateList"}, "ID", 10)
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, nil, err, "should match")
}
