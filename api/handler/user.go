package handler

import (
	"encoding/json"
	"keycloak-demo-5/middleware"
	"log"
	"net/http"
	"strings"
)

type Profile struct {
	Name      string `json:"name"`
	Scope     string `json:"scope"`
	Birthdate string `json:"birthdate,omitempty"`
}

func (h Handler) Profile(w http.ResponseWriter, r *http.Request) {
	token := middleware.GetToken(r.Context())
	if token == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if !strings.Contains(token.Scope, "profile") {
		log.Print("\ntoken doesnt contain profile scope\n")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	profile := Profile{
		Name:      token.Name,
		Scope:     token.Scope,
		Birthdate: token.Birthdate,
	}
	res, err := json.Marshal(profile)
	if err != nil {
		log.Printf("json.Marshal err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
