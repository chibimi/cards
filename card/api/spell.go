package api

import (
	"encoding/json"
	"net/http"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateSpell(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	spell := &card.Spell{}
	if err := json.NewDecoder(r.Body).Decode(spell); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.SaveSpell(spell)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusCreated)
}

func (s *Service) UpdateSpell(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.CreateSpell(w, r, p)
		return
	}
	spell := &card.Spell{}
	if err := json.NewDecoder(r.Body).Decode(spell); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.src.UpdateSpell(spell)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}
