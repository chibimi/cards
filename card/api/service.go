package api

import (
	"encoding/json"
	"net/http"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/inconshreveable/log15.v2"
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

func (s *Service) ListAbilities(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := s.src.ListAbilities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) ListSpells(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := s.src.ListSpells()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) ListFeats(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := s.src.ListFeats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func writeJson(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if body != nil {
		err := json.NewEncoder(w).Encode(body)
		if err != nil {
			log15.Error("Unable to send JSON body", "err", err.Error())
		}
	}
}
