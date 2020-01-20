package generator

import (
	"fmt"
	"strings"

	"github.com/chibimi/cards/card/reference"
)

func (g *Generator) PrintFront(X, Y float64, ref *reference.Reference, index int) {
	if index == 0 {
		g.pdf.Image(fmt.Sprintf("%s/images/front/%d_%d.png", g.assets, ref.PPID, index), X, Y, CardWidth, CardHeight, false, "", 0, "")
		return
	}
	g.pdf.Image(fmt.Sprintf("%s/images/front/0_0.png", g.assets), X, Y, CardWidth, CardHeight, false, "", 0, "")
}

func (g *Generator) PrintBack(X, Y float64, ref *reference.Reference) {
	g.pdf.Image(fmt.Sprintf("%s/images/back/%d.png", g.assets, ref.FactionID), X, Y, CardWidth, CardHeight, false, "", 0, "")
	g.pdf.SetFont("Abilities", "", 7)
	g.pdf.Text(X+9, Y+7, strings.ToUpper(ref.Title))
}
