package card

import "github.com/chibimi/cards/card/reference"

type Card struct {
	reference.Reference
	MainCardID int    `json:"main_card_id,omitempty,string" db:"main_card_id"`
	Models     string `json:"models_cnt,omitempty" db:"models_cnt"`
	ModelsMax  string `json:"models_max,omitempty" db:"models_max"`
	Cost       string `json:"cost,omitempty" db:"cost"`
	CostMax    string `json:"cost_max,omitempty" db:"cost_max"`
	FA         string `json:"fa,omitempty" db:"fa"`
	Text
}

type Text struct {
	Name       string `json:"name,omitempty" db:"name"`
	Properties string `json:"properties,omitempty" db:"properties"`
}
