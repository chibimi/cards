package card

type Faction struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Service) ListFactions() []Faction {
	return []Faction{
		{ID: 1, Name: "Cygnar"},
		{ID: 2, Name: "Protectorate of Menoth"},
		{ID: 3, Name: "Khador"},
		{ID: 4, Name: "Cryx"},
		{ID: 5, Name: "Retribution of Scyrah"},
		{ID: 6, Name: "Convergeance of Cyriss"},
		{ID: 7, Name: "Crucible Guard"},
		{ID: 8, Name: "Mercenaries"},
		{ID: 9, Name: "Trollbloods"},
		{ID: 10, Name: "Circle Orboros"},
		{ID: 11, Name: "Legion of Everblight"},
		{ID: 12, Name: "Skorne"},
		{ID: 13, Name: "Grymkin"},
		{ID: 14, Name: "Infernals"},
		{ID: 15, Name: "Minions"},
	}
}
