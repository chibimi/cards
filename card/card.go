package card

import (
	"strings"

	"github.com/pkg/errors"
)

type Card struct {
	ID         int    `json:"id,omitempty" db:"id"`
	FactionID  int    `json:"faction_id,omitempty" db:"faction_id"`
	CategoryID int    `json:"category_id,omitempty" db:"category_id"`
	MainCardID int    `json:"main_card_id,omitempty,string" db:"main_card_id"`
	Name       string `json:"name,omitempty" db:"name"`
	Properties string `json:"properties,omitempty" db:"properties"`
	Models     string `json:"models,omitempty" db:"models"`
	ModelsMax  string `json:"models_max,omitempty" db:"models_max"`
	Cost       string `json:"cost,omitempty" db:"cost"`
	CostMax    string `json:"cost_max,omitempty" db:"cost_max"`
	FA         string `json:"fa,omitempty" db:"fa"`
	Status     string `json:"status,omitempty" db:"status"`
}

func (s *Service) GetCard(id int) (*Card, error) {
	stmt, err := s.db.Preparex("SELECT * FROM cards WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Card{}
	if err := stmt.Get(res, id); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (s *Service) GetRelatedCards(id int) ([]Card, error) {
	stmt, err := s.db.Preparex("SELECT * FROM cards WHERE main_card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(id)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Card{}
	for rows.Next() {
		r := Card{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		res = append(res, r)
	}
	return res, nil
}

func (s *Service) GetCardAbilities(id int) ([]Ability, error) {
	stmt, err := s.db.Preparex("SELECT id, original_name, name, magical, description FROM card_ability AS l LEFT JOIN abilities AS a ON l.ability_id = a.id WHERE l.card_id = ?")
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

func (s *Service) GetCardFeat(id int) (*Feat, error) {
	stmt, err := s.db.Preparex("SELECT * FROM feats WHERE card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Feat{}
	if err := stmt.Get(res, id); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (s *Service) GetCardSpells(id int) ([]Spell, error) {
	stmt, err := s.db.Preparex("SELECT id, original_name, name, cost, rng, aoe, pow, dur, off, description FROM card_spell AS l LEFT JOIN spells AS a ON l.spell_id = a.id WHERE l.card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(id)
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

func (s *Service) GetCardModels(id int) ([]Model, error) {
	stmt, err := s.db.Preparex("SELECT * FROM models WHERE card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(id)
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

func (s *Service) ListCards(factionID, categoryID int) ([]Card, error) {
	stmt, err := s.db.Preparex("SELECT * FROM cards WHERE faction_id = ? AND category_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(factionID, categoryID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Card{}
	for rows.Next() {
		r := Card{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		res = append(res, r)
	}
	return res, nil
}

func (s *Service) DeleteCard(id int) error {
	stmt, err := s.db.Preparex("DELETE FROM cards WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}

func (s *Service) SaveCard(card *Card) (int, error) {
	stmt, err := s.db.PrepareNamed(`INSERT INTO cards VALUES(
		:id, :main_card_id, :faction_id, :category_id, :name, :properties, :models, :models_max, :cost, :cost_max, :fa, :status
	)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res, err := stmt.Exec(card)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	return int(id), errors.Wrap(err, "last index")
}

func (s *Service) UpdateCard(card *Card) error {
	stmt, err := s.db.PrepareNamed(`UPDATE cards SET 
	faction_id = :faction_id, category_id = :category_id, main_card_id = :main_card_id, name = :name, properties = :properties, 
	models = :models, models_max = :models_max, cost = :cost, cost_max = :cost_max, fa = :fa, status = :status
	WHERE id = :id`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(card)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
