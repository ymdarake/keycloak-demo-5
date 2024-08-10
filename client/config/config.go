package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT                   int
	AUTHORIZATION_ENDPOINT string
	TOKEN_ENDPOINT         string
	REVOKE_ENDPOINT        string
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
