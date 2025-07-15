package client

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/otelia/go-utils/client/utils"
)

type Sku struct {
	ID               uuid.UUID `json:"sku"`
	Description      string    `json:"description"`
	RoomsMin         int       `json:"rooms_min"`
	RoomsMax         int       `json:"rooms_max"`
	ShortDescription string    `json:"short_description"`
	SetupPrice       string    `json:"setup_price"`
	YearlyPrice      string    `json:"yearly_price"`
	MaxSetupDiscount string    `json:"maximum_setup_discount"`
	MaxDiscount      string    `json:"maximum_discount"`
	Hardware         bool      `json:"hardware"`
}

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ShortName   string    `json:"short_name"`
	Description string    `json:"description"`
	Vat         string    `json:"vat"`
	Public      bool      `json:"public"`
	Skus        []Sku     `json:"skus"`
}

func GetSKU(token string, sku uuid.UUID) (Sku, error) {
	url := utils.GetConfig("API_URL") + "/v0/products"
	response, err := utils.Request("POST", url, nil, token)
	if err != nil {
		return Sku{}, err
	}
	var products []Product
	err = json.Unmarshal(response, &products)
	if err != nil {
		return Sku{}, err
	}
	for _, p := range products {
		for _, s := range p.Skus {
			if s.ID == sku {
				return s, nil
			}
		}
	}
	return Sku{}, errors.New("None found")
}
