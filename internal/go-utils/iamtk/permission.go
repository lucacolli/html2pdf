package iamtk

import "github.com/google/uuid"

type Permission struct {
	Resource     uuid.UUID `json:"resource" cql:"resource" groups:"IAMSessionRead,SuperRead"`
	Capabilities []string  `json:"capabilities" cql:"capabilities" groups:"IAMSessionRead,SuperRead"`
}
