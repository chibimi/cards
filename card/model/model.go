package model

type Model struct {
	ID           int      `json:"id,omitempty" db:"id"`
	RefID        int      `json:"ref_id,omitempty" db:"ref_id"`
	Title        string   `json:"title,omitempty" db:"title"`
	Name         string   `json:"name,omitempty" db:"name"`
	SPD          string   `json:"spd,omitempty" db:"spd"`
	STR          string   `json:"str,omitempty" db:"str"`
	MAT          string   `json:"mat,omitempty" db:"mat"`
	RAT          string   `json:"rat,omitempty" db:"rat"`
	DEF          string   `json:"def,omitempty" db:"def"`
	ARM          string   `json:"arm,omitempty" db:"arm"`
	CMD          string   `json:"cmd,omitempty" db:"cmd"`
	BaseSize     string   `json:"base_size,omitempty" db:"base_size"`
	MagicAbility string   `json:"magic_ability,omitempty" db:"magic_ability"`
	Resource     string   `json:"resource,omitempty" db:"resource"`
	Threshold    string   `json:"threshold,omitempty" db:"threshold"`
	Damage       string   `json:"damage,omitempty" db:"damage"`
	Advantages   []string `json:"advantages" db:"-"`
}
