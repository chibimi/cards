package generator

import (
	"strconv"

	"github.com/chibimi/cards/card/weapon"
)

func (g *Generator) PrintWeapon(weapon *weapon.Weapon, str string, modelPosition, weaponPosition int) {
	X := g.x + 39.5
	Y := g.y + 22 + 36*float64(modelPosition) + 12*float64(weaponPosition)

	if weapon.Type == 1 {
		g.PrintMeeleWeapon(X, Y, weapon, str)
	} else if weapon.Type == 2 {
		g.PrintRangedWeapon(X, Y, weapon)
	} else if weapon.Type == 3 {
		g.PrintMountWeapon(X, Y, weapon)
	}

	g.PrintAdvantages(X+18.4, Y+8.7, weapon.Advantages)
}

func (g *Generator) PrintMeeleWeapon(X, Y float64, weapon *weapon.Weapon, str string) {
	if weapon.CNT != "" && weapon.CNT != "1" {
		g.pdf.Image("images/meelex.png", X, Y, 23, 10.5, false, "", 0, "")
		g.pdf.SetTextColor(255, 255, 255)
		g.pdf.SetFont("Arial", "B", 5)
		g.pdf.Text(X+3.6, Y+8.4, g.unicode(weapon.CNT))
		g.pdf.SetTextColor(0, 0, 0)
	} else {
		g.pdf.Image("images/meele.png", X, Y, 23, 10.5, false, "", 0, "")
	}
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+1, Y+3.4, g.unicode(weapon.Name))
	g.PrintMeeleStat(X, Y, 0, weapon.RNG)
	g.PrintMeeleStat(X, Y, 1, weapon.POW)
	p, _ := strconv.Atoi(weapon.POW)
	s, _ := strconv.Atoi(str)
	g.PrintMeeleStat(X, Y, 2, strconv.Itoa(p+s))
}
func (g *Generator) PrintMeeleStat(X, Y, index float64, value string) {
	X += 7.5
	Y += 8
	g.pdf.SetFont("Arial", "", 6)
	if len(value) > 1 {
		g.pdf.Text(X+6*index-1.2, Y, value)
	} else {
		g.pdf.Text(X+6*index, Y, value)
	}
}

func (g *Generator) PrintRangedWeapon(X, Y float64, weapon *weapon.Weapon) {
	if weapon.CNT != "" && weapon.CNT != "1" {
		g.pdf.Image("images/rangedx.png", X, Y, 23, 10.5, false, "", 0, "")
		g.pdf.SetTextColor(255, 255, 255)
		g.pdf.SetFont("Arial", "B", 5)
		g.pdf.Text(X+3.5, Y+8.7, g.unicode(weapon.CNT))
		g.pdf.SetTextColor(0, 0, 0)
	} else {
		g.pdf.Image("images/ranged.png", X, Y, 23, 10.5, false, "", 0, "")
	}
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+1, Y+3.4, g.unicode(weapon.Name))

	g.PrintRangedStat(X, Y, 0, weapon.RNG)
	g.PrintRangedStat(X, Y, 1, weapon.ROF)
	g.PrintRangedStat(X, Y, 2, weapon.AOE)
	g.PrintRangedStat(X, Y, 3, weapon.POW)
}
func (g *Generator) PrintRangedStat(X, Y, index float64, value string) {
	X += 7.5
	Y += 8
	g.pdf.SetFont("Arial", "", 6)
	if len(value) > 1 {
		g.pdf.Text(X+4*index-1.2, Y, value)
	} else {
		g.pdf.Text(X+4*index, Y, value)
	}
}

func (g *Generator) PrintMountWeapon(X, Y float64, weapon *weapon.Weapon) {
	g.pdf.Image("images/mount.png", X+5, Y, 18, 10, false, "", 0, "")

	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+6, Y+3.2, g.unicode(weapon.Name))

	g.PrintMountStat(X+5, Y, 0, weapon.RNG)
	g.PrintMountStat(X+5, Y, 1, weapon.POW)
}
func (g *Generator) PrintMountStat(X, Y, index float64, value string) {
	X += 8.2
	Y += 7.5
	g.pdf.SetFont("Arial", "", 6)
	if len(value) > 1 {
		g.pdf.Text(X+5.5*index-1.2, Y, value)
	} else {
		g.pdf.Text(X+5.5*index, Y, value)
	}
}
