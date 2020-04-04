package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/chibimi/cards/card/ability"
	multierror "github.com/hashicorp/go-multierror"
	"gitlab.com/golang-commonmark/markdown"
)

func (s *Service) Build(r Reference) (cards []Card, err error) {
	var errs *multierror.Error

	// Build the profile card.
	cards = append(cards, ProfileCard{
		Faction: Faction(r.Ref.FactionID),
		PPID:    r.Ref.PPID,
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
				ability.Description, err = s.Compile(fmt.Sprintf(`%s**%s** (%s) – %s`, bullet, a.Name, a.Title, translations[fmt.Sprintf("see_above_%s", r.Lang)]), r.Lang, a.Title, cardAbilities)
			} else {
				ability.Description, err = s.Compile(fmt.Sprintf(`%s**%s** (%s) – %s`, bullet, a.Name, a.Title, a.Description), r.Lang, a.Title, cardAbilities)
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
			addAbilities(w.Name, r.WeaponsAbilities[w.ID])
		}
	}

	cards = append(cards, rules)

	// Build the spells card. It is only a separate card for warboss-type
	// references.
	if r.HasSpells() {
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

		cards = append(cards, spells)
	}

	// Build the feat card, if the reference is a warboss-type reference.
	if r.HasFeat() {
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
		advantage, err := s.src.Advantage.Get(tag[1 : len(tag)-1])
		if err != nil {
			errs = multierror.Append(errs, wrap(err, "get advantage: %s", tag[1:len(tag)-1]))
			return tag
		}

		return fmt.Sprintf(`%s ![%s](%s/icons/%s.png)`, advantage.Name, advantage.ID, s.assets, advantage.ID)
	})

	return markdown.New().RenderToString([]byte(res)), errs.ErrorOrNil()
}

type Card interface {
	Type() string
}

type ProfileCard struct {
	Faction Faction
	PPID    int
}

func (ProfileCard) Type() string {
	return "profile"
}

type RulesCard struct {
	Faction   Faction
	Title     string
	Abilities []Abilities
}

func (RulesCard) Type() string {
	return "rules"
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
	Faction Faction
	Title   string
	Spells  []Spell
}

func (SpellsCard) Type() string {
	return "spells"
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
