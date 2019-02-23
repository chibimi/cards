package main

import (
	"fmt"
	"net/http"

	"github.com/chibimi/cards/api"
	log "github.com/inconshreveable/log15"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/cards", api.List) // list all card ?faction=""&type=?

	log.Info("Listening on port: 9901...")
	if err := http.ListenAndServe(":9901", router); err != nil {
		log.Crit(err.Error(), "Unable to start server")
	}
}
