package card

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Service) ListCategories() []Category {
	return []Category{
		{ID: 1, Name: "Warcaster"},
		{ID: 2, Name: "Warlock"},
		{ID: 3, Name: "Warjack"},
		{ID: 4, Name: "Warbeast"},
		{ID: 5, Name: "Unit"},
		{ID: 6, Name: "Solo"},
		{ID: 7, Name: "Attachments"},
		{ID: 8, Name: "Battle Engine"},
		{ID: 9, Name: "Structure"},
	}
}
