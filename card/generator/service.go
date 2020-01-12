package generator

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/chibimi/cards/card"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/pkg/errors"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type Service struct {
	src *card.SService
}

func NewService(cards *card.SService) *Service {
	return &Service{
		src: cards,
	}
}

func (s *Service) GeneratePDF(references []int, lang string) (string, error) {
	id := uuid.New()
	src := fmt.Sprintf("%s_source.pdf", id)

	// err := s.downloadFromPP(references, src)
	// if err != nil {
	// 	return "", err
	// }

	// images, err := s.convertToImages(src)
	// if err != nil {
	// 	return "", err
	// }

	// generator.new
	g := NewGenerator(s.src, references, lang)
	g.pdf = gofpdf.New("L", "mm", "letter", "")
	g.unicode = g.pdf.UnicodeTranslatorFromDescriptor("")
	g.pdf.AddPage()
	g.pdf.Image("pp/871c3261-9ddd-4bcf-9737-4b4fb9008105-0.png", 0, 0, 279.4, 215.9, false, "", 0, "")
	g.pdf.AddPage()
	g.pdf.Image("pp/871c3261-9ddd-4bcf-9737-4b4fb9008105-1.png", 0, 0, 279.4, 215.9, false, "", 0, "")
	g.pdf.SetPage(1)

	// generator.init
	// g.initializePDF(images)

	// generator.translate
	err := g.GeneratePDF()
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("%s.pdf", "hello_new")
	err = g.WritePDF(res)
	if err != nil {
		return "", err
	}

	os.Remove(src)
	return res, nil
}

func (s *Service) downloadFromPP(references []int, dest string) error {
	params := ""
	for _, id := range references {
		ref, err := s.src.Ref.Get(id, "")
		if err != nil {
			return errors.Wrap(err, "get ref")
		}
		params = fmt.Sprintf("%s$%s,1", params, ref.FA)
	}

	// Get the data
	resp, err := http.Get(fmt.Sprintf("http://cards.privateerpress.com/?card_items_to_pdf=%s", params))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (s *Service) convertToImages(src string) ([]string, error) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.SetOption("density", "300"); err != nil {
		return nil, errors.Wrap(err, "set density")
	}
	if err := mw.SetOption("quality", "100"); err != nil {
		return nil, errors.Wrap(err, "set quality")
	}
	if err := mw.ReadImage(src); err != nil {
		return nil, errors.Wrap(err, "read image")
	}
	if err := mw.SetImageFormat("png"); err != nil {
		return nil, errors.Wrap(err, "set format")
	}
	res := []string{}
	for i := 0; i < int(mw.GetNumberImages()); i++ {
		mw.SetIteratorIndex(i)
		path := fmt.Sprintf("%s-%d.png", src, i)
		if err := mw.WriteImage(path); err != nil {
			return nil, errors.Wrap(err, "write image")
		}
		res = append(res, path)

	}
	return res, nil
}
