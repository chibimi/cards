package weapon

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card/utils"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	wp := &Weapon{}
	if err := json.NewDecoder(r.Body).Decode(wp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	res, err := s.Create(wp, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusCreated)
}

func (s *Service) SaveEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.CreateEndpoint(w, r, p)
		return
	}
	wp := &Weapon{}
	if err := json.NewDecoder(r.Body).Decode(wp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if id != wp.ID {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	err = s.Save(wp, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusAccepted)
}

func (s *Service) ListEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	res, err := s.List(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}

func (s *Service) GetEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	res, err := s.Get(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}

func (s *Service) DeleteEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusNoContent)
}
