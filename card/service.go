package card

import (
	"database/sql"

	"gopkg.in/inconshreveable/log15.v2"
)

type Config struct {
}

type Service struct {
	db     *sql.DB
	logger log15.Logger
}

func NewService(db *sql.DB, l log15.Logger) *Service {
	return &Service{
		db:     db,
		logger: l,
	}
}
