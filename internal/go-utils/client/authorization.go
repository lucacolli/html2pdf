package client

import (
	"encoding/json"

	"github.com/otelia/go-utils/client/utils"
	"github.com/otelia/go-utils/config"
)

func GrantAuthorization(userID string, resource string, role string) error {
	// Perform Login
	token, err := Login(config.Get("API_USER"), config.Get("API_PASSWORD"))
	if err != nil {
		return err
	}

	url := utils.GetConfig("API_URL") + "/iam/v0/authorizations"
	data := make(map[string]string)
	data["user_id"] = userID
	data["resource"] = resource
	data["role"] = role
	payload, _ := json.Marshal(data)
	_, err = utils.Request("POST", url, payload, token)
	if err != nil {
		return err
	}
	return nil
}
