package config

import (
	"os"
)

// Get - gets specified variable from either environment or default one
func Get(variable string) string {

	var config = map[string]string{
		"API_IP":   "0.0.0.0",
		"API_PORT": "7979",
		"PORT":     "7979",

		"DB_HOST": "127.0.0.1",
		"DB_USER": "api",
		"DB_PASS": "apipass",
		"DB_NAME": "api",
		"DB_SSL":  "disable",

		"SECRET_KEY": "z9Cth9dOuNfCZtxtl8zERxgg79eIV67mM53CrjDISkNAnuvZWYOqoaJH2Hithe9f",

		"YUBIKEY_ID":  "31663",
		"YUBIKEY_KEY": "wwx2QiZOW0z5buRoj6yOTffI+eI=",

		"API_URL":      "https://api.otelia.io",
		"API_USER":     "api",
		"API_PASSWORD": "password",
	}

	for k, v := range config {
		if k == variable {
			if os.Getenv(k) != "" {
				return os.Getenv(k)
			}
			return v
		}
	}
	if os.Getenv(variable) != "" {
		return os.Getenv(variable)
	}
	return ""
}
