package generator

import (
	"fmt"
	"strings"

	"github.com/chibimi/cards/card/reference"
)

func (g *Generator) PrintFront(X, Y float64, ref *reference.Reference, index int) {
	g.pdf.Image(fmt.Sprintf("%s/images/front/%d.png", g.assets, ref.PPID), X, Y, CardWidth, CardHeight, false, "", 0, "")
}
func (g *Generator) PrintBack(X, Y float64, ref *reference.Reference) {
	g.pdf.Image(fmt.Sprintf("%s/images/back/%d.png", g.assets, ref.FactionID), X, Y, CardWidth, CardHeight, false, "", 0, "")
	g.pdf.SetFont("Abilities", "", 7)
	g.pdf.Text(X+9, Y+7, strings.ToUpper(ref.Title))
}
