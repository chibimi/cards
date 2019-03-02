package card

import (
	"github.com/pkg/errors"
)

type Spell struct {
	ID           int    `json:"id,omitempty" db:"id"`
	OriginalName string `json:"original_name,omitempty" db:"original_name"`
	Name         string `json:"name,omitempty" db:"name"`
	Cost         string `json:"cost,omitempty" db:"cost"`
	RNG          string `json:"rng,omitempty" db:"rng"`
	AOE          string `json:"aoe,omitempty" db:"aoe"`
	POW          string `json:"pow,omitempty" db:"pow"`
	DUR          string `json:"dur,omitempty" db:"dur"`
	OFF          string `json:"off,omitempty" db:"off"`
	Description  string `json:"description,omitempty" db:"description"`
}

func (s *Service) GetSpell(id int) (*Spell, error) {
	stmt, err := s.db.Preparex("SELECT * FROM spells WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Spell{}

	if err := stmt.Get(res, id); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (s *Service) DeleteSpell(id int) error {
	stmt, err := s.db.Preparex("DELETE FROM spells WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}

func (s *Service) SaveSpell(spell *Spell) (int, error) {
	stmt, err := s.db.PrepareNamed(`INSERT INTO spells VALUES(
		:id, :original_name, :name, :cost, :rng, :aoe, :pow, :dur, :off, :description
	)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res, err := stmt.Exec(spell)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	return int(id), errors.Wrap(err, "last index")
}

func (s *Service) UpdateSpell(spell *Spell) error {
	stmt, err := s.db.PrepareNamed(`UPDATE spells SET 
	name = :name, original_name = :original_name, cost = :cost, rng = :rng, aoe = :aoe, pow = :pow, dur = :dur, off = :off, description = :description WHERE id = :id`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(spell)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) ListSpells() ([]Spell, error) {
	stmt, err := s.db.Preparex("SELECT * FROM spells")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res := []Spell{}
	for rows.Next() {
		r := Spell{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
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
