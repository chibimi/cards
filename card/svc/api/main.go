package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/chibimi/cards/card"
	"github.com/chibimi/cards/card/api"
	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	cfg := struct {
		Login    string `envconfig:"db_login"`
		Password string `envconfig:"db_password"`
		Host     string `envconfig:"db_host"`
		DB       string `envconfig:"db"`
	}{}
	envconfig.Process("card_api", &cfg)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", cfg.Login, cfg.Password, cfg.Host, cfg.DB))
	if err != nil {
		log15.Crit("Unable to access db", "err", err.Error())
	}
	defer db.Close()

	s := api.NewService(card.NewService(db, log15.New()))

	router := httprouter.New()
	router.GET("/abilities", s.ListAbilities)
	router.GET("/abilities/:id", s.GetAbility)
	router.POST("/abilities", s.CreateAbility)
	router.PUT("/abilities/:id", s.UpdateAbility)
	router.DELETE("/abilities/:id", s.DeleteAbility)

	router.GET("/spells", s.ListSpells)
	router.POST("/spells", s.CreateSpell)
	router.PUT("/spells/:id", s.UpdateSpell)

	router.GET("/feats", s.ListFeats)
	router.POST("/feats", s.CreateFeat)
	router.PUT("/feats/:id", s.UpdateFeat)

	router.GET("/cards", s.ListCards)
	router.GET("/cards/:id", s.GetCard)
	router.GET("/cards/:id/related", s.GetRelatedCards)
	router.GET("/cards/:id/abilities", s.GetCardAbilities)
	router.GET("/cards/:id/models", s.GetCardModels)
	router.GET("/cards/:id/spells", s.GetCardSpells)
	router.GET("/cards/:id/feats", s.GetCardFeat)
	router.POST("/cards", s.CreateCard)
	router.PUT("/cards/:id", s.UpdateCard)
	router.DELETE("/cards/:id", s.DeleteCard)

	router.GET("/models/", s.ListModels)
	router.GET("/models/:id", s.GetModel)
	router.GET("/models/:id/abilities", s.GetModelAbilities)
	router.GET("/models/:id/weapons", s.GetModelWeapons)
	router.POST("/models", s.CreateModel)
	router.PUT("/models/:id", s.UpdateModel)
	router.DELETE("/models/:id", s.DeleteModel)

	router.GET("/weapons/", s.ListWeapons)
	router.GET("/weapons/:id", s.GetWeapon)
	router.GET("/weapons/:id/abilities", s.GetWeaponAbilities)
	router.POST("/weapons", s.CreateWeapon)
	router.PUT("/weapons/:id", s.UpdateWeapon)
	router.DELETE("/weapons/:id", s.DeleteWeapon)

	router.PUT("/cards/:id/spells/:spell_id", s.AddCardSpell)
	router.DELETE("/cards/:id/spells/:spell_id", s.DeleteCardSpell)

	router.PUT("/cards/:id/abilities/:ability_id", s.AddCardAbility)
	router.DELETE("/cards/:id/abilities/:ability_id", s.DeleteCardAbility)

	router.PUT("/models/:id/abilities/:ability_id", s.AddModelAbility)
	router.DELETE("/models/:id/abilities/:ability_id", s.DeleteModelAbility)

	router.PUT("/weapons/:id/abilities/:ability_id", s.AddWeaponAbility)
	router.DELETE("/weapons/:id/abilities/:ability_id", s.DeleteWeaponAbility)

	stack := negroni.New()
	stack.Use(cors.AllowAll())
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.UseHandler(router)

	log15.Info("Listening on port: 9901...")
	if err := http.ListenAndServe(":9901", stack); err != nil {
		log15.Crit("Unable to start server", "err", err.Error())
	}
}
