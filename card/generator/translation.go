package generator

type Translation struct {
	Advantages map[string]string
	Phrases    map[string]string
}

func GetTranslation(lang string) Translation {
	switch lang {
	case "FR":
		return langFR
	default:
		return langFR
	}
}

var langFR = Translation{
	Advantages: map[string]string{
		"advance_deploy":       "Déploiement Avancé",
		"amphibious":           "Amphibie",
		"arc_node":             "Arc Node",
		"assault":              "Assaut",
		"cavalry":              "Cavalerie",
		"cma":                  "CMA",
		"cra":                  "CRA",
		"construct":            "Construct",
		"eyeless_sight":        "Vision Aveugle",
		"flight":               "Vol",
		"gunfighter":           "Pistolero",
		"incorporeal":          "Incorporel",
		"immunity_corrosion":   "Immunité Corrosion ",
		"immunity_electricity": "Immunité Electricité",
		"immunity_fire":        "Immunité Feu",
		"immunity_frost":       "Immunité Froid",
		"jackmarshal":          "Jackmarshal",
		"officer":              "Officier",
		"parry":                "Parade",
		"pathfinder":           "Eclaireur",
		"soulless":             "Sans Âme",
		"stealth":              "Furtivité",
		"tough":                "Robustesse",
		"undead":               "Mort-Vivant",
		"blessed":              "Benie",
		"chain":                "Chaîne",
		"type_corrosion":       "Type: Corrosion",
		"continuous_corrosion": "Corrosion Continue",
		"crit_corrosion":       "Crit: Corrosion",
		"type_electricity":     "Type: Electrique",
		"disruption":           "Disruption",
		"crit_disruption":      "Crit: Disruption",
		"type_fire":            "Type: Feu",
		"continuous_fire":      "Feu Continu",
		"crit_fire":            "Crit: Feu",
		"type_frost":           "Type: Froid",
		"magical":              "Magique",
		"open_fist":            "Main Libre",
		"shield_1":             "Bouclier +1",
		"shield_2":             "Bouclier +2",
		"weapon_master":        "Maitre d'arme",
	},
	Phrases: map[string]string{
		"see_above": "Voir plus haut.",
	},
}
