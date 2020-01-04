package generator

import (
	"fmt"

	"github.com/chibimi/cards/card/reference"
	"github.com/pkg/errors"
)

func (g *Generator) PrintCard(ref *reference.Reference) error {
	g.PrintFront(ref)
	g.PrintBack(ref)

	models, err := g.src.Model.List(ref.ID, g.lang)
	if err != nil {
		return errors.Wrap(err, "get models")
	}
	ci, pi, x, y := g.cardIndex, g.pdf.PageNo(), g.x, g.y
	for i, model := range models {
		// If more that 2 models create a new card
		if i != 0 && i%2 == 0 {
			fmt.Println("GENERATE WEW CARD FOR MORE MODEL")
			g.nextCard()
			g.PrintFront(ref)
			g.PrintBack(ref)
		}

		err = g.PrintModel(&model, i%2)
		if err != nil {
			return errors.Wrap(err, "print model")
		}
	}
	g.pdf.SetPage(pi)
	g.cardIndex, g.x, g.y = ci, x, y

	err = g.PrintAbilities(models)
	if err != nil {
		return errors.Wrap(err, "print abilities")
	}

	return nil
}

func (g *Generator) PrintFront(ref *reference.Reference) {
	X, Y := g.x, g.y
	g.pdf.Image(fmt.Sprintf("images/front_%d.png", ref.FactionID), X, Y, CardWidth, CardHeight, false, "", 0, "")

	g.pdf.SetFont("Arial", "", 9)
	g.pdf.Text(X+14, Y+6, g.unicode(ref.Name))

	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+14, Y+8.5, g.unicode(ref.Properties))

	// TODO: print FA
	// TODO: print cost
}

func (g *Generator) PrintBack(ref *reference.Reference) {
	X, Y := g.x+CardWidth, g.y
	g.pdf.Image(fmt.Sprintf("images/back_%d.png", ref.FactionID), X, Y, CardWidth, CardHeight, false, "", 0, "")

	g.pdf.SetFont("Arial", "", 9)
	g.pdf.Text(X+10, Y+7, g.unicode(ref.Name))
}
