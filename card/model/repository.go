package model

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

type modelDB struct {
	Model
	Advantages []byte `db:"advantages"`
}

func (r *Repository) Create(m *Model, lang string) (int, error) {
	stmt := `
	INSERT INTO models (ref_id, spd, str, mat, rat, def, arm, cmd, base_size, magic_ability, resource, threshold, damage, advantages) 
	VALUES(:ref_id, :spd, :str, :mat, :rat, :def, :arm, :cmd, :base_size, :magic_ability, :resource, :threshold, :damage, :advantages)
	`
	adv, err := json.Marshal(m.Advantages)
	if err != nil {
		return 0, errors.Wrap(err, "marshal advantages")
	}
	mDB := &modelDB{
		Model:      *m,
		Advantages: adv,
	}
	res, err := r.db.NamedExec(stmt, mDB)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "last index")
	}

	stmt = `
	REPLACE INTO models_lang (model_id, lang, name)
	VALUES (?, ?, ?)
	`
	_, err = r.db.Exec(stmt, id, lang, m.Name)
	if err != nil {
		return int(id), errors.Wrap(err, "execute query")
	}

	return int(id), nil
}

func (r *Repository) List(ref int, lang string) ([]Model, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name FROM (
		SELECT * FROM models WHERE ref_id = ?
	) as r LEFT JOIN (
		SELECT model_id, name FROM models_lang WHERE lang = ?
	) as s ON r.id = s.model_id
	`
	res := []Model{}

	rows, err := r.db.Queryx(stmt, ref, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	for rows.Next() {
		tmp := modelDB{}
		err := rows.StructScan(&tmp)
		if err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		m := tmp.Model
		adv := []string{}
		err = json.Unmarshal(tmp.Advantages, &adv)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshall adv")
		}
		m.Advantages = adv
		res = append(res, m)
	}

	return res, nil
}

func (r *Repository) Get(id int, lang string) (*Model, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name FROM (
		SELECT * FROM models WHERE id = ?
	) as r LEFT JOIN (
		SELECT * FROM models_lang WHERE model_id = ? AND lang = ?
	) as s ON r.id = s.model_id
	`
	res := &modelDB{}
	err := r.db.Get(res, stmt, id, id, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	adv := []string{}
	err = json.Unmarshal(res.Advantages, &adv)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall adv")
	}
	res.Model.Advantages = adv
	return &res.Model, nil
}

func (r *Repository) Save(m *Model, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `
	UPDATE models SET ref_id = :ref_id, spd = :spd, str = :str, mat = :mat, rat = :rat, def = :def, arm = :arm, cmd = :cmd, 
	base_size = :base_size, magic_ability = :magic_ability, resource = :resource, threshold = :threshold, 
	damage = :damage, advantages = :advantages
	WHERE id = :id
	`
	adv, err := json.Marshal(m.Advantages)
	if err != nil {
		return errors.Wrap(err, "marshal merc")
	}

	mDB := &modelDB{
		Model:      *m,
		Advantages: adv,
	}
	_, err = tx.NamedExec(stmt, mDB)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `
	REPLACE INTO models_lang (model_id, lang, name)
	VALUES (?, ?, ?)
	`
	_, err = tx.Exec(stmt, m.ID, lang, m.Name)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "create transaction")
	}

	return nil
}

func (r *Repository) Delete(id int) error {
	stmt := `
	DELETE FROM models WHERE id = ?
	`
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
