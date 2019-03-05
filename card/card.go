package card

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Card struct {
	ID         int    `json:"id,omitempty"`
	FactionID  int    `json:"faction_id,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
	MainCardID int    `json:"main_card_id,omitempty,string"`
	Name       string `json:"name,omitempty"`
	Properties string `json:"properties,omitempty"`
	Models     string `json:"models_cnt,omitempty"`
	ModelsMax  string `json:"models_max,omitempty"`
	Cost       string `json:"cost,omitempty"`
	CostMax    string `json:"cost_max,omitempty"`
	FA         string `json:"fa,omitempty"`
	Status     string `json:"status,omitempty"`
}

func (s *Service) ListCards(factionID, categoryID int) ([]Card, error) {
	stmt, err := s.db.Prepare("SELECT data FROM cards WHERE faction_id = ? AND category_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(factionID, categoryID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Card{}
	for rows.Next() {
		r := Card{}
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

func (s *Service) GetCard(id int) (*Card, error) {
	stmt, err := s.db.Prepare("SELECT data FROM cards WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Card{}
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

func (s *Service) SaveCard(card *Card) (int, error) {
	stmt, err := s.db.Prepare(`INSERT INTO cards (data) VALUES(?)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(card)
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
	if _, err := s.db.Exec("UPDATE cards SET data = json_set(data, '$.id', id) WHERE id = ?", id); err != nil {
		return 0, errors.Wrap(err, "affect id to json")
	}
	return int(id), nil
}

func (s *Service) UpdateCard(card *Card) error {
	stmt, err := s.db.Prepare(`UPDATE cards SET data = ? WHERE id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(card)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}
	_, err = stmt.Exec(data, card.ID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) DeleteCard(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM cards WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}

func (s *Service) GetRelatedCards(id int) ([]Card, error) {
	stmt, err := s.db.Prepare("SELECT data FROM cards WHERE main_card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Card{}
	for rows.Next() {
		r := Card{}
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
