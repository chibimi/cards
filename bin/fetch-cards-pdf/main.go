package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gographics/imagick/imagick"
	log "github.com/inconshreveable/log15"
)

func main() {
	var (
		baseURL string
		destDir string
		workers int
	)

	flag.StringVar(&baseURL, "base-url", "http://cards.privateerpress.com", "URL of the Privateer Press Cards Database")
	flag.StringVar(&destDir, "dest-dir", ".", "directory to output the images into")
	flag.IntVar(&workers, "workers", 1, "number of cards to process in parallel")
	flag.Parse()

	log.Info("starting")

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
		Name    string
	}
	references := make(chan Ref)
	go func() {
		defer close(references)
		doc.Find("carditem").EachWithBreak(func(_ int, s *goquery.Selection) bool {
			references <- Ref{
				ID:      s.AttrOr(":card", ""),
				Faction: s.AttrOr("faction", ""),
				Job:     s.AttrOr("job", ""),
				Name:    s.AttrOr("title", ""),
			}
			return true
		})
	}()

	log.Debug("initializing imagick")

	imagick.Initialize()
	defer imagick.Terminate()

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Debug("starting worker", "w", w)

			mw := imagick.NewMagickWand()
			defer mw.Destroy()

			for ref := range references {
				log := log.New("ref", ref.ID, "name", ref.Name)
				log.Debug("retrieving card")

				res, err := http.Get(fmt.Sprintf("%s/?card_items_to_pdf=$%s,1", baseURL, ref.ID))
				if err != nil {
					log.Error("retrieving card", "err", err)
					continue
				}

				raw, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Error("reading card", "err", err)
					continue
				}

				mw.Clear()

				err = mw.SetOption("density", "300")
				if err != nil {
					log.Error("setting wand options", "err", err, "option", "density", "value", "300")
					return
				}

				err = mw.SetOption("quality", "100")
				if err != nil {
					log.Error("setting wand options", "err", err, "option", "quality", "value", "100")
					return
				}

				err = mw.ReadImageBlob(raw)
				if err != nil {
					log.Error("parsing card", "err", err)
					continue
				}

				err = mw.SetImageFormat("png")
				if err != nil {
					log.Error("setting wand format", "err", err)
					return
				}

				mw.SetIteratorIndex(0)
				// The coordinates here depends on the DPI used?
				err = mw.CropImage(748, 1050, 140, 227)
				if err != nil {
					log.Error("extracting card", "err", err)
					continue
				}

				err = mw.WriteImage(fmt.Sprintf("%s/%s.png", destDir, ref.ID))
				if err != nil {
					log.Error("writing card", "err", err)
					continue
				}
			}
		}()
	}

	wg.Wait()
}
