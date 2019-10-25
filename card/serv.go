package card

import (
	"github.com/chibimi/cards/card/feat"
	"github.com/chibimi/cards/card/reference"
	"github.com/jmoiron/sqlx"
)

type SConfig struct {
}

type SService struct {
	Ref *reference.Service
	// Card *card.Service
	Feat *feat.Service
}

func NewSService(db *sqlx.DB) *SService {
	return &SService{
		Ref: reference.NewService(reference.NewRepository(db)),
		// Card: card.NewService(card.NewRepository(db)),
		Feat: feat.NewService(feat.NewRepository(db)),
	}
}
