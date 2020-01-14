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

	X := X0 + float64(g.cardIndex)*(CardWidth+Separator) + 3
	Y := Y0 + 13

	spells, err := g.src.Spell.ListByRef(ref.ID, g.lang)
	if err != nil {
		return errors.Wrap(err, "list spells")
	}

	if len(spells) > 0 {
		g.pdf.Image(fmt.Sprintf("images/spell.png"), X-1.35, Y-2.5, CardWidth-3.3, 3, false, "", 0, "")
		Y += 4
	}
	for _, spell := range spells {
		Y += g.PrintSpell(X, Y, spell)
	}

	return nil
}
func (g *Generator) PrintFeat(ref *reference.Reference) error {
	g.PrintBack(ref, true)
	g.pdf.TransformBegin()
	X1 := g.x + CardWidth/2
	Y1 := g.y + CardHeight + 0.1
	g.pdf.TransformRotate(180, X1, Y1)
	defer g.pdf.TransformEnd()

	feat, err := g.src.Feat.Get(ref.ID, g.lang)
	if err != nil {
		return errors.Wrap(err, "get feat")
	}
	X := X0 + float64(g.cardIndex)*(CardWidth+Separator) + 3
	Y := Y0 + 13

	g.pdf.SetFont("Arial", "BU", 6)

	g.pdf.Text(X+10, Y+5, g.unicode(feat.Name))
	g.pdf.SetFont("Arial", "", 6)
	lineNb := 0.0
	_, lineHt := g.pdf.GetFontSize()
	data := g.pdf.SplitText(feat.Description, 50)
	for _, s := range data {
		g.pdf.Text(X+5, Y+10+lineNb*lineHt, g.unicode(s))
		lineNb++
	}
	return nil
}
