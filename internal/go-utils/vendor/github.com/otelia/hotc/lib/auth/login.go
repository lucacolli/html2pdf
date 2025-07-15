package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"

	"github.com/otelia/hotc/lib/utils"
)

var token string

type configs struct {
	Token string `json:"token"`
}

func Login(user string, pass string) (string, error) {
	url := utils.GetConfig("API_URL") + "/iam/v0/login"
	data := make(map[string]string)
	data["username"] = user
	data["password"] = pass
	payload, _ := json.Marshal(data)

	type Res struct {
		ID string `json:"id"`
	}
	var res Res
	req, err := utils.Request("POST", url, payload, "")
	if err != nil {
		return res.ID, err
	}
	json.Unmarshal(req, &res)
	token = res.ID
	return res.ID, nil
}

func SetToken(t string) {
	token = t
	usr, err := user.Current()
	if err != nil {
		log.Println("Impossible to retrieve user home")
		return
	}

	c := configs{Token: t}
	bytes, err := json.Marshal(c)
	if err != nil {
		log.Println("Impossible to marshal the token to save it in JSON file")
	}
	err = ioutil.WriteFile(usr.HomeDir+`/.hotc.json`, bytes, 0600)
	if err != nil {
		log.Println("Impossible save session to user home")
	}
}

func GetToken() (string, error) {
	if len(token) > 0 {
		return token, nil
	}
	usr, err := user.Current()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	raw, err := ioutil.ReadFile(usr.HomeDir + "/.hotc.json")
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	var c configs
	err = json.Unmarshal(raw, &c)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return c.Token, nil
}
