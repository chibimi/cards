package generator

import (
	"fmt"

	"github.com/chibimi/cards/card"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
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

func (s *Service) GeneratePDF(references []int, lang string) (string, error) {
	id := uuid.New()

	// generator.new
	g := NewGenerator(s.src, references, lang, s.assets)
	g.pdf = gofpdf.New("L", "mm", "letter", "")
	g.unicode = g.pdf.UnicodeTranslatorFromDescriptor("")
	// generator.translate
	err := g.GeneratePDF()
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("%s/%s.pdf", s.assets, id)
	err = g.WritePDF(res)
	if err != nil {
		return "", err
	}

	return res, nil
}
