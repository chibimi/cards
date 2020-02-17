package ability

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card/utils"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	repo *Repository
}

func NewAPI(repo *Repository) *API {
	return &API{repo}
}

/****************************
 ABILITIES
*****************************/

func (s *API) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.Create(w, r, p)
		return
	}
	ability := &Ability{}
	if err := json.NewDecoder(r.Body).Decode(ability); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if id != ability.ID {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	err = s.repo.Save(ability, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusAccepted)
}

func (s *API) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ability := &Ability{}
	if err := json.NewDecoder(r.Body).Decode(ability); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	id, err := s.repo.Create(ability, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, id, http.StatusCreated)
}

func (s *API) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	ability, err := s.repo.Get(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, ability, http.StatusOK)
}

func (s *API) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "require lang", http.StatusBadRequest)
		return
	}
	abilities, err := s.repo.List(lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, abilities, http.StatusOK)
}

/****************************
 ABILITIES BY MODEL
*****************************/

func (s *API) AddAbilityModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	relation := &Relation{}
	if err := json.NewDecoder(r.Body).Decode(relation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.repo.AddAbilityModel(relation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusCreated)
}

func (s *API) ListByModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	res, err := s.repo.ListByModel(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}

func (s *API) DeleteAbilityModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	err = s.repo.DeleteAbilityModel(rid, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusNoContent)
}

/****************************
 ABILITIES BY WEAPON
*****************************/

func (s *API) ListByWeapon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	res, err := s.repo.ListByWeapon(id, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, res, http.StatusOK)
}

func (s *API) AddAbilityWeapon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	relation := &Relation{}
	if err := json.NewDecoder(r.Body).Decode(relation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.repo.AddAbilityWeapon(relation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusCreated)
}

func (s *API) DeleteAbilityWeapon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	err = s.repo.DeleteAbilityWeapon(rid, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusNoContent)
}
