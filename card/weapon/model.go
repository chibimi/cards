package weapon

type Weapon struct {
	ID         int      `json:"id,omitempty" db:"id"`
	ModelID    int      `json:"model_id,omitempty" db:"model_id"`
	Type       int      `json:"type,omitempty,string" db:"type"`
	Name       string   `json:"name,omitempty" db:"name"`
	RNG        string   `json:"rng,omitempty" db:"rng"`
	POW        string   `json:"pow,omitempty" db:"pow"`
	ROF        string   `json:"rof,omitempty" db:"rof"`
	AOE        string   `json:"aoe,omitempty" db:"aoe"`
	LOC        string   `json:"loc,omitempty" db:"loc"`
	CNT        string   `json:"cnt,omitempty" db:"cnt"`
	Advantages []string `json:"advantages" db:"-"`
}
