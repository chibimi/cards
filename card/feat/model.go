package feat

type Feat struct {
	RefID       int    `json:"ref_id,omitempty" db:"ref_id"`
	Name        string `json:"name,omitempty" db:"name"`
	Description string `json:"description,omitempty" db:"description"`
	Fluff       string `json:"fluff,omitempty" db:"fluff"`
}
