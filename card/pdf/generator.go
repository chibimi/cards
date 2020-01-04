package pdf

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card"
	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/reference"
	"github.com/chibimi/cards/card/weapon"
	"github.com/jung-kurt/gofpdf"
	"github.com/pkg/errors"
)

const X0 float64 = 10.0
const Y0 float64 = 10.0
const CardWidth float64 = 64.0
const CardHeight float64 = 89.0

type Generator struct {
	pdf       *gofpdf.Fpdf
	src       *card.SService
	cardIndex int
	page      int
}

func NewGenerator(src *card.SService) *Generator {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	return &Generator{
		pdf:       pdf,
		src:       src,
		cardIndex: 0,
		page:      0,
	}
}

var src = map[int]*reference.Reference{
	1: &reference.Reference{
		ID:         1,
		FactionID:  11,
		Name:       "Absylonia, Terreur d'Everblight",
		Title:      "Absylonia,Terror of Everblight",
		Properties: "Sorcière corrompue de la légion",
	},
}

func (g *Generator) GeneratePDF(ids []int, lang string) error {
	for _, id := range ids {
		card, err := g.src.Ref.Get(id, lang)
		if err != nil {
			return errors.Wrap(err, "get ref")
		}
		if err := g.AddCard(card, lang); err != nil {
			return errors.Wrap(err, "add card")
		}
		g.cardIndex++

	}
	err := g.pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	return nil
}

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

