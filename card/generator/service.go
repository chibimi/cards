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
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/common/log"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

type Service struct {
	src    *card.SService
	assets string
}

func NewService(cards *card.SService, assets string) *Service {
	return &Service{
		src:    cards,
		assets: assets,
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
	var refs []int
	for _, v := range strings.Split(r.FormValue("cards"), ",") {
		ref, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		refs = append(refs, ref)
	}

	var lang = r.FormValue("lang")

	var cards = make([]Card, 0, len(refs)*2)
	for _, ref := range refs {
		c, err := s.Build(ref, lang)
		if err != nil {
			http.Error(w, fmt.Sprintf("building ref %d: %s", ref, err), http.StatusInternalServerError)
			return
		}
		cards = append(cards, c...)
	}

	t, err := template.New("cards.html").Funcs(template.FuncMap{
		"safe": func(html string) template.HTML {
			return template.HTML(html)
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

func (s *Service) Build(ref int, _ string) ([]Card, error) {
	return []Card{
		ProfileCard{
			Faction: 1,
		},
		RulesCard{
			Faction: 1,
			Header:  `<p>Comic Duet</p>`,
			Section: `
				<p class="title">Comic Duet</p>
				<p class="ability"><strong>Two types of comic</strong> Each time this model makes an attack, your can choose of the following special rules:</p>
				<ul class="attack_type">
					<li><strong>Heavy Comic</strong> Add one dice to the joke damage roll.</li>
					<li><strong>Repeat Comic</strong> Add one <em>repeat token</em> to the target model. When a model with <em>repeat tokens</em> suffers a joke damage roll, add +1 to the roll per <em>repeat token</em> on the model, up to +5.</li>
				</ul>
			`,
		},
	}, nil
}

type Card interface {
	Type() string
	Background() string
}

type ProfileCard struct {
	Faction int
}

func (c ProfileCard) Background() string {
	return strconv.Itoa(c.Faction)
}

func (ProfileCard) Type() string {
	return "profile"
}

type RulesCard struct {
	Faction int
	Header  string
	Section string
}

func (RulesCard) Type() string {
	return "rules"
}

func (c RulesCard) Background() string {
	return strconv.Itoa(c.Faction)
}
