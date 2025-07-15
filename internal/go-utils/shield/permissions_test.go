package shield_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/otelia/go-utils/shield"
)

func TestGetPermissions(t *testing.T) {
	var o []string

	// Internal Global Admin
	userAuthorizations := shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"ANY": "ADMIN"},
	}
	o = shield.GetPermissions("REALM", userAuthorizations, "r", "cadd3056-629c-4c79-8632-9be42d2af96d")
	assert.Equal(t, []string{"rY", "rD", "rE", "rF", "rG", "rL", "rM", "rN", "rO", "rP"}, o, "should match")

	// Non-internal Global Admin
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"ANY": "ADMIN"},
	}
	o = shield.GetPermissions("REALM", userAuthorizations, "r", "cadd3056-629c-4c79-8632-9be42d2af96d")
	assert.Equal(t, []string{"rY"}, o, "should match")

	// Internal Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"REALM": "ADMIN"},
	}
	o = shield.GetPermissions("REALM", userAuthorizations, "r", "cadd3056-629c-4c79-8632-9be42d2af96d")
	assert.Equal(t, []string{"rY", "rD", "rE", "rF", "rG", "rL", "rM", "rN", "rO", "rP"}, o, "should match")

	// Internal different Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    true,
		Permissions: map[string]string{"REALM2": "ADMIN"},
	}
	o = shield.GetPermissions("REALM", userAuthorizations, "r", "cadd3056-629c-4c79-8632-9be42d2af96d")
	assert.Equal(t, []string{"rY"}, o, "should match")

	// Non-internal Realm Admin
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"REALM": "ADMIN"},
	}
	o = shield.GetPermissions("REALM", userAuthorizations, "r", "cadd3056-629c-4c79-8632-9be42d2af96d")
	assert.Equal(t, []string{"rY"}, o, "should match")

	// Specific UUID
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"cadd3056-629c-4c79-8632-9be42d2af96d": "ADMINISTRATOR"},
	}
	o = shield.GetPermissions("REALM", userAuthorizations, "r", "cadd3056-629c-4c79-8632-9be42d2af96d")
	assert.Equal(t, []string{"rY", "rD", "rE", "rF", "rG"}, o, "should match")

	// Specific UUID - No power
	userAuthorizations = shield.Authorizations{
		Internal:    false,
		Permissions: map[string]string{"cadd3056-629c-4c79-8632-9be42d2af96d": "ADMINISTRATOR"},
	}
	o = shield.GetPermissions("REALM", userAuthorizations, "r", "6defce52-ee03-4b61-916c-eed032a75fcc")
	assert.Equal(t, []string{"rY"}, o, "should match")
}
