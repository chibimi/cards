package api

import (
	"encoding/json"
	"net/http"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateFeat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	feat := &card.Feat{}
	if err := json.NewDecoder(r.Body).Decode(feat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.SaveFeat(feat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusCreated)
}

func (s *Service) UpdateFeat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.CreateFeat(w, r, p)
		return
	}
	feat := &card.Feat{}
	if err := json.NewDecoder(r.Body).Decode(feat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.src.UpdateFeat(feat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}
