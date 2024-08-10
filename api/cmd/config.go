package cmd

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KEYCLOAK_CLIENT_ID          string
	KEYCLOAK_CLIENT_SECRET      string
	AUTH_SERVER_URL             string
	AUTH_INTROSPECTION_ENDPOINT string
}

func LoadConfig() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}
