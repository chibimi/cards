package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	model := &card.Model{}
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.SaveModel(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusCreated)
}

func (s *Service) UpdateModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.CreateModel(w, r, p)
		return
	}
	model := &card.Model{}
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := s.src.UpdateModel(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}

func (s *Service) GetModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := s.src.GetModel(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) ListModels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := s.src.ListModels()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) DeleteModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.src.DeleteModel(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusNoContent)
}

func (s *Service) GetModelAbilities(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.GetModelAbilities(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) AddModelAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cardID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	abilityID, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.src.AddModelAbility(cardID, abilityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}

func (s *Service) DeleteModelAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cardID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	abilityID, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.src.DeleteModelAbility(cardID, abilityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, nil, http.StatusNoContent)
}
