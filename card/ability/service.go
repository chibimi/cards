package ability

import "sort"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(sp *Ability, lang string) (int, error) {
	return s.repo.Create(sp, lang)
}

func (s *Service) List(lang string) ([]Ability, error) {
	return s.repo.List(lang)
}

func (s *Service) Save(sp *Ability, lang string) error {
	return s.repo.Save(sp, lang)
}

func (s *Service) ListByRef(ref int, lang string) ([]Ability, error) {
	return s.repo.ListByRef(ref, lang)
}
func (s *Service) AddAbilityRef(ref, ability, typ int) error {
	return s.repo.AddAbilityRef(ref, ability, typ)
}
func (s *Service) DeleteAbilityRef(ref, ability int) error {
	return s.repo.DeleteAbilityRef(ref, ability)
}

func (s *Service) ListByModel(model int, lang string) ([]Ability, error) {
	abilities, err := s.repo.ListByModel(model, lang)
	if err != nil {
		return nil, err
	}
	sort.Slice(abilities, func(i, j int) bool {
		if abilities[i].Type < abilities[j].Type {
			return true
		}
		if abilities[i].Type > abilities[j].Type {
			return false
		}
		return abilities[i].Title < abilities[j].Title
	})
	return abilities, err
}
func (s *Service) AddAbilityModel(model, ability, typ int) error {
	return s.repo.AddAbilityModel(model, ability, typ)
}
func (s *Service) DeleteAbilityModel(model, ability int) error {
	return s.repo.DeleteAbilityModel(model, ability)
}

func (s *Service) ListByWeapon(weapon int, lang string) ([]Ability, error) {
	abilities, err := s.repo.ListByWeapon(weapon, lang)
	if err != nil {
		return nil, err
	}
	sort.Slice(abilities, func(i, j int) bool {
		if abilities[i].Type < abilities[j].Type {
			return true
		}
		if abilities[i].Type > abilities[j].Type {
			return false
		}
		return abilities[i].Title < abilities[j].Title
	})
	return abilities, err
}
func (s *Service) AddAbilityWeapon(weapon, ability, typ int) error {
	return s.repo.AddAbilityWeapon(weapon, ability, typ)
}
func (s *Service) DeleteAbilityWeapon(weapon, ability int) error {
	return s.repo.DeleteAbilityWeapon(weapon, ability)
}

func (s *Service) Get(id int, lang string) (*Ability, error) {
	return s.repo.Get(id, lang)
}
