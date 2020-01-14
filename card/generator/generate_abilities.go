package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/reference"
	"github.com/chibimi/cards/card/spell"
	"github.com/chibimi/cards/card/weapon"
	"github.com/pkg/errors"
)

type ModelAbilities struct {
	Spell     []spell.Spell
	Normal    []ability.Ability
	Magic     []ability.Ability
	BatlePlan []ability.Ability
	Weapon    map[*weapon.Weapon]WeaponAbilities
}

type WeaponAbilities struct {
	Normal     []ability.Ability
	AttackType []ability.Ability
}

func (g *Generator) PrintAbilities(models []model.Model) error {
	// Used to not print 2 times a description
	abilityCache := map[int]ability.Ability{}

	// Map of all the abilities to print on a card for models and weapons
	cardAbilities := map[*model.Model]ModelAbilities{}

	for i, m := range models {
		// If more that 2 models fill the current card and create a new one
		if i != 0 && i%2 == 0 {
			g.printAbilities(cardAbilities, abilityCache)
			abilityCache = map[int]ability.Ability{}
			cardAbilities = map[*model.Model]ModelAbilities{}
			g.nextCard()
			g.PrintBack(&reference.Reference{}, true)
		}

		// Fill the abilities map and the cache with the model abilitie
		modelAbilities, abilityCache, err := g.getModelAbilities(m.ID, m.RefID, abilityCache)
		if err != nil {
			return errors.Wrap(err, "get model abilities")
		}

		// Fill the abilities map and the cache with the weapon abilities for this model
		modelAbilities.Weapon, abilityCache, err = g.getWeaponsAbilities(m.ID, abilityCache)
		if err != nil {
			return errors.Wrap(err, "get weapons abilities")
		}
		cardAbilities[&models[i]] = *modelAbilities

		if len(models)-1 == i || i%2 == 1 {
			g.printAbilities(cardAbilities, abilityCache)
		}
	}
	return nil
}

