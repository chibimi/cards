package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func loadSpells(db *sqlx.DB, src string) error {
	spells, err := openSpells(src)
	if err != nil {
		return err
	}

	cache := map[string]struct{}{}
	for _, a := range spells {
		a.Name = strings.ToLower(a.Name)
		if _, ok := cache[a.Name]; ok {
			continue
		}
		res, err := db.Exec("INSERT INTO spells (title,cost,rng,aoe,pow,dur,off) VALUES (?,?,?,?,?,?,?)", strings.Title(a.Name), a.Cost, strings.ToLower(a.Rng), strings.ToLower(a.Aoe), a.Pow, strings.ToLower(a.Duration), strings.ToLower(a.Off))
		if err != nil {
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		_, err = db.Exec("INSERT INTO spells_lang (spell_id, lang, name, description) VALUES (?,?,?,?)", id, "UK", strings.Title(a.Name), a.Text)
		if err != nil {
			return err
		}
		cache[a.Name] = struct{}{}
	}
	return nil
}

func openSpells(src string) ([]Spell, error) {
	res := []Spell{}

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
