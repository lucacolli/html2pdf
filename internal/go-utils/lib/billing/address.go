package billing

import (
	"encoding/json"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type Address struct {
	ID        string `json:"id"`
	To        string `json:"to"`
	CareOf    string `json:"care_of"`
	Line1     string `json:"line1"`
	Line2     string `json:"line2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Fax       string `json:"fax"`
	Website   string `json:"website"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	PlaceID   string `json:"place_id"`
	AccountID string `json:"account_id"`
}

func (a *Account) GetAddresses() ([]Address, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + a.ID + "/addresses"

	r, err := utils.Request("GET", url, []byte{}, token)

	var out []Address
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(r, &out)
	if err != nil {
		return out, err
	}
	for k, _ := range out {
		out[k].AccountID = a.ID
	}
	return out, err
}

func (a *Address) Update(data map[string]interface{}) (*Address, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + a.AccountID + "/addresses/" + a.ID

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	r, err := utils.Request("PATCH", url, payload, token)

	var out Address
	if err != nil {
		return &out, err
	}
	err = json.Unmarshal(r, &out)
	if err != nil {
		return &out, err
	}
	return &out, err
}

func (a *Account) CreateAddress(data map[string]interface{}) (*Address, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + a.ID + "/devices"

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	r, err := utils.Request("POST", url, payload, token)

	var out Address
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return &out, err
}
