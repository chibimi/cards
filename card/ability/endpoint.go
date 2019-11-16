package ability

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card/utils"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ab := &Ability{}
	if err := json.NewDecoder(r.Body).Decode(ab); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	res, err := s.Create(ab, lang)
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
	ab := &Ability{}
	if err := json.NewDecoder(r.Body).Decode(ab); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if id != ab.ID {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	err = s.Save(ab, lang)
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
func (s *Service) AddAbilityRefEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	typ, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AddAbilityRef(rid, id, typ)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusCreated)
}
func (s *Service) DeleteAbilityRefEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.DeleteAbilityRef(rid, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusNoContent)
}

func (s *Service) ListByModelEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	res, err := s.ListByModel(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}
func (s *Service) AddAbilityModelEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	typ, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AddAbilityModel(rid, id, typ)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusCreated)
}
func (s *Service) DeleteAbilityModelEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.DeleteAbilityModel(rid, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusNoContent)
}

func (s *Service) ListByWeaponEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	res, err := s.ListByWeapon(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}
func (s *Service) AddAbilityWeaponEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	typ, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AddAbilityWeapon(rid, id, typ)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusCreated)
}
func (s *Service) DeleteAbilityWeaponEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("ability_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.DeleteAbilityWeapon(rid, id)
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
	res, err := s.GetLang(id, "US")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}
