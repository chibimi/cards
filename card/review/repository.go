package review

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Save(review *Review) error {
	stmt := `
	INSERT INTO reviews_lang (ref_id, lang, ip, rating, comment, reviewer, created_at)
	VALUES(:ref_id, :lang, :ip, :rating, :comment, :reviewer, :created_at)
	`

	_, err := r.db.NamedExec(stmt, review)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	return nil
}
