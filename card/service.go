package card

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/inconshreveable/log15.v2"
)

type Config struct {
}

type Service struct {
	db     *sqlx.DB
	logger log15.Logger
}

func NewService(db *sqlx.DB, l log15.Logger) *Service {
	return &Service{
		db:     db,
		logger: l,
	}
}
