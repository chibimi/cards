package card

import (
	"github.com/pkg/errors"
)

type Ability struct {
	ID           int    `json:"id,omitempty" db:"id"`
	OriginalName string `json:"original_name,omitempty" db:"original_name"`
	Name         string `json:"name,omitempty" db:"name"`
	Magical      bool   `json:"magical,omitempty" db:"magical"`
	Description  string `json:"description,omitempty" db:"description"`
}

func (s *Service) GetAbility(id int) (*Ability, error) {
	stmt, err := s.db.Preparex("SELECT * FROM abilities WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Ability{}

	if err := stmt.Get(res, id); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (s *Service) DeleteAbility(id int) error {
	stmt, err := s.db.Preparex("DELETE FROM abilities WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}

func (s *Service) SaveAbility(ability *Ability) (int, error) {
	stmt, err := s.db.PrepareNamed(`INSERT INTO abilities VALUES(
		:id, :original_name, :name, :magical, :description
	)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res, err := stmt.Exec(ability)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	return int(id), errors.Wrap(err, "last index")
}

func (s *Service) UpdateAbility(ability *Ability) error {
	stmt, err := s.db.PrepareNamed(`UPDATE abilities SET 
	name = :name, original_name = :original_name, magical = :magical, description = :description WHERE id = :id`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(ability)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) ListAbilities() ([]Ability, error) {
	stmt, err := s.db.Preparex("SELECT * FROM abilities")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res := []Ability{}
	for rows.Next() {
		r := Ability{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		res = append(res, r)
	}
	return res, nil
}

func (s *Service) AddCardAbility(cardID, abilityID int) error {
	stmt, err := s.db.Prepare(`INSERT INTO card_ability VALUES(?, ?)`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(cardID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
func (s *Service) DeleteCardAbility(cardID, abilityID int) error {
	stmt, err := s.db.Prepare(`DELETE FROM card_ability WHERE card_id = ? AND ability_id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(cardID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) AddWeaponAbility(weaponID, abilityID int) error {
	stmt, err := s.db.Prepare(`INSERT INTO weapon_ability VALUES(?, ?)`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(weaponID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
func (s *Service) DeleteWeaponAbility(weaponID, abilityID int) error {
	stmt, err := s.db.Prepare(`DELETE FROM weapon_ability WHERE weapon_id = ? AND ability_id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(weaponID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) AddModelAbility(modelID, abilityID int) error {
	stmt, err := s.db.Prepare(`INSERT INTO model_ability VALUES(?, ?)`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(modelID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
func (s *Service) DeleteModelAbility(modelID, abilityID int) error {
	stmt, err := s.db.Prepare(`DELETE FROM model_ability WHERE model_id = ? AND ability_id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(modelID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) ListMagicAbilities(modelID int) ([]Ability, error) {
	stmt, err := s.db.Preparex("SELECT id, type, name, description FROM abilities AS a LEFT JOIN model_ability AS ma ON a.id = ma.ability_id WHERE ma.model_id = ? AND a.type = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(modelID, 3)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res := []Ability{}
	for rows.Next() {
		r := Ability{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		res = append(res, r)
	}
	return res, nil
}
func (s *Service) ListWeaponAbilities(weaponID int) ([]Ability, error) {
	stmt, err := s.db.Preparex("SELECT id, type, name, description FROM abilities AS a LEFT JOIN weapon_ability AS wa ON a.id = wa.ability_id WHERE wa.weapon_id = ? AND a.type = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(weaponID, 4)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res := []Ability{}
	for rows.Next() {
		r := Ability{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		res = append(res, r)
	}
	return res, nil
}
