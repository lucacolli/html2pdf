package shield

import (
	"github.com/google/uuid"

	"html2pdf/internal/go-utils/slice"
)

var Permissions = map[string][]string{
	"ADMIN":         []string{"Y", "D", "E", "F", "G", "L", "M", "N", "O", "P"},
	"SUPPORT2":      []string{"Y", "N", "O"},
	"SUPPORT1":      []string{"Y", "N"},
	"SALES2":        []string{"Y", "L", "M"},
	"SALES1":        []string{"Y", "L"},
	"ADMINISTRATOR": []string{"Y", "D", "E", "F", "G"},
	"MANAGER":       []string{"Y", "D", "E", "F"},
	"USER2":         []string{"Y", "D", "E"},
	"USER1":         []string{"Y", "D"},
	"IOT":           []string{"Y", "X"},
	"ANY":           []string{"Y"}, // NOTE: This could be implemented as a meta permission instead of a tagged permission
	//"AUDITOR":     []string{"Z"}, // NOTE: This could be implemented as a meta permission instead of a tagged permission
}

func HasPermissionsOn(realm string, authorizations Authorizations, resource string) bool {
	if authorizations.Internal {
		// If user is internal, check for capabilities on ALL realms
		if _, ok := authorizations.Permissions["ANY"]; ok {
			return true
		}
		// If user is internal, check for capabilities on the current realm
		if _, ok := authorizations.Permissions[realm]; ok {
			return true
		}
	}
	// If resource is an uuid, check for straight match
	if _, e := uuid.Parse(resource); e == nil {
		if _, ok := authorizations.Permissions[resource]; ok {
			return true
		}
	}
	return false
}

func Permitted(realm string, authorizations Authorizations, sufficientCapabilities []string, resource string) bool {
	if authorizations.Internal {
		// If user is internal, check for capabilities on ALL realms
		if userCapability, ok := authorizations.Permissions["ANY"]; ok {
			if slice.StringInSlice(userCapability, sufficientCapabilities) {
				return true
			}
		}
		// If user is internal, check for capabilities on the current realm
		if userCapability, ok := authorizations.Permissions[realm]; ok {
			if slice.StringInSlice(userCapability, sufficientCapabilities) {
				return true
			}
		}
	}
	// If resource is an uuid, check for straight match
	if _, e := uuid.Parse(resource); e == nil {
		if userCapability, ok := authorizations.Permissions[resource]; ok {
			if slice.StringInSlice(userCapability, sufficientCapabilities) {
				return true
			}
		}
	}
	return false
}

// ups == authorizations.Permissions
//  r: request
//  aps: map["CAPABILITY"]["rR","rP","cR","cC","uR"]
//  m: "c" == mode
//  k: "an-uuid-here" == resource
func GetPermissions(realm string, authorizations Authorizations, mode string, resource string) []string {
	var auths []string
	if ap, ok := Permissions["ANY"]; ok {
		auths = append(auths, ap...)
	}
	if authorizations.Internal {
		// If user is internal, check for capabilities on ALL realms
		if userCapability, ok := authorizations.Permissions["ANY"]; ok {
			auths = append(auths, Permissions[userCapability]...)
		}
		// If user is internal, check for capabilities on the current realm
		if userCapability, ok := authorizations.Permissions[realm]; ok {
			auths = append(auths, Permissions[userCapability]...)
		}
	}
	// If resource is an uuid, check for straight match
	if _, e := uuid.Parse(resource); e == nil {
		if userCapability, ok := authorizations.Permissions[resource]; ok {
			auths = append(auths, Permissions[userCapability]...)
		}
	}

	// Cleanup, add the modal and return
	slice.RemoveDuplicates(&auths)
	for k, v := range auths {
		auths[k] = mode + v
	}
	return auths
}
