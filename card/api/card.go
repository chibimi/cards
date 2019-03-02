package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) CreateCard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	card := &card.Card{}
	if err := json.NewDecoder(r.Body).Decode(card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := s.src.SaveCard(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusCreated)
}

func (s *Service) UpdateCard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "0" {
		s.CreateCard(w, r, p)
		return
	}
	card := &card.Card{}
	if err := json.NewDecoder(r.Body).Decode(card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := s.src.UpdateCard(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}

func (s *Service) GetCard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := s.src.GetCard(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) DeleteCard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.src.DeleteCard(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusNoContent)
}

func (s *Service) ListCards(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	res, err := s.src.ListCards(fid, cid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) GetRelatedCards(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := s.src.GetRelatedCards(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) GetCardAbilities(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.GetCardAbilities(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) GetCardFeat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.GetCardFeat(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}
func (s *Service) GetCardSpells(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.GetCardSpells(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) GetCardModels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.src.GetCardModels(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, res, http.StatusOK)
}

func (s *Service) AddCardAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	err = s.src.AddCardAbility(cardID, abilityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}

func (s *Service) DeleteCardAbility(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	err = s.src.DeleteCardAbility(cardID, abilityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, nil, http.StatusNoContent)
}

func (s *Service) AddCardSpell(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cardID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	spellID, err := strconv.Atoi(p.ByName("spell_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.src.AddCardSpell(cardID, spellID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, nil, http.StatusOK)
}

func (s *Service) DeleteCardSpell(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cardID, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	spellID, err := strconv.Atoi(p.ByName("spell_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.src.DeleteCardSpell(cardID, spellID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, nil, http.StatusNoContent)
}
