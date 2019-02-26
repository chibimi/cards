package card

import (
	"github.com/pkg/errors"
)

type Card struct {
	ID         int    `json:"id,omitempty" db:"id"`
	FactionID  int    `json:"faction_id,omitempty" db:"faction_id"`
	CategoryID int    `json:"category_id,omitempty" db:"category_id"`
	MainCardID int    `json:"main_card_id,omitempty" db:"main_card_id"`
	Name       string `json:"name,omitempty" db:"name"`
	Properties string `json:"properties,omitempty" db:"properties"`
	ModelsMin  string `json:"models_min,omitempty" db:"models_min"`
	Models     string `json:"models,omitempty" db:"models"`
	CostMin    string `json:"cost_min,omitempty" db:"cost_min"`
	Cost       string `json:"cost,omitempty" db:"cost"`
	FA         string `json:"fa,omitempty" db:"fa"`

	Fury      string `json:"fury,omitempty" db:"fury"`
	Focus     string `json:"focus,omitempty" db:"focus"`
	Threshold string `json:"threshold,omitempty" db:"threshold"`

	Damage       string `json:"damage,omitempty" db:"damage"`
	DamageGrid   string `json:"damage_grid,omitempty" db:"damage_grid"`
	DamageSpiral string `json:"damage_spiral,omitempty" db:"damage_spiral"`

	Status string `json:"status,omitempty" db:"status"`
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
		:id, :main_card_id, :faction_id, :category_id, :main_card_id:, :name, :properties, :models_min, :models, :cost_min, :cost, :fa, :status,
		:fury, :focus, :threshold, :damage, :damage_grid, :damage_spiral
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
	models_min = :models_min, models = :model, cost_min = :cost_min, cost = :cost, fa = :fa, status = :status, 
	fury = :fury, focus = :focus, threshold = :threshold, damage = :damage, damage_grid = :damage_grid, damage_spiral = :damage_spiral 
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
