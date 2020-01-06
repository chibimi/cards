package reference

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(ref *Reference) (int, error) {
	return s.repo.Create(ref)
}

func (s *Service) List(faction, category int, lang string) ([]Reference, error) {
	return s.repo.List(faction, category, lang)
}

func (s *Service) Get(id int, lang string) (*Reference, error) {
	return s.repo.Get(id, lang)
}

func (s *Service) Save(ref *Reference, lang string) error {
	return s.repo.Save(ref, lang)
}

func (s *Service) ListByStatus(lang, status string) ([]Reference, error) {
	return s.repo.ListByStatus(lang, status)
}
