package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chibimi/cards/card/model"
	"github.com/pkg/errors"
)

func (g *Generator) PrintModel(model *model.Model, position int) error {
	g.PrintStatline(model, position)
	weapons, err := g.src.Weapon.List(model.ID, g.lang)
	if err != nil {
		return errors.Wrap(err, "get models")
	}
	for i := len(weapons) - 1; i >= 0; i-- {
		g.PrintWeapon(&weapons[i], model.STR, position, i)

	}
	return nil
}

func (g *Generator) PrintStatline(model *model.Model, i int) {
	X := g.x + 28.7
	Y := g.y + 10.3 + 36*float64(i)

	g.pdf.Image("images/statline.png", X, Y, 34, 11.7, false, "", 0, "")

	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+2.3, Y+2.6, strings.ToUpper(g.unicode(model.Title)))

	g.pdf.SetFont("Arial", "", 7)
	g.PrintStat(X, Y, 0, model.SPD)
	g.PrintStat(X, Y, 1, model.STR)
	g.PrintStat(X, Y, 2, model.MAT)
	g.PrintStat(X, Y, 3, model.RAT)
	g.PrintStat(X, Y, 4, model.DEF)
	g.PrintStat(X, Y, 5, model.ARM)
	g.PrintStat(X, Y, 6, model.CMD)

	g.pdf.Text(X+29.8, Y+10.5, model.BaseSize)
	g.PrintAdvantages(X+24.8, Y+7.6, model.Advantages)
}

func (g *Generator) PrintStat(X, Y, position float64, value string) {
	X += 4
	Y += 7
	if len(value) > 1 {
		g.pdf.Text(X+4.3*position-1, Y, value)
	} else {
		g.pdf.Text(X+4.3*position, Y, value)
	}
}

func (g *Generator) PrintAdvantages(X, Y float64, advantages []string) {
	sort.Slice(advantages, func(i, j int) bool { return advantages[i] > advantages[j] })

	for i, a := range advantages {
		g.pdf.Image(fmt.Sprintf("images/advantages/%s.png", a), X-(4.2*float64(i)), Y, 4.2, 4.2, false, "", 0, "")
	}
}
