package reference

type Reference struct {
	ID         int     `json:"id,omitempty" db:"id"`
	PPID       int     `json:"ppid,omitempty" db:"ppid"`
	FactionID  Faction `json:"faction_id,omitempty" db:"faction_id"`
	CategoryID int     `json:"category_id,omitempty" db:"category_id"`
	Title      string  `json:"title,omitempty" db:"title"`
	MainCardID int     `json:"main_card_id,omitempty,string" db:"main_card_id"`
	Models     string  `json:"models_cnt,omitempty" db:"models_cnt"`
	ModelsMax  string  `json:"models_max,omitempty" db:"models_max"`
	Cost       string  `json:"cost,omitempty" db:"cost"`
	CostMax    string  `json:"cost_max,omitempty" db:"cost_max"`
	FA         string  `json:"fa,omitempty" db:"fa"`
	Name       string  `json:"name,omitempty" db:"name"`
	Properties string  `json:"properties,omitempty" db:"properties"`
	Status     string  `json:"status,omitempty" db:"status"`
	MercFor    []int   `json:"mercenary_for,omitempty"`
	MinFor     []int   `json:"minion_for,omitempty"`
	Special    string  `json:"special,omitempty" db:"special"`
	LinkedTo   *int    `json:"linked_to,omitempty" db:"linked_to"`
}

func (r Reference) HasFeat() bool {
	switch Category(r.CategoryID) {
	case CategoryWarcaster, CategoryWarlock, CategoryInfernalMaster:
		return true
	default:
		return false
	}
}

func (r Reference) HasSpells() bool {
	switch Category(r.CategoryID) {
	case CategoryWarcaster, CategoryWarlock, CategoryInfernalMaster:
		return true
	default:
		return false
	}
}

type Category int

const (
	CategoryInvalid Category = iota
	CategoryWarcaster
	CategoryWarlock
	CategoryWarjack
	CategoryWarbeast
	CategoryUnit
	CategorySolo
	CategoryAttachments
	CategoryBattleEngine
	CategoryStructure
	CategoryInfernalMaster
	CategoryHorror
)

type Faction int

const (
	FactionInvalid = iota
	FactionCygnar
	FactionProtectorateOfMenoth
	FactionKhador
	FactionCryx
	FactionRetributionOfScyrah
	FactionConvergeanceOfCyriss
	FactionCrucibleGuard
	FactionMercenaries
	FactionTrollbloods
	FactionCircleOrboros
	FactionLegionOfEverblight
	FactionSkorne
	FactionGrymkin
	FactionInfernals
	FactionMinions
)

func (f Faction) String() string {
	switch f {
	case FactionCygnar:
		return "cygnar"
	case FactionProtectorateOfMenoth:
		return "protectorate of menoth"
	case FactionKhador:
		return "khador"
	case FactionCryx:
		return "cryx"
	case FactionRetributionOfScyrah:
		return "retribution of scyrah"
	case FactionConvergeanceOfCyriss:
		return "convergeance of cyriss"
	case FactionCrucibleGuard:
		return "crucible guard"
	case FactionMercenaries:
		return "mercenaries"
	case FactionTrollbloods:
		return "trollbloods"
	case FactionCircleOrboros:
		return "circle orboros"
	case FactionLegionOfEverblight:
		return "legion of everblight"
	case FactionSkorne:
		return "skorne"
	case FactionGrymkin:
		return "grymkin"
	case FactionInfernals:
		return "infernals"
	case FactionMinions:
		return "minions"
	default:
		return "invalid"
	}
}