func (g *Generator) printAbilities(modelAbilities map[*model.Model]ModelAbilities, cache map[int]ability.Ability) {

	X := X0 + float64(g.cardIndex)*(CardWidth+Separator) + 3
	Y := Y0 + 13
	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

	lineNb := 0.0
	g.pdf.SetFont("Arial", "", 6)
	_, lineHt := g.pdf.GetFontSize()

	var validLink = regexp.MustCompile(`#[0-9]+:[^#]+#`)

	for _, abilities := range modelAbilities {
		//FIXME: do samething for spells, weapon, magic....
		for j, abi := range abilities.Normal {
			subskills := []string{}
			abi.Description = validLink.ReplaceAllStringFunc(abi.Description, func(src string) string {
				s := strings.SplitN(src, ":", 2)
				ids := s[0][1:]
				id, err := strconv.Atoi(ids)
				if err != nil {
					return s[1][:len(s[1])-1]
				}
				if a, ok := cache[id]; ok {
					return a.Name
				}
				// fetch ability
				a, err := g.src.Ability.Get(id, "FR")
				if err != nil {
					return s[1][:len(s[1])-1]
				}
				subskills = append(subskills, fmt.Sprintf("%s (%s): %s", a.Name, a.Title, a.Description))
				cache[a.ID] = *a
				return a.Name

				// return s[1][:len(s[1])-1]
			})
			if len(subskills) > 0 {
				abilities.Normal[j].Description = fmt.Sprintf("%s (%s)", abi.Description, strings.Join(subskills, ", "))

			}

		}

	}

	g.pdf.TransformBegin()
	X1 := g.x + CardWidth/2
	Y1 := g.y + CardHeight + 0.1
	fmt.Println("A", g.pdf.PageNo(), X1, Y1)
	g.pdf.TransformRotate(180, X1, Y1)

	for model, abilities := range modelAbilities {
		// if len(abilities.Spell) > 0 {
		// 	g.pdf.Image(fmt.Sprintf("images/spell.png"), X-1.35, Y-2.5, CardWidth-3.3, 3, false, "", 0, "")
		// 	Y += 4
		// }
		// for _, spell := range abilities.Spell {
		// 	Y += g.PrintSpell(X, Y, spell)
		// }
		// if len(abilities.Spell) > 0 {
		// 	g.pdf.Line(X, Y-3, X+57, Y-3)
		// }
		if len(abilities.Normal) > 0 || len(abilities.Magic) > 0 || len(abilities.BatlePlan) > 0 {
			g.pdf.SetFont("Arial", "BU", 6)
			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(strings.ToUpper(model.Title)))
			lineNb += 1.5
		}

		// for _, ability := range abilities.Spell {
		// 	g.pdf.SetFont("Arial", "", 5)
		// 	data := g.pdf.SplitLines([]byte(fmt.Sprintf("%s (%s) : %s", strings.ToUpper(ability.Name), ability.Title, ability.Description)), 60)
		// 	for _, s := range data {
		// 		g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
		// 		lineNb++
		// 	}
		// 	lineNb += 0.5
		// }
		for _, ability := range abilities.Normal {
			g.pdf.SetFont("Arial", "", 5)
			data := g.pdf.SplitLines([]byte(fmt.Sprintf("%s (%s) : %s", strings.ToUpper(ability.Name), ability.Title, ability.Description)), 60)
			for _, s := range data {
				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
				lineNb++
			}
			lineNb += 0.5
		}
		if len(abilities.Magic) > 0 {
			g.pdf.SetFont("Arial", "B", 6)
			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(fmt.Sprintf("Capacités Magiques [%s]", model.MagicAbility)))
			lineNb += 1.5
		}
		for _, ability := range abilities.Magic {
			g.pdf.SetFont("Arial", "", 5)
			data := g.pdf.SplitLines([]byte(fmt.Sprintf("    %s (%s) : %s", strings.ToUpper(ability.Name), ability.Title, ability.Description)), 60)
			for _, s := range data {
				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
				lineNb++
			}
			lineNb += 0.5
		}
		if len(abilities.BatlePlan) > 0 {
			g.pdf.SetFont("Arial", "B", 6)
			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(fmt.Sprintf("Plans de bataille")))
			lineNb += 1.5
		}
		for _, ability := range abilities.BatlePlan {
			g.pdf.SetFont("Arial", "", 5)
			data := g.pdf.SplitLines([]byte(fmt.Sprintf("    %s (%s) : %s", strings.ToUpper(ability.Name), ability.Title, ability.Description)), 60)
			for _, s := range data {
				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
				lineNb++
			}
			lineNb += 0.5
		}

		for weapon, abilities := range abilities.Weapon {
			if len(abilities.Normal) > 0 || len(abilities.AttackType) > 0 {
				g.pdf.SetFont("Arial", "BU", 6)
				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(strings.ToUpper(weapon.Title)))
				lineNb += 1.5
			}

			for _, ability := range abilities.Normal {
				g.pdf.SetFont("Arial", "", 5)
				data := g.pdf.SplitLines([]byte(fmt.Sprintf("%s (%s) : %s", strings.ToUpper(ability.Name), ability.Title, ability.Description)), 60)
				for _, s := range data {
					g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))

					lineNb++
				}
				lineNb += 0.5
			}
			if len(abilities.AttackType) > 0 {
				g.pdf.SetFont("Arial", "B", 6)
				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(fmt.Sprintf("Type d'attaque")))
				lineNb += 1.5
			}
			for _, ability := range abilities.AttackType {
				g.pdf.SetFont("Arial", "", 5)
				data := g.pdf.SplitLines([]byte(fmt.Sprintf("    %s (%s) : %s", strings.ToUpper(ability.Name), ability.Title, ability.Description)), 60)
				for _, s := range data {
					g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
					lineNb++
				}
				lineNb += 0.5
			}
		}
	}
	g.pdf.TransformEnd()
}

