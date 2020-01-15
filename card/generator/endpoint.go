package generator

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/chibimi/cards/card/utils"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) GenerateEndpoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cards := strings.Split(r.URL.Query().Get("cards"), ",")
	ids := make([]int, len(cards))
	for i, v := range cards {
		id, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ids[i] = id
	}

	filename, err := s.GeneratePDF(ids, "FR")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	// Open file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()
	defer os.Remove(filename)

	//Set header
	w.Header().Set("Content-type", "application/pdf")

	//Stream to response
	if _, err := io.Copy(w, f); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	utils.WriteJson(w, nil, http.StatusOK)
}
