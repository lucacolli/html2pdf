package client

import (
	"encoding/base64"
	"encoding/json"

	"github.com/otelia/go-utils/client/utils"
)

func Pdf(token string, in []byte) ([]byte, error) {
	url := utils.GetConfig("API_URL") + "/v0/pdfgen/fromhtml"
	data := make(map[string]string)
	data["html"] = base64.StdEncoding.EncodeToString(in)
	payload, _ := json.Marshal(data)
	response, err := utils.Request("POST", url, payload, token)
	if err != nil {
		return []byte{}, err
	}
	out := response
	return out, nil
}
