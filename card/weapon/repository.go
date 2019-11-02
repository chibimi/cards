package weapon

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

type weaponDB struct {
	Weapon
	Advantages []byte `db:"advantages"`
}

func (r *Repository) Create(wp *Weapon, lang string) (int, error) {
	stmt := `
	INSERT INTO weapons (model_id, type, rng, pow, rof, aoe, loc, cnt, advantages) 
	VALUES(:model_id, :type, :rng, :pow, :rof, :aoe, :loc, :cnt, :advantages)
	`

	adv, err := json.Marshal(wp.Advantages)
	if err != nil {
		return 0, errors.Wrap(err, "marshal advantages")
	}
	mDB := &weaponDB{
		Weapon:     *wp,
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
	REPLACE INTO weapons_lang (weapon_id, lang, name)
	VALUES (?, ?, ?)
	`
	_, err = r.db.Exec(stmt, id, lang, wp.Name)
	if err != nil {
		return int(id), errors.Wrap(err, "execute query")
	}

	return int(id), nil
}

func (r *Repository) List(ref int, lang string) ([]Weapon, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name FROM (
		SELECT * FROM weapons WHERE model_id = ?
	) as r LEFT JOIN (
		SELECT weapon_id, name FROM weapons_lang WHERE lang = ?
	) as s ON r.id = s.weapon_id
	`
	res := []Weapon{}

	rows, err := r.db.Queryx(stmt, ref, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	for rows.Next() {
		tmp := weaponDB{}
		err := rows.StructScan(&tmp)
		if err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		m := tmp.Weapon
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

func (r *Repository) Get(id int, lang string) (*Weapon, error) {
	stmt := `
	SELECT r.*, IFNULL(s.name, "") as name FROM (
		SELECT * FROM weapons WHERE id = ?
	) as r LEFT JOIN (
		SELECT * FROM weapons_lang WHERE weapon_id = ? AND lang = ?
	) as s ON r.id = s.weapon_id
	`
	res := &weaponDB{}
	err := r.db.Get(res, stmt, id, id, lang)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	adv := []string{}
	err = json.Unmarshal(res.Advantages, &adv)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshall adv")
	}
	res.Weapon.Advantages = adv
	return &res.Weapon, nil
}

func (r *Repository) Save(wp *Weapon, lang string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	stmt := `
	UPDATE weapons SET 
	model_id = :model_id, type = :type, rng = :rng, pow = :pow, rof = :rof , aoe = :aoe, loc = :loc, cnt = :cnt, advantages = :advantages
	WHERE id = :id
	`
	adv, err := json.Marshal(wp.Advantages)
	if err != nil {
		return errors.Wrap(err, "marshal merc")
	}

	mDB := &weaponDB{
		Weapon:     *wp,
		Advantages: adv,
	}
	_, err = tx.NamedExec(stmt, mDB)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}

	stmt = `
	REPLACE INTO weapons_lang (weapon_id, lang, name)
	VALUES (?, ?, ?)
	`
	_, err = tx.Exec(stmt, wp.ID, lang, wp.Name)
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
	DELETE FROM weapons WHERE id = ?
	`
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
