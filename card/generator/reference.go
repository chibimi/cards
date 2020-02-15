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

	r.RefAbilities, err = s.src.Ability.ListByRef(r.Ref.ID, lang)
	if err != nil {
		return r, fmt.Errorf(`fetching abilities for ref %d: %w`, r.Ref.ID, err)
	}

	r.Models, err = s.src.Model.List(id, lang)
	if err != nil {
		return r, fmt.Errorf(`fetching models for ref%d: %w`, r.Ref.ID, err)
	}

	r.ModelsAbilities = make(map[int][]ability.Ability)
	r.ModelsWeapons = make(map[int][]weapon.Weapon)
	r.WeaponsAbilities = make(map[int][]ability.Ability)

	for _, model := range r.Models {
		abilities, err := s.src.Ability.ListByModel(model.ID, lang)
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
			abilities, err := s.src.Ability.ListByWeapon(weapon.ID, lang)
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

func (r Reference) HasFeat() bool {
	switch Category(r.Ref.CategoryID) {
	case CategoryWarcaster, CategoryWarlock, CategoryInfernalMaster:
		return true
	default:
		return false
	}
}

func (r Reference) HasSpells() bool {
	switch Category(r.Ref.CategoryID) {
	case CategoryWarcaster, CategoryWarlock, CategoryInfernalMaster:
		return true
	default:
		return false
	}
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

func (f Faction) String() string {
	switch f {
	case FactionCygnar:
		return "cygnar"
	case FactionProtectorateOfMenoth:
		return "protectorate of menoth"
	case FactionKhador:
		return "khador"
	case FactionCryx:
		return "cryx"
	case FactionRetributionOfScyrah:
		return "retribution of scyrah"
	case FactionConvergeanceOfCyriss:
		return "convergeance of cyriss"
	case FactionCrucibleGuard:
		return "crucible guard"
	case FactionMercenaries:
		return "mercenaries"
	case FactionTrollbloods:
		return "trollbloods"
	case FactionCircleOrboros:
		return "circle orboros"
	case FactionLegionOfEverblight:
		return "legion of everblight"
	case FactionSkorne:
		return "skorne"
	case FactionGrymkin:
		return "grymkin"
	case FactionInfernals:
		return "infernals"
	case FactionMinions:
		return "minions"
	default:
		return "invalid"
	}
}
