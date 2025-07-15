package shield

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"html2pdf/internal/go-utils/config"
)

type Authorizations struct {
	UserID      string            `json:"user_id"`
	Internal    bool              `json:"internal"`
	Permissions map[string]string `json:"permissions"`
}

func ValidateAuthorization(authorization string) (Authorizations, error) {
	var rp Authorizations
	client := &http.Client{}
	req, _ := http.NewRequest("GET", config.Get("API_URL")+"/v0/user/authorizations", nil)
	req.Header.Set("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		return rp, err
	}
	if resp.StatusCode != 200 {
		return rp, errors.New("Unauthorized")
	}
	defer resp.Body.Close()
	rb, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(rb, &rp)
	return rp, nil
}
