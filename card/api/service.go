package api

import (
	"encoding/json"
	"net/http"
	"objenious/card"

	log "github.com/inconshreveable/log15"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
}

type Service struct {
	src *card.Service
}

func NewService(src *card.Service) *Service {
	return &Service{
		src: src,
	}
}

func (s *Service) ListFactions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeJson(w, s.src.ListFactions(), http.StatusOK)
}

func (s *Service) ListCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeJson(w, s.src.ListCategories(), http.StatusOK)
}

func writeJson(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if body != nil {
		err := json.NewEncoder(w).Encode(body)
		if err != nil {
			log.Error("Unable to send JSON body", "err", err.Error())
		}
	}
}
