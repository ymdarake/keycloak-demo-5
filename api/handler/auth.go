package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"keycloak-demo-5/config"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Handler struct {
	Config config.Config
}

type IntrospectionResponse struct {
	Active bool   `json:"active"`
	Error  string `json:"error"`
}

func (h Handler) Introspect(w http.ResponseWriter, r *http.Request) {
	token := r.PostFormValue("token")

	res, err := introspect(h.Config, token)
	if err != nil {
		log.Printf("ERROR: introspect: %v", err)
		return
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Printf("ERROR: json.Marshal: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(jsonRes)
}

func introspect(conf config.Config, token string) (*IntrospectionResponse, error) {
	endpoint := fmt.Sprintf("%s%s", conf.AUTH_SERVER_URL, conf.AUTH_INTROSPECTION_ENDPOINT)

	values := url.Values{}
	values.Set("token", token)

	req, err := http.NewRequest(
		"POST",
		endpoint,
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		log.Printf("=====ERROR: %+v\n", err)
		return nil, err
	}

	basicAuth := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", conf.KEYCLOAK_CLIENT_ID, conf.KEYCLOAK_CLIENT_SECRET))))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", basicAuth)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
		return nil, err
	}
	defer res.Body.Close()

	authRes := &IntrospectionResponse{}
	derr := json.NewDecoder(res.Body).Decode(authRes)
	if derr != nil {
		log.Printf("ERROR: json.NewDecoder: %+v\n", err)
		return nil, derr
	}

	return authRes, nil
}
