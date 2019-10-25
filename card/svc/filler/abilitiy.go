package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func loadAbilities(db *sqlx.DB, src string) error {
	abilities, err := openAbilities(src)
	if err != nil {
		return err
	}
	cache := map[string]struct{}{}
	for _, a := range abilities {
		a.Title = strings.ToLower(a.Title)
		if strings.HasPrefix(a.Title, "magic ability") || strings.HasPrefix(a.Title, "attack type") || strings.HasPrefix(a.Title, "battle plan") {
			continue
		}
		if _, ok := cache[a.Title]; ok {
			continue
		}
		res, err := db.Exec("INSERT INTO abilities (title) VALUES (?)", strings.Title(a.Title))
		if err != nil {
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		_, err = db.Exec("INSERT INTO abilities_lang (ability_id, lang, name, description) VALUES (?,?,?,?)", id, "UK", strings.Title(a.Title), a.Text)
		if err != nil {
			return err
		}
		cache[a.Title] = struct{}{}
	}
	return nil
}

func openAbilities(src string) ([]Capacity, error) {
	res := []Capacity{}

	file, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read refs file")
	}
	err = json.Unmarshal(file, &res)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to unmarshal refs file")
	}

	return res, nil
}
