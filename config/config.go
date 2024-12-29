package config

import "os"

func GetAppPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
