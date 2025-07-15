package iamtk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/otelia/go-utils/config"
)

type CurrentSession struct {
	ID          string       `json:"id"`
	UserID      uuid.UUID    `json:"user_id"`
	UserName    string       `json:"user_name"`
	PMSAlias    string       `json:"pms_alias"`
	CCCleared   bool         `json:"cc_cleared"`
	Internal    bool         `json:"internal"`
	NotBefore   time.Time    `json:"not_before"`
	NotAfter    time.Time    `json:"not_after"`
	Permissions []Permission `json:"permissions"`
}

func ValidateAuthorization(authorization string) (CurrentSession, error) {
	var rp CurrentSession
	client := &http.Client{}
	if len(authorization) < 6 {
		return rp, errors.New("Unrecognized authorization method")
	}
	// XXX: BEGIN --- LEGACY WORK AROUND FOR HW-WIFI <= 23
	if authorization[:6] != "Otelia" {
		if authorization[:6] != "Bearer" {
			return rp, errors.New("Unrecognized authorization method")
		}
	}
	// XXX:  END  --- LEGACY WORK AROUND FOR HW-WIFI <= 23
	if len(authorization) < 8 {
		return rp, errors.New("Unauthorized")
	}
	req, _ := http.NewRequest("GET", config.Get("API_URL")+"/iam/v0/sessions/"+authorization[7:], nil)
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
