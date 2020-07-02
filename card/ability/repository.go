package ability

import (
	"fmt"

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
	VALUES (?, ?, TRIM(?), TRIM(?))
	`
	_, err = r.db.Exec(stmt, id, lang, sp.Name, sp.Description)
	if err != nil {
		return int(id), errors.Wrap(err, "execute query")
	}

	return int(id), nil
}

func (r *Repository) Save(sp *Ability, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `UPDATE abilities SET title = :title WHERE id = :id`
	if sp.ID == 0 {
		stmt = `INSERT INTO abilities (title) VALUES(:title)`
	}

	res, err := tx.NamedExec(stmt, sp)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "last index")
	}
	fmt.Println(id)

	stmt = `
	REPLACE INTO abilities_lang (ability_id, lang, name, description)
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

func (r *Repository) ListByModel(model int, lang string) ([]Ability, error) {
	stmt := `
	SELECT id, title, COALESCE(name,"") as name, COALESCE(description,"") as description, star, header
	FROM  model_ability
	LEFT JOIN abilities ON model_ability.ability_id = abilities.id
	LEFT JOIN abilities_lang ON abilities.id = abilities_lang.ability_id
	LEFT JOIN (
		SELECT ability_id, name as header_name from abilities_lang where lang = ?
	) as h on h.ability_id = model_ability.header
	WHERE lang = ? AND model_id = ?
	ORDER BY COALESCE(header_name,name), header_name, name
	`
	res := []Ability{}
	err := r.db.Select(&res, stmt, lang, lang, model)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	return res, nil
}

func (r *Repository) AddAbilityModel(relation *Relation) error {
	stmt := `INSERT INTO model_ability VALUES(:relatedid, :abilityid, 0, :star, :header)
	ON DUPLICATE KEY UPDATE model_id = :relatedid, ability_id = :abilityid, star = :star, header = :header`

	_, err := r.db.NamedExec(stmt, relation)
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
	SELECT id, title, COALESCE(name,"") as name, COALESCE(description,"") as description, star, header
	FROM  weapon_ability
	LEFT JOIN abilities ON weapon_ability.ability_id = abilities.id
	LEFT JOIN abilities_lang ON abilities.id = abilities_lang.ability_id
	LEFT JOIN (
		SELECT ability_id, name as header_name from abilities_lang where lang = ?
	) as h on h.ability_id = weapon_ability.header
	WHERE lang = ? AND weapon_id = ?
	ORDER BY COALESCE(header_name,name), header_name, name
	`
	res := []Ability{}
	err := r.db.Select(&res, stmt, lang, lang, weapon)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	return res, nil
}

func (r *Repository) AddAbilityWeapon(relation *Relation) error {
	stmt := `INSERT INTO weapon_ability VALUES(:relatedid, :abilityid, 0, :star, :header)
	ON DUPLICATE KEY UPDATE weapon_id = :relatedid, ability_id = :abilityid, star = :star, header = :header`

	_, err := r.db.NamedExec(stmt, relation)
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
