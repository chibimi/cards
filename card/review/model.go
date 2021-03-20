package review

import "time"

type Review struct {
	RefID     int       `json:"ref_id,omitempty" db:"ref_id"`
	Lang      string    `json:"lang,omitempty" db:"lang"`
	IP        string    `json:"ip,omitempty" db:"ip"`
	Rating    string    `json:"rating,omitempty" db:"rating"`
	Comment   string    `json:"comment,omitempty" db:"comment"`
	Reviewer  string    `json:"reviewer,omitempty" db:"reviewer"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
