package spell

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(sp *Spell, lang string) (int, error) {
	return s.repo.Create(sp, lang)
}

func (s *Service) List(lang string) ([]Spell, error) {
	return s.repo.List(lang)
}

func (s *Service) ListByRef(ref int, lang string) ([]Spell, error) {
	return s.repo.ListByRef(ref, lang)
}

func (s *Service) Save(sp *Spell, lang string) error {
	return s.repo.Save(sp, lang)
}
func (s *Service) AddSpellRef(ref, spell int) error {
	return s.repo.AddSpellRef(ref, spell)
}
func (s *Service) DeleteSpellRef(ref, spell int) error {
	return s.repo.DeleteSpellRef(ref, spell)
}
