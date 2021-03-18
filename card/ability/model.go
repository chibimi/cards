package ability

type Ability struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Header      *int   `json:"header,omitempty"`
	Star        *int   `json:"star,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Relation struct {
	AbilityID int  `json:"ability_id"`
	RelatedID int  `json:"related_id"`
	Header    *int `json:"header,omitempty"`
	Star      *int `json:"star,omitempty"`
}

func (a Ability) GetStarText() string {
	if a.Star == nil {
		return ""
	}
	switch *a.Star {
	case 1:
		return " (\u2605Attaque)"
	case 2:
		return " (\u2605Action)"
	case 3:
		return " (\u2605Action ou \u2605Attaque)"
	default:
		return ""
	}
}
