package ability

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

func (r *Repository) Create(sp *Ability, lang string) (int, error) {
	stmt := `
	INSERT INTO abilities (title) 
	VALUES(:title)
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
	REPLACE INTO abilities_lang (ability_id, lang, name, description)
	VALUES (?, ?, ?, ?)
	`
	_, err = r.db.Exec(stmt, id, lang, sp.Name, sp.Description)
	if err != nil {
		return int(id), errors.Wrap(err, "execute query")
	}

	return int(id), nil
}

func (r *Repository) List(lang string) ([]Ability, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM abilities
	) as r LEFT JOIN (
		SELECT ability_id, name, description FROM abilities_lang WHERE lang = ?
	) as s ON r.id = s.ability_id
	`
	res := []Ability{}
	err := r.db.Select(&res, stmt, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (r *Repository) Save(sp *Ability, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `
	UPDATE abilities SET 
	title = :title
	WHERE id = :id
	`
	_, err = tx.NamedExec(stmt, sp)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `
	REPLACE INTO abilities_lang (ability_id, lang, name, description)
	VALUES (?, ?, ?, ?)
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

func (r *Repository) ListByRef(ref int, lang string) ([]Ability, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM ref_ability WHERE ref_id = ?
	) as a LEFT JOIN (
		SELECT * FROM abilities
	) as r ON a.ability_id = r.id LEFT JOIN (
		SELECT ability_id, name, description FROM abilities_lang WHERE lang = ?
	) as s ON r.id = s.ability_id
	`
	res := []Ability{}
	err := r.db.Select(&res, stmt, ref, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	return res, nil
}

func (r *Repository) AddAbilityRef(ref, ability, typ int) error {
	stmt := `INSERT INTO ref_ability VALUES(?, ?, ?)`

	_, err := r.db.Exec(stmt, ref, ability, typ)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) DeleteAbilityRef(ref, ability int) error {
	stmt := `DELETE FROM ref_ability WHERE ref_id = ? AND ability_id = ?`

	_, err := r.db.Exec(stmt, ref, ability)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) ListByModel(model int, lang string) ([]Ability, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM model_ability WHERE model_id = ?
	) as a LEFT JOIN (
		SELECT * FROM abilities
	) as r ON a.ability_id = r.id LEFT JOIN (
		SELECT ability_id, name, description FROM abilities_lang WHERE lang = ?
	) as s ON r.id = s.ability_id
	`
	res := []Ability{}
	err := r.db.Select(&res, stmt, model, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	return res, nil
}

func (r *Repository) AddAbilityModel(model, ability, typ int) error {
	stmt := `INSERT INTO model_ability VALUES(?, ?, ?)`

	_, err := r.db.Exec(stmt, model, ability, typ)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) DeleteAbilityModel(model, ability int) error {
	stmt := `DELETE FROM model_ability WHERE model_id = ? AND ability_id = ?`

	_, err := r.db.Exec(stmt, model, ability)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) ListByWeapon(weapon int, lang string) ([]Ability, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM weapon_ability WHERE weapon_id = ?
	) as a LEFT JOIN (
		SELECT * FROM abilities
	) as r ON a.ability_id = r.id LEFT JOIN (
		SELECT ability_id, name, description FROM abilities_lang WHERE lang = ?
	) as s ON r.id = s.ability_id
	`
	res := []Ability{}
	err := r.db.Select(&res, stmt, weapon, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	return res, nil
}

func (r *Repository) AddAbilityWeapon(weapon, ability, typ int) error {
	stmt := `INSERT INTO weapon_ability VALUES(?, ?, ?)`

	_, err := r.db.Exec(stmt, weapon, ability, typ)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) DeleteAbilityWeapon(weapon, ability int) error {
	stmt := `DELETE FROM weapon_ability WHERE weapon_id = ? AND ability_id = ?`

	_, err := r.db.Exec(stmt, weapon, ability)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (r *Repository) GetLang(id int, lang string) (*Ability, error) {
	stmt := `
	SELECT name, description FROM abilities_lang WHERE ability_id = ? AND lang = ?
	`
	res := &Ability{}
	err := r.db.Get(res, stmt, id, lang)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}

func (r *Repository) Get(id int, lang string) (*Ability, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name, IFNULL(s.description, "") as description FROM (
		SELECT * FROM abilities WHERE id = ?
	) as r LEFT JOIN (
		SELECT * FROM abilities_lang WHERE ability_id = ? AND lang = ?
	) as s ON r.id = s.ability_id
	`
	res := &Ability{}
	err := r.db.Get(res, stmt, id, id, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return res, nil
}
