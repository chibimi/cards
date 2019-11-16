package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	db, err := sqlx.Open("mysql", "cards_api:cards_api@/cards_db")
	if err != nil {
		log15.Crit("Unable to access db", "err", err.Error())
	}
	defer db.Close()
	// err = loadAbilities(db, "abilities.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = loadSpells(db, "spells.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err = loadRefs(db, "models.json")
	if err != nil {
		fmt.Println(err)
	}

}
func loadRefs(db *sqlx.DB, src string) error {
	refs, err := openRefs(src)
	if err != nil {
		return errors.Wrap(err, "open ref")
	}
	toAttach := map[int]string{}
	ids := map[string]int{}
	cache := map[string]struct{}{}
	for _, ref := range refs {
		if _, ok := cache[ref.FullName]; ok {
			continue
		}
		fid, ok := factions[ref.Faction]
		if !ok {
			return errors.New("faction " + ref.FullName)
		}
		cid, ok := categories[ref.Type]
		if !ok {
			return errors.New("category " + ref.FullName)
		}

		if ref.MaxSize == "-" {
			ref.MaxSize = ""
		}
		if ref.Cost != "" && ref.Cost != "-" {
			ref.MinCost = ref.Cost
		}
		if ref.WJPoints != "" && ref.WJPoints != "-" {
			ref.MinCost = ref.WJPoints
		}
		if ref.WBPoints != "" && ref.WBPoints != "-" {
			ref.MinCost = ref.WBPoints
		}
		minion := []int{}
		merc := []int{}

		for _, v := range ref.WorksFor {
			id, ok := factions[v.ID]
			if !ok {
				return errors.New("faction work " + ref.FullName + v.ID)
			}
			if id == 9 || id == 10 || id == 11 || id == 12 {
				minion = append(minion, id)
			} else {
				merc = append(merc, id)
			}
		}
		minionB, err := json.Marshal(minion)
		if err != nil {
			return errors.Wrap(err, "minion")
		}
		mercB, err := json.Marshal(merc)
		if err != nil {
			return errors.Wrap(err, "merc")
		}
		res, err := db.Exec("INSERT INTO refs (faction_id, category_id, title, main_card_id, models_cnt, models_max, cost, cost_max, fa, minion_for, mercenary_for) VALUES (?,?,?,0,?,?,?,?,?,?,?)",
			fid, cid, ref.FullName, ref.MinSize, ref.MaxSize, ref.MinCost, ref.MaxCost, ref.Fa, minionB, mercB)
		if err != nil {
			return errors.Wrap(err, "refs")
		}
		refID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		ids[ref.ID] = int(refID)
		_, err = db.Exec("INSERT INTO refs_lang (ref_id, lang, status, name, properties) VALUES (?,?,?,?,?)", refID, "UK", "wip", ref.FullName, ref.Qualification)
		if err != nil {
			return errors.Wrap(err, "refs lang")
		}

		_, err = db.Exec("INSERT INTO feats (ref_id, lang, name, description, fluff) VALUES (?,?,?,?,?)", refID, "UK", strings.Title(strings.ToLower(ref.Feat.Title)), ref.Feat.Text, "")
		if err != nil {
			return errors.Wrap(err, "feats")
		}

		if len(ref.RestrictedTo) == 1 {
			toAttach[int(refID)] = ref.RestrictedTo[0].ID
		}
		cache[ref.FullName] = struct{}{}

		for _, ability := range ref.Capacities {
			title := strings.Title(strings.ToLower(ability.Title))
			var abilityID int64
			err = db.Get(&abilityID, "SELECT id FROM abilities WHEre title = ?", title)
			if err != nil && err != sql.ErrNoRows {
				return errors.Wrap(err, "select ability id")
			}
			if err == sql.ErrNoRows {
				fmt.Println("not found", title)
				continue
			}
			_, err = db.Exec("INSERT INTO ref_ability (ref_id, ability_id) VALUES (?,?)", refID, abilityID)
			if err != nil {
				return errors.Wrap(err, "ref abi")
			}
		}

		for _, model := range ref.Models {
			advantages, err := json.Marshal(getModelAdvantages(model))
			if err != nil {
				return err
			}
			magic := getModelMagic(model)
			rsc := getModelResource(model)
			hp := getModelHP(model)
			bs := model.Basestats
			res, err := db.Exec("INSERT INTO models (title, ref_id, spd, str, mat, rat, def, arm, cmd, magic_ability, damage, resource, threshold, base_size, advantages) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
				strings.Title(strings.ToLower(bs.Name)), refID, bs.Spd, bs.Str, bs.Mat, bs.Rat, bs.Def, bs.Arm, bs.Cmd, magic, hp, rsc, bs.Thr, "", advantages)
			if err != nil {
				return errors.Wrap(err, "models")
			}
			modelID, err := res.LastInsertId()
			if err != nil {
				return err
			}
			_, err = db.Exec("INSERT INTO models_lang (model_id, lang, name) VALUES (?,?,?)", modelID, "UK", strings.Title(strings.ToLower(model.Basestats.Name)))
			if err != nil {
				return errors.Wrap(err, "models_lang")
			}

			for _, ability := range model.Capacities {
				title := strings.Title(strings.ToLower(ability.Title))
				var abilityID int64
				err = db.Get(&abilityID, "SELECT id FROM abilities WHEre title = ?", title)
				if err != nil && err != sql.ErrNoRows {
					return errors.Wrap(err, "select ability id")
				}
				if err == sql.ErrNoRows {
					fmt.Println("not found", title)
					continue
				}
				_, err = db.Exec("INSERT INTO model_ability (model_id, ability_id) VALUES (?,?)", modelID, abilityID)
				if err != nil {
					return errors.Wrap(err, "model abi")
				}
			}

			for _, spell := range model.Spells {
				title := strings.Title(strings.ToLower(spell.Name))
				var spellID int64
				err = db.Get(&spellID, "SELECT id FROM spells WHEre title = ?", title)
				if err != nil && err != sql.ErrNoRows {
					return errors.Wrap(err, "select spell id")
				}
				if err == sql.ErrNoRows {
					fmt.Println("not found", title)
					continue
				}
				_, err = db.Exec("INSERT INTO ref_spell (ref_id, spell_id) VALUES (?,?)", refID, spellID)
				if err != nil {
					return errors.Wrap(err, "ref spell")
				}
			}

			for _, wp := range model.Weapons.RangedWeapon {
				advantages, err := json.Marshal(getWeaponAdvantages(wp))
				if err != nil {
					return err
				}

				res, err := db.Exec("INSERT INTO weapons (title, model_id, type, rng, pow, rof, aoe, loc, cnt, advantages) VALUES (?,?,?,?,?,?,?,?,?,?)",
					strings.Title(strings.ToLower(wp.Name)), modelID, "2", wp.Rng, wp.Pow, wp.Rof, wp.Aoe, wp.Location, wp.Count, advantages)
				if err != nil {
					return errors.Wrap(err, "ranged weapons")
				}
				weaponID, err := res.LastInsertId()
				if err != nil {
					return err
				}
				_, err = db.Exec("INSERT INTO weapons_lang (weapon_id, lang, name) VALUES (?,?,?)", weaponID, "UK", strings.Title(strings.ToLower(wp.Name)))
				if err != nil {
					return errors.Wrap(err, "weapons_lang")
				}

				for _, ability := range wp.Capacities {
					title := strings.Title(strings.ToLower(ability.Title))
					var abilityID int64
					err = db.Get(&abilityID, "SELECT id FROM abilities WHEre title = ?", title)
					if err != nil && err != sql.ErrNoRows {
						return errors.Wrap(err, "select ability id")
					}
					if err == sql.ErrNoRows {
						fmt.Println("not found", title)
						continue
					}
					_, err = db.Exec("INSERT INTO weapon_ability (weapon_id, ability_id) VALUES (?,?)", weaponID, abilityID)
					if err != nil {
						return errors.Wrap(err, "wp abi")
					}
				}
			}
			for _, wp := range model.Weapons.MeleeWeapon {
				advantages, err := json.Marshal(getWeaponAdvantages(wp))
				if err != nil {
					return err
				}
				typ := "1"
				if strings.ToLower(wp.Name) == "mount" {
					typ = "3"
				}
				res, err := db.Exec("INSERT INTO weapons (title, model_id, type, rng, pow, rof, aoe, loc, cnt, advantages) VALUES (?,?,?,?,?,?,?,?,?,?)",
					strings.Title(strings.ToLower(wp.Name)), modelID, typ, wp.Rng, wp.Pow, wp.Rof, wp.Aoe, wp.Location, wp.Count, advantages)
				if err != nil {
					return errors.Wrap(err, "melee weapons")
				}
				weaponID, err := res.LastInsertId()
				if err != nil {
					return err
				}
				_, err = db.Exec("INSERT INTO weapons_lang (weapon_id, lang, name) VALUES (?,?,?)", weaponID, "UK", strings.Title(strings.ToLower(wp.Name)))
				if err != nil {
					return errors.Wrap(err, "weapons_lang")
				}

				for _, ability := range wp.Capacities {
					title := strings.Title(strings.ToLower(ability.Title))
					var abilityID int64
					err = db.Get(&abilityID, "SELECT id FROM abilities WHEre title = ?", title)
					if err != nil && err != sql.ErrNoRows {
						return errors.Wrap(err, "select ability id")
					}
					if err == sql.ErrNoRows {
						fmt.Println("not found", title)
						continue
					}
					_, err = db.Exec("INSERT INTO weapon_ability (weapon_id, ability_id) VALUES (?,?)", weaponID, abilityID)
					if err != nil {
						return errors.Wrap(err, "wp abi")
					}
				}
			}
		}
	}

	for k, v := range toAttach {
		mainID, ok := ids[v]
		if !ok {
			continue
		}
		_, err := db.Exec("UPDATE refs SET main_card_id = ? WHERE id = ?", k, mainID)
		if err != nil {
			return err
		}
	}

	return nil
}

