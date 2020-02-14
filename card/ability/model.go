package ability

import "fmt"

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

func (a Ability) Text() string {
	return fmt.Sprintf("**%s** – %s", a.Title, a.Description)
}
