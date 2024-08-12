package authz

import (
	"encoding/base64"
	"encoding/json"
)

type TokenPayload struct {
	Sub       string `json:"sub"`
	Name      string `json:"name"`
	Scope     string `json:"scope"`
	Exp       int    `json:"exp"`
	Iat       int    `json:"iat"`
	Birthdate string `json:"birthdate"`
}

func TokenPayloadFromString(str string) (*TokenPayload, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	tok := TokenPayload{}
	json.Unmarshal(decoded, &tok)
	return &tok, nil
}
