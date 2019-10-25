package feat

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Save(f *Feat, lang string) error {
	return s.repo.Save(f, lang)
}

func (s *Service) Get(id int, lang string) (*Feat, error) {
	return s.repo.Get(id, lang)
}
