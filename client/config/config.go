package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT                              int
	API_SERVER_URL                    string
	API_SERVER_INTROSPECTION_ENDPOINT string
	AUTHORIZATION_ENDPOINT            string
	TOKEN_ENDPOINT                    string
	REVOKE_ENDPOINT                   string
	KEYCLOAK_CLIENT_ID                string
	KEYCLOAK_CLIENT_SECRET            string
}

func LoadConfig() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
