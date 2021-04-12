package main

func main() {}

// import (
// 	"fmt"

// 	"github.com/chibimi/cards/card"
// 	"github.com/chibimi/cards/card/generator"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
// 	"gopkg.in/inconshreveable/log15.v2"
// )

// func main() {
// 	db, err := sqlx.Open("mysql", "cards_api:cards_api@/cards_db")
// 	if err != nil {
// 		log15.Crit("Unable to access db", "err", err.Error())
// 	}
// 	defer db.Close()

// 	s := generator.NewService(card.NewSService(db))
// 	res, err := s.GeneratePDF([]int{3, 91, 91, 82, 91}, "FR")
// 	fmt.Println(res, err)
// }
