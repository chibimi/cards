package main

import (
	"fmt"
	"net/http"

	"github.com/chibimi/cards/card"
	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	cfg := struct {
		Login       string `envconfig:"db_login"`
		Password    string `envconfig:"db_password"`
		Host        string `envconfig:"db_host"`
		DB          string `envconfig:"db"`
		Port        int    `envconfig:"port" default:"4203"`
		EditorFront string `envconfig:"editor_front"`
		PDFAssets   string `envconfig:"pdf_assets"`
	}{}
	envconfig.Process("card_api", &cfg)

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", cfg.Login, cfg.Password, cfg.Host, cfg.DB))
	if err != nil {
		log15.Crit("Unable to access db", "err", err.Error())
	}
	defer db.Close()

	// s := api.NewService(card.NewService(db, log15.New()))
	ss := card.NewSService(db)

	// generator := generator.NewService(ss, cfg.PDFAssets)

	router := httprouter.New()
	router.POST("/ref", ss.Ref.CreateRef)
	router.GET("/ref", ss.Ref.ListRef)
	router.GET("/status/ref", ss.Ref.ListRefByStatus)
	router.GET("/ref/:id", ss.Ref.GetRef)
	router.PUT("/ref/:id", ss.Ref.SaveRef)
	router.GET("/ref/:id/spell", ss.Spell.ListByRefEndpoint)
	router.PUT("/ref/:id/spell/:spell_id", ss.Spell.AddSpellRefEndpoint)
	router.DELETE("/ref/:id/spell/:spell_id", ss.Spell.DeleteSpellRefEndpoint)
	router.PUT("/ref/:id/feat", ss.Feat.SaveFeat)
	router.GET("/ref/:id/feat", ss.Feat.GetFeat)
	router.GET("/ref/:id/model", ss.Model.ListEndpoint)

	router.POST("/model", ss.Model.CreateEndpoint)
	router.PUT("/model/:id", ss.Model.SaveEndpoint)
	router.DELETE("/model/:id", ss.Model.DeleteEndpoint)
	router.GET("/model/:id/ability", ss.Ability.ListByModel)
	router.PUT("/model/:id/ability/:ability_id", ss.Ability.AddAbilityModel)
	router.DELETE("/model/:id/ability/:ability_id", ss.Ability.DeleteAbilityModel)
	router.GET("/model/:id/weapon", ss.Weapon.ListEndpoint)

	router.POST("/weapon", ss.Weapon.CreateEndpoint)
	router.PUT("/weapon/:id", ss.Weapon.SaveEndpoint)
	router.DELETE("/weapon/:id", ss.Weapon.DeleteEndpoint)
	router.GET("/weapon/:id/ability", ss.Ability.ListByWeapon)
	router.PUT("/weapon/:id/ability/:ability_id", ss.Ability.AddAbilityWeapon)
	router.DELETE("/weapon/:id/ability/:ability_id", ss.Ability.DeleteAbilityWeapon)

	router.GET("/spells/:id", ss.Spell.GetEndpoint)
	router.GET("/spells", ss.Spell.ListEndpoint)
	router.POST("/spells", ss.Spell.CreateEndpoint)
	router.PUT("/spells/:id", ss.Spell.SaveEndpoint)

	router.GET("/abilities/:id", ss.Ability.Get)
	router.GET("/abilities", ss.Ability.List)
	router.POST("/abilities", ss.Ability.Create)
	router.PUT("/abilities/:id", ss.Ability.Save)

	// router.GET("/pdf/generate", generator.GenerateEndpoint)

	router.ServeFiles("/editor/*filepath", http.Dir(cfg.EditorFront))

	stack := negroni.New()
	stack.Use(cors.AllowAll())
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.UseHandler(router)

	log15.Info("Listening...", "port", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), stack); err != nil {
		log15.Crit("Unable to start server", "err", err.Error())
	}
}
