package card

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Model struct {
	ID           int      `json:"id,omitempty"`
	CardID       int      `json:"card_id,omitempty"`
	Name         string   `json:"name,omitempty"`
	SPD          string   `json:"spd,omitempty"`
	STR          string   `json:"str,omitempty"`
	MAT          string   `json:"mat,omitempty"`
	RAT          string   `json:"rat,omitempty"`
	DEF          string   `json:"def,omitempty"`
	ARM          string   `json:"arm,omitempty"`
	CMD          string   `json:"cmd,omitempty"`
	BaseSize     string   `json:"base_size,omitempty"`
	MagicAbility string   `json:"magic_ability,omitempty"`
	Resource     string   `json:"resource,omitempty"`
	Threshold    string   `json:"threshold,omitempty"`
	Damage       string   `json:"damage,omitempty"`
	Advantages   []string `json:"advantages"`
	AdvantagesDB string   `json:"-"`
}

func (s *Service) ListModels() ([]Model, error) {
	stmt, err := s.db.Prepare("SELECT data FROM models")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Model{}
	for rows.Next() {
		r := Model{}
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

func (s *Service) GetModel(id int) (*Model, error) {
	stmt, err := s.db.Prepare("SELECT data FROM models WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Model{}
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

func (s *Service) SaveModel(model *Model) (int, error) {
	stmt, err := s.db.Prepare(`INSERT INTO models (data) VALUES(?)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(model)
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
	if _, err := s.db.Exec("UPDATE models SET data = json_set(data, '$.id', id) WHERE id = ?", id); err != nil {
		return 0, errors.Wrap(err, "affect id to json")
	}
	return int(id), nil
}

func (s *Service) UpdateModel(model *Model) error {
	stmt, err := s.db.Prepare(`UPDATE models SET data = ? WHERE id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(model)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}
	_, err = stmt.Exec(data, model.ID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) DeleteModel(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM models WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) GetCardModels(cardID int) ([]Model, error) {
	stmt, err := s.db.Prepare("SELECT data FROM models WHERE card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(cardID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Model{}
	for rows.Next() {
		r := Model{}
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
