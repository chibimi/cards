package card

import (
	"strings"

	"github.com/pkg/errors"
)

type Weapon struct {
	ID           int      `json:"id,omitempty" db:"id"`
	ModelID      int      `json:"model_id,omitempty" db:"model_id"`
	Type         int      `json:"type,omitempty" db:"type"`
	Name         string   `json:"name,omitempty" db:"name"`
	RNG          string   `json:"rng,omitempty" db:"rng"`
	POW          string   `json:"pow,omitempty" db:"pow"`
	ROF          string   `json:"rof,omitempty" db:"rof"`
	AOE          string   `json:"aoe,omitempty" db:"aoe"`
	LOC          string   `json:"loc,omitempty" db:"loc"`
	CNT          string   `json:"cnt,omitempty" db:"cnt"`
	Advantages   []string `json:"advantages" db:"-"`
	AdvantagesDB string   `json:"-" db:"advantages"`
}

func (s *Service) GetWeapon(id int) (*Weapon, error) {
	stmt, err := s.db.Preparex("SELECT * FROM weapons WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Weapon{}

	if err := stmt.Get(res, id); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res.Advantages = strings.Split(res.AdvantagesDB, ",")

	return res, nil
}

func (s *Service) DeleteWeapon(id int) error {
	stmt, err := s.db.Preparex("DELETE FROM weapons WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}

func (s *Service) SaveWeapon(weapon *Weapon) (int, error) {
	stmt, err := s.db.PrepareNamed(`INSERT INTO weapons VALUES(
		:id, :model_id, :type, :name, :rng, :pow, :rof, :aoe, :loc, :cnt, :advantages
	)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	weapon.AdvantagesDB = strings.Join(weapon.Advantages, ",")

	res, err := stmt.Exec(weapon)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	return int(id), errors.Wrap(err, "last index")
}

func (s *Service) UpdateWeapon(weapon *Weapon) error {
	stmt, err := s.db.PrepareNamed(`UPDATE abilities SET 
	model_id = :model_id, type = :type, name = :name, 
	rng = :rng, pow = :pow, rof = :rof, aoe = :aoe, loc = :loc, cnt = :cnt, advantages = :advantages WHERE id = :id`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	weapon.AdvantagesDB = strings.Join(weapon.Advantages, ",")

	_, err = stmt.Exec(weapon)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) ListWeapons(modelID int) ([]Weapon, error) {
	stmt, err := s.db.Preparex("SELECT * FROM weapons WHERE model_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(modelID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res := []Weapon{}
	for rows.Next() {
		r := Weapon{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		r.Advantages = strings.Split(r.AdvantagesDB, ",")

		res = append(res, r)
	}
	return res, nil
}
