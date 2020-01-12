package generator

import (
	"fmt"

	"github.com/chibimi/cards/card/reference"
	"github.com/pkg/errors"
)

func (g *Generator) PrintCard(ref *reference.Reference) error {
	g.PrintBack(ref, true)

	models, err := g.src.Model.List(ref.ID, g.lang)
	if err != nil {
		return errors.Wrap(err, "get models")
	}

	err = g.PrintAbilities(models)
	if err != nil {
		return errors.Wrap(err, "print abilities")
	}

	return nil
}

func (g *Generator) PrintBack(ref *reference.Reference, rotate bool) {
	if rotate {
		g.pdf.TransformBegin()
		X1 := g.x + CardWidth/2
		Y1 := g.y + CardHeight + 0.1
		g.pdf.TransformRotate(180, X1, Y1)
		defer g.pdf.TransformEnd()
	}

	X, Y := g.x, g.y
	g.pdf.Image(fmt.Sprintf("images/back_%s.png", "wm"), X+2, Y+10.3, CardWidth-4, CardHeight-18, false, "", 0, "")

}

func (g *Generator) PrintSpells(ref *reference.Reference) error {
	g.PrintBack(ref, false)
	return nil
}
func (g *Generator) PrintFeat(ref *reference.Reference) error {
	g.PrintBack(ref, true)
	return nil
}
