package review

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Save(review *Review) error {
	return s.repo.Save(review)
}
