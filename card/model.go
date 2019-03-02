package card

import (
	"strings"

	"github.com/pkg/errors"
)

type Model struct {
	ID           int      `json:"id,omitempty" db:"id"`
	CardID       int      `json:"card_id,omitempty" db:"card_id"`
	Name         string   `json:"name,omitempty" db:"name"`
	SPD          string   `json:"spd,omitempty" db:"spd"`
	STR          string   `json:"str,omitempty" db:"str"`
	MAT          string   `json:"mat,omitempty" db:"mat"`
	RAT          string   `json:"rat,omitempty" db:"rat"`
	DEF          string   `json:"def,omitempty" db:"def"`
	ARM          string   `json:"arm,omitempty" db:"arm"`
	CMD          string   `json:"cmd,omitempty" db:"cmd"`
	BaseSize     string   `json:"base_size,omitempty" db:"base_size"`
	MagicAbility string   `json:"magic_ability,omitempty" db:"magic_ability"`
	Resource     string   `json:"resource,omitempty" db:"resource"`
	Threshold    string   `json:"threshold,omitempty" db:"threshold"`
	Damage       string   `json:"damage,omitempty" db:"damage"`
	Advantages   []string `json:"advantages" db:"-"`
	AdvantagesDB string   `json:"-" db:"advantages"`
}

func (s *Service) GetModel(id int) (*Model, error) {
	stmt, err := s.db.Preparex("SELECT * FROM models WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Model{}

	if err := stmt.Get(res, id); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res.Advantages = strings.Split(res.AdvantagesDB, ",")

	return res, nil
}

func (s *Service) GetModelAbilities(id int) ([]Ability, error) {
	stmt, err := s.db.Preparex("SELECT id, original_name, name, magical, description FROM model_ability AS l LEFT JOIN abilities AS a ON l.ability_id = a.id WHERE l.model_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(id)
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

func (s *Service) DeleteModel(id int) error {
	stmt, err := s.db.Preparex("DELETE FROM models WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}

func (s *Service) SaveModel(model *Model) (int, error) {
	stmt, err := s.db.PrepareNamed(`INSERT INTO models VALUES(
		:id, :card_id, :name, :spd, :str, :mat, :rat, :def, :arm, :cmd, :magic_ability, :damage, :resource, :threshold, :base_size, :advantages
	)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	model.AdvantagesDB = strings.Join(model.Advantages, ",")

	res, err := stmt.Exec(model)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	return int(id), errors.Wrap(err, "last index")
}

func (s *Service) UpdateModel(model *Model) error {
	stmt, err := s.db.PrepareNamed(`UPDATE models SET 
	card_id = :card_id, name = :name, 
	spd = :spd, str = :str, mat = :mat, rat = :rat, def = :def, arm = :arm, cmd = :cmd, magic_ability = :magic_ability, 
	advantages = :advantages, base_size = :base_size, resource = :resource, threshold = :threshold, damage = :damage WHERE id = :id`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	model.AdvantagesDB = strings.Join(model.Advantages, ",")

	_, err = stmt.Exec(model)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) ListModels() ([]Model, error) {
	stmt, err := s.db.Preparex("SELECT * FROM models")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res := []Model{}
	for rows.Next() {
		r := Model{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		r.Advantages = strings.Split(r.AdvantagesDB, ",")

		res = append(res, r)
	}
	return res, nil
}

func (s *Service) GetModelWeapons(id int) ([]Weapon, error) {
	stmt, err := s.db.Preparex("SELECT * FROM weapons WHERE model_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(id)
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