func (g *Generator) PrintSpell(X, Y float64, spell spell.Spell) float64 {
	g.pdf.SetFont("Arial", "B", 5)
	g.pdf.Text(X, Y-1.5, g.unicode(fmt.Sprintf("%s", strings.ToUpper(spell.Name))))
	g.pdf.SetFont("Arial", "", 5)
	g.pdf.Text(X, Y+0.5, g.unicode(fmt.Sprintf("(%s)", spell.Title)))
	g.pdf.SetFont("Arial", "B", 5)
	g.pdf.Text(X+30, Y-1.5, g.unicode(spell.Cost))
	g.pdf.Text(X+35, Y-1.5, g.unicode(spell.RNG))
	g.pdf.Text(X+40, Y-1.5, g.unicode(spell.AOE))
	g.pdf.Text(X+45, Y-1.5, g.unicode(spell.POW))
	g.pdf.Text(X+50, Y-1.5, g.unicode(spell.DUR))
	g.pdf.Text(X+55, Y-1.5, g.unicode(spell.OFF))

	g.pdf.SetFont("Arial", "", 5)
	lineNb := 0.0
	_, lineHt := g.pdf.GetFontSize()
	data := g.pdf.SplitText(spell.Description, 60)
	for _, s := range data {
		g.pdf.Text(X, Y+3+lineNb*lineHt, g.unicode(s))
		lineNb++
	}
	return 5.5 + lineNb*lineHt

}

func (g *Generator) getModelAbilities(id, refId int, cache map[int]ability.Ability) (*ModelAbilities, map[int]ability.Ability, error) {
	modelAbilities := &ModelAbilities{
		Spell:     []spell.Spell{},
		Normal:    []ability.Ability{},
		Magic:     []ability.Ability{},
		BatlePlan: []ability.Ability{},
		Weapon:    map[*weapon.Weapon]WeaponAbilities{},
	}
	spells, err := g.src.Spell.ListByRef(refId, g.lang)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get model spells")
	}
	for j, _ := range spells {
		modelAbilities.Spell = append(modelAbilities.Spell, spells[j])
	}

	mAbilities, err := g.src.Ability.ListByModel(id, g.lang)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get model abilities")
	}
	for j, ability := range mAbilities {
		if _, ok := cache[ability.ID]; ok {
			mAbilities[j].Description = "Voir plus haut"
		} else {
			cache[ability.ID] = ability
		}
		switch ability.Type {
		case 0:
			modelAbilities.Normal = append(modelAbilities.Normal, mAbilities[j])
		case 1:
			modelAbilities.Magic = append(modelAbilities.Magic, mAbilities[j])
		case 2:
			modelAbilities.BatlePlan = append(modelAbilities.BatlePlan, mAbilities[j])
		}
	}
	return modelAbilities, cache, nil
}

func (g *Generator) getWeaponsAbilities(modelID int, cache map[int]ability.Ability) (map[*weapon.Weapon]WeaponAbilities, map[int]ability.Ability, error) {
	res := map[*weapon.Weapon]WeaponAbilities{}
	weapons, err := g.src.Weapon.List(modelID, g.lang)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get weapons")
	}
	for j, weapon := range weapons {
		var weaponAbilities *WeaponAbilities
		weaponAbilities, cache, err = g.getWeaponAbilities(weapon.ID, cache)
		if err != nil {
			return nil, nil, errors.Wrap(err, "get weapon abilities")
		}
		res[&weapons[j]] = *weaponAbilities
	}
	return res, cache, nil
}

func (g *Generator) getWeaponAbilities(id int, cache map[int]ability.Ability) (*WeaponAbilities, map[int]ability.Ability, error) {
	weaponAbilities := &WeaponAbilities{
		Normal:     []ability.Ability{},
		AttackType: []ability.Ability{},
	}

	wAbilities, err := g.src.Ability.ListByWeapon(id, g.lang)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get weapon abilities")
	}
	for k, ability := range wAbilities {
		if _, ok := cache[ability.ID]; ok {
			wAbilities[k].Description = "Voir plus haut"
		} else {
			cache[ability.ID] = ability
		}
		switch ability.Type {
		case 0:
			weaponAbilities.Normal = append(weaponAbilities.Normal, wAbilities[k])
		case 3:
			weaponAbilities.AttackType = append(weaponAbilities.AttackType, wAbilities[k])
		}
	}
	return weaponAbilities, cache, nil
}
