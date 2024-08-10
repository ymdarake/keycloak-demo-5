package handler

import (
	"encoding/json"
	"fmt"
	"keycloak-demo-5-client/config"
	"net/http"
	"net/url"
	"strings"

	"log"
)

// NOTE: DBがわりにメモリで管理する
var token = ""

type Handler struct {
	Config config.Config
}

type AuthResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	IDToken          string `json:"id_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	Scope            string `json:"scope"`
	Error            string `json:"error,omitempty"`
}

type IntrospectionResponse struct {
	Active bool `json:"active"`
}

func (h Handler) StartAuth(w http.ResponseWriter, r *http.Request) {
	authURL := fmt.Sprintf(
		"%s?response_type=code&client_id=%s&redirect_uri=http%%3A%%2F%%2Flocalhost%%3A8081%%2Fauth%%2Fcallback",
		h.Config.AUTHORIZATION_ENDPOINT,
		h.Config.KEYCLOAK_CLIENT_ID,
	)
	w.Header().Set("Content-Type", "text/json")
	http.Redirect(w, r, authURL, http.StatusFound)
}

func (h Handler) Callback(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")
	cliendID := h.Config.KEYCLOAK_CLIENT_ID
	clientSecret := h.Config.KEYCLOAK_CLIENT_SECRET
	endpoint := h.Config.TOKEN_ENDPOINT

	values := url.Values{}
	values.Set("code", authCode)
	values.Set("grant_type", "authorization_code")
	values.Set("redirect_uri", "http://localhost:8081/auth/callback")
	values.Set("scope", "openid")
	values.Add("client_id", cliendID)
	values.Add("client_secret", clientSecret)

	req, err := http.NewRequest(
		"POST",
		endpoint,
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		log.Printf("=====ERROR: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	authRes := &AuthResponse{}
	derr := json.NewDecoder(res.Body).Decode(authRes)
	if derr != nil {
		log.Printf("ERROR: json.NewDecoder: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonRes, err := json.Marshal(authRes)
	if err != nil {
		log.Printf("ERROR: json.Marshal: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(jsonRes)

	token = authRes.AccessToken
}

func (h Handler) Introspect(w http.ResponseWriter, r *http.Request) {
	endpoint := h.Config.API_SERVER_INTROSPECTION_ENDPOINT
	values := url.Values{}
	values.Set("token", token)

	req, err := http.NewRequest(
		"POST",
		endpoint,
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		log.Printf("=====ERROR: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	authRes := &IntrospectionResponse{}
	derr := json.NewDecoder(res.Body).Decode(authRes)
	if derr != nil {
		log.Printf("ERROR: json.NewDecoder: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonRes, err := json.Marshal(authRes)
	if err != nil {
		log.Printf("ERROR: json.Marshal: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(jsonRes)
}
