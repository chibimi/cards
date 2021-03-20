package reference

type Reference struct {
	ID         int    `json:"id,omitempty" db:"id"`
	PPID       int    `json:"ppid,omitempty" db:"ppid"`
	FactionID  int    `json:"faction_id,omitempty" db:"faction_id"`
	CategoryID int    `json:"category_id,omitempty" db:"category_id"`
	Title      string `json:"title,omitempty" db:"title"`
	MainCardID int    `json:"main_card_id,omitempty,string" db:"main_card_id"`
	Models     string `json:"models_cnt,omitempty" db:"models_cnt"`
	ModelsMax  string `json:"models_max,omitempty" db:"models_max"`
	Cost       string `json:"cost,omitempty" db:"cost"`
	CostMax    string `json:"cost_max,omitempty" db:"cost_max"`
	FA         string `json:"fa,omitempty" db:"fa"`
	Name       string `json:"name,omitempty" db:"name"`
	Properties string `json:"properties,omitempty" db:"properties"`
	Status     string `json:"status,omitempty" db:"status"`
	MercFor    []int  `json:"mercenary_for,omitempty"`
	MinFor     []int  `json:"minion_for,omitempty"`
	Special    string `json:"special,omitempty" db:"special"`
	LinkedTo   *int   `json:"linked_to,omitempty" db:"linked_to"`
	ReviewGood int    `json:"review_good" db:"good"`
	ReviewBad  int    `json:"review_bad" db:"bad"`
}

type Rating struct {
	Good int `json:"good" db:"good"`
	Bad  int `json:"bad" db:"bad"`
}