func openRefs(src string) ([]Reference, error) {
	refs := []Reference{}

	file, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read refs file")
	}
	err = json.Unmarshal(file, &refs)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to unmarshal refs file")
	}

	return refs, nil
}

func getModelMagic(m Model) string {
	for _, a := range m.Capacities {
		if strings.HasPrefix(strings.ToLower(a.Title), "magic ability") && len(a.Title) > 15 {
			// fmt.Println("found magic", a.Title, a.Title[15])
			return string(a.Title[15])
		}
	}
	return ""
}
func getModelResource(m Model) string {
	if m.Basestats.Fury != "" {
		return m.Basestats.Fury
	}
	if m.Basestats.Foc != "" {
		return m.Basestats.Foc
	}
	return ""
}

func getModelAdvantages(m Model) []string {
	res := []string{}
	if m.Basestats.AD {
		res = append(res, "advance_deploy")
	}
	if m.Basestats.Amphibious {
		res = append(res, "amphibious")
	}
	if m.Basestats.ArcNode {
		res = append(res, "arc_node")
	}
	if m.Basestats.Assault {
		res = append(res, "assault")
	}
	if m.Basestats.Cavalry {
		res = append(res, "cavalry")
	}
	if m.Basestats.CMA {
		res = append(res, "cma")
	}
	if m.Basestats.CRA {
		res = append(res, "cra")
	}
	if m.Basestats.Construct {
		res = append(res, "construct")
	}
	if m.Basestats.EyelessSight {
		res = append(res, "eyeless_sight")
	}
	if m.Basestats.Flight {
		res = append(res, "flight")
	}
	if m.Basestats.GunFighter {
		res = append(res, "gunfighter")
	}
	if m.Basestats.Incoporeal {
		res = append(res, "incorporeal")
	}
	if m.Basestats.ImmunityCorrosion {
		res = append(res, "immunity_corrosion")
	}
	if m.Basestats.ImmunityElec {
		res = append(res, "immunity_electricity")
	}
	if m.Basestats.ImmunityFire {
		res = append(res, "immunity_fire")
	}
	if m.Basestats.ImmunityFrost {
		res = append(res, "immunity_frost")
	}
	if m.Basestats.JackMarshal {
		res = append(res, "jackmarshal")
	}
	if m.Basestats.Officer {
		res = append(res, "officer")
	}
	if m.Basestats.Parry {
		res = append(res, "parry")
	}
	if m.Basestats.Pathfinder {
		res = append(res, "pathfinder")
	}
	if m.Basestats.Soulless {
		res = append(res, "soulless")
	}
	if m.Basestats.Stealth {
		res = append(res, "stealth")
	}
	if m.Basestats.Tough {
		res = append(res, "tought")
	}
	if m.Basestats.Undead {
		res = append(res, "undead")
	}
	return res
}
func getWeaponAdvantages(w Weapon) []string {
	res := []string{}
	if w.Blessed {
		res = append(res, "blessed")
	}
	if w.Chain {
		res = append(res, "chain")
	}
	if w.TypeCorrosion {
		res = append(res, "type_corrosion")
	}
	if w.ContCorrosion {
		res = append(res, "continuous_corrosion")
	}
	if w.CritCorrosion {
		res = append(res, "crit_corrotion")
	}
	if w.TypeElec {
		res = append(res, "type_electricity")
	}
	if w.Disruption {
		res = append(res, "disruption")
	}
	if w.CritDisruption {
		res = append(res, "crit_disruption")
	}
	if w.TypeFire {
		res = append(res, "type_fire")
	}
	if w.ContFire {
		res = append(res, "continuous_fire")
	}
	if w.CritFire {
		res = append(res, "crit_fire")
	}
	if w.Magical {
		res = append(res, "magical")
	}
	if w.OpenFist {
		res = append(res, "open_fist")
	}
	if w.Shield1 {
		res = append(res, "shield_1")
	}
	if w.Shield2 {
		res = append(res, "shield_2")
	}
	if w.WeaponMaster {
		res = append(res, "weapon_master")
	}

	return res
}

