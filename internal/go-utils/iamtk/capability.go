package iamtk

import (
	"github.com/google/uuid"
	"github.com/otelia/go-utils/slice"
)

var WildCard, _ = uuid.Parse("00000000-0000-0000-0000-000000000001")

func HasCapability(ps *[]Permission, id uuid.UUID, desiredCapabilities []string) bool {
	for _, p := range *ps {
		if p.Resource == WildCard || p.Resource == id {
			for _, dc := range desiredCapabilities {
				for _, ac := range p.Capabilities {
					if ac == dc {
						return true
					}
				}
			}
		}
	}
	return false
}

func PertinentCapabilities(ps *[]Permission, id uuid.UUID, desiredCapabilities []string) []string {
	ocs := []string{}
	for _, p := range *ps {
		if p.Resource == WildCard || p.Resource == id {
			for _, dc := range desiredCapabilities {
				for _, ac := range p.Capabilities {
					if ac == dc {
						ocs = append(ocs, ac)
					}
				}
			}
		}
	}
	slice.RemoveDuplicates(&ocs)
	return ocs
}
