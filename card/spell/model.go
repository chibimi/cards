package spell

type Spell struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Cost        string `json:"cost,omitempty"`
	RNG         string `json:"rng,omitempty"`
	AOE         string `json:"aoe,omitempty"`
	POW         string `json:"pow,omitempty"`
	DUR         string `json:"dur,omitempty"`
	OFF         string `json:"off,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