func getModelHP(m Model) string {
	if m.Basestats.Spiral != "" {
		return m.Basestats.Spiral
	}
	if m.Basestats.Grid != "" {
		return m.Basestats.Grid
	}
	return m.Basestats.Hitpoints
}

var factions = map[string]int{
	"faction_everblight":  11,
	"faction_orboros":     10,
	"faction_minions":     15,
	"faction_skorne":      12,
	"faction_trollblood":  9,
	"faction_retribution": 5,
	"faction_cyriss":      6,
	"faction_mercs":       8,
	"faction_cryx":        4,
	"faction_khador":      3,
	"faction_cygnar":      1,
	"faction_menoth":      2,
}

var categories = map[string]int{
	"warbeast":      4,
	"CA":            7,
	"WA":            7,
	"unit":          5,
	"solo":          6,
	"warjack":       3,
	"warlock":       2,
	"warcaster":     1,
	"battle engine": 8,
	"colossal":      3,
}

type Reference struct {
	ID            string     `json:"_id"`
	Name          string     `json:"name"`
	Status        string     `json:"status"`
	Qualification string     `json:"qualification"`
	Type          string     `json:"type"`
	Faction       string     `json:"faction"`
	FullName      string     `json:"full_name"`
	Fa            string     `json:"fa"`
	Cost          string     `json:"cost"`
	MinSize       string     `json:"minSize"`
	MaxSize       string     `json:"maxSize"`
	MinCost       string     `json:"minCost"`
	MaxCost       string     `json:"maxCost"`
	WJPoints      string     `json:"wj_points"`
	WBPoints      string     `json:"wb_points"`
	WorksFor      []ID       `json:"works_for"`
	RestrictedTo  []ID       `json:"restricted_to"`
	Models        []Model    `json:"models"`
	Feat          Feat       `json:"feat"`
	Capacities    []Capacity `json:"capacities"`
}

