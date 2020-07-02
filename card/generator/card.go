package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/chibimi/cards/card/ability"
	"github.com/chibimi/cards/card/model"
	"github.com/chibimi/cards/card/weapon"
	multierror "github.com/hashicorp/go-multierror"
	"gitlab.com/golang-commonmark/markdown"
)

func (s *Service) Build(r Reference) (cards []Card, err error) {
	translations := GetTranslation(r.Lang)
	var errs *multierror.Error

	if r.Ref.Special == "colossal" {
		secondRef := Reference{
			Lang:   r.Lang,
			Ref:    r.Ref,
			FileID: fmt.Sprintf("%s_1", r.FileID),
		}
		secondRef.Ref.Special = ""
		defer func() {
			secondCards, err := s.Build(secondRef)
			if err != nil {
				errs = multierror.Append(errs, wrap(err, "build second ref"))
			}
			cards = append(cards, secondCards...)
		}()
	}

	if r.Ref.Special == "charunit" || r.Ref.Special == "dragoon" {
		firstModel := r.Models[0]
		secondRef := Reference{
			Lang:             r.Lang,
			Ref:              r.Ref,
			FileID:           fmt.Sprintf("%s_1", r.FileID),
			RefAbilities:     r.RefAbilities,
			ModelsAbilities:  map[int][]ability.Ability{},
			ModelsWeapons:    map[int][]weapon.Weapon{},
			WeaponsAbilities: map[int][]ability.Ability{},
			Models:           make([]model.Model, len(r.Models)-1),
		}

		copy(secondRef.Models, r.Models[1:])
		r.Models = r.Models[:1]
		for k, v := range r.ModelsAbilities {
			if k == firstModel.ID {
				continue
			}
			secondRef.ModelsAbilities[k] = v
			delete(r.ModelsAbilities, k)
		}
		for k, v := range r.ModelsWeapons {
			if k == firstModel.ID {
				continue
			}
			secondRef.ModelsWeapons[k] = v
			delete(r.ModelsWeapons, k)
			for _, w := range v {
				secondRef.WeaponsAbilities[w.ID] = r.WeaponsAbilities[w.ID]
				delete(r.WeaponsAbilities, w.ID)
			}
		}

		secondRef.Ref.Special = ""
		defer func() {
			secondCards, err := s.Build(secondRef)
			if err != nil {
				errs = multierror.Append(errs, wrap(err, "build second ref"))
			}
			cards = append(cards, secondCards...)
		}()
	}

	// Used later for various checks.
	var category = Category(r.Ref.CategoryID)

	// Build the profile card.
	cards = append(cards, ProfileCard{
		Faction: Faction(r.Ref.FactionID),
		FileID:  r.FileID,
	})

	// Build the rules card.
	// This start with the abilities that refer to the reference itself,
	// which don't have a source displayed on the card. Example: partisan,
	// etc. Then, for each model of the reference, the model abilities then
	// the model weapons' abilities, each set being a different set.
	rules := RulesCard{
		Faction: Faction(r.Ref.FactionID),
		Title:   r.Ref.Title,
	}

	// abilityCache is used to keep track of abilities already on the card
	// to change the description of subsequent iteration by "see_above"
	abilityCache := map[int]ability.Ability{}

	// cardAbilities contains the list of all abilities present on the card
	// this is used to not add the description of a linked ability that is already on the card
	cardAbilities := &sync.Map{}
	for _, a := range r.RefAbilities {
		cardAbilities.Store(a.ID, nil)
	}
	for _, m := range r.Models {
		for _, a := range r.ModelsAbilities[m.ID] {
			cardAbilities.Store(a.ID, nil)
		}
		for _, w := range r.ModelsWeapons[m.ID] {
			for _, a := range r.WeaponsAbilities[w.ID] {
				cardAbilities.Store(a.ID, nil)
			}
		}
	}
	addAbilities := func(source string, abilities []ability.Ability) {
		if len(abilities) == 0 {
			return
		}

		var err error
		list := Abilities{
			Source: source,
		}
		for _, a := range abilities {
			ability := Ability{
				Name: a.Name,
			}
			bullet := ""
			if a.Header != nil {
				bullet = "- "
			}
			if _, ok := abilityCache[a.ID]; ok {
				ability.Description, err = s.Compile(fmt.Sprintf(`%s**%s** (%s)%s – %s`, bullet, a.Name, a.Title, a.GetStarText(), translations.Phrases["see_above"]), r.Lang, a.Name, cardAbilities)
			} else {
				ability.Description, err = s.Compile(fmt.Sprintf(`%s**%s** (%s)%s – %s`, bullet, a.Name, a.Title, a.GetStarText(), a.Description), r.Lang, a.Name, cardAbilities)
				abilityCache[a.ID] = a
			}
			if err != nil {
				errs = multierror.Append(errs, wrap(err, "compile description"))
			}

			list.Abilities = append(list.Abilities, ability)
		}
		rules.Abilities = append(rules.Abilities, list)
	}

	addAbilities("", r.RefAbilities)
	for _, m := range r.Models {
		addAbilities(m.Title, r.ModelsAbilities[m.ID])
		for _, w := range r.ModelsWeapons[m.ID] {
			addAbilities(w.Title, r.WeaponsAbilities[w.ID])
		}
	}

	switch category {
	case CategorySolo, CategoryWarbeast:
		if len(r.Spells) == 0 {
			break
		}

		for _, sp := range r.Spells {
			spell := Spell{
				Name: sp.Name,
				Stats: map[string]string{
					"COST": sp.Cost,
					"RNG":  sp.RNG,
					"AOE":  sp.AOE,
					"POW":  sp.POW,
					"DUR":  sp.DUR,
					"OFF":  sp.OFF,
				},
			}
			spell.Description, err = s.Compile(sp.Description, r.Lang, sp.Name, &sync.Map{})
			if err != nil {
				errs = multierror.Append(errs, wrap(err, "compile spell description"))
			}

			rules.Spells = append(rules.Spells, spell)
		}
	}

	rules.SetFontSize()
	cards = append(cards, rules)

	// Build the spells card. It is only a separate card for warboss-type
	// references.
	switch category {
	case CategoryWarcaster, CategoryWarlock, CategoryInfernalMaster:
		spells := SpellsCard{
			Faction: Faction(r.Ref.FactionID),
			Title:   r.Ref.Title,
		}

		for _, sp := range r.Spells {
			spell := Spell{
				Name: sp.Name,
				Stats: map[string]string{
					"COST": sp.Cost,
					"RNG":  sp.RNG,
					"AOE":  sp.AOE,
					"POW":  sp.POW,
					"DUR":  sp.DUR,
					"OFF":  sp.OFF,
				},
			}
			spell.Description, err = s.Compile(sp.Description, r.Lang, sp.Name, &sync.Map{})
			if err != nil {
				errs = multierror.Append(errs, wrap(err, "compile spell description"))
			}

			spells.Spells = append(spells.Spells, spell)
		}
		spells.SetFontSize()
		cards = append(cards, spells)
	}

	// Build the feat card, if the reference is a warboss-type reference.
	switch Category(r.Ref.CategoryID) {
	case CategoryWarcaster, CategoryWarlock, CategoryInfernalMaster:
		feat := FeatCard{
			Faction: Faction(r.Ref.FactionID),
			Title:   r.Ref.Title,
			Name:    r.Feat.Name,
			Fluff:   r.Feat.Fluff,
		}
		feat.Description, err = s.Compile(r.Feat.Description, r.Lang, r.Feat.Name, &sync.Map{})
		if err != nil {
			errs = multierror.Append(errs, wrap(err, "compile feat description"))
		}

		cards = append(cards, feat)
	}

	for _, attachment := range r.Attachments {
		attachmentCards, err := s.Build(attachment)
		if err != nil {
			errs = multierror.Append(errs, wrap(err, "build attachment"))
		}
		cards = append(cards, attachmentCards...)
	}
	return cards, errs.ErrorOrNil()
}

