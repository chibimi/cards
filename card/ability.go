package card

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Ability struct {
	ID           int    `json:"id,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Magical      bool   `json:"magical,omitempty"`
	Description  string `json:"description,omitempty"`
}

func (s *Service) ListAbilities() ([]Ability, error) {
	stmt, err := s.db.Prepare("SELECT data FROM abilities")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Ability{}
	for rows.Next() {
		r := Ability{}
		data := []byte{}
		if err := rows.Scan(&data); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, errors.Wrap(err, "unmarshal data")
		}
		res = append(res, r)
	}
	return res, nil
}

func (s *Service) GetAbility(id int) (*Ability, error) {
	stmt, err := s.db.Prepare("SELECT data FROM abilities WHERE id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	res := &Ability{}
	data := []byte{}
	row := stmt.QueryRow(id)

	if err := row.Scan(&data); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, errors.Wrap(err, "unmarshal data")
	}

	return res, nil
}

func (s *Service) SaveAbility(ability *Ability) (int, error) {
	stmt, err := s.db.Prepare(`INSERT INTO abilities (data) VALUES(?)`)
	if err != nil {
		return 0, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(ability)
	if err != nil {
		return 0, errors.Wrap(err, "marshal data")
	}
	res, err := stmt.Exec(data)
	if err != nil {
		return 0, errors.Wrap(err, "execute query")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "last index")
	}
	if _, err := s.db.Exec("UPDATE abilities SET data = json_set(data, '$.id', id) WHERE id = ?", id); err != nil {
		return 0, errors.Wrap(err, "affect id to json")
	}
	return int(id), nil
}

func (s *Service) UpdateAbility(ability *Ability) error {
	stmt, err := s.db.Prepare(`UPDATE abilities SET data = ? WHERE id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	data, err := json.Marshal(ability)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}
	_, err = stmt.Exec(data, ability.ID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) DeleteAbility(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM abilities WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) GetCardAbilities(cardID int) ([]Ability, error) {
	stmt, err := s.db.Prepare("SELECT data FROM card_ability AS l LEFT JOIN abilities AS a ON l.ability_id = a.id WHERE l.card_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(cardID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Ability{}
	for rows.Next() {
		r := Ability{}
		data := []byte{}
		if err := rows.Scan(&data); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		if len(data) == 0 {
			return res, nil
		}
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, errors.Wrap(err, "unmarshal data")
		}
		res = append(res, r)
	}
	return res, nil
}
func (s *Service) AddCardAbility(cardID, abilityID int) error {
	stmt, err := s.db.Prepare(`INSERT INTO card_ability VALUES(?, ?)`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(cardID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
func (s *Service) DeleteCardAbility(cardID, abilityID int) error {
	stmt, err := s.db.Prepare(`DELETE FROM card_ability WHERE card_id = ? AND ability_id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(cardID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) GetWeaponAbilities(id int) ([]Ability, error) {
	stmt, err := s.db.Prepare("SELECT data FROM weapon_ability AS l LEFT JOIN abilities AS a ON l.ability_id = a.id WHERE l.weapon_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Ability{}
	for rows.Next() {
		r := Ability{}
		data := []byte{}
		if err := rows.Scan(&data); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		if len(data) == 0 {
			return res, nil
		}
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, errors.Wrap(err, "unmarshal data")
		}
		res = append(res, r)
	}
	return res, nil
}
func (s *Service) AddWeaponAbility(weaponID, abilityID int) error {
	stmt, err := s.db.Prepare(`INSERT INTO weapon_ability VALUES(?, ?)`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(weaponID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
func (s *Service) DeleteWeaponAbility(weaponID, abilityID int) error {
	stmt, err := s.db.Prepare(`DELETE FROM weapon_ability WHERE weapon_id = ? AND ability_id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(weaponID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

func (s *Service) GetModelAbilities(modelID int) ([]Ability, error) {
	stmt, err := s.db.Prepare("SELECT data FROM model_ability AS l LEFT JOIN abilities AS a ON l.ability_id = a.id WHERE l.model_id = ?")
	if err != nil {
		return nil, errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(modelID)
	if err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	res := []Ability{}
	for rows.Next() {
		r := Ability{}
		data := []byte{}
		if err := rows.Scan(&data); err != nil {
			return nil, errors.Wrap(err, "struct scan")
		}
		if len(data) == 0 {
			return res, nil
		}
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, errors.Wrap(err, "unmarshal data")
		}
		res = append(res, r)
	}
	return res, nil
}
func (s *Service) AddModelAbility(modelID, abilityID int) error {
	stmt, err := s.db.Prepare(`INSERT INTO model_ability VALUES(?, ?)`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(modelID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}
func (s *Service) DeleteModelAbility(modelID, abilityID int) error {
	stmt, err := s.db.Prepare(`DELETE FROM model_ability WHERE model_id = ? AND ability_id = ?`)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(modelID, abilityID)
	if err != nil {
		return errors.Wrap(err, "execute query")
	}
	return nil
}

// func (s *Service) ListMagicAbilities(modelID int) ([]Ability, error) {
// 	stmt, err := s.db.Preparex("SELECT id, type, name, description FROM abilities AS a LEFT JOIN model_ability AS ma ON a.id = ma.ability_id WHERE ma.model_id = ? AND a.type = ?")
// 	if err != nil {
// 		return nil, errors.Wrap(err, "prepare statement")
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Queryx(modelID, 3)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "execute query")
// 	}
// 	res := []Ability{}
// 	for rows.Next() {
// 		r := Ability{}
// 		if err := rows.StructScan(&r); err != nil {
// 			return nil, errors.Wrap(err, "struct scan")
// 		}
// 		res = append(res, r)
// 	}
// 	return res, nil
// }
// func (s *Service) ListWeaponAbilities(weaponID int) ([]Ability, error) {
// 	stmt, err := s.db.Preparex("SELECT id, type, name, description FROM abilities AS a LEFT JOIN weapon_ability AS wa ON a.id = wa.ability_id WHERE wa.weapon_id = ? AND a.type = ?")
// 	if err != nil {
// 		return nil, errors.Wrap(err, "prepare statement")
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Queryx(weaponID, 4)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "execute query")
// 	}
// 	res := []Ability{}
// 	for rows.Next() {
// 		r := Ability{}
// 		if err := rows.StructScan(&r); err != nil {
// 			return nil, errors.Wrap(err, "struct scan")
// 		}
// 		res = append(res, r)
// 	}
// 	return res, nil
// }
