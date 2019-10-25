package card

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

func (r *Repository) Save(card *Card, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `
	REPLACE INTO cards (ref_id, main_card_id, models_cnt, models_max, cost, cost_max, fa)
	VALUES (:id, :main_card_id, :models_cnt, :models_max, :cost, :cost_max, :fa)
	`
	_, err = tx.NamedExec(stmt, card)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `UPDATE refs SET faction_id = :faction_id, category_id = :category_id, title = :title WHERE id = :id`
	_, err = r.db.NamedExec(stmt, card)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `
	REPLACE INTO refs_status (ref_id, lang, status)
	VALUES (?, ?, ?)
	`
	_, err = tx.Exec(stmt, card.ID, lang, card.Status)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `
	REPLACE INTO cards_lang (ref_id, lang, name, properties)
	VALUES (?, ?, ?, ?)
	`
	_, err = tx.Exec(stmt, card.ID, lang, card.Name, card.Properties)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "create transaction")
	}

	return nil
}

func (r *Repository) Get(id int, lang string) (*Card, error) {
	stmt := `
	SELECT r.*, IFNULL(s.status, "wip") as status, IFNULL(l.name, "") as name, IFNULL(l.properties, "") as properties,
	IFNULL(c.main_card_id, 0) as main_card_id, IFNULL(c.models_cnt, "") as models_cnt, IFNULL(c.models_max, "") as models_max, 
	IFNULL(c.cost, "") as cost, IFNULL(c.cost_max, "") as cost_max, IFNULL(c.fa, "") as fa
	FROM (
		SELECT * FROM refs WHERE id = ?
	) as r LEFT JOIN (
		SELECT * FROM refs_status WHERE ref_id = ? AND lang = ?
	) as s ON r.id = s.ref_id LEFT JOIN (
		SELECT * FROM cards WHERE ref_id = ?
	) as c ON r.id = c.ref_id LEFT JOIN (
		SELECT * FROM cards_lang WHERE ref_id = ? AND lang = ?
	) as l ON r.id = l.ref_id
	`
	res := &Card{}
	err := r.db.Get(res, stmt, id, id, lang, id, id, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	return res, nil
}
