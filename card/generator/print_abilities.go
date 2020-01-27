package generator

import (
	"fmt"
	"strings"

	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/reference"
	"github.com/chibimi/cards/card/weapon"
)

func (g *Generator) PrintAbilities(X, Y float64, ref *reference.Reference, models []model.Model) error {
	g.PrintBack(X, Y, ref)
	X += 3
	Y += 13

	present := map[int]bool{}
	cache, err := g.buildAbilityCache(models)
	if err != nil {
		return err
	}
	for _, model := range models {
		abilities, err := g.src.Ability.ListByModel(model.ID, g.lang)
		if err != nil {
			return err
		}
		preparedAbilities := g.prepareAbilities(abilities, cache, &present)
		Y += g.PrintModelAbilities(X, Y, model, preparedAbilities)

		weapons, err := g.src.Weapon.List(model.ID, g.lang)
		if err != nil {
			return err
		}
		for _, weapon := range weapons {
			abilities, err := g.src.Ability.ListByWeapon(weapon.ID, g.lang)
			if err != nil {
				return err
			}
			preparedAbilities := g.prepareAbilities(abilities, cache, &present)
			Y += g.PrintWeaponAbilities(X, Y, weapon, preparedAbilities)
		}
	}
	return nil
}

func (g *Generator) PrintModelAbilities(X, Y float64, model model.Model, abilitiesByType map[int][]ability.Ability) float64 {
	startingY := Y

	if len(abilitiesByType) == 0 {
		return 0
	}

	g.pdf.SetFont("Abilities", "BU", 6)
	g.pdf.Text(X, Y, strings.ToUpper(model.Title))
	Y += 3
	for _, ability := range abilitiesByType[0] {
		Y += g.printAbility(X, Y, 60, ability)
	}

	for i, ability := range abilitiesByType[1] {
		if i == 0 {
			g.pdf.SetFont("Abilities", "B", 5)
			g.pdf.Text(X, Y, fmt.Sprintf("Capacités Magiques [%s]", model.MagicAbility))
			Y += 2.5
		}
		g.pdf.Circle(X+1, Y-0.5, 0.4, "F")
		Y += g.printAbility(X+2, Y, 58, ability)
	}

	for i, ability := range abilitiesByType[2] {
		if i == 0 {
			g.pdf.SetFont("Abilities", "B", 5)
			g.pdf.Text(X, Y, "Plans de bataille")
			Y += 2.5
		}
		g.pdf.Circle(X+1, Y-0.5, 0.4, "F")
		Y += g.printAbility(X+2, Y, 58, ability)
	}

	return Y - startingY
}

func (g *Generator) PrintWeaponAbilities(X, Y float64, weapon weapon.Weapon, abilitiesByType map[int][]ability.Ability) float64 {
	startingY := Y

	if len(abilitiesByType) == 0 {
		return 0
	}

	g.pdf.SetFont("Abilities", "BU", 6)
	g.pdf.Text(X, Y, strings.ToUpper(weapon.Title))
	Y += 3
	for _, ability := range abilitiesByType[0] {
		Y += g.printAbility(X, Y, 60, ability)
	}

	for i, ability := range abilitiesByType[3] {
		if i == 0 {
			g.pdf.SetFont("Abilities", "B", 5)
			g.pdf.Text(X, Y, "Attaque typée")
			Y += 2.5
		}
		g.pdf.Circle(X+1, Y-0.5, 0.4, "F")
		Y += g.printAbility(X+2, Y, 58, ability)
	}

	return Y - startingY
}

func (g *Generator) prepareAbilities(abilities []ability.Ability, cache map[int]ability.Ability, present *map[int]bool) map[int][]ability.Ability {
	abilitiesByType := map[int][]ability.Ability{}
	for _, ability := range abilities {
		temp := abilitiesByType[ability.Type]
		ability.Description = g.replaceLinks(ability.Name, ability.Description, cache)
		if _, ok := (*present)[ability.ID]; ok {
			ability.Description = "Voir plus haut"
		} else {
			(*present)[ability.ID] = true
		}
		temp = append(temp, ability)
		abilitiesByType[ability.Type] = temp
	}
	return abilitiesByType
}

func (g *Generator) printAbility(X, Y, w float64, ability ability.Ability) float64 {

	g.pdf.SetFont("Abilities", "", 5)
	lineNb := 0.0
	_, lineHt := g.pdf.GetFontSize()
	data := g.pdf.SplitLines([]byte(fmt.Sprintf("%s (%s)%s: %s", strings.ToUpper(ability.Name), ability.Title, star(ability.Star), ability.Description)), w)
	for _, s := range data {
		g.pdf.Text(X, Y+float64(lineNb)*lineHt, string(s))
		lineNb++
	}
	return 2 + lineNb*lineHt
}

func star(typ int) string {
	switch typ {
	case 1:
		return " (*Attaque)"
	case 2:
		return " (*Action)"
	default:
		return ""
	}
}

func (g *Generator) buildAbilityCache(models []model.Model) (map[int]ability.Ability, error) {
	cache := map[int]ability.Ability{}
	for _, model := range models {
		abilities, err := g.src.Ability.ListByModel(model.ID, g.lang)
		if err != nil {
			return nil, err
		}
		for _, ability := range abilities {
			cache[ability.ID] = ability
		}
		weapons, err := g.src.Weapon.List(model.ID, g.lang)
		if err != nil {
			return nil, err
		}
		for _, weapon := range weapons {
			abilities, err := g.src.Ability.ListByWeapon(weapon.ID, g.lang)
			if err != nil {
				return nil, err
			}
			for _, ability := range abilities {
				cache[ability.ID] = ability
			}
		}
	}
	return cache, nil
}
