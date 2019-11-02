package ability

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
func (s *Service) AddAbilityRef(ref, spell int) error {
	return s.repo.AddAbilityRef(ref, spell)
}
func (s *Service) DeleteAbilityRef(ref, spell int) error {
	return s.repo.DeleteAbilityRef(ref, spell)
}

func (s *Service) ListByModel(model int, lang string) ([]Ability, error) {
	return s.repo.ListByModel(model, lang)
}
func (s *Service) AddAbilityModel(model, spell int) error {
	return s.repo.AddAbilityModel(model, spell)
}
func (s *Service) DeleteAbilityModel(model, spell int) error {
	return s.repo.DeleteAbilityModel(model, spell)
}

func (s *Service) ListByWeapon(weapon int, lang string) ([]Ability, error) {
	return s.repo.ListByWeapon(weapon, lang)
}
func (s *Service) AddAbilityWeapon(weapon, spell int) error {
	return s.repo.AddAbilityWeapon(weapon, spell)
}
func (s *Service) DeleteAbilityWeapon(weapon, spell int) error {
	return s.repo.DeleteAbilityWeapon(weapon, spell)
}

func (s *Service) GetLang(id int, lang string) (*Ability, error) {
	return s.repo.GetLang(id, lang)
}
