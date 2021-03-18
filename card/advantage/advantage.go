package advantage

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Advantage struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r Repository) Get(ID string) (*Advantage, error) {
	name, ok := advantages[ID]
	if !ok {
		return nil, errors.Wrap(sql.ErrNoRows, `execute query`)
	}

	return &Advantage{
		ID:   ID,
		Name: name,
	}, nil
}

var advantages = map[string]string{
	"advance_deploy":       "Advance Deployement",
	"amphibious":           "Amphibious",
	"arc_node":             "Arc Node",
	"assault":              "Assault",
	"cavalry":              "Cavalry",
	"cma":                  "CMA",
	"cra":                  "CRA",
	"construct":            "Construct",
	"eyeless_sight":        "Eyeless Sight",
	"flight":               "Flight",
	"gunfighter":           "Gunfighter",
	"incorporeal":          "Incorporeal",
	"immunity_corrosion":   "Immune corrosion ",
	"immunity_electricity": "Immune electricity",
	"immunity_fire":        "Immune fire",
	"immunity_frost":       "Immune frost",
	"jackmarshal":          "Jackmarshal",
	"officer":              "Officer",
	"parry":                "Parry",
	"pathfinder":           "Pathfinder",
	"soulless":             "Soulless",
	"stealth":              "Stealth",
	"tough":                "Tough",
	"undead":               "Undead",
	"blessed":              "Blessed",
	"chain":                "Chain",
	"type_corrosion":       "Type: Corrosion",
	"continuous_corrosion": "Cont. Corrosion",
	"crit_corrosion":       "Crit. Corrosion",
	"type_electricity":     "Type: Electricity",
	"disruption":           "Dusruption",
	"crit_disruption":      "Crit. Disruption",
	"type_fire":            "Type: Fire",
	"continuous_fire":      "Cont. Fire",
	"crit_fire":            "Crit. Fire",
	"type_frost":           "Type: Frost",
	"magical":              "Magical",
	"open_fist":            "Open Fist",
	"shield_1":             "Shield +1",
	"shield_2":             "Shield +2",
	"weapon_master":        "Weapon Master",
}
