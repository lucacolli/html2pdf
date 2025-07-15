package billing

import (
	"encoding/json"
	"time"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type Order struct {
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	ContractDate   time.Time  `json:"contract_date"`
	EndDate        *time.Time `json:"end_date"`
	Description    string     `json:"description"`
	Product        string     `json:"product"`
	SKU            string     `json:"sku"`
	ListedPriceDec string     `json:"listed_price"`
	PriceDec       string     `json:"price"`
	Quantity       int        `json:"quantity"`
	Discount       int        `json:"discount"`
	Recursive      bool       `json:"recursive"`
	AddressID      string     `json:"address_id"`
	AccountID      string     `json:"account_id"`
}

func (a *Account) GetOrders() ([]Order, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + a.ID + "/addresses"

	r, err := utils.Request("GET", url, []byte{}, token)

	var out []Order
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

func (o *Order) Update(data map[string]interface{}) (*Order, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("API_URL") + "/v0/accounts/" + o.AccountID + "/addresses/" + o.ID

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	r, err := utils.Request("PATCH", url, payload, token)

	var out Order
	if err != nil {
		return &out, err
	}
	err = json.Unmarshal(r, &out)
	if err != nil {
		return &out, err
	}
	return &out, err
}

func (a *Account) CreateOrder(data map[string]interface{}) (*Order, error) {
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

	var out Order
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return &out, err
}
