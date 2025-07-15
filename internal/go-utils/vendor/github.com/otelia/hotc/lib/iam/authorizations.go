package iam

import (
	"encoding/json"
	"time"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type Authorization struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	ResourceID string    `json:"resource_id"`
	RoleID     string    `json:"role_id"`
	GrantedOn  time.Time `json:"granted_on"`
	GrantedBy  string    `json:"granted_by"`
}

func (u *User) GetAuthorizations() ([]Authorization, error) {
	token, err := auth.GetToken()
	url := utils.GetConfig("API_URL") + "/iam/v0/users/" + u.ID + "/authorizations"
	r, err := utils.Request("GET", url, nil, token)

	var out []Authorization
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}

func (a *Authorization) Delete() error {
	token, err := auth.GetToken()
	url := utils.GetConfig("API_URL") + "/iam/v0/authorizations/" + a.ID
	_, err = utils.Request("DELETE", url, nil, token)
	return err
}

func CreateAuthorization(data map[string]interface{}) (Authorization, error) {
	token, err := auth.GetToken()
	url := utils.GetConfig("API_URL") + "/iam/v0/authorizations"

	payload, err := json.Marshal(data)
	if err != nil {
		return Authorization{}, err
	}
	r, err := utils.Request("POST", url, payload, token)

	var out Authorization
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}

func DeleteAuthorization(id string) error {
	token, err := auth.GetToken()
	url := utils.GetConfig("API_URL") + "/iam/v0/authorizations/" + id
	_, err = utils.Request("DELETE", url, nil, token)
	return err
}

func ListAuthorizationsByObject(id string) ([]Authorization, error) {
	token, err := auth.GetToken()
	url := utils.GetConfig("API_URL") + "/iam/v0/authorizations?resource=" + id
	r, err := utils.Request("GET", url, nil, token)

	var out []Authorization
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}
