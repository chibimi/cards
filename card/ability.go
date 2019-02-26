package card

import (
	"github.com/pkg/errors"
)

type Ability struct {
	ID          int    `json:"id,omitempty" db:"id"`
	Type        int    `json:"type,omitempty" db:"type"`
	Name        string `json:"name,omitempty" db:"name"`
	Description string `json:"description,omitempty" db:"description"`
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
		:id, :type, :name, :description
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
	type = :type, name = :name, description = :description WHERE id = :id`)
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

func (s *Service) ListCardAbilities(cardID int) ([]Ability, error) {
	stmt, err := s.db.Preparex("SELECT id, type, name, description FROM abilities AS a LEFT JOIN card_ability AS ca ON a.id = ca.ability_id WHERE ca.card_id = ? AND a.type = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(cardID, 1)
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
func (s *Service) ListModelAbilities(modelID int) ([]Ability, error) {
	stmt, err := s.db.Preparex("SELECT id, type, name, description FROM abilities AS a LEFT JOIN model_ability AS ma ON a.id = ma.ability_id WHERE ma.model_id = ? AND a.type = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(modelID, 2)
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
