package cmd

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	API_SERVER_URL         string
	KEYCLOAK_CLIENT_ID     string
	KEYCLOAK_CLIENT_SECRET string
}

func LoadConfig() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
