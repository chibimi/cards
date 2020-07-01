package generator

import (
	"fmt"

	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/feat"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/reference"
	"github.com/chibimi/cards/card/spell"
	"github.com/chibimi/cards/card/weapon"
)

type Reference struct {
	Lang             string
	Ref              *reference.Reference
	RefAbilities     []ability.Ability
	Models           []model.Model
	ModelsAbilities  map[int][]ability.Ability
	ModelsWeapons    map[int][]weapon.Weapon
	WeaponsAbilities map[int][]ability.Ability
	Spells           []spell.Spell
	Feat             *feat.Feat
}

func (s *Service) Get(id int, lang string) (r Reference, err error) {
	r.Lang = lang

	r.Ref, err = s.src.Ref.Get(id, lang)
	if err != nil {
		return r, fmt.Errorf(`fetching reference: %w`, err)
	}

	r.Models, err = s.src.Model.List(id, lang)
	if err != nil {
		return r, fmt.Errorf(`fetching models for ref%d: %w`, r.Ref.ID, err)
	}

	r.ModelsAbilities = make(map[int][]ability.Ability)
	r.ModelsWeapons = make(map[int][]weapon.Weapon)
	r.WeaponsAbilities = make(map[int][]ability.Ability)

	for _, model := range r.Models {
		abilities, err := s.ability.ListByModel(model.ID, lang)
		if err != nil {
			return r, fmt.Errorf(`fetching abilities for model %d: %w`, model.ID, err)
		}
		r.ModelsAbilities[model.ID] = abilities

		weapons, err := s.src.Weapon.List(model.ID, lang)
		if err != nil {
			return r, fmt.Errorf(`fetching weapons for model %d: %w`, model.ID, err)
		}
		r.ModelsWeapons[model.ID] = weapons

		for _, weapon := range weapons {
			abilities, err := s.ability.ListByWeapon(weapon.ID, lang)
			if err != nil {
				return r, fmt.Errorf(`fetching abilities for weapon %d: %w`, weapon.ID, err)
			}
			r.WeaponsAbilities[weapon.ID] = abilities
		}
	}

	r.Spells, err = s.src.Spell.ListByRef(r.Ref.ID, lang)
	if err != nil {
		return r, fmt.Errorf(`fetching spells for ref %d: %w`, r.Ref.ID, err)
	}

	r.Feat, err = s.src.Feat.Get(r.Ref.ID, lang)
	if err != nil {
		return r, fmt.Errorf(`fetching feat: %w`, err)
	}

	return r, nil
}

type Category int

const (
	CategoryInvalid Category = iota
	CategoryWarcaster
	CategoryWarlock
	CategoryWarjack
	CategoryWarbeast
	CategoryUnit
	CategorySolo
	CategoryAttachments
	CategoryBattleEngine
	CategoryStructure
	CategoryInfernalMaster
	CategoryHorror
)

type Faction int

const (
	FactionInvalid = iota
	FactionCygnar
	FactionProtectorateOfMenoth
	FactionKhador
	FactionCryx
	FactionRetributionOfScyrah
	FactionConvergeanceOfCyriss
	FactionCrucibleGuard
	FactionMercenaries
	FactionTrollbloods
	FactionCircleOrboros
	FactionLegionOfEverblight
	FactionSkorne
	FactionGrymkin
	FactionInfernals
	FactionMinions
)

var factionsNames = map[Faction]string{
	FactionInvalid:              "invalid",
	FactionCygnar:               "cygnar",
	FactionProtectorateOfMenoth: "protectorate of menoth",
	FactionKhador:               "khador",
	FactionCryx:                 "cryx",
	FactionRetributionOfScyrah:  "retribution of scyrah",
	FactionConvergeanceOfCyriss: "convergeance of cyriss",
	FactionCrucibleGuard:        "crucible guard",
	FactionMercenaries:          "mercenaries",
	FactionTrollbloods:          "trollbloods",
	FactionCircleOrboros:        "circle orboros",
	FactionLegionOfEverblight:   "legion of everblight",
	FactionSkorne:               "skorne",
	FactionGrymkin:              "grymkin",
	FactionInfernals:            "infernals",
	FactionMinions:              "minions",
}

func (f Faction) String() string {
	return factionsNames[f]
}

var translations = map[string]string{
	"see_above_FR": "Voir plus haut.",
	"see_above_US": "See above.",
}
