package card

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Weapon struct {
	ID           int      `json:"id,omitempty"`
	ModelID      int      `json:"model_id,omitempty"`
	Type         int      `json:"type,omitempty,string"`
	Name         string   `json:"name,omitempty"`
	RNG          string   `json:"rng,omitempty"`
	POW          string   `json:"pow,omitempty"`
	ROF          string   `json:"rof,omitempty"`
	AOE          string   `json:"aoe,omitempty"`
	LOC          string   `json:"loc,omitempty"`
	CNT          string   `json:"cnt,omitempty"`
	Advantages   []string `json:"advantages"`
	AdvantagesDB string   `json:"-"`
}

func (s *Service) ListWeapons() ([]Weapon, error) {
	stmt, err := s.db.Prepare("SELECT data FROM weapons")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Weapon{}
	for rows.Next() {
		r := Weapon{}
		data := []byte{}
		if err := rows.Scan(&data); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, errors.Wrap(err, "unmarshal data")
		}
		res = append(res, r)
	}
	return res, nil
}
func (s *Service) GetWeapon(id int) (*Weapon, error) {
	stmt, err := s.db.Prepare("SELECT data FROM weapons WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Weapon{}
	data := []byte{}
	row := stmt.QueryRow(id)

	if err := row.Scan(&data); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, errors.Wrap(err, "unmarshal data")
	}

	return res, nil
}
func (s *Service) SaveWeapon(weapon *Weapon) (int, error) {
	stmt, err := s.db.Prepare(`INSERT INTO weapons (data) VALUES(?)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(weapon)
	if err != nil {
		return 0, errors.Wrap(err, "marshal data")
	}
	res, err := stmt.Exec(data)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "last index")
	}
	if _, err := s.db.Exec("UPDATE weapons SET data = json_set(data, '$.id', id) WHERE id = ?", id); err != nil {
		return 0, errors.Wrap(err, "affect id to json")
	}
	return int(id), nil
}

func (s *Service) UpdateWeapon(weapon *Weapon) error {
	stmt, err := s.db.Prepare(`UPDATE weapons SET data = ? WHERE id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(weapon)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}
	_, err = stmt.Exec(data, weapon.ID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) DeleteWeapon(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM weapons WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) GetModelWeapons(modelID int) ([]Weapon, error) {
	stmt, err := s.db.Prepare("SELECT data FROM weapons WHERE model_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(modelID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Weapon{}
	for rows.Next() {
		r := Weapon{}
		data := []byte{}
		if err := rows.Scan(&data); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, errors.Wrap(err, "unmarshal data")
		}
		res = append(res, r)
	}
	return res, nil
}
