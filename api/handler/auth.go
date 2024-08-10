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
	endpoint := fmt.Sprintf("%s%s", h.Config.AUTH_SERVER_URL, h.Config.AUTH_INTROSPECTION_ENDPOINT)

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

	basicAuth := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", h.Config.KEYCLOAK_CLIENT_ID, h.Config.KEYCLOAK_CLIENT_SECRET))))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", basicAuth)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

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

	fmt.Println(endpoint, authRes.Error)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(jsonRes)
}
