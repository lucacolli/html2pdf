package iam

import (
	"encoding/json"
	"time"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type User struct {
	ID                    string    `json:"id"`
	Name                  string    `json:"name"`
	RealName              string    `json:"real_name"`
	TaxCode               string    `json:"tax_code"`
	PMSAlias              string    `json:"pms_alias"`
	MainEmail             string    `json:"main_email"`
	PasswordChangedOn     time.Time `json:"password_changed_on"`
	PasswordExpiresOn     time.Time `json:"password_expires_on"`
	Yubikey               string    `json:"yubikey"`
	OTPActive             bool      `json:"otp_active"`
	AuthMethod            int       `json:"auth_method"`
	LastLoginOn           time.Time `json:"last_login_on"`
	Language              string    `json:"language"`
	TimeZone              string    `json:"time_zone"`
	CCWaitingForClearance bool      `json:"cc_waiting_for_clearance"`
	CCCleared             bool      `json:"cc_cleared"`
}

func UserCreate(data map[string]interface{}) (User, error) {
	token, err := auth.GetToken()
	if err != nil {
		return User{}, err
	}
	url := utils.GetConfig("API_URL") + "/iam/v0/users"
	payload, _ := json.Marshal(data)

	r, err := utils.Request("POST", url, payload, token)
	var out User
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}

func UserRead(id string, token string) (User, error) {
	url := utils.GetConfig("API_URL") + "/v0/users/" + id

	r, err := utils.Request("GET", url, []byte{}, token)

	var out User
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}

func UsersByName(name string) ([]User, error) {
	token, err := auth.GetToken()
	if err != nil {
		return []User{}, err
	}
	url := utils.GetConfig("API_URL") + "/iam/v0/users?limit=10000&search=" + name

	r, err := utils.Request("GET", url, []byte{}, token)

	var out []User
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}

// TODO: Implement queries
func UserReadMany() ([]User, error) {
	token, err := auth.GetToken()
	if err != nil {
		return []User{}, err
	}
	url := utils.GetConfig("API_URL") + "/iam/v0/users?limit=10000"

	r, err := utils.Request("GET", url, []byte{}, token)

	var out []User
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}

func UserUpdate(id string, data User, token string) (User, error) {
	url := utils.GetConfig("API_URL") + "/v0/users/" + id
	payload, _ := json.Marshal(data)

	r, err := utils.Request("PATCH", url, payload, token)

	var out User
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}

func UserDelete(id string, token string) (User, error) {
	url := utils.GetConfig("API_URL") + "/v0/users/" + id

	r, err := utils.Request("DELETE", url, []byte{}, token)

	var out User
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}
