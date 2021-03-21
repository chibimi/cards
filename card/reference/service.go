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

func (s *Service) List(faction, category int, lang, status string) ([]Reference, error) {
	return s.repo.List(faction, category, lang, status)
}

func (s *Service) ListIDs(faction int, lang, status string) ([]int, error) {
	return s.repo.ListIDs(faction, lang, status)
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

func (s *Service) ListRefAttachments(lang string, linked_to int) ([]Reference, error) {
	return s.repo.ListRefAttachments(lang, linked_to)
}

func (s *Service) GetRating(refID int, lang string) (*Rating, error) {
	return s.repo.GetRating(refID, lang)
}
