package main

import (
	"flag"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/inconshreveable/log15"
	"github.com/jmoiron/sqlx"
)

func main() {
	var (
		baseURL string
		dsn     string
		workers int
	)

	flag.StringVar(&baseURL, "base-url", "http://cards.privateerpress.com", "URL of the Privateer Press Cards Database")
	flag.StringVar(&dsn, "dsn", "jackmarshall:coconuts@tcp(localhost:3306)/jackmarshall", "DSN of the Jackmarshall Database")
	flag.IntVar(&workers, "workers", 1, "number of cards to process in parallel")
	flag.Parse()

	log.Info("starting")

	log.Debug("connecting to back-end")
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Error("connecting to the back-end database", "err", err)
		os.Exit(1)
	}

	log.Debug("retrieving database")
	res, err := http.Get(baseURL)
	if err != nil {
		log.Error("retrieving database", "err", err)
		os.Exit(1)
	}

	log.Debug("parsing database")
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Error("reading database", "err", err)
		os.Exit(1)
	}

	log.Debug("looking for refs")
	type Ref struct {
		ID      string
		Faction string
		Job     string
		Title   string
	}
	references := make(chan Ref)
	go func() {
		defer close(references)
		doc.Find("carditem").EachWithBreak(func(_ int, s *goquery.Selection) bool {
			references <- Ref{
				ID:      s.AttrOr(":card", ""),
				Faction: s.AttrOr("faction", ""),
				Job:     s.AttrOr("job", ""),
				Title:   s.AttrOr("title", ""),
			}
			return true
		})
	}()

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Debug("starting worker", "w", w)

			for ref := range references {
				log := log.New("ref", ref.ID, "title", ref.Title)
				log.Debug("matching ref")

				_, err := db.Exec("update refs set ppid = ? where title sounds like ?", ref.ID, ref.Title)
				if err != nil {
					log.Error("matching ref", "err", err)
					continue
				}
			}
		}()
	}

	wg.Wait()
}
