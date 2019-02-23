package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	val := r.URL.Query()
	fmt.Println(val.Get("faction"))
	fmt.Println(val.Get("category"))
}
