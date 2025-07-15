package products

import (
	"encoding/json"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
	Vat         string `json:"vat"`
	Public      bool   `json:"public"`
	Skus        []Sku  `json:"skus"`
}

type Sku struct {
	ID               string `json:"sku"`
	Description      string `json:"description"`
	RoomsMin         int    `json:"rooms_min"`
	RoomsMax         int    `json:"rooms_max"`
	ShortDescription string `json:"short_description"`
	SetupPrice       string `json:"setup_price"`
	YearlyPrice      string `json:"yearly_price"`
	MaxSetupDiscount string `json:"maximum_setup_discount"`
	MaxDiscount      string `json:"maximum_discount"`
	Hardware         bool   `json:"hardware"`
}

func List() ([]Product, error) {
	var products []Product
	token, err := auth.GetToken()
	if err != nil {
		return products, err
	}
	url := utils.GetConfig("API_URL") + "/v0/products"

	r, err := utils.Request("GET", url, []byte{}, token)
	if err != nil {
		return products, err
	}

	err = json.Unmarshal(r, &products)
	return products, err
}

func Get(productid string) (Product, error) {
	var product Product
	token, err := auth.GetToken()
	if err != nil {
		return product, err
	}
	url := utils.GetConfig("API_URL") + "/v0/products/" + productid

	r, err := utils.Request("GET", url, []byte{}, token)
	if err != nil {
		return product, err
	}

	err = json.Unmarshal(r, &product)
	return product, err
}
