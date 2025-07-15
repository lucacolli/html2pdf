package pms

import (
	"encoding/json"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p *Pms) GetProperties() ([]Property, error) {
	var properties []Property
	token, err := auth.GetToken()
	if err != nil {
		return properties, err
	}
	url := utils.GetConfig("PMS_URL") + "/" + p.ID + "/api/resources/properties/v1"

	r, err := utils.Request("GET", url, []byte{}, token)
	if err != nil {
		return properties, err
	}

	err = json.Unmarshal(r, &properties)
	return properties, err
}
