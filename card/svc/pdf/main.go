package main

import (
	"fmt"

	"github.com/chibimi/cards/card"
	"github.com/chibimi/cards/card/pdf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	db, err := sqlx.Open("mysql", "cards_api:cards_api@/cards_db")
	if err != nil {
		log15.Crit("Unable to access db", "err", err.Error())
	}
	defer db.Close()
	g := pdf.NewGenerator(card.NewSService(db))
	err = g.GeneratePDF([]int{884, 884, 884, 884, 884, 884}, "FR")
	fmt.Println(err)
}
