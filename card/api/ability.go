package api

import (
	"encoding/json"
	"net/http"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ability := &card.Ability{}
	if err := json.NewDecoder(r.Body).Decode(ability); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.SaveAbility(ability)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusCreated)
}

func (s *Service) UpdateAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.CreateAbility(w, r, p)
		return
	}
	ability := &card.Ability{}
	if err := json.NewDecoder(r.Body).Decode(ability); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.src.UpdateAbility(ability)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}
