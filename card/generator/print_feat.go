package generator

import (
	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/feat"
	"github.com/chibimi/cards/card/reference"
)

func (g *Generator) PrintFeat(X, Y float64, ref *reference.Reference, feat *feat.Feat) {
	g.PrintBack(X, Y, ref)
	feat.Description = g.replaceLinks(feat.Description, map[int]ability.Ability{})

	// Feat name
	X += 4
	Y += 15
	g.pdf.SetFont("Abilities", "BU", 6)
	g.pdf.Text(X, Y, feat.Name)

	// Feat description
	Y += 5
	g.pdf.SetFont("Abilities", "", 6)
	lineNb := 0.0
	_, lineHt := g.pdf.GetFontSize()
	data := g.pdf.SplitText(feat.Description, 55)
	for _, s := range data {
		g.pdf.Text(X, Y+lineNb*lineHt, s)
		lineNb++
	}
}
