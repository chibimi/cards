package weapon

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(wp *Weapon, lang string) (int, error) {
	return s.repo.Create(wp, lang)
}

func (s *Service) List(model int, lang string) ([]Weapon, error) {
	return s.repo.List(model, lang)
}

func (s *Service) Get(id int, lang string) (*Weapon, error) {
	return s.repo.Get(id, lang)
}

func (s *Service) Save(wp *Weapon, lang string) error {
	return s.repo.Save(wp, lang)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetLang(id int, lang string) (*Weapon, error) {
	return s.repo.GetLang(id, lang)
}
