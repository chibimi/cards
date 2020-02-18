package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card/ability"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"gitlab.com/golang-commonmark/markdown"
)

func (s *Service) Build(r Reference) (cards []Card, err error) {
	var errs *multierror.Error

	// Build the profile card.
	profile := ProfileCard{
		Faction:  Faction(r.Ref.FactionID),
		Title:    r.Ref.Title,
		Tagline:  r.Ref.Properties,
		Portrait: strconv.Itoa(r.Ref.ID),
	}

	for _, m := range r.Models {
		p := Profile{
			Name: m.Title,
			Stats: map[string]string{
				"SPD":      m.SPD,
				"STR":      m.STR,
				"MAT":      m.MAT,
				"RAT":      m.RAT,
				"DEF":      m.DEF,
				"ARM":      m.ARM,
				"CMD":      m.CMD,
				"Resource": m.Resource,
				"Base":     m.BaseSize,
			},
			Advantages: makeAdvantages(m.Advantages, r.Lang),
		}

		str, err := strconv.Atoi(m.STR)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "conv STR"))
		}

		for _, w := range r.ModelsWeapons[m.ID] {
			pow, err := strconv.Atoi(w.POW)
			if err != nil {
				errs = multierror.Append(errs, errors.Wrap(err, "conv POW"))
			}

			p.Weapons = append(p.Weapons, Weapon{
				Name:     w.Name,
				Type:     WeaponType(w.Type),
				Number:   w.CNT,
				Location: w.LOC,
				Stats: map[string]string{
					"RNG": w.RNG,
					"POW": w.POW,
					"AOE": w.AOE,
					"ROF": w.ROF,
					"PS":  strconv.Itoa(str + pow),
				},
				Advantages: makeAdvantages(w.Advantages, r.Lang),
			})
		}

		profile.Profiles = append(profile.Profiles, p)
	}

	cards = append(cards, profile)

	// Build the rules card.
	// This start with the abilities that refer to the reference itself,
	// which don't have a source displayed on the card. Example: partisan,
	// etc. Then, for each model of the reference, the model abilities then
	// the model weapons' abilities, each set being a different set.
	rules := RulesCard{
		Faction: Faction(r.Ref.FactionID),
		Title:   r.Ref.Title,
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
			ability.Description, err = s.Compile(fmt.Sprintf(`%s**%s** (%s) – %s`, bullet, a.Name, a.Title, a.Description), r.Lang, a.Title)
			errs = multierror.Append(errs, errors.Wrap(err, "compile description"))

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
			spell.Description, err = s.Compile(sp.Description, r.Lang, sp.Name)
			errs = multierror.Append(errs, errors.Wrap(err, "compile spell description"))

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
		feat.Description, err = s.Compile(r.Feat.Description, r.Lang, r.Feat.Name)
		errs = multierror.Append(errs, errors.Wrap(err, "compile feat description"))

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
func (s *Service) Compile(src, lang, this string) (string, error) {
	var abilities []string
	var errs *multierror.Error

	src = strings.Replace(src, `#this#`, this, -1)

	var buf strings.Builder
	buf.WriteString(reAbility.ReplaceAllStringFunc(src, func(tag string) string {
		id, err := strconv.Atoi(tag[1:strings.IndexRune(tag, ':')])
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "conv link ID"))
			return tag
		}

		ability, err := s.ability.Get(id, lang)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "get linked ability"))
			return tag
		}

		abilities = append(abilities, fmt.Sprintf("%s: %s", ability.Name, ability.Description))
		return ability.Name
	}))
	for _, ability := range abilities {
		buf.WriteString(fmt.Sprintf(" _(%s)_", ability))
	}

	res := reAdvantage.ReplaceAllStringFunc(buf.String(), func(tag string) string {
		advantage, err := s.src.Advantage.Get(tag[1 : len(tag)-1])
		if err != nil {
			errs = multierror.Append(errs, errors.Wrapf(err, "get advantage: %s", tag[1:len(tag)-1]))
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
	Faction  Faction
	Title    string
	Tagline  string
	Portrait string
	Profiles []Profile
}

func (ProfileCard) Type() string {
	return "profile"
}

type Profile struct {
	Name       string
	Stats      map[string]string
	Advantages []Advantage
	Weapons    []Weapon
}

type Weapon struct {
	Name       string
	Type       WeaponType
	Number     string
	Location   string
	Stats      map[string]string
	Advantages []Advantage
}

type WeaponType int

const (
	WeaponTypeInvalid = iota
	WeaponTypeMelee
	WeaponTypeRanged
	WeaponTypeMount
)

var weaponTypesNames = map[WeaponType]string{
	WeaponTypeInvalid: "invalid",
	WeaponTypeMelee:   "melee",
	WeaponTypeRanged:  "ranged",
	WeaponTypeMount:   "mount",
}

func (wt WeaponType) String() string {
	return weaponTypesNames[wt]
}

type Advantage struct {
	ID   string
	Name string
}

func makeAdvantages(advantages []string, lang string) []Advantage {
	var res []Advantage
	for _, id := range advantages {
		res = append(res, Advantage{
			ID: id,
		})
	}
	return res
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
