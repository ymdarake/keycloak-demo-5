package authz

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Token struct {
	Sub       string `json:"sub"`
	Name      string `json:"name"`
	Scope     string `json:"scope"`
	Exp       int    `json:"exp"`
	Iat       int    `json:"iat"`
	Birthdate string `json:"birthdate"`
}

func TokenFromString(str string) (*Token, error) {
	// TODO: verify with key
	parts := strings.Split(str, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("token must be separated into 3 parts by dot")
	}
	decoded, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	tok := Token{}
	json.Unmarshal(decoded, &tok)
	return &tok, nil
}
