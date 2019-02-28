package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateWeapon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	weapon := &card.Weapon{}
	if err := json.NewDecoder(r.Body).Decode(weapon); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.SaveWeapon(weapon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusCreated)
}

func (s *Service) UpdateWeapon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.CreateWeapon(w, r, p)
		return
	}
	weapon := &card.Weapon{}
	if err := json.NewDecoder(r.Body).Decode(weapon); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := s.src.UpdateWeapon(weapon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}

func (s *Service) GetWeapon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := s.src.GetWeapon(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) ListWeapons(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := s.src.ListWeapons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) DeleteWeapon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.src.DeleteWeapon(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusNoContent)
}

func (s *Service) GetWeaponAbilities(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.GetWeaponAbilities(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) AddWeaponAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	err = s.src.AddWeaponAbility(cardID, abilityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}

func (s *Service) DeleteWeaponAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	err = s.src.DeleteWeaponAbility(cardID, abilityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, nil, http.StatusNoContent)
}
