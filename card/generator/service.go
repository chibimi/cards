package generator

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card"
	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/feat"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/reference"
	"github.com/chibimi/cards/card/spell"
	"github.com/chibimi/cards/card/weapon"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/common/log"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

type Service struct {
	src    *card.SService
	assets string
}

func NewService(cards *card.SService, assets string) *Service {
	return &Service{
		src:    cards,
		assets: assets,
	}
}

func (s *Service) GenerateEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var refs []int
	for _, v := range strings.Split(r.FormValue("cards"), ",") {
		ref, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		refs = append(refs, ref)
	}

	res, err := s.Generate(refs, r.FormValue("lang"))
	if err != nil {
		log15.Error("generating cards", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/pdf")
	_, err = io.Copy(w, res)
	if err != nil {
		log15.Error("sending cards", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) Generate(refs []int, lang string) (io.Reader, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) DisplayEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var ids []int
	for _, v := range strings.Split(r.FormValue("cards"), ",") {
		id, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ids = append(ids, id)
	}
	var lang = r.FormValue("lang")

	// References generally have 1 cards. Casters, Colossals, Character
	// Units & other edge cases are the notable exceptions, but cards
	// always have 2 faces anyway, and we consider each face as an
	// independent card for the purpose of generating them as it's way
	// simpler to handle.
	var cards = make([]Card, 0, len(ids)*2)
	for _, id := range ids {
		c, err := s.Build(id, lang)
		if err != nil {
			http.Error(w, fmt.Sprintf("building ref %d: %s", id, err), http.StatusInternalServerError)
			return
		}
		cards = append(cards, c...)
	}

	t, err := template.New("cards.html").Funcs(template.FuncMap{
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
	}).ParseFiles(path.Join(s.assets, "templates/cards.html"))
	if err != nil {
		log.Error("parsing card template", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, cards)
	if err != nil {
		log.Error("generating output", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) Build(id int, lang string) (cards []Card, err error) {
	ref, err := s.src.Ref.Get(id, lang)
	if err != nil {
		return nil, fmt.Errorf(`fetching reference: %w`, err)
	}

	profile := ProfileCard{
		Ref:           ref,
		ModelsWeapons: make(map[int][]weapon.Weapon),
	}

	rules := RulesCard{
		Ref:              ref,
		ModelsAbilities:  make(map[int][]ability.Ability),
		ModelsWeapons:    make(map[int][]weapon.Weapon),
		WeaponsAbilities: make(map[int][]ability.Ability),
	}

	rules.RefAbilities, err = s.src.Ability.ListByRef(ref.ID, lang)
	if err != nil {
		return nil, fmt.Errorf(`fetching abilities for ref %d: %w`, ref.ID, err)
	}

	models, err := s.src.Model.List(id, lang)
	if err != nil {
		return nil, fmt.Errorf(`fetching models for ref%d: %w`, ref.ID, err)
	}
	profile.Models = models
	rules.Models = models

	for _, model := range models {
		abilities, err := s.src.Ability.ListByModel(model.ID, lang)
		if err != nil {
			return nil, fmt.Errorf(`fetching abilities for model %d: %w`, model.ID, err)
		}
		rules.ModelsAbilities[model.ID] = abilities

		weapons, err := s.src.Weapon.List(model.ID, lang)
		if err != nil {
			return nil, fmt.Errorf(`fetching weapons for model %d: %w`, model.ID, err)
		}
		profile.ModelsWeapons[model.ID] = weapons
		rules.ModelsWeapons[model.ID] = weapons

		for _, weapon := range weapons {
			abilities, err := s.src.Ability.ListByWeapon(weapon.ID, lang)
			if err != nil {
				return nil, fmt.Errorf(`fetching abilities for weapon %d: %w`, weapon.ID, err)
			}
			rules.WeaponsAbilities[model.ID] = abilities
		}
	}

	cards = append(cards, profile, rules)

	if ref.HasSpells() {
		spells, err := s.src.Spell.ListByRef(ref.ID, lang)
		if err != nil {
			return nil, fmt.Errorf(`fetching spells for ref %d: %w`, ref.ID, err)
		}

		cards = append(cards, SpellsCard{
			Ref:    ref,
			Spells: spells,
		})
	}

	if ref.HasFeat() {
		feat, err := s.src.Feat.Get(ref.ID, lang)
		if err != nil {
			return nil, fmt.Errorf(`fetching feat: %w`, err)
		}

		cards = append(cards, FeatCard{
			Ref:  ref,
			Feat: feat,
		})
	}

	return cards, nil
}

type Card interface {
	Type() string
	Background() string
}

type ProfileCard struct {
	Ref           *reference.Reference
	Models        []model.Model
	ModelsWeapons map[int][]weapon.Weapon
}

func (c ProfileCard) Background() string {
	return strconv.Itoa(c.Ref.FactionID)
}

func (ProfileCard) Type() string {
	return "profile"
}

type RulesCard struct {
	Ref              *reference.Reference
	RefAbilities     []ability.Ability
	Models           []model.Model
	ModelsAbilities  map[int][]ability.Ability
	ModelsWeapons    map[int][]weapon.Weapon
	WeaponsAbilities map[int][]ability.Ability
}

func (RulesCard) Type() string {
	return "rules"
}

func (c RulesCard) Background() string {
	return strconv.Itoa(c.Ref.FactionID)
}

type SpellsCard struct {
	Ref    *reference.Reference
	Spells []spell.Spell
}

func (c SpellsCard) Background() string {
	return strconv.Itoa(c.Ref.FactionID)
}

func (SpellsCard) Type() string {
	return "spells"
}

type FeatCard struct {
	Ref  *reference.Reference
	Feat *feat.Feat
}

func (c FeatCard) Background() string {
	return strconv.Itoa(c.Ref.FactionID)
}

func (FeatCard) Type() string {
	return "feat"
}
