package generator

import (
	"fmt"

	"github.com/chibimi/cards/card"
	"github.com/jung-kurt/gofpdf"
	"github.com/pkg/errors"
)

const X0 float64 = 10.0
const Y0 float64 = 10.0
const CardWidth float64 = 64.0
const CardHeight float64 = 89.0

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
		x:          10,
		y:          10,
		cardIndex:  -1,
	}
}

func (g *Generator) GeneratePDF() error {
	g.pdf = gofpdf.New("L", "mm", "A4", "")
	g.unicode = g.pdf.UnicodeTranslatorFromDescriptor("")
	g.pdf.AddPage()

	for _, id := range g.references {
		g.nextCard()

		ref, err := g.src.Ref.Get(id, g.lang)
		if err != nil {
			return errors.Wrap(err, "get ref")
		}
		err = g.PrintCard(ref)
		if err != nil {
			return errors.Wrap(err, "add card")
		}
	}

	err := g.pdf.OutputFileAndClose("hello_new.pdf")
	fmt.Println(err)
	return errors.Wrap(err, "write output file")
}

func (g *Generator) nextCard() {
	if g.cardIndex == 3 {
		g.nextPage()
		return
	}
	g.cardIndex++
	g.x = 10 + float64(g.cardIndex%2)*CardWidth*2
	g.y = 10 + float64(g.cardIndex/2)*CardHeight
}

func (g *Generator) nextPage() {
	if g.pdf.PageCount() > g.pdf.PageNo() {
		g.pdf.SetPage(g.pdf.PageNo() + 1)
	} else {
		g.pdf.AddPage()
		g.currentPage++
	}
	g.cardIndex = 0
	g.x = 10
	g.y = 10
}