type Model struct {
	Basestats  BaseStats  `json:"basestats"`
	Weapons    Weapons    `json:"weapons"`
	Spells     []Spell    `json:"spells"`
	Capacities []Capacity `json:"capacities"`
}

type BaseStats struct {
	Name              string `json:"_name"`
	Spd               string `json:"_spd"`
	Str               string `json:"_str"`
	Mat               string `json:"_mat"`
	Rat               string `json:"_rat"`
	Def               string `json:"_def"`
	Arm               string `json:"_arm"`
	Cmd               string `json:"_cmd"`
	Hitpoints         string `json:"_hitpoints"`
	Spiral            string `json:"_damage_spiral"`
	Grid              string `json:"_damage_grid"`
	Foc               string `json:"_foc"`
	Fury              string `json:"_fur"`
	Thr               string `json:"_thr"`
	AD                bool   `json:"_advance_deployment"`
	Amphibious        bool   `json:"_amphibious"`
	ArcNode           bool   `json:"_arc_node"`
	Assault           bool   `json:"_assault"`
	Cavalry           bool   `json:"_cavalry"`
	CMA               bool   `json:"_cma"`
	CRA               bool   `json:"_cra"`
	Construct         bool   `json:"_construct"`
	EyelessSight      bool   `json:"_eyelesssight"`
	Flight            bool   `json:"_flight"`
	GunFighter        bool   `json:"_gunfighter"`
	Incoporeal        bool   `json:"_incorporeal"`
	ImmunityElec      bool   `json:"_immunity_electricity"`
	ImmunityFire      bool   `json:"_immunity_fire"`
	ImmunityFrost     bool   `json:"_immunity_frost"`
	ImmunityCorrosion bool   `json:"_immunity_corrosion"`
	JackMarshal       bool   `json:"_jack_marshal"`
	Officer           bool   `json:"_officer"`
	Parry             bool   `json:"_parry"`
	Pathfinder        bool   `json:"_pathfinder"`
	Soulless          bool   `json:"_soulless"`
	Stealth           bool   `json:"_stealth"`
	Tough             bool   `json:"_tough"`
	Undead            bool   `json:"_undead"`
}

