package generator

import (
	"fmt"
	"strings"

	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/reference"
	"github.com/chibimi/cards/card/spell"
)

func (g *Generator) PrintSpells(X, Y float64, ref *reference.Reference, spells []spell.Spell) {
	g.PrintBack(X, Y, ref)

	// Spell header
	if len(spells) > 0 {
		g.pdf.Image(fmt.Sprintf("%s/images/icon/spell.png", g.assets), X+1.55, Y+11, 60, 3, false, "", 0, "")
	}
	X += 3
	Y += 16

	// Spells
	for _, spell := range spells {
		spell.Description = g.replaceLinks(spell.Description, map[int]ability.Ability{})
		Y += g.PrintSpell(X, Y, spell)
	}
}

func (g *Generator) PrintSpell(X, Y float64, spell spell.Spell) float64 {
	// Name and cost
	g.pdf.SetFont("Arial", "B", 5)
	g.pdf.Text(X, Y, g.unicode(strings.ToUpper(spell.Name)))
	g.pdf.SetFont("Arial", "", 5)
	g.pdf.Text(X, Y+1.5, fmt.Sprintf("(%s)", spell.Title))
	g.pdf.SetFont("Arial", "B", 5)
	g.pdf.Text(X+30, Y, spell.Cost)
	g.pdf.Text(X+35, Y, spell.RNG)
	g.pdf.Text(X+40, Y, spell.AOE)
	g.pdf.Text(X+45, Y, spell.POW)
	g.pdf.Text(X+50, Y, spell.DUR)
	g.pdf.Text(X+55, Y, spell.OFF)

	// Spell description
	Y += 3.4
	g.pdf.SetFont("Abilities", "", 5)
	lineNb := 0.0
	_, lineHt := g.pdf.GetFontSize()
	data := g.pdf.SplitText(spell.Description, 60)
	for _, s := range data {
		g.pdf.Text(X, Y+lineNb*lineHt, s)
		lineNb++
	}
	return 4.5 + lineNb*lineHt

}
