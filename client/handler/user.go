package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name      string `json:"name"`
	Scope     string `json:"scope"`
	Birthdate string `json:"birthdate,omitempty"`
}

func (h Handler) Profile(w http.ResponseWriter, r *http.Request) {
	endpoint := h.Config.API_SERVER_USER_PROFILE_ENDPOINT

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("=====ERROR: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("ERROR: status is not OK: %+v\n", res.Status)
		w.WriteHeader(res.StatusCode)
		w.Write([]byte(`<body><div>forbidden</div><div><a href="/">Home</a></div></body>`))
		return
	}
	defer res.Body.Close()

	user := User{}
	derr := json.NewDecoder(res.Body).Decode(&user)
	if derr != nil {
		log.Printf("ERROR: json.NewDecoder: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userRes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(userRes)
}
