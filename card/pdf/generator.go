package pdf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card"
	"github.com/jung-kurt/gofpdf"
	"github.com/pkg/errors"
)

const X0 float64 = 10.0
const Y0 float64 = 10.0
const CardWidth float64 = 64.0
const CardHeight float64 = 89.0

type Generator struct {
	pdf       *gofpdf.Fpdf
	src       *card.Service
	cardIndex int
	page      int
}

func NewGenerator(src *card.Service) *Generator {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	return &Generator{
		pdf:       pdf,
		src:       src,
		cardIndex: 0,
		page:      0,
	}
}

func (g *Generator) GeneratePDF(ids []int) error {
	for _, id := range ids {
		card, err := g.src.GetCard(id)
		if err != nil {
			return errors.Wrap(err, "get card")
		}
		if err := g.AddCard(card); err != nil {
			return errors.Wrap(err, "add card")
		}
		g.cardIndex++

	}
	err := g.pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	return nil
}

func (g *Generator) AddCard(card *card.Card) error {
	if err := g.PrintCard(card); err != nil {
		return errors.Wrap(err, "print card")
	}
	models, err := g.src.ListModels(card.ID)
	if err != nil {
		return errors.Wrap(err, "get models")
	}

	abilities := []Abilities{}
	for i, model := range models {
		if i != 0 && i%2 == 0 {
			// g.PrintAbilities(abilities)
			// abilities = []Abilities{}
			g.cardIndex++
			if err := g.PrintCard(card); err != nil {
				return errors.Wrap(err, "print card")
			}
		}
		g.PrintStatline(&model, card.FactionID, i%2)

		cardAbilities, err := g.src.ListCardAbilities(card.ID)
		if err != nil {
			return errors.Wrap(err, "get card abilities")
		}
		temp := map[string]string{}
		for _, a := range cardAbilities {
			if contains(abilities, a.Name) {
				temp[a.Name] = "Voir plus haut."
			} else {
				temp[a.Name] = a.Description
			}
		}
		abilities = append(abilities, Abilities{Label: "", Values: temp})

		modelAbilities, err := g.src.ListModelAbilities(model.ID)
		if err != nil {
			return errors.Wrap(err, "get model abilities")
		}
		temp = map[string]string{}
		for _, a := range modelAbilities {
			if contains(abilities, a.Name) {
				temp[a.Name] = "Voir plus haut."
			} else {
				temp[a.Name] = a.Description
			}
		}

		magicAbilities, err := g.src.ListMagicAbilities(model.ID)
		if err != nil {
			return errors.Wrap(err, "get magic abilities")
		}
		tempM := map[string]string{}
		for _, a := range magicAbilities {
			if contains(abilities, a.Name) {
				tempM[a.Name] = "Voir plus haut."
			} else {
				tempM[a.Name] = a.Description
			}
		}
		abilities = append(abilities, Abilities{Label: model.Name, Values: temp, ValuesMagic: tempM, Magic: model.MagicAbility})

		weapons, err := g.src.ListWeapons(model.ID)
		if err != nil {
			return errors.Wrap(err, "get weapons")
		}

		for j, weapon := range weapons {
			g.PrintWeapon(&weapon, model.STR, card.FactionID, i%2, j)
			weaponAbilities, err := g.src.ListWeaponAbilities(weapon.ID)
			if err != nil {
				return errors.Wrap(err, "get weapon abilities")
			}
			fmt.Println("WA", weapon.Name, len(weaponAbilities))
			temp := map[string]string{}
			for _, a := range weaponAbilities {
				if contains(abilities, a.Name) {
					temp[a.Name] = "Voir plus haut."
				} else {
					temp[a.Name] = a.Description
				}
			}
			abilities = append(abilities, Abilities{Label: weapon.Name, Values: temp})
		}

		if len(models)-1 == i || i%2 == 1 {
			g.PrintAbilities(abilities)
			abilities = []Abilities{}
		}
		// abilityCursor, abilities = g.PrintAbilities(&model, 1, abilityCursor, abilities)
		// if model.MagicAbility != "" {
		// 	g.pdf.SetFont("Arial", "", 8)
		// 	g.pdf.Text(X0+float64(g.cardIndex%2)*CardWidth*2, abilityCursor, fmt.Sprintf("Score de Magie [%s]", model.MagicAbility))
		// 	abilityCursor, abilities = g.PrintAbilities(&model, 2, abilityCursor, abilities)
		// }

	}
	return nil
}

type Abilities struct {
	Label       string
	Values      map[string]string
	Magic       string
	ValuesMagic map[string]string
}

func contains(abilities []Abilities, ability string) bool {
	for _, v := range abilities {
		if _, ok := v.Values[ability]; ok {
			return true
		}
		if _, ok := v.ValuesMagic[ability]; ok {
			return true
		}
	}
	return false
}