type Spell struct {
	Name     string `json:"_name"`
	Cost     string `json:"_cost"`
	Rng      string `json:"_rng"`
	Pow      string `json:"_pow"`
	Aoe      string `json:"_aoe"`
	Duration string `json:"_duration"`
	Off      string `json:"_off"`
	Text     string `json:"__text"`
}

type ID struct {
	ID string `json:"_id"`
}
type Capacity struct {
	Title string `json:"_title"`
	Type  string `json:"_type"`
	Text  string `json:"__text"`
}

type Feat struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Weapons struct {
	MeleeWeapon  []Weapon `json:"melee_weapon"`
	RangedWeapon []Weapon `json:"ranged_weapon"`
}

type Weapon struct {
	Capacities     []Capacity `json:"capacities"`
	Name           string     `json:"_name"`
	Pow            string     `json:"_pow"`
	Count          string     `json:"_count"`
	Rof            string     `json:"_rof"`
	Aoe            string     `json:"_aoe"`
	Rng            string     `json:"_rng"`
	Location       string     `json:"_location"`
	Blessed        bool       `json:"_blessed"`
	Chain          bool       `json:"_chain"`
	TypeCorrosion  bool       `json:"_corrosion"`
	ContCorrosion  bool       `json:"_continuous_corrosion"`
	CritCorrosion  bool       `json:"_critical_corrosion"`
	TypeElec       bool       `json:"_electricity"`
	Disruption     bool       `json:"_disrupt"`
	CritDisruption bool       `json:"_critical_disrupt"`
	TypeFire       bool       `json:"_fire"`
	ContFire       bool       `json:"_continuous_fire"`
	CritFire       bool       `json:"_critical_fire"`
	Magical        bool       `json:"_magical"`
	OpenFist       bool       `json:"_open_fist"`
	Shield1        bool       `json:"_buckler"`
	Shield2        bool       `json:"_shield"`
	WeaponMaster   bool       `json:"_weapon_master"`
}
