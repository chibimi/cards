package main

import (
	"net/http"
	"objenious/card"

	"github.com/jmoiron/sqlx"

	"objenious/card/api"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	db, err := sqlx.Open("mysql", "cards_api:cards_api@/cards_db")
	if err != nil {
		log15.Crit("Unable to access db", "err", err.Error())
	}
	defer db.Close()

	s := api.NewService(card.NewService(db, log15.New()))

	router := httprouter.New()
	router.GET("/factions", s.ListFactions)
	router.GET("/categories", s.ListCategories)
	router.POST("/cards", s.CreateCard)
	router.PUT("/cards/:id", s.UpdateCard)
	router.GET("/cards/:id", s.GetCard)
	router.GET("/cards/", s.ListCards)
	router.GET("/cards/:/related", s.GetRelatedCards)
	router.DELETE("/cards/:id", s.DeleteCard)
	router.POST("/models", s.CreateModel)
	router.PUT("/models/:id", s.UpdateModel)
	router.GET("/models/", s.ListModels)
	router.GET("/models/:id", s.GetModel)
	router.DELETE("/models/:id", s.DeleteModel)
	// router.GET("/models", s.ListModels)
	// router.POST("/models", s.CreateModel)
	// router.PUT("/models/:id", s.UpdateModel)
	// router.GET("/models/:id", s.GetModel)

	handler := cors.AllowAll().Handler(router)

	log15.Info("Listening on port: 9901...")
	if err := http.ListenAndServe(":9901", handler); err != nil {
		log15.Crit("Unable to start server", "err", err.Error())
	}
}