func (g *Generator) PrintCard(card *card.Card) error {
	fmt.Println("PRINT", card.ID, g.cardIndex)
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
	g.pdf.SetFont("Arial", "", 10)
	g.pdf.Text(frontX+14, frontY+6, tr(card.Name))
	g.pdf.Text(frontX+CardWidth+10, frontY+7.5, tr(card.Name))
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(frontX+14, frontY+8.5, tr(card.Properties))
	return nil
}

func (g *Generator) PrintWeapon(weapon *card.Weapon, str string, faction, i, j int) error {
	X := X0 + float64(g.cardIndex%2)*CardWidth*2 + 39.5
	Y := Y0 + float64(g.cardIndex/2)*CardHeight + 23 + 36*float64(i) + 12*float64(j)

	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	if weapon.CNT == "" {
		weapon.CNT = "1"
	}
	g.pdf.ImageOptions(fmt.Sprintf("images/weapon_%d_%d_%s.png", faction, weapon.Type, weapon.CNT), X, Y, 23, 10.5, false, opt, 0, "")

	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
	g.pdf.SetFont("Arial", "", 7)
	g.pdf.Text(X+1, Y+3.4, tr(weapon.Name))

	if weapon.Type == 1 {
		g.PrintMeeleStat(X, Y, 0, weapon.RNG)
		g.PrintMeeleStat(X, Y, 1, weapon.POW)
		pow, _ := strconv.Atoi(weapon.POW)
		str, _ := strconv.Atoi(str)
		g.PrintMeeleStat(X, Y, 2, strconv.Itoa(pow+str))

	}
	// g.PrintStat(X, Y, 1, model.STR)
	// g.PrintStat(X, Y, 2, model.MAT)
	// g.PrintStat(X, Y, 3, model.RAT)
	// g.PrintStat(X, Y, 4, model.DEF)
	// g.PrintStat(X, Y, 5, model.ARM)
	// g.PrintStat(X, Y, 6, model.CMD)

	// g.pdf.Text(X+30.1, Y+11.2, model.BaseSize)

	return nil
}

func (g *Generator) PrintStatline(model *card.Model, faction, i int) error {
	X := X0 + float64(g.cardIndex%2)*CardWidth*2 + 29
	Y := Y0 + float64(g.cardIndex/2)*CardHeight + 10
	if i == 1 {
		Y += 36
	}
	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	g.pdf.ImageOptions(fmt.Sprintf("images/stat_%d.png", faction), X, Y, 35, 13, false, opt, 0, "")

	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
	g.pdf.SetFont("Arial", "", 8)
	g.pdf.Text(X+2, Y+3.2, tr(model.Name))
	g.pdf.SetFont("Arial", "", 7)
	g.PrintStat(X, Y, 0, model.SPD)
	g.PrintStat(X, Y, 1, model.STR)
	g.PrintStat(X, Y, 2, model.MAT)
	g.PrintStat(X, Y, 3, model.RAT)
	g.PrintStat(X, Y, 4, model.DEF)
	g.PrintStat(X, Y, 5, model.ARM)
	g.PrintStat(X, Y, 6, model.CMD)

	g.pdf.Text(X+30.1, Y+11.2, model.BaseSize)

	return nil
}

func (g *Generator) PrintStat(X, Y, index float64, value string) {
	X += 4.3
	Y += 8
	if len(value) > 1 {
		g.pdf.Text(X+4.3*index-1.2, Y, value)
	} else {
		g.pdf.Text(X+4.3*index, Y, value)
	}
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

func (g *Generator) PrintAbilities(abilities []Abilities) {
	X := X0 + float64(g.cardIndex%2)*CardWidth*2 + CardWidth + 3
	Y := Y0 + float64(g.cardIndex/2)*CardHeight + 13
	tr := g.pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

	lineNb := 0.0
	g.pdf.SetFont("Arial", "", 6)
	_, lineHt := g.pdf.GetFontSize()

	for _, a := range abilities {
		if a.Label != "" && (len(a.Values) > 0 || len(a.ValuesMagic) > 0) {
			g.pdf.SetFont("Arial", "BU", 6)
			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(strings.ToUpper(a.Label)))
			lineNb += 1.5
		}
		for k, v := range a.Values {
			g.pdf.SetFont("Arial", "", 5)
			data := g.pdf.SplitLines([]byte(fmt.Sprintf("%s - %s", strings.ToUpper(k), v)), 62)
			for _, s := range data {
				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
				lineNb++
			}
			lineNb += 0.5
		}
		if a.Magic != "" {
			g.pdf.SetFont("Arial", "B", 6)
			g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(fmt.Sprintf("Précision Magique [%s]", a.Magic)))
			lineNb += 1.5
		}
		for k, v := range a.ValuesMagic {
			g.pdf.SetFont("Arial", "", 5)
			data := g.pdf.SplitLines([]byte(fmt.Sprintf("\u2022 %s - %s", strings.ToUpper(k), v)), 62)
			for _, s := range data {
				g.pdf.Text(X, Y+float64(lineNb)*lineHt, tr(string(s)))
				lineNb++
			}
			lineNb += 0.5
		}
	}
}
