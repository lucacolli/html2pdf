package utils

import (
	"os"
)

// Get - gets specified variable from either environment or default one
func GetConfig(variable string) string {

	var config = map[string]string{
		"API_URL": "https://api.otelia.io",
		"PMS_URL": "https://pms.otelia.io",
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
