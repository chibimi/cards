package feat

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Save(f *Feat, lang string) error {
	stmt := `
	REPLACE INTO feats (ref_id, lang, name, description, fluff) 
	VALUES(?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(stmt, f.RefID, lang, f.Name, f.Description, f.Fluff)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) Get(id int, lang string) (*Feat, error) {
	stmt := `
	SELECT ref_id, name, description, fluff FROM feats WHERE ref_id = ? AND lang = ?
	`
	res := &Feat{}
	err := r.db.Get(res, stmt, id, lang)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}
