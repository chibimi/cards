package reference

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

type referenceDB struct {
	Reference
	Merc   []byte `db:"mercenary_for"`
	Minion []byte `db:"minion_for"`
}

func (r *Repository) Create(ref *Reference) (int, error) {
	stmt := `
	INSERT INTO refs (ppid, faction_id, category_id, title, main_card_id, models_cnt, models_max, cost, cost_max, fa, mercenary_for, minion_for, special, linked_to)
	VALUES(:ppid, :faction_id, :category_id, :title, :main_card_id, :models_cnt, :models_max, :cost, :cost_max, :fa, :mercenary_for, :minion_for, :special, :linked_to)
	`
	merc, err := json.Marshal(ref.MercFor)
	if err != nil {
		return 0, errors.Wrap(err, "marshal merc")
	}
	minion, err := json.Marshal(ref.MinFor)
	if err != nil {
		return 0, errors.Wrap(err, "marshal minion")
	}
	refDB := &referenceDB{
		Reference: *ref,
		Minion:    minion,
		Merc:      merc,
	}
	res, err := r.db.NamedExec(stmt, refDB)
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
		SELECT id, ppid, faction_id, category_id, title FROM refs WHERE faction_id = ? AND category_id = ?
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
	res := &referenceDB{}
	err := r.db.Get(res, stmt, id, id, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	merc := []int{}
	err = json.Unmarshal(res.Merc, &merc)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall merc")
	}
	minion := []int{}
	err = json.Unmarshal(res.Merc, &minion)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall minion")
	}
	res.MercFor = merc
	res.MinFor = minion
	return &res.Reference, nil
}

func (r *Repository) Save(ref *Reference, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `
	UPDATE refs SET ppid = :ppid, faction_id = :faction_id, category_id = :category_id, title = :title, main_card_id = :main_card_id,
	models_cnt = :models_cnt, models_max = :models_max, cost = :cost, cost_max = :cost_max, fa = :fa,
	mercenary_for = :mercenary_for, minion_for = :minion_for, special = :special, linked_to = :linked_to
	WHERE id = :id
	`
	merc, err := json.Marshal(ref.MercFor)
	if err != nil {
		return errors.Wrap(err, "marshal merc")
	}
	minion, err := json.Marshal(ref.MinFor)
	if err != nil {
		return errors.Wrap(err, "marshal minion")
	}
	refDB := &referenceDB{
		Reference: *ref,
		Minion:    minion,
		Merc:      merc,
	}
	_, err = tx.NamedExec(stmt, refDB)
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

func (r *Repository) ListByStatus(lang, status string) ([]Reference, error) {
	stmt := `
	SELECT r.*, IFNULL(s.status, "wip") as status FROM (
		SELECT id, faction_id, category_id, title FROM refs
	) as r INNER JOIN (
		SELECT ref_id, status FROM refs_lang WHERE lang = ? AND status = ?
	) as s ON r.id = s.ref_id
	`
	res := []Reference{}
	err := r.db.Select(&res, stmt, lang, status)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (r *Repository) ListRefAttachments(lang string, linked_to int) ([]Reference, error) {
	stmt := `
	SELECT id, faction_id, category_id, title
	FROM refs
	WHERE linked_to = ?
	`
	res := []Reference{}
	err := r.db.Select(&res, stmt, linked_to)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}
