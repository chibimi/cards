package spell

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

func (r *Repository) Create(sp *Spell, lang string) (int, error) {
	stmt := `
	INSERT INTO spells (title, cost, rng, aoe, pow, dur, off)
	VALUES(:title, :cost, :rng, :aoe, :pow, :dur, :off)
	`

	res, err := r.db.NamedExec(stmt, sp)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "last index")
	}

	stmt = `
	REPLACE INTO spells_lang (spell_id, lang, name, description)
	VALUES (?, ?, TRIM(?), TRIM(?))
	`
	_, err = r.db.Exec(stmt, id, lang, sp.Name, sp.Description)
	if err != nil {
		return int(id), errors.Wrap(err, "execute query")
	}

	return int(id), nil
}

func (r *Repository) List(lang string) ([]Spell, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM spells
	) as r LEFT JOIN (
		SELECT spell_id, name, description FROM spells_lang WHERE lang = ?
	) as s ON r.id = s.spell_id
	`
	res := []Spell{}
	err := r.db.Select(&res, stmt, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (r *Repository) Save(sp *Spell, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `
	UPDATE spells SET
	title = :title, cost = :cost, rng = :rng, aoe = :aoe, pow = :pow, dur = :dur, off = :off
	WHERE id = :id
	`
	_, err = tx.NamedExec(stmt, sp)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `
	REPLACE INTO spells_lang (spell_id, lang, name, description)
	VALUES (?, ?, TRIM(?), TRIM(?))
	`
	_, err = tx.Exec(stmt, sp.ID, lang, sp.Name, sp.Description)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "create transaction")
	}

	return nil
}

func (r *Repository) ListByRef(ref int, lang string) ([]Spell, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM ref_spell WHERE ref_id = ?
	) as a LEFT JOIN (
		SELECT * FROM spells
	) as r ON a.spell_id = r.id LEFT JOIN (
		SELECT spell_id, name, description FROM spells_lang WHERE lang = ?
	) as s ON r.id = s.spell_id
	`
	res := []Spell{}
	err := r.db.Select(&res, stmt, ref, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	return res, nil
}

func (r *Repository) AddSpellRef(ref, spell int) error {
	stmt := `INSERT INTO ref_spell VALUES(?, ?)`

	_, err := r.db.Exec(stmt, ref, spell)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) DeleteSpellRef(ref, spell int) error {
	stmt := `DELETE FROM ref_spell WHERE ref_id = ? AND spell_id = ?`

	_, err := r.db.Exec(stmt, ref, spell)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) Get(id int, lang string) (*Spell, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM spells WHERE id = ?
	) as r LEFT JOIN (
		SELECT * FROM spells_lang WHERE spell_id = ? AND lang = ?
	) as s ON r.id = s.spell_id
	`
	res := &Spell{}
	err := r.db.Get(res, stmt, id, id, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}
