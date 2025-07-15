package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Request(method string, url string, payload []byte, token string) ([]byte, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, err
	}
	if token != "" {
		req.Header.Set("Authorization", "Otelia "+token)
	}
	response, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	if response.StatusCode != 200 {
		return bodyBytes, errors.New(strconv.Itoa(response.StatusCode))
	}
	return bodyBytes, nil
}
