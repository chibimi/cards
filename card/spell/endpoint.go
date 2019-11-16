package spell

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card/utils"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sp := &Spell{}
	if err := json.NewDecoder(r.Body).Decode(sp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	res, err := s.Create(sp, lang)
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
	sp := &Spell{}
	if err := json.NewDecoder(r.Body).Decode(sp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if id != sp.ID {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	err = s.Save(sp, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusAccepted)
}

func (s *Service) ListEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	res, err := s.List(lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}

func (s *Service) ListByRefEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	res, err := s.ListByRef(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}

func (s *Service) AddSpellRefEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("spell_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AddSpellRef(rid, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusCreated)
}

func (s *Service) DeleteSpellRefEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("spell_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.DeleteSpellRef(rid, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusNoContent)
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
func (s *Service) GetVO(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.GetLang(id, "UK")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}
