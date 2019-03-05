package card

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Feat struct {
	ID           int    `json:"id,omitempty"`
	CardID       int    `json:"card_id,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Fluff        string `json:"fluff,omitempty"`
}

func (s *Service) ListFeats() ([]Feat, error) {
	stmt, err := s.db.Prepare("SELECT data FROM feats")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Feat{}
	for rows.Next() {
		r := Feat{}
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

func (s *Service) GetFeat(id int) (*Feat, error) {
	stmt, err := s.db.Prepare("SELECT data FROM feats WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Feat{}
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

func (s *Service) SaveFeat(feat *Feat) (int, error) {
	stmt, err := s.db.Prepare(`INSERT INTO feats (data) VALUES(?)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(feat)
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
	if _, err := s.db.Exec("UPDATE feats SET data = json_set(data, '$.id', id) WHERE id = ?", id); err != nil {
		return 0, errors.Wrap(err, "affect id to json")
	}
	return int(id), nil
}

func (s *Service) UpdateFeat(feat *Feat) error {
	stmt, err := s.db.Prepare(`UPDATE feats SET data = ? WHERE id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(feat)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}
	_, err = stmt.Exec(data, feat.ID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) DeleteFeat(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM feats WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) GetCardFeat(cardID int) (*Feat, error) {
	stmt, err := s.db.Prepare("SELECT data FROM feats WHERE card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Feat{}
	data := []byte{}
	row := stmt.QueryRow(cardID)

	if err := row.Scan(&data); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, errors.Wrap(err, "unmarshal data")
	}

	return res, nil
}
