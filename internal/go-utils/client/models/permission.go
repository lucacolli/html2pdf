package models

type Permission struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	ResourceID string `json:"resource_id"`
	RoleID     string `json:"role_id"`
}
