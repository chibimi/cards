package generator

import (
	"os"

	"github.com/chibimi/cards/card"
	"github.com/jung-kurt/gofpdf"
	"github.com/pkg/errors"
)

const X0 float64 = 10.0 + 1.9
const Y0 float64 = 10.0 + 9.2
const CardWidth float64 = 63.3
const CardHeight float64 = 88.9
const Separator float64 = 0.55

type Generator struct {
	src         *card.SService
	references  []int
	lang        string
	pdf         *gofpdf.Fpdf
	unicode     func(string) string
	cardIndex   int
	currentPage int
	x, y        float64
}

func NewGenerator(src *card.SService, references []int, lang string) *Generator {
	return &Generator{
		src:        src,
		references: references,
		lang:       lang,
		x:          X0,
		y:          Y0,
		cardIndex:  -1,
	}
}

func (g *Generator) initializePDF(images []string) {
	g.pdf = gofpdf.New("L", "mm", "letter", "")
	g.unicode = g.pdf.UnicodeTranslatorFromDescriptor("")
	for _, v := range images {
		g.pdf.AddPage()
		g.pdf.Image(v, 0, 0, 279.4, 215.9, false, "", 0, "")
		os.Remove(v)
	}
	g.pdf.SetPage(1)
}

func (g *Generator) GeneratePDF() error {
	for _, id := range g.references {
		g.nextCard()
		ref, err := g.src.Ref.Get(id, g.lang)
		if err != nil {
			return errors.Wrap(err, "get ref")
		}

		err = g.PrintCard(ref)
		if err != nil {
			return errors.Wrap(err, "print abilities")
		}

		if ref.CategoryID == 1 || ref.CategoryID == 2 || ref.CategoryID == 10 {
			g.nextCard()
			err = g.PrintSpells(ref)
			if err != nil {
				return errors.Wrap(err, "print spells")
			}
			err = g.PrintFeat(ref)
			if err != nil {
				return errors.Wrap(err, "print feat")
			}
		}
	}

	return nil
}

func (g *Generator) WritePDF(dest string) error {
	return g.pdf.OutputFileAndClose(dest)
}

func (g *Generator) nextCard() {
	if g.cardIndex == 3 {
		g.nextPage()
		return
	}
	g.cardIndex++
	g.x = X0 + float64(g.cardIndex)*(CardWidth+Separator)*1
}

func (g *Generator) nextPage() {
	g.pdf.SetPage(g.pdf.PageNo() + 1)

	g.cardIndex = 0
	g.x = X0
	g.y = Y0
}
