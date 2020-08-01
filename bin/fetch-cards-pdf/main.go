package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

type Service struct {
	cfg      config
	db       *sqlx.DB
	jobQueue chan DownloadJob
	errQueue chan error
}

type DownloadJob struct {
	RefID string
	Index int
}

func main() {
	cfg := config{}

	flag.StringVar(&cfg.dsn, "dsn", "cards_api:cards_api@tcp(localhost:3306)/cards_db", "DSN of the Jackmarshall Database")
	flag.StringVar(&cfg.ppURL, "pp-url", "http://cards.privateerpress.com", "URL of the Privateer Press Cards Database")
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

	s := &Service{
		cfg:      cfg,
		db:       db,
		jobQueue: make(chan DownloadJob),
		errQueue: make(chan error),
	}

	// fill job queue
	go func() {
		// queue front cards for all refs
		if !s.cfg.secondaryOnly {
			s.queueFirstCards()
		}

		// queue front cards for special cases (colossal, character unit, dragoon)
		s.queueSpecialCaseCards()

		// queue front cards for attachments (makeda & exalted court)
		s.queueAttachmentCards()

		close(s.jobQueue)
		close(s.errQueue)
	}()

	// process jobQueue
	var wg sync.WaitGroup
	for w := 0; w < cfg.workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			log.Debug("starting worker", "w", w)

			for job := range s.jobQueue {
				err := s.downloadCard(job)
				if err != nil {
					log.Error("downloading card", "err", err, "refId", job.RefID, "index", job.Index)
					continue
				}
			}
		}(w)
	}

	// process errorQueue
	wg.Add(1)
	go func() {
		for err := range s.errQueue {
			log.Error("error while queuing refs", "err", err)
		}
		wg.Done()
	}()

	wg.Wait()
	log.Info("job done")
}

func (s *Service) queueFirstCards() {
	log.Debug("retrieving PP cards database")
	res, err := http.Get(s.cfg.ppURL)
	if err != nil {
		s.errQueue <- errors.Wrap(err, "retrieving pp cards database")
		return
	}

	log.Debug("parsing PP cards database")
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		s.errQueue <- errors.Wrap(err, "parsing pp cards database")
		return
	}

	log.Debug("looking for refs")
	doc.Find("carditem").EachWithBreak(func(_ int, block *goquery.Selection) bool {
		s.jobQueue <- DownloadJob{
			RefID: block.AttrOr(":card", ""),
			Index: 0,
		}
		return true
	})
}

func (s *Service) queueSpecialCaseCards() {
	log.Debug("looking for special cases")
	stmt := `SELECT id, ppid, special FROM refs WHERE special != ""`
	refs := []reference.Reference{}
	err := s.db.Select(&refs, stmt)
	if err != nil {
		s.errQueue <- errors.Wrap(err, "fetching special cases")
		return
	}

	for _, ref := range refs {
		index := 1
		if strings.HasPrefix(ref.Special, "index") {
			index, err = strconv.Atoi(ref.Special[5:])
			if err != nil {
				s.errQueue <- errors.Wrap(err, "converting index")
				continue
			}
		}
		s.jobQueue <- DownloadJob{
			RefID: strconv.Itoa(ref.PPID),
			Index: index,
		}
	}
	log.Info("done queuing special cases", "count", len(refs))
}

func (s *Service) queueAttachmentCards() {
	log.Debug("looking for attachements ")
	stmt := `
		SELECT parent.ppid, parent.category_id, parent.special, ref.id
		FROM refs AS ref
		JOIN refs AS parent ON ref.linked_to = parent.id
		`
	refs := []reference.Reference{}
	err := s.db.Select(&refs, stmt)
	if err != nil {
		s.errQueue <- errors.Wrap(err, "fetching attachements")
		return
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
		if strings.HasPrefix(ref.Special, "index") {
			index, err = strconv.Atoi(ref.Special[5:])
			if err != nil {
				s.errQueue <- errors.Wrap(err, "converting index")
				continue
			}
		}
		s.jobQueue <- DownloadJob{
			RefID: strconv.Itoa(ref.PPID),
			Index: index,
		}
	}
	log.Info("done queuing attachements", "count", len(refs))
}

func (s *Service) downloadCard(job DownloadJob) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	log := log.New("ref", job.RefID)
	log.Debug("retrieving card")

	res, err := http.Get(fmt.Sprintf("%s/?card_items_to_pdf=$%s,1", s.cfg.ppURL, job.RefID))
	if err != nil {
		return errors.Wrap(err, "retrieving card")
	}

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "reading card")
	}

	err = mw.SetOption("density", "300")
	if err != nil {
		return errors.Wrap(err, "wand option density")
	}

	err = mw.SetOption("quality", "100")
	if err != nil {
		return errors.Wrap(err, "wand option quality")
	}

	err = mw.ReadImageBlob(raw)
	if err != nil {
		return errors.Wrap(err, "parsing card")
	}

	err = mw.SetImageFormat("png")
	if err != nil {
		return errors.Wrap(err, "setting wand format")
	}

	mw.SetIteratorIndex(0)
	// The coordinates here depends on the DPI used.
	card_h := 1050
	card_w := 748
	space := 7
	err = mw.CropImage(uint(card_w), uint(card_h), 140+(card_w+space)*job.Index, 227)
	if err != nil {
		return errors.Wrap(err, "extracting card")
	}

	var path string
	if job.Index == 0 {
		path = fmt.Sprintf("%s/%s.png", s.cfg.destDir, job.RefID)
	} else {
		path = fmt.Sprintf("%s/%s_%d.png", s.cfg.destDir, job.RefID, job.Index)
	}
	err = mw.WriteImage(path)
	if err != nil {
		return errors.Wrap(err, "writing card")
	}
	return nil
}
