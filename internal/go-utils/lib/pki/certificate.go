package pki

import (
	"encoding/json"

	"github.com/otelia/hotc/lib/pki/models"
	"github.com/otelia/hotc/lib/utils"
)

func Create(token string, host string) (models.Certificate, error) {
	url := utils.GetConfig("API_URL") + "/v0/pki"

	data := map[string]interface{}{"host": host}
	payload, err := json.Marshal(data)
	if err != nil {
		return models.Certificate{}, err
	}
	r, err := utils.Request("POST", url, payload, token)

	var out models.Certificate
	if err == nil {
		json.Unmarshal(r, &out)
	}
	return out, err
}
