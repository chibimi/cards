package utils

import (
	"encoding/json"
	"net/http"

	"gopkg.in/inconshreveable/log15.v2"
)

func WriteJson(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if body != nil {
		err := json.NewEncoder(w).Encode(body)
		if err != nil {
			log15.Error("Unable to send JSON body", "err", err.Error())
		}
	}
}