func (g *Generator) AddCard(card *reference.Reference, lang string) error {
	if err := g.PrintCard(card); err != nil {
		return errors.Wrap(err, "print card")
	}
	models, err := g.src.Model.List(card.ID, lang)
	if err != nil {
		return errors.Wrap(err, "get models")
	}
	abilityCache := map[int]ability.Ability{}
	cardAbilities := map[*model.Model]ModelAbilities{}
	for i, model := range models {
		// If more that 2 models create a new card
		if i != 0 && i%2 == 0 {
			g.PrintAbilities(cardAbilities, abilityCache)
			abilityCache = map[int]ability.Ability{}
			g.cardIndex++
			if err := g.PrintCard(card); err != nil {
				return errors.Wrap(err, "print card")
			}
		}

		modelAbilities := ModelAbilities{
			Normal:    []ability.Ability{},
			Magic:     []ability.Ability{},
			BatlePlan: []ability.Ability{},
			Weapon:    map[*weapon.Weapon]WeaponAbilities{},
		}

		mAbilities, err := g.src.Ability.ListByModel(model.ID, lang)
		if err != nil {
			return errors.Wrap(err, "get model abilities")
		}
		sort.Slice(mAbilities, func(i, j int) bool {
			if mAbilities[i].Type < mAbilities[j].Type {
				return true
			}
			if mAbilities[i].Type > mAbilities[j].Type {
				return false
			}
			return mAbilities[i].Title < mAbilities[j].Title

		})
		for j, ability := range mAbilities {
			if _, ok := abilityCache[ability.ID]; ok {
				mAbilities[j].Description = "Voir plus haut"
			} else {
				abilityCache[ability.ID] = ability
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

		weapons, err := g.src.Weapon.List(model.ID, lang)
		if err != nil {
			return errors.Wrap(err, "get weapons")
		}
		sort.Slice(weapons, func(i, j int) bool {
			if weapons[i].Type == 2 {
				return true
			}
			return weapons[i].Type < weapons[j].Type
		})

		for j, weapon := range weapons {
			g.PrintWeapon(&weapon, model.STR, card.FactionID, i%2, j)
			weaponAbilities := WeaponAbilities{
				Normal:     []ability.Ability{},
				AttackType: []ability.Ability{},
			}

			wAbilities, err := g.src.Ability.ListByWeapon(weapon.ID, lang)
			if err != nil {
				return errors.Wrap(err, "get weapon abilities")
			}
			sort.Slice(wAbilities, func(i, j int) bool {
				if wAbilities[i].Type < wAbilities[j].Type {
					return true
				}
				if wAbilities[i].Type > wAbilities[j].Type {
					return false
				}
				return wAbilities[i].Title < wAbilities[j].Title

			})
			for k, ability := range wAbilities {
				if _, ok := abilityCache[ability.ID]; ok {
					wAbilities[j].Description = "Voir plus haut"
				} else {
					abilityCache[ability.ID] = ability
				}
				switch ability.Type {
				case 0:
					weaponAbilities.Normal = append(weaponAbilities.Normal, wAbilities[k])
				case 3:
					weaponAbilities.AttackType = append(weaponAbilities.AttackType, wAbilities[k])
				}
			}
			modelAbilities.Weapon[&weapons[j]] = weaponAbilities
		}
		cardAbilities[&models[i]] = modelAbilities

		g.PrintStatline(&model, card.FactionID, i%2)
		if len(models)-1 == i || i%2 == 1 {
			g.PrintAbilities(cardAbilities, abilityCache)
			abilityCache = map[int]ability.Ability{}
		}
	}

	// abilities := []Abilities{}
	// for i, model := range models {
	// 	if i != 0 && i%2 == 0 {
	// 		// g.PrintAbilities(abilities)
	// 		// abilities = []Abilities{}
	// 		g.cardIndex++
	// 		if err := g.PrintCard(card); err != nil {
	// 			return errors.Wrap(err, "print card")
	// 		}
	// 	}
	// 		g.PrintStatline(&model, card.FactionID, i%2)

	// 		cardAbilities, err := g.src.ListCardAbilities(card.ID)
	// 		if err != nil {
	// 			return errors.Wrap(err, "get card abilities")
	// 		}
	// 		temp := map[string]string{}
	// 		for _, a := range cardAbilities {
	// 			if contains(abilities, a.Name) {
	// 				temp[a.Name] = "Voir plus haut."
	// 			} else {
	// 				temp[a.Name] = a.Description
	// 			}
	// 		}
	// 		abilities = append(abilities, Abilities{Label: "", Values: temp})

	// 		modelAbilities, err := g.src.ListModelAbilities(model.ID)
	// 		if err != nil {
	// 			return errors.Wrap(err, "get model abilities")
	// 		}
	// 		temp = map[string]string{}
	// 		for _, a := range modelAbilities {
	// 			if contains(abilities, a.Name) {
	// 				temp[a.Name] = "Voir plus haut."
	// 			} else {
	// 				temp[a.Name] = a.Description
	// 			}
	// 		}

	// 		magicAbilities, err := g.src.ListMagicAbilities(model.ID)
	// 		if err != nil {
	// 			return errors.Wrap(err, "get magic abilities")
	// 		}
	// 		tempM := map[string]string{}
	// 		for _, a := range magicAbilities {
	// 			if contains(abilities, a.Name) {
	// 				tempM[a.Name] = "Voir plus haut."
	// 			} else {
	// 				tempM[a.Name] = a.Description
	// 			}
	// 		}
	// 		abilities = append(abilities, Abilities{Label: model.Name, Values: temp, ValuesMagic: tempM, Magic: model.MagicAbility})

	// 		weapons, err := g.src.ListWeapons(model.ID)
	// 		if err != nil {
	// 			return errors.Wrap(err, "get weapons")
	// 		}

	// 		for j, weapon := range weapons {
	// 			g.PrintWeapon(&weapon, model.STR, card.FactionID, i%2, j)
	// 			weaponAbilities, err := g.src.ListWeaponAbilities(weapon.ID)
	// 			if err != nil {
	// 				return errors.Wrap(err, "get weapon abilities")
	// 			}
	// 			fmt.Println("WA", weapon.Name, len(weaponAbilities))
	// 			temp := map[string]string{}
	// 			for _, a := range weaponAbilities {
	// 				if contains(abilities, a.Name) {
	// 					temp[a.Name] = "Voir plus haut."
	// 				} else {
	// 					temp[a.Name] = a.Description
	// 				}
	// 			}
	// 			abilities = append(abilities, Abilities{Label: weapon.Name, Values: temp})
	// 		}

	// 		if len(models)-1 == i || i%2 == 1 {
	// 			g.PrintAbilities(abilities)
	// 			abilities = []Abilities{}
	// 		}
	// 		// abilityCursor, abilities = g.PrintAbilities(&model, 1, abilityCursor, abilities)
	// 		// if model.MagicAbility != "" {
	// 		// 	g.pdf.SetFont("Arial", "", 8)
	// 		// 	g.pdf.Text(X0+float64(g.cardIndex%2)*CardWidth*2, abilityCursor, fmt.Sprintf("Score de Magie [%s]", model.MagicAbility))
	// 		// 	abilityCursor, abilities = g.PrintAbilities(&model, 2, abilityCursor, abilities)
	// 		// }

	// 	}
	return nil
}

// type Abilities struct {
// 	Label       string
// 	Values      map[string]string
// 	Magic       string
// 	ValuesMagic map[string]string
// }

// func contains(abilities []Abilities, ability string) bool {
// 	for _, v := range abilities {
// 		if _, ok := v.Values[ability]; ok {
// 			return true
// 		}
// 		if _, ok := v.ValuesMagic[ability]; ok {
// 			return true
// 		}
// 	}
// 	return false
// }

func (g *Generator) PrintCard(card *reference.Reference) error {
	if g.cardIndex > 3 {
		g.pdf.AddPage()
		g.page++
		g.cardIndex = 0
	}
	frontX := X0 + float64(g.cardIndex%2)*CardWidth*2
	frontY := Y0 + float64(g.cardIndex/2)*CardHeight
	backX := X0 + float64(g.cardIndex%2)*CardWidth*2 + CardWidth
	backY := Y0 + float64(g.cardIndex/2)*CardHeight
	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	g.pdf.ImageOptions(fmt.Sprintf("images/front_%d.png", card.FactionID), frontX, frontY, CardWidth, CardHeight, false, opt, 0, "")
	g.pdf.ImageOptions(fmt.Sprintf("images/back_%d.png", card.FactionID), backX, backY, CardWidth, CardHeight, false, opt, 0, "")

	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
	g.pdf.SetFont("Arial", "", 9)
	g.pdf.Text(frontX+14, frontY+6, tr(card.Name))
	g.pdf.Text(frontX+CardWidth+10, frontY+7.5, tr(card.Name))
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(frontX+14, frontY+8.5, tr(card.Properties))
	return nil
}

func (g *Generator) PrintWeapon(weapon *weapon.Weapon, str string, faction, i, j int) {
	X := X0 + float64(g.cardIndex%2)*CardWidth*2 + 39.5
	Y := Y0 + float64(g.cardIndex/2)*CardHeight + 22 + 36*float64(i) + 12*float64(j)

	if weapon.Type == 1 {
		g.PrintMeeleWeapon(X, Y, weapon, str)
	} else if weapon.Type == 2 {
		g.PrintRangedWeapon(X, Y, weapon)
	} else if weapon.Type == 3 {
		g.PrintMountWeapon(X, Y, weapon)
	}

	g.PrintAdvantages(X+18.4, Y+8.7, weapon.Advantages)
}

func (g *Generator) PrintStatline(model *model.Model, faction, i int) {
	X := X0 + float64(g.cardIndex%2)*CardWidth*2 + 28.7
	Y := Y0 + float64(g.cardIndex/2)*CardHeight + 10.3
	if i == 1 {
		Y += 36
	}
	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	g.pdf.ImageOptions(fmt.Sprintf("images/statline.png"), X, Y, 34, 11.7, false, opt, 0, "")

	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+2.3, Y+2.6, strings.ToUpper(tr(model.Title)))
	g.pdf.SetFont("Arial", "", 7)
	g.PrintStat(X, Y, 0, model.SPD)
	g.PrintStat(X, Y, 1, model.STR)
	g.PrintStat(X, Y, 2, model.MAT)
	g.PrintStat(X, Y, 3, model.RAT)
	g.PrintStat(X, Y, 4, model.DEF)
	g.PrintStat(X, Y, 5, model.ARM)
	g.PrintStat(X, Y, 6, model.CMD)

	g.pdf.Text(X+29.8, Y+10.5, model.BaseSize)
	g.PrintAdvantages(X+28.8, Y+7.6, model.Advantages)
}

func (g *Generator) PrintStat(X, Y, index float64, value string) {
	X += 4
	Y += 7
	if len(value) > 1 {
		g.pdf.Text(X+4.3*index-1, Y, value)
	} else {
		g.pdf.Text(X+4.3*index, Y, value)
	}
}

func (g *Generator) PrintAdvantages(X, Y float64, advantages []string) {
	sort.Slice(advantages, func(i, j int) bool { return advantages[i] > advantages[j] })
	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	for i, a := range advantages {
		g.pdf.ImageOptions(fmt.Sprintf("images/advantages/%s.png", a), X-(4.2*float64(i)), Y, 4.2, 4.2, false, opt, 0, "")
	}
}

func (g *Generator) PrintMeeleWeapon(X, Y float64, weapon *weapon.Weapon, str string) {
	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

	if weapon.CNT != "" && weapon.CNT != "1" {
		g.pdf.ImageOptions(fmt.Sprintf("images/meelex.png"), X, Y, 23, 10.5, false, opt, 0, "")
		g.pdf.SetTextColor(255, 255, 255)
		g.pdf.SetFont("Arial", "B", 5)
		g.pdf.Text(X+3.6, Y+8.4, tr(weapon.CNT))
		g.pdf.SetTextColor(0, 0, 0)
	} else {
		g.pdf.ImageOptions(fmt.Sprintf("images/meele.png"), X, Y, 23, 10.5, false, opt, 0, "")
	}
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+1, Y+3.4, tr(weapon.Name))
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
	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

	if weapon.CNT != "" && weapon.CNT != "1" {
		g.pdf.ImageOptions(fmt.Sprintf("images/rangedx.png"), X, Y, 23, 10.5, false, opt, 0, "")
		g.pdf.SetTextColor(255, 255, 255)
		g.pdf.SetFont("Arial", "B", 5)
		g.pdf.Text(X+3.5, Y+8.7, tr(weapon.CNT))
		g.pdf.SetTextColor(0, 0, 0)
	} else {
		g.pdf.ImageOptions(fmt.Sprintf("images/ranged.png"), X, Y, 23, 10.5, false, opt, 0, "")
	}
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+1, Y+3.4, tr(weapon.Name))

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
	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

	g.pdf.ImageOptions(fmt.Sprintf("images/mount.png"), X+5, Y, 18, 10, false, opt, 0, "")

	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+6, Y+3.2, tr(weapon.Name))

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

