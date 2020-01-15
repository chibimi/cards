package generator

import (
	"fmt"

	"github.com/chibimi/cards/card"
	"github.com/jung-kurt/gofpdf"
)

type Service struct {
	src *card.SService
}

func NewService(cards *card.SService) *Service {
	return &Service{
		src: cards,
	}
}

func (s *Service) GeneratePDF(references []int, lang string) (string, error) {
	// id := uuid.New()

	// generator.new
	g := NewGenerator(s.src, references, lang)
	g.pdf = gofpdf.New("L", "mm", "letter", "")
	g.unicode = g.pdf.UnicodeTranslatorFromDescriptor("")
	// generator.translate
	err := g.GeneratePDF()
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("%s.pdf", "hello_new")
	err = g.WritePDF(res)
	if err != nil {
		return "", err
	}

	return res, nil
}
