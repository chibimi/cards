package review

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chibimi/cards/card/utils"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) SaveReview(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	review := &Review{}
	if err := json.NewDecoder(r.Body).Decode(review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	review.CreatedAt = time.Now()
	review.IP = r.Header.Get("X-Forwarded-For")

	err := s.Save(review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, nil, http.StatusCreated)
}
