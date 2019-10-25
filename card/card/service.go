package card

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Save(card *Card, lang string) error {
	return s.repo.Save(card, lang)
}

func (s *Service) Get(id int, lang string) (*Card, error) {
	return s.repo.Get(id, lang)
}
