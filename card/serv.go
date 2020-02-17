package card

import (
	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/feat"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/reference"
	"github.com/chibimi/cards/card/spell"
	"github.com/chibimi/cards/card/weapon"
	"github.com/jmoiron/sqlx"
)

type SConfig struct {
}

type SService struct {
	Ref     *reference.Service
	Feat    *feat.Service
	Model   *model.Service
	Weapon  *weapon.Service
	Spell   *spell.Service
	Ability *ability.API
}

func NewSService(db *sqlx.DB) *SService {
	return &SService{
		Ref:     reference.NewService(reference.NewRepository(db)),
		Feat:    feat.NewService(feat.NewRepository(db)),
		Model:   model.NewService(model.NewRepository(db)),
		Weapon:  weapon.NewService(weapon.NewRepository(db)),
		Spell:   spell.NewService(spell.NewRepository(db)),
		Ability: ability.NewAPI(ability.NewRepository(db)),
	}
}
