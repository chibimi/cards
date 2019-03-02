package card

import (
	"github.com/pkg/errors"
)

type Feat struct {
	ID           int    `json:"id,omitempty" db:"id"`
	CardID       int    `json:"card_id,omitempty" db:"card_id"`
	OriginalName string `json:"original_name,omitempty" db:"original_name"`
	Name         string `json:"name,omitempty" db:"name"`
	Description  string `json:"description,omitempty" db:"description"`
	Fluff        string `json:"fluff,omitempty" db:"fluff"`
}

func (s *Service) GetFeat(id int) (*Feat, error) {
	stmt, err := s.db.Preparex("SELECT * FROM feats WHERE id = ?")
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

func (s *Service) DeleteFeat(id int) error {
	stmt, err := s.db.Preparex("DELETE FROM feats WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}

func (s *Service) SaveFeat(feat *Feat) (int, error) {
	stmt, err := s.db.PrepareNamed(`INSERT INTO feats VALUES(
		:id, :card_id, :original_name, :name, :description, :fluff
	)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res, err := stmt.Exec(feat)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	return int(id), errors.Wrap(err, "last index")
}

func (s *Service) UpdateFeat(feat *Feat) error {
	stmt, err := s.db.PrepareNamed(`UPDATE feats SET 
	card_id = :card_id, name = :name, original_name = :original_name, description = :description, fluff = :fluff WHERE id = :id`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(feat)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) ListFeats() ([]Feat, error) {
	stmt, err := s.db.Preparex("SELECT * FROM feats")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Queryx()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	res := []Feat{}
	for rows.Next() {
		r := Feat{}
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		res = append(res, r)
	}
	return res, nil
}