// reAbility is the regex for abilities tags in texts.
var reAbility = regexp.MustCompile(`#[0-9]+:[^#]+#`)

// reAdvantage is the regex for advantages tags in texts.
var reAdvantage = regexp.MustCompile(`:[^: ]+:`)

// Compile takes a strings and returns a HTML version of it
// with abilities and advantages tags replaced by their textual
// versions.
func (s *Service) Compile(src, lang, this string, cardAbilities *sync.Map) (string, error) {
	translations := GetTranslation(lang)
	var abilities []string
	var errs *multierror.Error

	src = strings.Replace(src, `#this#`, this, -1)

	var buf strings.Builder
	buf.WriteString(reAbility.ReplaceAllStringFunc(src, func(tag string) string {
		id, err := strconv.Atoi(tag[1:strings.IndexRune(tag, ':')])
		if err != nil {
			errs = multierror.Append(errs, wrap(err, "conv link ID"))
			return tag
		}

		ability, err := s.ability.Get(id, lang)
		if err != nil {
			errs = multierror.Append(errs, wrap(err, "get linked ability"))
			return tag
		}
		if _, ok := cardAbilities.LoadOrStore(id, nil); !ok {
			abilities = append(abilities, fmt.Sprintf("%s: %s", ability.Name, ability.Description))
		}

		return ability.Name
	}))
	for _, ability := range abilities {
		buf.WriteString(fmt.Sprintf(" _(%s)_", ability))
	}

	res := reAdvantage.ReplaceAllStringFunc(buf.String(), func(tag string) string {
		tag = tag[1 : len(tag)-1]
		translation, ok := translations.Advantages[tag]
		if !ok {
			errs = multierror.Append(errs, fmt.Errorf("advantage no found: %s", tag[1:len(tag)-1]))
			return tag
		}

		return fmt.Sprintf(`%s ![%s](%s/icons/%s.png)`, translation, tag, s.assets, tag)
	})

	return markdown.New().RenderToString([]byte(res)), errs.ErrorOrNil()
}

