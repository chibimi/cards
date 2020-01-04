package weapon

import "sort"

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
	weapons, err := s.repo.List(model, lang)
	if err != nil {
		return nil, err
	}
	sort.Slice(weapons, func(i, j int) bool {
		if weapons[i].Type == 2 {
			return true
		}
		return weapons[i].Type < weapons[j].Type
	})
	return weapons, nil
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
