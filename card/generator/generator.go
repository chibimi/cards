package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card"
	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/model"
	"github.com/jung-kurt/gofpdf"
	"github.com/pkg/errors"
)

const X0 float64 = 10.0 + 1.9
const Y0 float64 = 10.0 + 9.2
const CardWidth float64 = 63.3
const CardHeight float64 = 88.9
const SeparatorW float64 = 0.55
const SeparatorH float64 = 0.4

type Generator struct {
	src         *card.SService
	references  []int
	lang        string
	pdf         *gofpdf.Fpdf
	unicode     func(string) string
	cardIndex   int
	currentPage int
	x, y        float64
}

func NewGenerator(src *card.SService, references []int, lang string) *Generator {
	pdf := gofpdf.New("L", "mm", "letter", "")
	unicode := pdf.UnicodeTranslatorFromDescriptor("")
	return &Generator{
		src:        src,
		references: references,
		lang:       lang,
		x:          X0,
		y:          Y0,
		cardIndex:  0,
		pdf:        pdf,
		unicode:    unicode,
	}
}

func (g *Generator) LoadFont() {
	g.pdf.AddUTF8Font("Abilities", "", "fonts/CALIST.ttf")
	g.pdf.AddUTF8Font("Abilities", "B", "fonts/CALISTB.ttf")
}

func (g *Generator) GeneratePDF() error {
	g.LoadFont()

	for _, id := range g.references {
		ref, err := g.src.Ref.Get(id, g.lang)
		if err != nil {
			return errors.Wrap(err, "get ref")
		}

		models, err := g.src.Model.List(ref.ID, g.lang)
		modelsByCard := split(models)
		for i, models := range modelsByCard {
			g.nextCard()
			g.PrintFront(g.x, g.y, ref, i)
			g.y += SeparatorH
			g.reverse()
			g.PrintAbilities(g.x, g.y, ref, models)
			g.pdf.TransformEnd()
		}

		if ref.CategoryID == 1 || ref.CategoryID == 2 || ref.CategoryID == 10 {
			g.nextCard()
			spells, err := g.src.Spell.ListByRef(ref.ID, g.lang)
			if err != nil {
				return errors.Wrap(err, "get spells")
			}

			g.PrintSpells(g.x, g.y, ref, spells)

			feat, err := g.src.Feat.Get(ref.ID, g.lang)
			if err != nil {
				return errors.Wrap(err, "get feat")
			}

			g.reverse()
			g.PrintFeat(g.x, g.y, ref, feat)
			g.pdf.TransformEnd()
		}
	}

	return nil
}

func (g *Generator) reverse() {
	g.pdf.TransformBegin()
	g.pdf.TransformRotate(180, g.x+CardWidth/2, g.y+CardHeight)
}

func split(models []model.Model) [][]model.Model {
	var divided [][]model.Model

	for i := 0; i < len(models); i += 2 {
		end := i + 2
		if end > len(models) {
			end = len(models)
		}
		divided = append(divided, models[i:end])
	}
	return divided
}

func (g *Generator) WritePDF(dest string) error {
	return g.pdf.OutputFileAndClose(dest)
}

func (g *Generator) nextCard() {
	if g.cardIndex%4 == 0 {
		g.nextPage()
	}
	g.x = X0 + float64(g.cardIndex)*(CardWidth+SeparatorW)
	g.cardIndex++
}

func (g *Generator) nextPage() {
	g.pdf.AddPage()
	g.pdf.Image("pp/871c3261-9ddd-4bcf-9737-4b4fb9008105-0.png", 0, 0, 279.4, 215.9, false, "", 0, "")

	g.cardIndex = 0
	g.x = X0
	g.y = Y0
}

var validLink = regexp.MustCompile(`#[0-9]+:[^#]+#`)

func (g *Generator) replaceLinks(description string, abilities map[int]ability.Ability) string {
	subskills := []string{}
	description = validLink.ReplaceAllStringFunc(description, func(src string) string {
		s := strings.SplitN(src, ":", 2)
		ids := s[0][1:]
		id, err := strconv.Atoi(ids)
		if err != nil {
			return s[1][:len(s[1])-1]
		}
		if a, ok := abilities[id]; ok {
			return a.Name
		}
		// fetch ability
		a, err := g.src.Ability.Get(id, "FR")
		if err != nil {
			return s[1][:len(s[1])-1]
		}
		subskills = append(subskills, fmt.Sprintf("%s (%s): %s", a.Name, a.Title, a.Description))
		abilities[a.ID] = *a
		return a.Name

	})
	if len(subskills) > 0 {
		description = fmt.Sprintf("%s (%s)", description, strings.Join(subskills, ", "))
	}
	return description
}