type Card interface {
	Type() string
}

type ProfileCard struct {
	Faction Faction
	FileID  string
}

func (ProfileCard) Type() string {
	return "profile"
}

type RulesCard struct {
	Faction   Faction
	Title     string
	Spells    []Spell
	Abilities []Abilities
	FontClass string
}

func (RulesCard) Type() string {
	return "rules"
}
func (c *RulesCard) SetFontSize() {
	nbChar := 0
	for _, s := range c.Spells {
		nbChar += len(s.Name)
		nbChar += len(s.Description)
	}
	for _, abilities := range c.Abilities {
		for _, a := range abilities.Abilities {
			nbChar += len(a.Name)
			nbChar += len(a.Description)
		}
	}

	c.FontClass = getFontClass(nbChar)
}

type Abilities struct {
	Source    string
	Abilities []Ability
}

type Ability struct {
	Name        string
	Description string
}

type AbilityType int

type SpellsCard struct {
	Faction   Faction
	Title     string
	Spells    []Spell
	FontClass string
}

func (SpellsCard) Type() string {
	return "spells"
}

func (c *SpellsCard) SetFontSize() {
	nbChar := 0
	for _, s := range c.Spells {
		nbChar += len(s.Name)
		nbChar += len(s.Description)
	}
	c.FontClass = getFontClass(nbChar)
}

type Spell struct {
	Name        string
	Stats       map[string]string
	Description string
}

type FeatCard struct {
	Faction     Faction
	Title       string
	Name        string
	Fluff       string
	Description string
}

func (FeatCard) Type() string {
	return "feat"
}

func getFontClass(nbChar int) string {
	switch {
	case nbChar > 2250:
		return "font-seriously-pp"
	case nbChar > 1900:
		return "font-xs"
	case nbChar > 1500:
		return "font-s"
	default:
		return "font-m"
	}
}
