package model

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(m *Model, lang string) (int, error) {
	return s.repo.Create(m, lang)
}

func (s *Service) List(ref int, lang string) ([]Model, error) {
	return s.repo.List(ref, lang)
}

func (s *Service) Get(id int, lang string) (*Model, error) {
	return s.repo.Get(id, lang)
}

func (s *Service) Save(m *Model, lang string) error {
	return s.repo.Save(m, lang)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
