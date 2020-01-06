package ability

type Ability struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        int    `json:"type"`
	Star        int    `json:"star"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