func (g *Generator) PrintAbilities(modelAbilities map[*model.Model]ModelAbilities, cache map[int]ability.Ability) {
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
			fmt.Println(validLink.FindAllString(abi.Description, -1))
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
				fmt.Println(a)
				subskills = append(subskills, fmt.Sprintf("%s (%s): %s", a.Name, a.Title, a.Description))
				cache[a.ID] = *a
				return a.Name

				// return s[1][:len(s[1])-1]
			})
			if len(subskills) > 0 {
				abilities.Normal[j].Description = fmt.Sprintf("%s (%s)", abi.Description, strings.Join(subskills, ", "))

			}
			fmt.Println(abi.Description)

		}

	}

	for model, abilities := range modelAbilities {
		fmt.Println("MODEL", model.Title)
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
			fmt.Println("WEAPON", weapon.Title)
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

// func (g *Generator) PrintAbilities(abilities []Abilities) {
// 	X := X0 + float64(g.cardIndex%2)*CardWidth*2 + CardWidth + 3
// 	Y := Y0 + float64(g.cardIndex/2)*CardHeight + 13
// 	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

// 	lineNb := 0.0
// 	g.pdf.SetFont("Arial", "", 6)
// 	_, lineHt := g.pdf.GetFontSize()

// 	for _, a := range abilities {
// 		if a.Label != "" && (len(a.Values) > 0 || len(a.ValuesMagic) > 0) {
// 			g.pdf.SetFont("Arial", "BU", 6)
// 			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(strings.ToUpper(a.Label)))
// 			lineNb += 1.5
// 		}
// 		for k, v := range a.Values {
// 			g.pdf.SetFont("Arial", "", 5)
// 			data := g.pdf.SplitLines([]byte(fmt.Sprintf("%s - %s", strings.ToUpper(k), v)), 62)
// 			for _, s := range data {
// 				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
// 				lineNb++
// 			}
// 			lineNb += 0.5
// 		}
// 		if a.Magic != "" {
// 			g.pdf.SetFont("Arial", "B", 6)
// 			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(fmt.Sprintf("Précision Magique [%s]", a.Magic)))
// 			lineNb += 1.5
// 		}
// 		for k, v := range a.ValuesMagic {
// 			g.pdf.SetFont("Arial", "", 5)
// 			data := g.pdf.SplitLines([]byte(fmt.Sprintf("\u2022 %s - %s", strings.ToUpper(k), v)), 62)
// 			for _, s := range data {
// 				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
// 				lineNb++
// 			}
// 			lineNb += 0.5
// 		}
// 	}
// }
