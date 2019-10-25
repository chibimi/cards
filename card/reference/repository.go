package reference

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

func (r *Repository) Create(ref *Reference) (int, error) {
	stmt := `
	INSERT INTO refs (faction_id, category_id, title, main_card_id, models_cnt, models_max, cost, cost_max, fa) 
	VALUES(:faction_id, :category_id, :title, :main_card_id, :models_cnt, :models_max, :cost, :cost_max, :fa)
	`

	res, err := r.db.NamedExec(stmt, ref)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "last index")
	}
	return int(id), nil
}

func (r *Repository) List(faction, category int, lang string) ([]Reference, error) {
	stmt := `
	SELECT r.*, IFNULL(s.status, "wip") as status FROM (
		SELECT * FROM refs WHERE faction_id = ? AND category_id = ? 
	) as r LEFT JOIN (
		SELECT ref_id, status FROM refs_lang WHERE lang = ?
	) as s ON r.id = s.ref_id
	`
	res := []Reference{}
	err := r.db.Select(&res, stmt, faction, category, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (r *Repository) Get(id int, lang string) (*Reference, error) {
	stmt := `
	SELECT r.*, IFNULL(s.status, "wip") as status, IFNULL(s.name, "") as name, IFNULL(s.properties, "") as properties FROM (
		SELECT * FROM refs WHERE id = ?
	) as r LEFT JOIN (
		SELECT * FROM refs_lang WHERE ref_id = ? AND lang = ?
	) as s ON r.id = s.ref_id
	`
	res := &Reference{}
	err := r.db.Get(res, stmt, id, id, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (r *Repository) Save(ref *Reference, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `
	UPDATE refs SET faction_id = :faction_id, category_id = :category_id, title = :title, main_card_id = :main_card_id, 
	models_cnt = :models_cnt, models_max = :models_max, cost = :cost, cost_max = :cost_max, fa = :fa
	WHERE id = :id
	`
	_, err = tx.NamedExec(stmt, ref)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `
	REPLACE INTO refs_lang (ref_id, lang, name, properties, status)
	VALUES (?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(stmt, ref.ID, lang, ref.Name, ref.Properties, ref.Status)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "create transaction")
	}

	return nil
}
