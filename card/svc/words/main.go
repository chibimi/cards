package main

func main() {}

// import (
// 	"fmt"
// 	"sort"
// 	"strings"

// 	"github.com/chibimi/cards/card"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
// 	"gopkg.in/inconshreveable/log15.v2"
// )

// type WordCount struct {
// 	word  string
// 	count int
// }

// func (w WordCount) String() string {
// 	return fmt.Sprintf("%s(%d),  ", w.word, w.count)
// }

// func main() {
// 	db, err := sqlx.Open("mysql", "cards_api:cards_api@/cards_db")
// 	if err != nil {
// 		log15.Crit("Unable to access db", "err", err.Error())
// 	}
// 	defer db.Close()
// 	s := card.NewSService(db)
// 	spells, err := s.Spell.List("UK")
// 	if err != nil {
// 		panic(err)
// 	}
// 	abilities, err := s.Ability.List("UK")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(len(spells))
// 	fmt.Println(len(abilities))
// 	wordCount := map[string]int{}

// 	for _, spell := range spells {

// 		words := strings.Fields(spell.Description)
// 		for _, word := range words {
// 			wordCount[word]++
// 		}
// 	}
// 	for _, ability := range abilities {
// 		words := strings.Fields(ability.Description)
// 		for _, word := range words {
// 			wordCount[word]++
// 		}
// 	}
// 	res := []WordCount{}
// 	for k, v := range wordCount {
// 		res = append(res, WordCount{word: k, count: v})
// 	}
// 	sort.Slice(res, func(i, j int) bool { return res[i].count > res[j].count })

// 	fmt.Println(res)
// }
