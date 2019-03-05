package card

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Spell struct {
	ID           int    `json:"id,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Cost         string `json:"cost,omitempty"`
	RNG          string `json:"rng,omitempty"`
	AOE          string `json:"aoe,omitempty"`
	POW          string `json:"pow,omitempty"`
	DUR          string `json:"dur,omitempty"`
	OFF          string `json:"off,omitempty"`
	Description  string `json:"description,omitempty"`
}

func (s *Service) ListSpells() ([]Spell, error) {
	stmt, err := s.db.Prepare("SELECT data FROM spells")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Spell{}
	for rows.Next() {
		r := Spell{}
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

func (s *Service) GetSpell(id int) (*Spell, error) {
	stmt, err := s.db.Prepare("SELECT data FROM spells WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Spell{}
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

func (s *Service) SaveSpell(spell *Spell) (int, error) {
	stmt, err := s.db.Prepare(`INSERT INTO spells (data) VALUES(?)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(spell)
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
	if _, err := s.db.Exec("UPDATE spells SET data = json_set(data, '$.id', id) WHERE id = ?", id); err != nil {
		return 0, errors.Wrap(err, "affect id to json")
	}
	return int(id), nil
}

func (s *Service) UpdateSpell(spell *Spell) error {
	stmt, err := s.db.Prepare(`UPDATE spells SET data = ? WHERE id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(spell)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}
	_, err = stmt.Exec(data, spell.ID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) DeleteSpell(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM spells WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) GetCardSpells(cardID int) ([]Spell, error) {
	stmt, err := s.db.Prepare("SELECT data FROM card_spell AS l LEFT JOIN spells AS a ON l.spell_id = a.id WHERE l.card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(cardID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Spell{}
	for rows.Next() {
		r := Spell{}
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
func (s *Service) AddCardSpell(cardID, spellID int) error {
	stmt, err := s.db.Prepare(`INSERT INTO card_spell VALUES(?, ?)`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(cardID, spellID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
func (s *Service) DeleteCardSpell(cardID, spellID int) error {
	stmt, err := s.db.Prepare(`DELETE FROM card_spell WHERE card_id = ? AND spell_id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(cardID, spellID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
