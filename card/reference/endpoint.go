package reference

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card/utils"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateRef(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ref := &Reference{}
	if err := json.NewDecoder(r.Body).Decode(ref); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.Create(ref)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusCreated)
}

func (s *Service) SaveRef(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ref := &Reference{}
	if err := json.NewDecoder(r.Body).Decode(ref); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if id != ref.ID || id == 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	err = s.Save(ref, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusAccepted)
}

func (s *Service) ListRef(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fid, err := strconv.Atoi(r.URL.Query().Get("faction_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cid, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	res, err := s.List(fid, cid, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}

func (s *Service) GetRef(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
