package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/chibimi/cards/card/reference"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gographics/imagick/imagick"
	log "github.com/inconshreveable/log15"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type config struct {
	ppURL         string
	destDir       string
	dsn           string
	workers       int
	secondaryOnly bool
}

type DownloadJob struct {
	RefID string
	Index int
}

func main() {
	cfg := config{}

	flag.StringVar(&cfg.dsn, "dsn", "cards_api:cards_api@tcp(localhost:3306)/cards_db", "DSN of the Jackmarshall Database")
	flag.StringVar(&cfg.ppURL, "base-url", "http://cards.privateerpress.com", "URL of the Privateer Press Cards Database")
	flag.StringVar(&cfg.destDir, "dest-dir", ".", "directory to output the images into")
	flag.IntVar(&cfg.workers, "workers", 10, "number of cards to process in parallel")
	flag.BoolVar(&cfg.secondaryOnly, "secondary-only", false, "disable download of the first cards")
	flag.Parse()

	log.Info("starting")
	imagick.Initialize()
	defer imagick.Terminate()

	log.Debug("connecting to database")
	db, err := sqlx.Connect("mysql", cfg.dsn)
	if err != nil {
		log.Error("connecting to database", "err", err)
		os.Exit(1)
	}

	// create job queue
	jobQueue := make(chan DownloadJob)

	// queue front cards for all refs
	if !cfg.secondaryOnly {
		go queueFirstCards(cfg.ppURL, jobQueue)
	}

	// queue front cards for special cases (colossal, character unit, dragoon)
	go queueSpecialCaseCards(db, jobQueue)

	// queue front cards for attachments (makeda & exalted court)
	go queueAttachmentCards(db, jobQueue)

	// process jobQueue
	var wg sync.WaitGroup
	for w := 0; w < cfg.workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			log.Debug("starting worker", "w", w)

			for job := range jobQueue {
				if err := downloadCard(cfg.ppURL, cfg.destDir, job.RefID, job.Index); err != nil {
					log.Error("extract card", "err", err)
					continue
				}
			}
		}(w)
	}
	wg.Wait()
}

func queueFirstCards(ppURL string, jobQueue chan DownloadJob) {
	log.Debug("retrieving PP cards database")
	res, err := http.Get(ppURL)
	if err != nil {
		log.Error("retrieving pp cards database", "err", err)
		os.Exit(1)
	}

	log.Debug("parsing PP cards database")
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Error("parsing pp cards database", "err", err)
		os.Exit(1)
	}

	log.Debug("looking for refs")
	doc.Find("carditem").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		jobQueue <- DownloadJob{
			RefID: s.AttrOr(":card", ""),
			Index: 0,
		}
		return true
	})
}

func queueSpecialCaseCards(db *sqlx.DB, jobQueue chan DownloadJob) {
	log.Debug("looking for special cases")

	stmt := `SELECT id, ppid, special FROM refs WHERE special != ""`
	refs := []reference.Reference{}
	err := db.Select(&refs, stmt)
	if err != nil {
		log.Error("fetch special cases", "err", err)
		os.Exit(1)
	}

	for _, ref := range refs {
		jobQueue <- DownloadJob{
			RefID: strconv.Itoa(ref.PPID),
			Index: 1,
		}
	}
}

func queueAttachmentCards(db *sqlx.DB, jobQueue chan DownloadJob) {
	log.Debug("looking for attachements ")

	stmt := `
		SELECT parent.ppid, parent.category_id, parent.special, ref.id
		FROM (select id, title, linked_to FROM refs WHERE linked_to is not null AND linked_to != 0) as ref
		LEFT JOIN (select id, ppid, category_id, special FROM refs) AS parent ON ref.linked_to = parent.id
		`
	refs := []reference.Reference{}
	err := db.Select(&refs, stmt)
	if err != nil {
		log.Error("fetch attachementss", "err", err)
		os.Exit(1)
	}

	for _, ref := range refs {
		index := 1
		if ref.CategoryID == 1 || ref.CategoryID == 2 || ref.CategoryID == 10 {
			// if the ref is a warcaster/warlock/master there is one more card before the attachement
			index++
		}
		if ref.Special != "" {
			// if the ref if a special case there is one more card before the attachement
			index++
		}
		for _, ref := range refs {
			jobQueue <- DownloadJob{
				RefID: strconv.Itoa(ref.PPID),
				Index: index,
			}
		}
	}
}

func downloadCard(ppURL, destDir, id string, index int) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	log := log.New("ref", id)
	log.Debug("retrieving card")

	res, err := http.Get(fmt.Sprintf("%s/?card_items_to_pdf=$%s,1", ppURL, id))
	if err != nil {
		return errors.Wrap(err, "retrieving card")
	}

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "reading card")
	}

	err = mw.SetOption("density", "300")
	if err != nil {
		log.Error("setting wand options", "err", err, "option", "density", "value", "300")
		return errors.Wrap(err, "wand option density")
	}

	err = mw.SetOption("quality", "100")
	if err != nil {
		log.Error("setting wand options", "err", err, "option", "quality", "value", "100")
		return errors.Wrap(err, "wand option quality")
	}

	err = mw.ReadImageBlob(raw)
	if err != nil {
		errors.Wrap(err, "parsing card")
	}

	err = mw.SetImageFormat("png")
	if err != nil {
		errors.Wrap(err, "setting wand format")
	}

	mw.SetIteratorIndex(0)
	// The coordinates here depends on the DPI used?
	card_h := 1050
	card_w := 748
	space := 7
	err = mw.CropImage(uint(card_w), uint(card_h), 140+(card_w+space)*index, 227)
	if err != nil {
		errors.Wrap(err, "extracting card")
	}

	var path string
	if index == 0 {
		path = fmt.Sprintf("%s/%s.png", destDir, id)
	} else {
		path = fmt.Sprintf("%s/%s_%d.png", destDir, id, index)
	}
	err = mw.WriteImage(path)
	if err != nil {
		errors.Wrap(err, "writing card")
	}
	return nil
}
