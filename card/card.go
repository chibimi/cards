package card

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Card struct {
	ID         int    `json:"id,omitempty" db:"id"`
	FactionID  int    `json:"faction_id,omitempty" db:"faction_id"`
	CategoryID int    `json:"category_id,omitempty" db:"category_id"`
	MainCardID int    `json:"main_card_id,omitempty,string" db:"main_card_id"`
	Name       string `json:"name,omitempty" db:"name"`
	Properties string `json:"properties,omitempty" db:"properties"`
	Models     string `json:"models_cnt,omitempty" db:"models_cnt"`
	ModelsMax  string `json:"models_max,omitempty" db:"models_max"`
	Cost       string `json:"cost,omitempty" db:"cost"`
	CostMax    string `json:"cost_max,omitempty" db:"cost_max"`
	FA         string `json:"fa,omitempty" db:"fa"`
	Status     string `json:"status,omitempty" db:"status"`
}

func (s *Service) ListCards(factionID, categoryID int, lang string) ([]Card, error) {
	stmt, err := s.db.Preparex(`
	SELECT c.*, en.name as name, IFNULL(l.status, "wip") as status
	FROM (
		SELECT * FROM cards WHERE faction_id = ? AND category_id = ?
	) as c INNER JOIN (
		SELECT card_id, name FROM cards_lang WHERE lang = "UK"
	) as en ON c.id = en.card_id LEFT JOIN (
		SELECT card_id, status FROM cards_lang WHERE lang = ?
	) as l ON c.id = l.card_id
	`)
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(factionID, categoryID, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Card{}
	for rows.Next() {
		r := Card{}
		err := rows.StructScan(&r)
		if err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		res = append(res, r)
	}
	return res, nil
}

func (s *Service) GetCard(id int, lang string) (*Card, error) {
	stmt, err := s.db.Preparex(`
	SELECT c.*, IFNULL(name, "") as name, IFNULL(properties, "") as properties, IFNULL(status, "wip") as status
	FROM (
		SELECT * from cards WHERE id = ?
	) AS c LEFT JOIN (
		SELECT * FROM cards_lang WHERE card_id = ? AND lang = ?
	) AS l ON c.id=l.card_id
	`)
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Card{}
	row := stmt.QueryRowx(id, id, lang)
	if err := row.StructScan(res); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (s *Service) SaveCard(card *Card, lang string) (int, error) {
	stmt := `
	INSERT INTO cards (faction_id, category_id, main_card_id, models_cnt, models_max, cost, cost_max, fa) 
	VALUES(:faction_id, :category_id, :main_card_id, :models_cnt, :models_max, :cost, :cost_max, :fa)
	`
	res, err := s.db.NamedExec(stmt, card)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "last index")
	}
	card.ID = int(id)

	stmt = `
	INSERT INTO cards_lang (card_id, lang, name, properties, status) 
	VALUES(?, ?, ?, ?, ?)
	`
	res, err = s.db.Exec(stmt, card.ID, lang, card.Name, card.Properties, card.Status)
	if err != nil {
		return 0, errors.Wrap(err, "execute query lang")
	}

	return card.ID, nil
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
