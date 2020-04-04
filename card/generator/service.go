package generator

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card"
	"github.com/chibimi/cards/card/ability"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/common/log"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

type Service struct {
	src     *card.SService
	ability *ability.Repository
	assets  string
}

func NewService(cards *card.SService, db *sqlx.DB, assets string) *Service {
	return &Service{
		src:     cards,
		ability: ability.NewRepository(db),
		assets:  assets,
	}
}

func (s *Service) GenerateEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var refs []int
	for _, v := range strings.Split(r.FormValue("cards"), ",") {
		ref, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		refs = append(refs, ref)
	}

	res, err := s.Generate(refs, r.FormValue("lang"))
	if err != nil {
		log15.Error("generating cards", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/pdf")
	_, err = io.Copy(w, res)
	if err != nil {
		log15.Error("sending cards", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) Generate(refs []int, lang string) (io.Reader, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) DisplayEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var ids []int
	for _, v := range strings.Split(r.FormValue("cards"), ",") {
		id, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ids = append(ids, id)
	}
	var lang = r.FormValue("lang")

	// References generally have 1 cards. Casters, Colossals, Character
	// Units & other edge cases are the notable exceptions, but cards
	// always have 2 faces anyway, and we consider each face as an
	// independent card for the purpose of generating them as it's way
	// simpler to handle.
	var cards = make([]Card, 0, len(ids)*2)
	for _, id := range ids {
		r, err := s.Get(id, lang)
		if err != nil {
			http.Error(w, fmt.Sprintf("getting ref %d: %s", id, err), http.StatusInternalServerError)
			return
		}

		c, err := s.Build(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("building ref %d: %s", id, err), http.StatusInternalServerError)
			return
		}

		cards = append(cards, c...)
	}

	t, err := template.New("cards.html").Funcs(template.FuncMap{
		// safe is for printing HTML directly into the template. Might
		// be useful if we use markdown somewhere.
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		// slug takes a string and returns a slug-case version of if
		// (spaces replaced by dashes), which is useful for using
		// strings as class name, like for the multi-word factions
		// names.
		"slug": func(s string) string {
			return strings.Replace(s, " ", "-", -1)
		},
	}).ParseFiles(path.Join(s.assets, "templates/cards.html"))
	if err != nil {
		log.Error("parsing card template", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, cards)
	if err != nil {
		log.Error("generating output", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// wrap an error using the provided message and arguments.
func wrap(err error, msg string, args ...interface{}) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(msg, args...), err)
}
