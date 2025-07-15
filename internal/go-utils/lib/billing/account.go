package billing

import (
	"encoding/json"
	"time"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type Account struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Number        string     `json:"number"`
	Vat           string     `json:"vat"`
	SecureEmail   string     `json:"email_secure"`
	LastBilled    *time.Time `json:"last_billed"`
	Periodicity   int        `json:"periodicity"`
	PaymentMethod string     `json:"payment_method"`
	SDDIBAN       string     `json:"payment_sdd_iban"`
	SDDFlag       string     `json:"payment_sdd_flag"`
	Intrastat     bool       `json:"intrastat"`
	// Company address
	Line1   string `json:"address_line1"`
	Line2   string `json:"address_line2"`
	City    string `json:"address_city"`
	State   string `json:"address_state"`
	Zip     string `json:"address_zip"`
	Country string `json:"address_country"`
	// Easy contact
	POCName  string `json:"poc_name"`
	POCPhone string `json:"poc_phone"`
	POCEmail string `json:"poc_email"`
}

func AccountCreate(data Account) (*Account, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts"
	payload, _ := json.Marshal(data)

	r, err := utils.Request("POST", url, payload, token)

	var out Account
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return &out, err
}

func AccountRead(id string) (*Account, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + id

	r, err := utils.Request("GET", url, []byte{}, token)

	var out Account
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return &out, err
}

// TODO: Implement queries
func AccountList() ([]Account, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts?limit=10000"

	r, err := utils.Request("GET", url, []byte{}, token)

	var accounts []Account
	if err != nil {
		return accounts, err
	}
	err = json.Unmarshal(r, &accounts)
	if err != nil {
		return accounts, err
	}
	return accounts, err
}

func AccountUpdate(id string, data Account) (*Account, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + id
	payload, _ := json.Marshal(data)

	r, err := utils.Request("PATCH", url, payload, token)

	var out Account
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return &out, err
}

func AccountDelete(id string) (*Account, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + id

	r, err := utils.Request("DELETE", url, []byte{}, token)

	var out Account
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return &out, err
}
