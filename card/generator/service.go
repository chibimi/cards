package generator

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card"
	"github.com/julienschmidt/httprouter"
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
	_, err := io.Copy(w, res)
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

	for _, ref := range refs {
		c, err := s.Build(ref)
		if err != nil {
			return nil, fmt.Errorf("building ref %d: %w", ref, err)
		}
		cards = append(cards, c...)
	}

	return nil, nil
}
