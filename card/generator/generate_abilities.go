package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/weapon"
	"github.com/pkg/errors"
)

type ModelAbilities struct {
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
	abilityCache := map[int]ability.Ability{}
	cardAbilities := map[*model.Model]ModelAbilities{}
	for i, m := range models {
		// If more that 2 models create a new card
		if i != 0 && i%2 == 0 {
			g.printAbilities(cardAbilities, abilityCache)
			abilityCache = map[int]ability.Ability{}
			cardAbilities = map[*model.Model]ModelAbilities{}
			g.nextCard()

		}
		modelAbilities, abilityCache, err := g.getModelAbilities(m.ID, abilityCache)
		if err != nil {
			return errors.Wrap(err, "get model abilities")
		}

		weapons, err := g.src.Weapon.List(m.ID, g.lang)
		if err != nil {
			return errors.Wrap(err, "get weapons")
		}
		for j, weapon := range weapons {
			var weaponAbilities *WeaponAbilities
			weaponAbilities, abilityCache, err = g.getWeaponAbilities(weapon.ID, abilityCache)
			if err != nil {
				return errors.Wrap(err, "get weapon abilities")
			}
			modelAbilities.Weapon[&weapons[j]] = *weaponAbilities
		}
		cardAbilities[&models[i]] = *modelAbilities

		if len(models)-1 == i || i%2 == 1 {
			g.printAbilities(cardAbilities, abilityCache)
		}
	}
	return nil
}

func (g *Generator) printAbilities(modelAbilities map[*model.Model]ModelAbilities, cache map[int]ability.Ability) {
	X := X0 + float64(g.cardIndex%2)*CardWidth*2 + CardWidth + 3
	Y := Y0 + float64(g.cardIndex/2)*CardHeight + 13
	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

	lineNb := 0.0
	g.pdf.SetFont("Arial", "", 6)
	_, lineHt := g.pdf.GetFontSize()

	var validLink = regexp.MustCompile(`#[0-9]+:[^#]+#`)

	for _, abilities := range modelAbilities {
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

	for model, abilities := range modelAbilities {
		if len(abilities.Normal) > 0 || len(abilities.Magic) > 0 || len(abilities.BatlePlan) > 0 {
			g.pdf.SetFont("Arial", "BU", 6)
			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(strings.ToUpper(model.Title)))
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
}

func (g *Generator) getModelAbilities(id int, cache map[int]ability.Ability) (*ModelAbilities, map[int]ability.Ability, error) {
	modelAbilities := &ModelAbilities{
		Normal:    []ability.Ability{},
		Magic:     []ability.Ability{},
		BatlePlan: []ability.Ability{},
		Weapon:    map[*weapon.Weapon]WeaponAbilities{},
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
