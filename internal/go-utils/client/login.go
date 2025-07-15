package client

import (
	"encoding/json"

	"github.com/otelia/go-utils/client/models"
	"github.com/otelia/go-utils/client/utils"
)

func Login(user string, pass string) (string, error) {
	url := utils.GetConfig("API_URL") + "/iam/v0/login"
	data := make(map[string]string)
	data["username"] = user
	data["password"] = pass
	payload, _ := json.Marshal(data)

	var res models.Token
	req, err := utils.Request("POST", url, payload, "")
	if err != nil {
		return "", err
	}
	json.Unmarshal(req, &res)
	return res.Token, nil
}
