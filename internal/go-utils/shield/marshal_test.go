package shield_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/shield"
)

func TestMarshalSingleItems(t *testing.T) {
	type Device struct {
		ID       string `json:"id" groups:"rY"`
		Public   string `json:"public" groups:"rY"`
		Private  string `json:"private" groups:"rD"`
		Reserved string `json:"reserved" groups:"rP"`
	}
	item := Device{
		ID:       "cadd3056-629c-4c79-8632-9be42d2af96d",
		Public:   "test",
		Private:  "test",
		Reserved: "test",
	}
	var o interface{}
	var expected map[string]interface{}
	var err error

	// Internal Global Admin
	userAuthorizations := shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"ANY": "ADMIN"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test", "reserved": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Non-internal Global Admin
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"ANY": "ADMIN"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Internal Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"REALM": "ADMIN"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test", "reserved": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Internal different Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"REALM2": "ADMIN"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Non-internal Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"REALM": "ADMIN"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Specific UUID
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"cadd3056-629c-4c79-8632-9be42d2af96d": "ADMINISTRATOR"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Specific UUID - Simple user
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"cadd3056-629c-4c79-8632-9be42d2af96d": "USER2"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Specific UUID - No power
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"6defce52-ee03-4b61-916c-eed032a75fcc": "ADMINISTRATOR"},
	}
	expected = map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", item, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")
}

func TestMarshalSlicesOfItems(t *testing.T) {
	type Device struct {
		ID       string `json:"id" groups:"rY"`
		Public   string `json:"public" groups:"rY"`
		Private  string `json:"private" groups:"rD"`
		Reserved string `json:"reserved" groups:"rP"`
	}
	items := []Device{
		Device{
			ID:       "cadd3056-629c-4c79-8632-9be42d2af96d",
			Public:   "test",
			Private:  "test",
			Reserved: "test",
		},
		Device{
			ID:       "6defce52-ee03-4b61-916c-eed032a75fcc",
			Public:   "test2",
			Private:  "test2",
			Reserved: "test2",
		},
	}
	var o interface{}
	var expected []map[string]interface{}
	var err error

	// Internal Global Admin
	userAuthorizations := shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"ANY": "ADMIN"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test", "reserved": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2", "private": "test2", "reserved": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Non-internal Global Admin
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"ANY": "ADMIN"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Internal Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"REALM": "ADMIN"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test", "reserved": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2", "private": "test2", "reserved": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Internal different Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"REALM2": "ADMIN"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Non-internal Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"REALM": "ADMIN"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Specific UUID
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"cadd3056-629c-4c79-8632-9be42d2af96d": "ADMINISTRATOR"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Specific UUID - Simple user
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"cadd3056-629c-4c79-8632-9be42d2af96d": "USER2"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test", "private": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")

	// Specific UUID - No power
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"6defce52-ee03-4b61-916c-eed032a75fcc": "ADMINISTRATOR"},
	}
	expected = []map[string]interface{}{
		map[string]interface{}{"id": "cadd3056-629c-4c79-8632-9be42d2af96d", "public": "test"},
		map[string]interface{}{"id": "6defce52-ee03-4b61-916c-eed032a75fcc", "public": "test2", "private": "test2"},
	}
	o, err = shield.Marshal("REALM", userAuthorizations, "r", items, "ID")
	assert.Equal(t, expected, o, "should match")
	assert.Equal(t, err, nil, "should match")
}
